package asm

import (
	"fmt"

	"github.com/keen-cp/compiler-project-c/ir"
)

type ProgramAsm struct {
	Globals   map[string]struct{}
	Functions map[string]*Function
}

type Function struct {
	paramOffset  int
	localOffset  int
	calleeOffset int
	spillOffset  int
	Blocks       []*Block
}

var globals = map[string]struct{}{}

var constants = map[string]string{
	"$_print":   "%ld \\000",
	"$_println": "%ld\\012\\000",
	"$_scan":    "%ld\\000",
}

func (p ProgramAsm) ToAsm() string {
	// Add target details
	ret := ".arch armv8-a\n"

	// Set main as global
	ret += ".global main\n\n"

	// Switch to the rodata section
	ret += ".section .rodata\n"
	ret += fmt.Sprintf(".align %v\n", dataSize)

	// Define constants
	for k, v := range constants {
		ret += k + ":\n"
		ret += fmt.Sprintf("    .ascii \"%v\"\n", v)
	}
	ret += "\n"

	// Switch to the bss section
	ret += ".section .bss\n"
	ret += fmt.Sprintf(".align %v\n", dataSize)

	// Define globals
	for glob := range p.Globals {
		ret += fmt.Sprintf(".lcomm %v, %v\n", glob, dataSize)
	}
	ret += "\n"

	// Switch to the text section
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

	// Populate globals map
	for _, v := range program.Globals {
		globals["$"+v.Name] = struct{}{}
	}

	/* TODO: Why is this here?
	for k, v := range constants {
		globals[k] =
	}
	*/

	// Create a Go routine for each function
	for k, v := range program.Functions {
		go processFunction(v, k, funcChan)
	}

	// Synchronize completed routines
	fns := make(map[string]*Function, len(program.Functions))
	for k := range program.Functions {
		fns[k] = <-funcChan
	}

	return &ProgramAsm{
		Globals:   globals,
		Functions: fns,
	}
}
