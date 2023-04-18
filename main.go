package main

import (
	"fmt"
	"os"

	// "github.com/alecthomas/repr"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/keen-cp/compiler-project-c/ast"
	"github.com/keen-cp/compiler-project-c/parser"
	"github.com/keen-cp/compiler-project-c/parser/mantlr"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := mantlr.NewMiniLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := mantlr.NewMiniParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	prog := p.Program()

	root := parser.MiniToAst(prog)
	/*
		fmt.Println("===== AST =====")
		repr.Println(root, repr.Indent("   "))
	*/

	err := ast.TypeCheck(&root)
	/*
		fmt.Println("===== Struct table =====")
		repr.Println(ast.StructTable)
		fmt.Println("\n===== Symbol table =====")
		repr.Println(ast.SymbolTable)
		fmt.Println("\n===== Errors =====")
	*/
	fmt.Println(err)
}
