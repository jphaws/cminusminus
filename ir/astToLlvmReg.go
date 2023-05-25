package ir

import (
	"github.com/keen-cp/compiler-project-c/ast"
)

// === Statements ===
func assignmentStatementToLlvmReg(asgn *ast.AssignmentStatement,
	curr *Block, locals map[string]*Register) []Instr {

	// Process lvalue
	instrs, ptr, name := lValueToLlvmReg(asgn.Target, curr, locals)

	// Assignment is a read
	if asgn.Source == nil {
		noPointer := false

		// Create a read pointer for use with local name lvalues
		if ptr == nil {
			noPointer = true
			ptr = allocRead(locals)
		}

		// Read into pointer location
		instrs = append(instrs, readToLlvm(ptr)...)

		if noPointer {
			// Create read result register
			reg := &Register{
				Name: nextRegName(),
				Type: &IntType{64},
			}
			locals[reg.Name] = reg

			// Load scanned value from pointer to register
			load := &LoadInstr{
				Reg: reg,
				Mem: ptr,
			}
			instrs = append(instrs, load)
			addDefUse(load)

			// Update block context
			curr.context[name] = reg
		}

		// Assignment is not a read
	} else {
		// Evaluate expression
		exprInstrs, val := expressionToLlvm(asgn.Source, curr, locals, false)
		instrs = append(instrs, exprInstrs...)

		// Store to pointer (if possible: dot lvalues, globals)
		if ptr != nil {
			store := &StoreInstr{
				Mem: ptr,
				Reg: val,
			}
			addDefUse(store)

			instrs = append(instrs, store)

			// Otherwise, update block context
		} else {
			curr.context[name] = val
		}
	}

	return instrs
}

func lValueToLlvmReg(lval ast.LValue, curr *Block,
	locals map[string]*Register) (instrs []Instr, ptr Value, name string) {

	switch v := lval.(type) {
	case *ast.NameLValue:
		name = v.Name

		// Look up globals in the symbol table
		if _, ok := locals[name]; !ok {
			ptr = symbolTable[name]
			return
		}

	case *ast.DotLValue:
		var dotInstrs []Instr
		dotInstrs, ptr = dotLValueToLlvmReg(v, curr, locals, 0)
		instrs = append(instrs, dotInstrs...)
	}

	return
}

func dotLValueToLlvmReg(dot *ast.DotLValue, curr *Block,
	locals map[string]*Register, level int) (instrs []Instr, val Value) {

	var base Value

	switch v := dot.Left.(type) {
	case *ast.NameLValue:
		// Do a simple lookup for final name lvalue
		instrs, base = lookupSymbolReg(v.Name, curr, locals)

	case *ast.DotLValue:
		// Otherwise, recurse until a name lvalue is reached
		instrs, base = dotLValueToLlvmReg(v, curr, locals, level+1)
	}

	// Get field pointer and load (except on the outermost recursive level)
	var loadInstrs []Instr
	loadInstrs, val = loadField(base, dot.Name, locals, level)
	instrs = append(instrs, loadInstrs...)

	return
}

func loadField(base Value, fieldName string,
	locals map[string]*Register, level int) (instrs []Instr, field *Register) {

	// Get field pointer
	gepInstr, ptr := getFieldPointer(base, fieldName, locals)
	instrs = append(instrs, gepInstr)

	// Allow overriding the load
	if level == 0 {
		field = ptr
		return
	}

	// Create field register
	name := nextRegName()
	field = &Register{
		Name: name,
		Type: ptr.GetType().(*PointerType).TargetType,
	}
	locals[name] = field

	// Load field from field pointer
	load := &LoadInstr{
		Reg: field,
		Mem: ptr,
	}
	addDefUse(load)

	instrs = append(instrs, load)

	return
}

func allocRead(locals map[string]*Register) *Register {
	// Create read pointer (if it doesn't already exit)
	if _, ok := locals[readPtrName]; !ok {
		reg := &Register{
			Name: "%" + readPtrName,
			Type: &PointerType{&IntType{64}},
		}
		locals[readPtrName] = reg
	}

	return locals[readPtrName]
}

func returnStatementToLlvmReg(ret *ast.ReturnStatement,
	curr *Block, funcExit *Block, locals map[string]*Register) []Instr {

	jump := createJump(funcExit)

	// Check for a void return
	if ret.Expression == nil {
		return []Instr{jump}
	}

	// Process expression
	instrs, val := expressionToLlvm(ret.Expression, curr, locals, false)
	curr.context[retValName] = val

	return append(instrs, jump)
}

// === Expressions ===
func identifierExpressionToLlvmReg(ident *ast.IdentifierExpression,
	curr *Block, locals map[string]*Register) (instrs []Instr, val Value) {

	return lookupSymbolReg(ident.Name, curr, locals)
}

