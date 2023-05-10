package ir

import (
	"fmt"
	"strings"

	"github.com/keen-cp/compiler-project-c/ast"

	om "github.com/wk8/go-ordered-map/v2"
)

// (struct) 'Id' -> [field 'Name' -> Type]
var structTable = make(map[string]*om.OrderedMap[string, Type])

// 'Name' -> Register
var symbolTable = make(map[string]*Register)

// 'Name' -> Function
var functionTable = make(map[string]*Function)

type ProgramIr struct {
	Structs   map[string]*om.OrderedMap[string, Type]
	Globals   map[string]*Register
	Functions map[string]*Function
}

type Function struct {
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
		ret += structToLlvm(k, v) + "\n"
	}
	ret += "\n"

	// Declare/define globals
	for k, v := range p.Globals {
		ret += globalToLlvm(k, v) + "\n"
	}
	ret += "\n"

	// Define functions
	for k, v := range p.Functions {
		ret += functionToLlvm(k, v)
	}

	return ret
}

func structToLlvm(id string, fields *om.OrderedMap[string, Type]) string {
	// Start with "struct" type declaration
	ret := fmt.Sprintf("%%struct.%v = type {", id)

	// Create strings for each field type
	fieldStrs := make([]string, fields.Len())

	i := 0
	for pair := fields.Oldest(); pair != nil; pair = pair.Next() {
		fieldStrs[i] = fmt.Sprintf("%v", pair.Value)
		i++
	}

	// Add all field types to declaration
	ret += strings.Join(fieldStrs, ", ")
	ret += "}"

	return ret
}

func globalToLlvm(name string, reg *Register) string {
	typ := reg.GetType().(*PointerType)

	// Choose default value based on type
	init := "0"

	if _, ok := typ.TargetType.(*PointerType); ok {
		init = "inttoptr(i64 0 to ptr)"
	}

	// Create declaration
	return fmt.Sprintf("@%v = global %v %v, align 8", name, typ.TargetType, init)
}

func functionToLlvm(name string, fn *Function) string {
	// Create strings for each parameter
	params := make([]string, len(fn.Parameters))

	for i, param := range fn.Parameters {
		pTyp := param.Type
		pName := param.Name
		params[i] = fmt.Sprintf("%v %v", pTyp, pName)
	}

	// Create declaration
	ret := fmt.Sprintf("define %v @%v(%v) {\n",
		fn.ReturnType, name, strings.Join(params, ", "))
	ret += fn.Cfg.toLlvm()
	ret += "}"

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
		// Pre-populate return types for functions (in the function table)
		if fn, ok := v.(*ast.FunctionType); ok {
			functionTable[k] = &Function{
				ReturnType: typeToLlvm(fn.ReturnType),
			}
			continue
		}

		// Add non-function types to the symbol table
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
	for range root.Functions {
		fn := <-funcChan
		functionTable[fn.Cfg.function] = fn
	}

	return &ProgramIr{
		Structs:   structTable,
		Globals:   symbolTable,
		Functions: functionTable,
	}
}
