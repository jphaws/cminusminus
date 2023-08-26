package ir

import (
	"fmt"
	"strings"

	"github.com/jphaws/cminusminus/ast"

	om "github.com/wk8/go-ordered-map/v2"
)

// (struct) 'Id' -> Struct
var structTable = map[string]*Struct{}

// 'Name' -> Register
var symbolTable = map[string]*Register{}

// 'Name' -> Function
var functionTable = map[string]*Function{}

type ProgramIr struct {
	Structs   map[string]*Struct
	Globals   map[string]*Register
	Functions map[string]*Function
}

type Struct struct {
	Fields *om.OrderedMap[string, *Field]
	Size   int
}

type Field struct {
	Type  Type
	Index int
}

type Function struct {
	Parameters []*Register
	ReturnType Type
	Registers  map[string]*Register
	Instrs     map[Instr]*Block
	Cfg        *Block
}

var stackLlvm = false

const (
	intSize        = 8
	pointerSize    = 8
	PrintStrName   = "_print"
	PrintlnStrName = "_println"
	ScanStrName    = "_scan"
)

func (p ProgramIr) ToLlvm() string {
	// Add target details
	ret := "target triple = \"x86_64-pc-linux-gnu\"\n\n"

	// Declare library functions
	ret += "declare ptr @malloc(i64)\n"
	ret += "declare void @free(ptr)\n"
	ret += "declare i32 @printf(ptr, ...)\n"
	ret += "declare i32 @scanf(ptr, ...)\n\n"

	// Define format strings
	ret += "@" + PrintStrName +
		" = private unnamed_addr constant [5 x i8] c\"%ld \\00\", align 1\n"

	ret += "@" + PrintlnStrName +
		" = private unnamed_addr constant [5 x i8] c\"%ld\\0a\\00\", align 1\n"

	ret += "@" + ScanStrName +
		" = private unnamed_addr constant [4 x i8] c\"%ld\\00\", align 1\n\n"

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
		ret += functionToLlvm(k, v) + "\n\n"
	}

	return ret
}

func structToLlvm(id string, st *Struct) string {
	fields := st.Fields

	// Start with "struct" type declaration
	ret := fmt.Sprintf("%%struct.%v = type {", id)

	// Create strings for each field type
	fieldStrs := make([]string, 0, fields.Len())

	for pair := fields.Oldest(); pair != nil; pair = pair.Next() {
		fieldStrs = append(fieldStrs, fmt.Sprintf("%v", pair.Value.Type))
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
		init = "null"
	}

	// Create declaration
	return fmt.Sprintf("@%v = global %v %v, align 8", name, typ.TargetType, init)
}

func functionToLlvm(name string, fn *Function) string {
	// Create strings for each parameter
	params := make([]string, 0, len(fn.Parameters))

	for _, param := range fn.Parameters {
		pTyp := param.Type
		pName := param.Name
		params = append(params, fmt.Sprintf("%v %%%v", pTyp, pName))
	}

	// Create declaration
	ret := fmt.Sprintf("define %v @%v(%v) {\n",
		fn.ReturnType, name, strings.Join(params, ", "))
	ret += fn.Cfg.toLlvm()
	ret += "}"

	return ret
}

var visitedLlvm = map[*Block]bool{}

func (b *Block) toLlvm() string {
	visitedLlvm[b] = true

	// Process the current block
	ret := b.Label() + ":\n"

	for _, v := range b.Phis {
		ret += fmt.Sprintf("    %v\n", v)
	}

	for _, v := range b.Allocs {
		ret += fmt.Sprintf("    %v\n", v)
	}

	for _, v := range b.Instrs {
		ret += fmt.Sprintf("    %v\n", v)
	}

	// Process next block
	if b.Next != nil && !visitedLlvm[b.Next] {
		ret += b.Next.toLlvm()
	}

	// Process else block
	if b.Els != nil && !visitedLlvm[b.Els] {
		ret += b.Els.toLlvm()
	}

	return ret
}

func (p ProgramIr) UseDef() string {
	var ret string

	// Loop through each function, printing local use-def information
	for name, fn := range p.Functions {
		ret += "=== " + name + " ===\n"

		for _, reg := range fn.Registers {
			ret += regUseDefLocal(reg, fn.Parameters)
		}
	}

	// Loop through globals, printing their use-def information
	ret += "=== Globals ===\n"

	for _, reg := range p.Globals {
		ret += regUseDefGlobal(reg)
	}

	return ret
}

func regUseDefLocal(reg *Register, params []*Register) string {
	var ret string

	// Handle register definition
	if reg.Def != nil {
		ret += fmt.Sprintf("%v\n", reg.Def)
	} else {
		isParam := false

		// Check if this local is a parameter
		for _, p := range params {
			if reg == p {
				isParam = true
				break
			}
		}

		// Determine "not defined" message (parameter vs regular local)
		if isParam {
			ret += fmt.Sprintf("%v is parameter (not defined)\n", reg.Name)
		} else {
			ret += fmt.Sprintf("%v not defined\n", reg.Name)
		}
	}

	// Handle register uses
	for use := range reg.Uses {
		ret += fmt.Sprintf("    %v\n", use)
	}

	return ret + "\n"
}

func regUseDefGlobal(reg *Register) string {
	var ret string

	ret += fmt.Sprintf("%v is global (not defined)\n", reg.Name)

	for use := range reg.Uses {
		ret += fmt.Sprintf("    %v\n", use)
	}

	return ret + "\n"
}

func CreateIr(root *ast.Root, tables *ast.Tables, stack bool) *ProgramIr {
	funcChan := make(chan *Function)
	stackLlvm = stack

	// Convert struct table (llvm-ify)
	for k, v := range tables.StructTable {
		omap := om.New[string, *Field]()

		i := 0
		for pair := v.Oldest(); pair != nil; pair = pair.Next() {
			omap.Set(pair.Key, &Field{
				Type:  typeToLlvm(pair.Value),
				Index: i,
			})

			i++
		}

		structTable[k] = &Struct{
			Fields: omap,
			Size:   computeStructSize(omap),
		}
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
			Name:   k,
			Type:   &PointerType{typeToLlvm(v)},
			Global: true,
		}
	}

	// Create a Go routine to generate each CFG
	for _, v := range root.Functions {
		go processFunction(v, funcChan)
	}

	// Synchronize completed routines
	fns := make([]*Function, 0, len(root.Functions))
	for range root.Functions {
		fns = append(fns, <-funcChan)
	}

	// Create function table
	for _, v := range fns {
		functionTable[v.Cfg.function] = v
	}

	return &ProgramIr{
		Structs:   structTable,
		Globals:   symbolTable,
		Functions: functionTable,
	}
}

func computeStructSize(omap *om.OrderedMap[string, *Field]) int {
	size := 0

	for pair := omap.Oldest(); pair != nil; pair = pair.Next() {
		typ := pair.Value.Type

		if _, ok := typ.(*IntType); ok {
			size += intSize
		} else {
			size += pointerSize
		}
	}

	return size
}
