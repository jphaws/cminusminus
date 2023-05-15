package ir // Mini

import (
	"fmt"

	"github.com/keen-cp/compiler-project-c/ast"
)

type Block struct {
	function string
	types    []blockType
	sealed   bool
	context  map[string]Value
	Prev     []*Block
	Next     *Block
	Els      *Block
	Phis     []*PhiInstr
	Instrs   []Instr
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

func processFunction(fn *ast.Function, ch chan *Function) {
	// Create entry block
	entry := &Block{
		function: fn.Name,
		types:    make([]blockType, 0, 2),
		Prev:     []*Block{},
		Instrs:   []Instr{},
	}

	entry.types = append(entry.types, &fnEntryBlock{})

	// Initialize entry variables
	instrs, locals, params := functionInitLlvm(fn)
	entry.Instrs = append(entry.Instrs, instrs...)

	// Create exit block
	exit := &Block{
		function: fn.Name,
		types:    make([]blockType, 0, 1),
		Prev:     []*Block{},
		Instrs:   []Instr{},
	}

	exit.types = append(exit.types, &fnExitBlock{})

	// Add a dummy return statement to the end of void functions
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		fn.Body = append(fn.Body, &ast.ReturnStatement{}) // No position required
	}

	// Process function body
	end, _ := processStatements(fn.Body, entry, exit, locals, 0)

	// Compress final and end blocks (if needed)
	if len(exit.Prev) == 1 {
		// Add extra fnexit type to the final block
		end = exit.Prev[0]
		end.types = append(end.types, &fnExitBlock{})

		// Remove unneeded branch instruction
		if len(end.Instrs) > 0 {
			end.Instrs = end.Instrs[:len(end.Instrs)-1]
		}

		// Remove links to old exit block
		end.Next = nil
		exit.Prev = []*Block{}
		exit = end
	}

	// Add return instructions
	instrs = functionFiniLlvm(fn, locals)
	exit.Instrs = append(exit.Instrs, instrs...)

	// Create IR function wrapper
	ret := &Function{
		Parameters: params,
		ReturnType: typeToLlvm(fn.ReturnType),
		Registers:  locals,
		Cfg:        entry,
	}

	ch <- ret
}

func processStatements(stmts []ast.Statement, curr *Block,
	funcExit *Block, locals map[string]*Register, count int) (b *Block, rcount int) {

	rcount = count

	for _, s := range stmts {
		switch stmt := s.(type) {
		case *ast.IfStatement:
			curr, count = processIfStatement(stmt, curr, funcExit, locals, count+1)
			rcount = count

			// Bail out if return equivalent
			if curr == nil {
				b = curr
				return
			}

		case *ast.WhileStatement:
			curr, count = processWhileStatement(stmt, curr, funcExit, locals, count+1)

		case *ast.ReturnStatement:
			curr.Next = funcExit
			funcExit.Prev = append(funcExit.Prev, curr)

			instrs := returnStatementToLlvm(stmt, funcExit, locals)
			curr.Instrs = append(curr.Instrs, instrs...)

			b = nil
			return

		case *ast.BlockStatement:
			curr, count = processStatements(stmt.Statements, curr, funcExit, locals, count)

			// Bail out if return equivalent
			if curr == nil {
				b = curr
				return
			}

		default:
			curr.Instrs = append(curr.Instrs, statementToLlvm(stmt, locals)...)
		}
	}

	b = curr
	return
}

