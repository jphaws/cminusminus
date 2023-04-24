package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "github.com/alecthomas/repr"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/keen-cp/compiler-project-c/ast"
	"github.com/keen-cp/compiler-project-c/parser"
	"github.com/keen-cp/compiler-project-c/parser/mantlr"
)

func main() {
	// Check for correct number of arguments
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "No input file specified")
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
	lines := make([]string, 1)

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

	// Create token stream
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

	err = ast.TypeCheck(root, lines)
	/*
		fmt.Println("===== Struct table =====")
		repr.Println(ast.StructTable)
		fmt.Println("\n===== Symbol table =====")
		repr.Println(ast.SymbolTable)
		fmt.Println("\n===== Errors =====")
	*/
	if err != nil {
		fmt.Println(err)
	}
}
