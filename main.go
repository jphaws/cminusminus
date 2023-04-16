package main

import (
	"os"

	"github.com/alecthomas/repr"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/keen-cp/compiler-project-c/parser/mantlr"
	"github.com/keen-cp/compiler-project-c/parser"
	"github.com/keen-cp/compiler-project-c/ast"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := mantlr.NewMiniLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := mantlr.NewMiniParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	prog := p.Program()

	var mr ast.Root
	mr = parser.MiniToAst(prog)
	repr.Println(mr, repr.Indent("   "))
}
