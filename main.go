package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/keen-cp/compiler-project-c/ast"
	"github.com/keen-cp/compiler-project-c/color"
	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/parser"
	"github.com/keen-cp/compiler-project-c/parser/mantlr"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var lines []string
var syntaxErrors = false

func main() {
	// Parse options/arguments
	opts, args := parseArgs()

	// Read input file
	var err error
	lines, err = readFile(args[0])
	if err != nil {
		fmt.Printf("mc: %v\n", err)
		os.Exit(3)
	}

	// Create ANTLR input stream
	input := antlr.NewInputStream(strings.Join(lines, "\n"))

	// Create lexer
	lexer := mantlr.NewMiniLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(miniErrorListener{})

	// Create token stream (from lexer)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	parsr := mantlr.NewMiniParser(stream)
	parsr.RemoveErrorListeners()
	parsr.AddErrorListener(miniErrorListener{})

	// Build parse tree
	parsr.BuildParseTrees = true
	prog := parsr.Program()

	if syntaxErrors {
		os.Exit(1)
	}

	// Convert parse tree to AST
	root := parser.MiniToAst(prog)

	// Type check AST
	tables, err := ast.TypeCheck(root, lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create IR
	rep := ir.CreateIr(root, tables)

	// Generate output string
	var output string
	if opts.graph {
		output = rep.ToDot()
	} else if opts.defUse {
		output = rep.UseDef()
	} else {
		output = rep.ToLlvm()
	}

	// Write output
	if opts.outputFile != "" {
		err = writeOutputFile(opts.outputFile, output)
	} else {
		err = writeString(os.Stdout, output)
	}

	if err != nil {
		fmt.Printf("mc: %v\n", err)
		os.Exit(3)
	}
}

type Options struct {
	outputFile string
	graph      bool
	defUse     bool
	stackIr    bool
}

func parseArgs() (opts Options, args []string) {
	opts = Options{}

	// Set optional flags
	flags := flag.NewFlagSet("mc", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: mc [options] file\nOptions:\n")
		flags.PrintDefaults()
	}

	flags.StringVar(&opts.outputFile, "o", "", "output to `filename`")
	flags.BoolVar(&opts.stackIr, "stack", false, "use a stack-based intermediate representation")
	flags.BoolVar(&opts.graph, "graph", false, "output a control flow graph in the dot language")
	flags.BoolVar(&opts.defUse, "def-use", false, "output def-use chains for each function")

	// Parse flags
	err := flags.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	args = flags.Args()

	// Check for correct number of positional arguments
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "mc: no input file specified")
		os.Exit(2)
	}

	if opts.graph && opts.defUse {
		fmt.Fprintln(os.Stderr, "mc: only one output mode allowed")
		os.Exit(2)
	}

	return
}

func readFile(filename string) (lines []string, err error) {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	// Close file (at end of function)
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for read error
	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func writeOutputFile(filename string, str string) error {
	// Create file
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	// Close file (at end of function)
	defer file.Close()

	// Write to file
	if err = writeString(file, str); err != nil {
		return err
	}

	return nil
}

func writeString(w io.Writer, str string) error {
	// Write to file
	writer := bufio.NewWriter(w)
	_, err := writer.WriteString(str)

	if err != nil {
		return err
	}

	// Flush output
	if err = writer.Flush(); err != nil {
		return err
	}

	return nil
}

type miniErrorListener struct {
	*antlr.DefaultErrorListener
}

func (m miniErrorListener) SyntaxError(rec antlr.Recognizer,
	sym interface{}, line, co int, msg string, e antlr.RecognitionException) {

	col := co + 1
	syntaxErrors = true

	fmt.Printf("%v:%v: syntax error: %v", line, col, msg)
	fmt.Printf("\n %4v | %s%s%s%s\n      |\n",
		line, color.Red, color.Bright, lines[line-1], color.Reset)
}
