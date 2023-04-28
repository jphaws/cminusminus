package cfg // Mini

import (
	"fmt"

	"github.com/keen-cp/compiler-project-c/ast"
)

type Block struct {
	function string
	types    []blockType
	Prev     []*Block
	Next     *Block
	Els      *Block
	Instrs   []ast.Statement
}

func (b *Block) Label() string {
	ret := fmt.Sprintf("%v", b.function)

	for _, typ := range b.types {
		ret += fmt.Sprintf("_%v", typ)
	}

	return ret
}

type blockType interface {
	blockTypeFunc()
}

type fnEntryBlock struct{}

func (f fnEntryBlock) blockTypeFunc() {}

func (f fnEntryBlock) String() string {
	return fmt.Sprintf("entry")
}

type ifEntryBlock struct {
	num int
}

func (i ifEntryBlock) blockTypeFunc() {}

func (i ifEntryBlock) String() string {
	return fmt.Sprintf("ifentry%v", i.num)
}

type thenBlock struct {
	num int
}

func (t thenBlock) blockTypeFunc() {}

func (t thenBlock) String() string {
	return fmt.Sprintf("then%v", t.num)
}

type elseBlock struct {
	num int
}

func (e elseBlock) blockTypeFunc() {}

func (e elseBlock) String() string {
	return fmt.Sprintf("else%v", e.num)
}

type ifExitBlock struct {
	num int
}

func (i ifExitBlock) blockTypeFunc() {}

func (i ifExitBlock) String() string {
	return fmt.Sprintf("ifexit%v", i.num)
}

type whileGuardBlock struct {
	num int
}

func (w whileGuardBlock) blockTypeFunc() {}

func (w whileGuardBlock) String() string {
	return fmt.Sprintf("whlguard%v", w.num)
}

type whileBlock struct {
	num int
}

func (w whileBlock) blockTypeFunc() {}

func (w whileBlock) String() string {
	return fmt.Sprintf("whlentry%v", w.num)
}

type whileExitBlock struct {
	num int
}

func (w whileExitBlock) blockTypeFunc() {}

func (w whileExitBlock) String() string {
	return fmt.Sprintf("whlexit%v", w.num)
}

type fnExitBlock struct{}

func (f fnExitBlock) blockTypeFunc() {}

func (f fnExitBlock) String() string {
	return fmt.Sprintf("exit")
}

func CreateCfg(root *ast.Root) []*Block {
	blockChan := make(chan *Block)
	ret := make([]*Block, len(root.Functions))

	// Generate CFG for each function
	for i := range root.Functions {
		go functionCfg(root.Functions[i], blockChan)
	}

	// Synchronize completed functions
	for i := range root.Functions {
		ret[i] = <-blockChan
	}

	return ret
}

func functionCfg(fn *ast.Function, ch chan *Block) {
	entry := &Block{
		function: fn.Name,
		types:    make([]blockType, 1),
		Prev:     []*Block{},
		Instrs:   make([]ast.Statement, 0),
	}

	entry.types[0] = &fnEntryBlock{}

	exit := &Block{
		function: fn.Name,
		types:    make([]blockType, 1),
		Prev:     make([]*Block, 0),
		Instrs:   make([]ast.Statement, 0),
	}

	exit.types[0] = &fnExitBlock{}

	processStatements(fn.Body, entry, exit, 0)
	ch <- entry
}

func processStatements(stmts []ast.Statement, curr *Block,
	funcExit *Block, count int) (b *Block, rcount int) {

	rcount = count

	// fmt.Printf("%v: stmts %v\n", curr.Label(), stmts)

	for _, s := range stmts {
		switch stmt := s.(type) {
		case *ast.IfStatement:
			// fmt.Printf("%T at %v\n", stmt, stmt.Position)
			curr, count = processIfStatement(stmt, curr, funcExit, count+1)
			rcount = count

			// Bail out if return equivalent
			if curr == nil {
				b = curr
				return
			}

		case *ast.WhileStatement:
			curr, count = processWhileStatement(stmt, curr, funcExit, count+1)

		case *ast.ReturnStatement:
			// fmt.Printf("%T at %v\n", stmt, stmt.Position)
			curr.Next = funcExit
			funcExit.Prev = append(funcExit.Prev, curr)
			b = nil
			return

		default:
			// fmt.Printf("%T\n", stmt)
			curr.Instrs = append(curr.Instrs, stmt)
		}
	}

	b = curr
	return
}

func processIfStatement(fi *ast.IfStatement, curr *Block,
	funcExit *Block, count int) (b *Block, rcount int) {

	var thenExit, elseExit, ifExit *Block
	fn := curr.function

	// Add extra ifentry type to the current block
	curr.types = append(curr.types, &ifEntryBlock{count})

	// TODO: Add guard instructions

	// Create initial then block (prev: curr)
	thenEntry := createBlock(fn, &thenBlock{count}, []*Block{curr})
	curr.Next = thenEntry

	// Process then statements
	thenExit, rcount = processStatements(fi.Then.Statements, thenEntry, funcExit, count)

	// When there is no else block, link current block to ifexit
	if fi.Else == nil {
		// Create ifexit block
		ifExit = createBlock(fn, &ifExitBlock{count}, []*Block{thenExit, curr})
		curr.Els = ifExit

	} else {
		// Otherwise, create initial else block (prev: curr)
		elseEntry := createBlock(fn, &elseBlock{count}, []*Block{curr})
		curr.Els = elseEntry

		// Process else statements
		elseExit, rcount = processStatements(fi.Else.Statements, elseEntry, funcExit, rcount)

		// Create ifexit block
		ifExit = createBlock(fn, &ifExitBlock{count}, []*Block{thenExit, elseExit})
	}

	// Check if both bodies are return equivalent
	if thenExit == nil && elseExit == nil {
		return
	}

	// Link then/else exit blocks to overall ifexit block
	if thenExit != nil {
		thenExit.Next = ifExit
	}

	if elseExit != nil {
		elseExit.Next = ifExit
	}

	b = ifExit
	return
}

func processWhileStatement(whl *ast.WhileStatement, curr *Block,
	funcExit *Block, count int) (whileExit *Block, rcount int) {

	fn := curr.function

	// Add extra whlguard type to the current block
	curr.types = append(curr.types, &whileGuardBlock{count})

	// TODO: Add guard instructions

	// Create whlexit block (prev: curr, dynamic)
	whileExit = createBlock(fn, &whileExitBlock{count}, make([]*Block, 1))
	whileExit.Prev[0] = curr
	curr.Els = whileExit

	// Create initial while block (prev: curr, dynamic)
	whileEntry := createBlock(fn, &whileBlock{count}, make([]*Block, 1))
	whileEntry.Prev[0] = curr
	curr.Next = whileEntry

	// Process while statements
	whileGuard, rcount := processStatements(whl.Body.Statements, whileEntry, funcExit, count)

	// Check if body is return equivalent
	if whileGuard != nil {
		// Add extra whlguard type to the final body block
		whileGuard.types = append(whileGuard.types, &whileGuardBlock{count})

		// Create backedge
		whileGuard.Next = whileEntry
		whileEntry.Prev = append(whileEntry.Prev, whileGuard)

		// Link guard to whlexit
		whileGuard.Els = whileExit
		whileExit.Prev = append(whileExit.Prev, whileGuard)
	}

	return
}

func createBlock(function string, typ blockType, prev []*Block) *Block {
	ret := &Block{
		function: function,
		types:    make([]blockType, 1),
		Prev:     prev,
		Instrs:   make([]ast.Statement, 0),
	}

	ret.types[0] = typ

	return ret
}
