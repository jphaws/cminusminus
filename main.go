package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/keen-cp/compiler-project-c/ast"
	"github.com/keen-cp/compiler-project-c/color"
	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/parser"
	"github.com/keen-cp/compiler-project-c/parser/mantlr"

	// "github.com/alecthomas/repr"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var lines []string
var syntaxErrors = false

type MiniErrorListener struct {
	*antlr.DefaultErrorListener
}

func (m MiniErrorListener) SyntaxError(rec antlr.Recognizer,
	sym interface{}, line, co int, msg string, e antlr.RecognitionException) {

	col := co + 1
	syntaxErrors = true

	fmt.Printf("%v:%v: syntax error: %v", line, col, msg)
	fmt.Printf("\n %4v | %s%s%s%s\n      |\n", line, color.Red, color.Bright, lines[line-1], color.Reset)
}

func main() {
	// Check for correct number of arguments
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "error: no input file specified")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	// Close file (at end of function)
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for read error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Create ANTLR input stream
	input := antlr.NewInputStream(strings.Join(lines, "\n"))

	// Create lexer
	lexer := mantlr.NewMiniLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(MiniErrorListener{})

	// Create token stream (from lexer)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	parsr := mantlr.NewMiniParser(stream)
	parsr.RemoveErrorListeners()
	parsr.AddErrorListener(MiniErrorListener{})

	// Build parse tree
	parsr.BuildParseTrees = true
	prog := parsr.Program()

	if syntaxErrors {
		os.Exit(3)
	}

	// Convert parse tree to AST
	root := parser.MiniToAst(prog)

	// Type check AST
	tables, err := ast.TypeCheck(root, lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	// Create IR
	rep := ir.CreateIr(root, tables)

	// fmt.Println(rep.ToDot())

	fmt.Println(rep.ToLlvm())
}
