package ir

import (
	"github.com/keen-cp/compiler-project-c/ast"
)

// === Statements ===
func assignmentStatementToLlvmStack(asgn *ast.AssignmentStatement,
	curr *Block, locals map[string]*Register) []Instr {

	instrs, target := lValueToLlvmStack(asgn.Target, locals)

	// Read from stdin (if read assignment)
	if asgn.Source == nil {
		readInstrs := readToLlvm(target)
		instrs = append(instrs, readInstrs...)
		return instrs
	}

	// Otherwise, process expression
	expInstrs, val := expressionToLlvm(asgn.Source, curr, locals, false)
	instrs = append(instrs, expInstrs...)

	store := &StoreInstr{
		Mem: target,
		Reg: val,
	}
	addDefUse(store)

	return append(instrs, store)
}

func lValueToLlvmStack(lval ast.LValue,
	locals map[string]*Register) (instrs []Instr, reg *Register) {

	switch v := lval.(type) {
	case *ast.NameLValue:
		// Base case
		reg = lookupSymbolStack(v.Name, locals)

	case *ast.DotLValue:
		// Recurse on left side
		var ptr *Register
		instrs, ptr = lValueToLlvmStack(v.Left, locals)

		// Create base register (pointer to start of struct)
		name := nextRegName()
		base := &Register{
			Name: name,
			Type: ptr.GetType().(*PointerType).TargetType,
		}
		locals[name] = base

		// Load base pointer
		load := &LoadInstr{
			Reg: base,
			Mem: ptr,
		}
		addDefUse(load)

		instrs = append(instrs, load)

		// Get field pointer
		var gepInstr *GepInstr
		gepInstr, reg = getFieldPointer(base, v.Name, locals)
		instrs = append(instrs, gepInstr)
	}

	return
}

func returnStatementToLlvmStack(ret *ast.ReturnStatement, curr *Block,
	funcExit *Block, locals map[string]*Register) []Instr {

	jump := createJump(funcExit)

	// Check for a void return
	if ret.Expression == nil {
		return []Instr{jump}
	}

	// Process expression
	instrs, val := expressionToLlvm(ret.Expression, curr, locals, false)

	// Store expression value and jump to exit block
	store := &StoreInstr{
		Mem: locals[retPtrName],
		Reg: val,
	}
	addDefUse(store)

	return append(instrs, store, jump)
}

// === Expressions ===
func identifierExpressionToLlvmStack(ident *ast.IdentifierExpression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	mem := lookupSymbolStack(ident.Name, locals)

	name := nextRegName()
	reg := &Register{
		Name: name,
		Type: mem.GetType().(*PointerType).TargetType,
	}
	locals[name] = reg

	load := &LoadInstr{
		Reg: reg,
		Mem: mem,
	}
	addDefUse(load)

	instrs = []Instr{load}
	val = reg

	return
}

func lookupSymbolStack(name string, locals map[string]*Register) *Register {
	if r, ok := locals[name]; ok {
		return r
	} else {
		return symbolTable[name]
	}
}

// === Functions ===
func functionInitLlvmStack(fn *ast.Function) (instrs []Instr,
	locals map[string]*Register, params []*Register) {

	// Allocate space for local variables and parameters
	instrs = make([]Instr, 0, 2*len(fn.Parameters)+len(fn.Locals))
	locals = make(map[string]*Register, len(fn.Parameters)+len(fn.Locals))
	params = make([]*Register, 0, len(fn.Parameters))

	// Handle parameters
	for _, v := range fn.Parameters {
		// Create parameter register
		pReg := &Register{
			Name: "%_p" + v.Name,
			Type: typeToLlvm(v.Type),
		}
		locals["_p"+v.Name] = pReg

		// Create parameter stack pointer
		reg := &Register{
			Name: "%" + v.Name,
			Type: &PointerType{typeToLlvm(v.Type)},
		}
		locals[v.Name] = reg
		params = append(params, pReg)

		// Allocate space for parameter on the stack
		alloc := &AllocInstr{reg}
		addDefUse(alloc)

		// Store parameter on the stack
		store := &StoreInstr{
			Mem: reg,
			Reg: pReg,
		}
		addDefUse(store)

		instrs = append(instrs, alloc, store)
	}

	// Handle locals
	for _, v := range fn.Locals {
		reg := &Register{
			Name: "%" + v.Name,
			Type: &PointerType{typeToLlvm(v.Type)},
		}
		locals[v.Name] = reg

		alloc := &AllocInstr{reg}
		addDefUse(alloc)

		instrs = append(instrs, alloc)
	}

	// Create return register (if needed)
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		return
	}

	retPtr := &Register{
		Name: "%" + retPtrName,
		Type: &PointerType{typeToLlvm(fn.ReturnType)},
	}
	locals[retPtrName] = retPtr

	alloc := &AllocInstr{retPtr}
	addDefUse(alloc)

	instrs = append(instrs, alloc)

	return
}

func functionFiniLlvmStack(fn *ast.Function, locals map[string]*Register) []Instr {
	// Store the return value in a register (if needed)
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		return []Instr{&RetInstr{}}
	}

	retVal := &Register{
		Name: "%" + retValName,
		Type: typeToLlvm(fn.ReturnType),
	}
	locals[retVal.Name] = retVal

	// Create load and return instructions
	load := &LoadInstr{
		Reg: retVal,
		Mem: locals[retPtrName],
	}
	addDefUse(load)

	ret := &RetInstr{retVal}
	addDefUse(ret)

	return []Instr{load, ret}
}
