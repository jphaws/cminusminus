package asm

import (
	"fmt"

	"github.com/keen-cp/compiler-project-c/ir"
)

type ProgramAsm struct {
	Globals   map[string]string
	Functions map[string]*Function
}

type Function struct {
	paramOffset  int
	localOffset  int
	calleeOffset int
	spillOffset  int
	Blocks       []*Block
}

var globals = map[string]string{}

var constantMap = map[string]string{}

func (p ProgramAsm) ToAsm() string {
	// Add target details
	ret := ".arch armv8-a\n"

	// Set main as global
	ret += ".global main\n\n"

	// Define constants in the rodata section
	ret += ".section .rodata\n"
	ret += ".align 8\n"
	ret += "$" + ir.PrintStrName + ":\n    .asciz \"%ld \"\n"
	ret += "$" + ir.PrintlnStrName + ":\n    .asciz \"%ld\\n\"\n"
	ret += "$" + ir.ScanStrName + ":\n    .asciz \"%ld\"\n\n"

	// Switch to the bss section
	ret += ".section .bss\n"
	ret += fmt.Sprintf(".align %v\n", dataSize)

	// Define globals
	for _, glob := range p.Globals {
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

	for _, v := range b.PhiOuts {
		ret += fmt.Sprintf("    %v\n", v)
	}

	for _, v := range b.Terminals {
		ret += fmt.Sprintf("    %v\n", v)
	}

	return ret
}

func CreateAsm(program *ir.ProgramIr) *ProgramAsm {
	funcChan := make(chan *Function)

	// Populate globals map
	for _, v := range program.Globals {
		globals[v.Name] = "$" + v.Name
	}

	// Add constants to globals map
	globals["@"+ir.PrintStrName] = "$_print"
	globals["@"+ir.PrintlnStrName] = "$_println"
	globals["@"+ir.ScanStrName] = "$_scan"

	// Create a Go routine for each function
	for k, v := range program.Functions {
		go processFunction(v, k, funcChan)
	}

	// Synchronize completed routines
	fns := make(map[string]*Function, len(program.Functions))
	for k := range program.Functions {
		fns[k] = <-funcChan
	}

	// Delete constants from globals map
	delete(globals, "@"+ir.PrintStrName)
	delete(globals, "@"+ir.PrintlnStrName)
	delete(globals, "@"+ir.ScanStrName)

	return &ProgramAsm{
		Globals:   globals,
		Functions: fns,
	}
}
