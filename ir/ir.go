package ir

import (
	"fmt"
	"strings"

	"github.com/keen-cp/compiler-project-c/ast"

	om "github.com/wk8/go-ordered-map/v2"
)

// (struct) 'ID' -> [field 'Name' -> Type]
var structTable = make(map[string]*om.OrderedMap[string, Type])

// 'Name' -> Type
var symbolTable = make(map[string]*Register)

type ProgramIr struct {
	Structs   map[string]*om.OrderedMap[string, Type]
	Globals   map[string]*Register
	Functions []*Function
}

type Function struct {
	Name       string
	Parameters []*Register
	ReturnType Type
	Cfg        *Block
}

func (p ProgramIr) ToLlvm() string {
	// Add target details
	ret := "target triple = \"x86_64-pc-linux-gnu\"\n\n"

	// Declare library functions
	ret += "declare ptr @malloc(i64)\n"
	ret += "declare void @free(ptr)\n\n"

	// Declare structs
	for k, v := range p.Structs {
		ret += fmt.Sprintf("%%struct.%v = type {", k)
		fields := make([]string, v.Len())

		i := 0
		for pair := v.Oldest(); pair != nil; pair = pair.Next() {
			fields[i] = fmt.Sprintf("%v", pair.Value)
			i++
		}

		ret += strings.Join(fields, ", ")
		ret += "}\n"
	}
	ret += "\n"

	// Declare/define globals
	for k, v := range p.Globals {
		typ := v.GetType().(*PointerType)

		// Skip functions
		if _, ok := typ.TargetType.(*FunctionType); ok {
			continue
		}

		init := "0"

		if _, ok := typ.TargetType.(*PointerType); ok {
			init = "inttoptr(i64 0 to ptr)"
		}
		ret += fmt.Sprintf("@%v = global %v %v, align 8\n", k, typ.TargetType, init)
	}
	ret += "\n"

	// Define functions
	for _, v := range p.Functions {
		params := make([]string, len(v.Parameters))
		for i, param := range v.Parameters {
			pTyp := param.Type
			pName := param.Name
			params[i] = fmt.Sprintf("%v %v", pTyp, pName)
		}

		ret += fmt.Sprintf("define %v @%v(%v) {\n",
			v.ReturnType, v.Name, strings.Join(params, ", "))
		ret += v.Cfg.toLlvm()
		ret += "}\n\n"
	}

	return ret
}

func (b *Block) toLlvm() string {
	ret := b.Label() + ":\n"

	for _, v := range b.Instrs {
		ret += fmt.Sprintf("  %v\n", v)
	}

	ret += "\n"
	return ret
}

func CreateIr(root *ast.Root, tables *ast.Tables) *ProgramIr {
	funcChan := make(chan *Function)
	fns := make([]*Function, len(root.Functions))

	// Convert struct table (llvm-ify)
	for k, v := range tables.StructTable {
		omap := om.New[string, Type]()

		for pair := v.Oldest(); pair != nil; pair = pair.Next() {
			omap.Set(pair.Key, typeToLlvm(pair.Value))
		}

		structTable[k] = omap
	}

	// Convert symbol table
	for k, v := range tables.SymbolTable {
		symbolTable[k] = &Register{
			Name: "@" + k,
			Type: &PointerType{typeToLlvm(v)},
		}
	}

	// Create a Go routine to generate each CFG
	for _, v := range root.Functions {
		go processFunction(v, funcChan)
	}

	// Synchronize completed routines
	for i := range root.Functions {
		fns[i] = <-funcChan
	}

	return &ProgramIr{
		Structs:   structTable,
		Globals:   symbolTable,
		Functions: fns,
	}
}

func (p ProgramIr) ToDot() string {
	// return CreateGraph(p.Cfgs)
	return "a"
}
