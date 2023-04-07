package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
   "github.com/keen-cp/compiler-project-c/parser"
	"fmt"
)

type TreeShapeListener struct {
	*parser.BaseMiniListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewMiniLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewMiniParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.Program())
}
