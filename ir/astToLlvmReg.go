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
	instrs, val := expressionToLlvm(ret.Expression, locals, false)
	curr.context[retValName] = val

	return append(instrs, jump)
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

func functionFiniLlvmReg(fn *ast.Function, curr *Block, locals map[string]*Register) []Instr {
	// Return nothing for a void function
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		return []Instr{&RetInstr{}}
	}

	// Get return value
	retVal := findValue(retValName, curr)

	// Create return instruction
	ret := &RetInstr{retVal}
	addDefUse(ret)

	return []Instr{ret}
}

// The start of THE ALGORITHM
func findValue(name string, curr *Block) Value {
	if v, ok := curr.context[name]; ok {
		return v
	}

	panic("Not found in current context")
}
