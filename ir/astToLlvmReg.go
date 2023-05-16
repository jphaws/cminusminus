package ir

import (
	"github.com/keen-cp/compiler-project-c/ast"
)

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

	// Get local identifiers from this block or previous block
	if v, ok := locals[ident.Name]; ok {
		val = findValue(ident.Name, curr, locals, v.GetType())

		// Otherwise, get and load a global
	} else {
		mem := symbolTable[ident.Name]

		// Create load target register
		name := nextRegName()
		reg := &Register{
			Name: name,
			Type: mem.GetType().(*PointerType).TargetType,
		}
		locals[name] = reg

		// Load global into register
		load := &LoadInstr{
			Reg: reg,
			Mem: mem,
		}
		addDefUse(load)

		instrs = []Instr{load}
		val = reg
	}

	return
}

/*
// TODO: Rename
func lookupSymbolReg(name string, curr *Block, locals map[string]*Register, typ Type) Value {
	var ret Value

	if _, ok := locals[name]; ok {
		ret = findValue(name, curr, locals, typ)
	} else {
		ret = symbolTable[name]
	}

	return ret
}
*/

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

func functionFiniLlvmReg(fn *ast.Function, curr *Block, locals map[string]*Register) []Instr {
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

// The start of THE ALGORITHM
func findValue(name string, curr *Block, locals map[string]*Register, typ Type) Value {
	if v, ok := curr.context[name]; ok {
		return v
	}

	return findValueFromPrevs(name, curr, locals, typ)

}

func findValueFromPrevs(name string, curr *Block, locals map[string]*Register, typ Type) Value {
	var val Value

	if curr.sealed { // TODO: Change to `!curr.sealed`
		// Handle unsealed blocks
		panic("Block not sealed!")

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

		// Create target register (type not set yet)
		name := nextRegName()
		reg := &Register{
			Name: name,
			Type: typ,
		}
		locals[name] = reg

		// Create phi instruction
		phi := &PhiInstr{
			Target: reg,
		}
		curr.Phis = append(curr.Phis, phi)

		// Update context first to prevent loops
		curr.context[name] = reg

		// Get possible values from previous blocks (recursively)
		for _, v := range curr.Prev {
			phiVal := findValue(name, v, locals, typ)

			phi.Values = append(phi.Values, &PhiVal{
				Value: phiVal,
				Block: v,
			})
		}

		addDefUse(phi)
		val = reg
	}

	curr.context[name] = val

	return val
}