func lookupSymbolReg(name string, curr *Block,
	locals map[string]*Register) (instrs []Instr, val Value) {

	// Get local identifiers from this block or previous block
	if v, ok := locals[name]; ok {
		val = findValue(name, curr, locals, v.GetType())

		// Otherwise, get and load a global
	} else {
		val = symbolTable[name]

		// Create load target register
		reg := &Register{
			Name: nextRegName(),
			Type: val.GetType().(*PointerType).TargetType,
		}
		locals[reg.Name] = reg

		// Load global into register
		load := &LoadInstr{
			Reg: reg,
			Mem: val,
		}
		addDefUse(load)

		instrs = []Instr{load}
		val = reg
	}

	return
}

// === Functions ===
func functionInitLlvmReg(fn *ast.Function,
	curr *Block) (locals map[string]*Register, params []*Register) {

	// Allocate space for local variables and parameters
	locals = make(map[string]*Register, len(fn.Parameters)+len(fn.Locals))
	params = make([]*Register, 0, len(fn.Parameters))

	// Handle parameters
	for _, v := range fn.Parameters {
		// Create parameter register
		reg := &Register{
			Name: "%" + v.Name,
			Type: typeToLlvm(v.Type),
		}
		locals[v.Name] = reg
		params = append(params, reg)

		curr.context[v.Name] = reg
	}

	// Handle locals
	for _, v := range fn.Locals {
		reg := &Register{
			Name: "%" + v.Name,
			Type: typeToLlvm(v.Type),
		}
		locals[v.Name] = reg
	}

	return
}

func functionFiniLlvmReg(fn *ast.Function, funcEntry *Block,
	curr *Block, locals map[string]*Register) []Instr {

	// Allocate read pointer (if needed)
	if v, ok := locals[readPtrName]; ok {
		alloc := &AllocInstr{v}
		funcEntry.Allocs = append(funcEntry.Allocs, alloc)
		addDefUse(alloc)
	}

	// Remove registers for local variables (these are unused in the future, at least in reg-based)
	for _, v := range fn.Locals {
		delete(locals, v.Name)
	}

	// Return nothing for a void function
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		return []Instr{&RetInstr{}}
	}

	// Get return value
	retVal := findValue(retValName, curr, locals, typeToLlvm(fn.ReturnType))

	// Create return instruction
	ret := &RetInstr{retVal}
	addDefUse(ret)

	return []Instr{ret}
}

func findValue(name string, curr *Block, locals map[string]*Register, typ Type) Value {
	if v, ok := curr.context[name]; ok {
		return v
	}

	return findValueFromPrevs(name, curr, locals, typ)
}

func findValueFromPrevs(name string, curr *Block, locals map[string]*Register, typ Type) Value {
	var val Value

	if !curr.isSealed() {
		// Handle unsealed blocks

		// Create target register
		reg := &Register{
			Name: nextRegName(),
			Type: typ,
		}
		locals[reg.Name] = reg

		// Create phi instruction
		phi := &PhiInstr{
			Target: reg,
		}
		curr.incompletePhis[name] = phi

		val = reg

	} else if len(curr.Prev) == 0 {
		// Handle sealed blocks (0 previous) by using a default value
		def := "0"
		if _, ok := typ.(*PointerType); ok {
			def = "null"
		}

		val = &Literal{
			Value: def,
			Type:  typ,
		}

	} else if len(curr.Prev) == 1 {
		// Handle sealed blocks (1 previous) by recursing into prev
		val = findValue(name, curr.Prev[0], locals, typ)

	} else {
		// Handle sealed blocks (2+ previous) by adding a phi and recursing into all prev

		// Create target register
		reg := &Register{
			Name: nextRegName(),
			Type: typ,
		}
		locals[reg.Name] = reg

		// Create phi instruction
		phi := &PhiInstr{
			Target: reg,
		}
		curr.Phis = append(curr.Phis, phi)

		// Update context first to prevent loops
		curr.context[name] = reg

		addPhiValues(phi, name, curr, locals, typ)
		val = reg
	}

	curr.context[name] = val

	return val
}

func addPhiValues(phi *PhiInstr, name string, curr *Block, locals map[string]*Register, typ Type) {
	// Get possible values from previous blocks (recursively)
	for _, v := range curr.Prev {
		phiVal := findValue(name, v, locals, typ)

		phi.Values = append(phi.Values, &PhiVal{
			Value: phiVal,
			Block: v,
		})
	}

	addDefUse(phi)
}

func completePhis(curr *Block, locals map[string]*Register) {
	for name, phi := range curr.incompletePhis {
		addPhiValues(phi, name, curr, locals, phi.Target.GetType())

		curr.Phis = append(curr.Phis, phi)
	}
}
