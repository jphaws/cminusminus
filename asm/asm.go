package asm

import (
	"fmt"

	"github.com/keen-cp/compiler-project-c/ir"
)

type ProgramAsm struct {
	// Globals   map[string]*Register // TODO Handle globals
	Functions map[string]*Function
}

type Function struct {
	paramOffset  int
	localOffset  int
	calleeOffset int
	spillOffset  int
	Blocks       []*Block
}

func (p ProgramAsm) ToAsm() string {
	// Add target details
	ret := ".arch armv8-a\n"

	// Mark functions as global
	for k := range p.Functions {
		ret += ".global " + k + "\n"
	}

	ret += "\n"

	// Set section
	ret += ".section .text\n"
	ret += ".align 4\n\n"

	// Define functions
	for _, v := range p.Functions {
		for _, block := range v.Blocks {
			ret += block.toAsm()
		}

		ret += "\n"
	}

	return ret
}

func (b *Block) toAsm() string {
	ret := b.Label + ":\n"

	for _, v := range b.Instrs {
		ret += fmt.Sprintf("    %v\n", v)
	}

	return ret
}

func CreateAsm(program *ir.ProgramIr) *ProgramAsm {
	funcChan := make(chan *Function)

	// Create a Go routine for each function
	for k, v := range program.Functions {
		go processFunction(v, k, funcChan)
	}

	// Synchronize completed routines
	fns := make(map[string]*Function, len(program.Functions))
	for k := range program.Functions {
		fns[k] = <-funcChan
	}

	return &ProgramAsm{fns} // TODO Add named fields if handling globals
}