func processIfStatement(fi *ast.IfStatement, curr *Block,
	funcExit *Block, locals map[string]*Register, count int) (b *Block, rcount int) {

	var thenEntry, thenExit, elseEntry, elseExit, ifExit *Block
	fn := curr.function

	// Add extra ifentry type to the current block
	curr.types = append(curr.types, &ifEntryBlock{count})

	// Create initial then block (prev: curr)
	thenEntry = createBlock(fn, &thenBlock{count}, []*Block{curr})
	curr.Next = thenEntry

	// Add guard instructions
	guardInstrs, guardVal := createGuardLlvm(fi.Guard, locals)
	curr.Instrs = append(curr.Instrs, guardInstrs...)

	// Process then statements
	thenExit, rcount = processStatements(fi.Then.Statements, thenEntry, funcExit, locals, count)

	// When there is no else block, link current block to ifexit
	if fi.Else == nil {
		// Create ifexit block
		ifExit = createBlock(fn, &ifExitBlock{count}, []*Block{thenExit, curr})
		elseEntry = ifExit
		curr.Els = ifExit

	} else {
		// Otherwise, create initial else block (prev: curr)
		elseEntry = createBlock(fn, &elseBlock{count}, []*Block{curr})
		curr.Els = elseEntry

		// Process else statements
		elseExit, rcount = processStatements(fi.Else.Statements, elseEntry, funcExit, locals, rcount)

		// Create ifexit block
		ifExit = createBlock(fn, &ifExitBlock{count}, []*Block{thenExit, elseExit})
	}

	// Create guard branch
	curr.Instrs = append(curr.Instrs, createBranch(guardVal, thenEntry, elseEntry))

	// Check if both bodies are return equivalent
	if thenExit == nil && fi.Else != nil && elseExit == nil {
		return
	}

	// Link then/else exit blocks to overall ifexit block
	if thenExit != nil {
		thenExit.Next = ifExit

		thenExit.Instrs = append(thenExit.Instrs, createJump(ifExit))
	}

	if elseExit != nil {
		elseExit.Next = ifExit

		elseExit.Instrs = append(elseExit.Instrs, createJump(ifExit))
	}

	b = ifExit
	return
}

func processWhileStatement(whl *ast.WhileStatement, curr *Block,
	funcExit *Block, locals map[string]*Register, count int) (whileExit *Block, rcount int) {

	fn := curr.function

	// Add extra whlguard type to the current block
	curr.types = append(curr.types, &whileGuardBlock{count})

	// Add guard instructions
	guardInstrs, guardVal := createGuardLlvm(whl.Guard, locals)
	curr.Instrs = append(curr.Instrs, guardInstrs...)

	// Create whlexit block (prev: curr, dynamic)
	whileExit = createBlock(fn, &whileExitBlock{count}, make([]*Block, 0, 2))
	whileExit.Prev = append(whileExit.Prev, curr)
	curr.Els = whileExit

	// Create initial while block (prev: curr, dynamic)
	whileEntry := createBlock(fn, &whileBlock{count}, make([]*Block, 0, 2))
	whileEntry.Prev = append(whileEntry.Prev, curr)
	curr.Next = whileEntry

	// Add branch instruction
	curr.Instrs = append(curr.Instrs, createBranch(guardVal, whileEntry, whileExit))

	// Process while statements
	whileGuard, rcount := processStatements(whl.Body.Statements, whileEntry, funcExit, locals, count)

	// Check if body is return equivalent
	if whileGuard != nil {
		// Add extra whlguard type to the final body block
		whileGuard.types = append(whileGuard.types, &whileGuardBlock{count})

		// Add guard instructions
		guardInstrs, guardVal = createGuardLlvm(whl.Guard, locals)
		whileGuard.Instrs = append(whileGuard.Instrs, guardInstrs...)

		// Create backedge
		whileGuard.Next = whileEntry
		whileEntry.Prev = append(whileEntry.Prev, whileGuard)

		// Link guard to whlexit
		whileGuard.Els = whileExit
		whileExit.Prev = append(whileExit.Prev, whileGuard)

		whileGuard.Instrs = append(whileGuard.Instrs, createBranch(guardVal, whileEntry, whileExit))
	}

	return
}

func createBlock(function string, typ blockType, prev []*Block) *Block {
	ret := &Block{
		function: function,
		types:    make([]blockType, 0, 2),
		Prev:     prev,
		Instrs:   []Instr{},
	}

	ret.types = append(ret.types, typ)

	return ret
}
