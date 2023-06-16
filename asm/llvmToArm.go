package asm

import (
	"fmt"
	"strconv"

	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/util"
)

const (
	dataSize = 8

	paramOffset = 16 // Offset from fp
	localOffset = 0  // Offset from sp

	maxRegisterParams = 8

	arithImmediateMin = -4096
	arithImmediateMax = 4096
	movImmediateMin   = -65536
	movImmediateMax   = 65536
)

var tmpNum = 0

type functionInfo struct {
	epiBlock    *Block
	blockMap    map[*ir.Block]*Block
	registers   map[string]*Register
	stackVars   map[string]stackOffset
	spillOffset int
}

type stackOffset struct {
	base   *Register
	offset int
}

func processFunction(fn *ir.Function, name string, doRegAlloc bool, ch chan *Function) {
	info := &functionInfo{
		blockMap:  map[*ir.Block]*Block{},
		registers: map[string]*Register{},
		stackVars: map[string]stackOffset{},
	}

	// Pre-populate general registers
	genRegs, specRegs := populatePhysicalRegs(info)

	// Pre-populate register and stack maps with parameters
	paramInstrs := populateParams(fn.Parameters, info)

	// Pre-populate stack map with local stack variables
	stackPointerOffset := populateLocals(fn.Cfg.Allocs, info)

	// Set spill offset to top of local section
	info.spillOffset = stackPointerOffset

	// Collect a slice of all ARM blocks and map CFG -> ARM blocks
	blocks := collectBlock(fn.Cfg, info)

	// Process all blocks
	for irBlock, armBlock := range info.blockMap {
		processBlock(irBlock, armBlock, info)
	}

	// Create parameter block
	paramBlock := &Block{
		Label:  "." + name + "_param",
		Instrs: paramInstrs,
		Next:   []*Block{blocks[0]},
	}
	blocks[0].Prev = append(blocks[0].Prev, paramBlock)

	// Allocate physical registers (unless not requested)
	var callees []*Register
	if doRegAlloc {
		coloring := allocateRegisters(paramBlock, genRegs, specRegs)

		// TODO: Make this better?
		for old, nw := range coloring {
			old.Name = nw.Name
		}

		// Get callee register slice
		callees = findCallees(coloring, genRegs)
		stackPointerOffset += len(callees) * dataSize
	}

	// Rectify stack pointer to a 16-byte boundary
	if stackPointerOffset%16 != 0 {
		stackPointerOffset += dataSize
	}

	// Create prologue block
	proBlock := &Block{
		Label:  name,
		Instrs: createPrologue(callees, info, stackPointerOffset),
		Next:   []*Block{paramBlock},
	}
	proBlock.Prev = append(proBlock.Prev, paramBlock)

	// Add epilogue instructions to return block
	epiBlock := info.epiBlock
	epiBlock.Instrs = append(epiBlock.Instrs, createEpilogue(callees, info)...)

	// Create function wrapper
	ret := &Function{
		Blocks: []*Block{proBlock, paramBlock},
	}
	ret.Blocks = append(ret.Blocks, blocks...)

	ch <- ret
}

func populateParams(params []*ir.Register, info *functionInfo) []Instr {
	var instrs []Instr
	numReg := util.Min(len(params), maxRegisterParams)

	// Partition register parameters (x0-x7) vs. stack parameters
	regParams := params[:numReg]
	stackParams := params[numReg:]

	// Move register parameters into temporary registers
	for i, v := range regParams {
		dst := nextTempReg(info.registers)

		mov := &MovInstr{
			Dst: dst,
			Src: findRegister(fmt.Sprintf("x%v", i), info),
		}
		addUses(mov)
		instrs = append(instrs, mov)

		//Create name->register mappings for register parameters
		info.registers[v.Name] = dst
	}

	// Create name->offset mappings for stack parameters
	for i, v := range stackParams {
		info.stackVars[v.Name] = stackOffset{
			base:   findRegister("fp", info),
			offset: paramOffset + i*dataSize,
		}
	}

	return instrs
}

func populateLocals(allocs []*ir.AllocInstr, info *functionInfo) int {
	// Create name->offset mappings for stack locals
	for i, v := range allocs {
		info.stackVars[v.Target.Name] = stackOffset{
			base:   findRegister("sp", info),
			offset: localOffset + i*dataSize,
		}
	}

	return len(allocs) * dataSize
}

func collectBlock(b *ir.Block, info *functionInfo) []*Block {
	// Create new ASM block and add to the block map
	curr := &Block{
		Label: "." + b.Label(),
	}
	info.blockMap[b] = curr

	// Create block slice
	blocks := []*Block{curr}

	// Collect next block (if it exists)
	if b.Next != nil {
		if _, ok := info.blockMap[b.Next]; !ok {
			blocks = append(blocks, collectBlock(b.Next, info)...)
		}

		// Add next block to ARM block successors (and link this block as a predecessor)
		nextArm := info.blockMap[b.Next]
		curr.Next = append(curr.Next, nextArm)
		nextArm.Prev = append(nextArm.Prev, curr)
	}

	// Process else block (if it exists)
	if b.Els != nil {
		if _, ok := info.blockMap[b.Els]; !ok {
			blocks = append(blocks, collectBlock(b.Els, info)...)
		}

		// Add else block to ARM block successors
		nextArm := info.blockMap[b.Els]
		curr.Next = append(curr.Next, nextArm)
		nextArm.Prev = append(nextArm.Prev, curr)
	}

	return blocks
}

func processBlock(irBlock *ir.Block, armBlock *Block, info *functionInfo) {
	// Translate phi instructions to ARM
	for _, v := range irBlock.Phis {
		armBlock.Instrs = append(armBlock.Instrs, phiInstrToArm(v, info)...)
	}

	// Translate all (non-magic) LLVM instructions to ARM
	for _, v := range irBlock.Instrs {
		// Use EndInstrs slice for branch-related instruction
		if branch, ok := v.(*ir.BranchInstr); ok {
			armBlock.Terminals = append(armBlock.Terminals, branchInstrToArm(branch, info)...)
			break
		}

		// Use Instrs slice for all other instructions
		armBlock.Instrs = append(armBlock.Instrs, instrToArm(v, armBlock, info)...)
	}
}

func instrToArm(instr ir.Instr, curr *Block, info *functionInfo) []Instr {
	switch v := instr.(type) {
	case *ir.LoadInstr:
		return loadInstrToArm(v, info)
	case *ir.StoreInstr:
		return storeInstrToArm(v, info)
	case *ir.CallInstr:
		return callInstrToArm(v, info)
	case *ir.RetInstr:
		info.epiBlock = curr
		return retInstrToArm(v, info)
	case *ir.CompInstr:
		return compInstrToArm(v, info)
	case *ir.BinaryInstr:
		return binaryInstrToArm(v, info)
	case *ir.ConvInstr:
		return convInstrToArm(v, info)
	}

	return nil
}

func loadInstrToArm(ld *ir.LoadInstr, info *functionInfo) []Instr {
	dst := findRegister(ld.Reg.Name, info)

	// Check if the load is from a global
	if glob, ok := ld.Mem.(*ir.Register); ok && glob.Global {
		return loadGlobalToArm(dst, glob, info)
	}

	var instrs []Instr
	var base *Register
	var offset int

	// Find base register and offset for load
	switch v := ld.Mem.(type) {
	case *ir.Literal:
		// For immediates, mov/load a new base register and use 0 offset
		instrs, base = movLoadImmediate(nil, v, info)
		offset = 0

	case *ir.Register:
		// For registers, look for a GepInstr to check for its base and offset
		var baseInstrs []Instr
		baseInstrs, base, offset = getLoadStoreBase(v, info)
		instrs = append(instrs, baseInstrs...)
	}

	// Create load instruction
	load := &LoadInstr{
		Dst:    dst,
		Base:   base,
		Offset: offset,
	}
	addUses(load)

	return append(instrs, load)
}

func loadGlobalToArm(dst *Register, glob *ir.Register, info *functionInfo) []Instr {
	// Create temporary base register to hold page address
	base := nextTempReg(info.registers)

	// Get page aligned-address of global
	label := globals[glob.Name]
	adrp := &PageAddressInstr{
		Dst:   base,
		Label: label,
	}
	addUses(adrp)

	// Load global using page offset
	load := &LoadInstr{
		Dst:        dst,
		Base:       base,
		PageOffset: label,
	}
	addUses(load)

	return []Instr{adrp, load}
}

func globalAddrToArm(dst *Register, glob *ir.Register, info *functionInfo) []Instr {
	// Create temporary base register to hold page address
	base := nextTempReg(info.registers)

	// Get page aligned-address of global
	label := globals[glob.Name]
	adrp := &PageAddressInstr{
		Dst:   base,
		Label: label,
	}
	addUses(adrp)

	// Add page offset to get global address
	add := &ArithInstr{
		Operator:   AddOperator,
		Dst:        dst,
		Src1:       base,
		PageOffset: label,
	}
	addUses(add)

	return []Instr{adrp, add}
}

func storeInstrToArm(st *ir.StoreInstr, info *functionInfo) []Instr {
	var instrs []Instr
	var src *Register

	// Set source for the store depending on register vs. literal
	switch v := st.Reg.(type) {
	case *ir.Register:
		instrs, src = useLoadRegister(v.Name, info)
	case *ir.Literal:
		instrs, src = movLoadImmediate(nil, v, info)
	}

	// Check if the store is to a global
	if glob, ok := st.Mem.(*ir.Register); ok && glob.Global {
		globInstrs := storeGlobalToArm(src, glob, info)
		return append(instrs, globInstrs...)
	}

	var base *Register
	var offset int

	// Find base register and offset for store
	switch v := st.Mem.(type) {
	case *ir.Literal:
		// For immediates, mov/load a new base register and use 0 offset
		var baseInstrs []Instr
		baseInstrs, base = movLoadImmediate(nil, v, info)
		instrs = append(instrs, baseInstrs...)

		offset = 0

	case *ir.Register:
		// For registers, look for a GepInstr to check for its base and offset
		var baseInstrs []Instr
		baseInstrs, base, offset = getLoadStoreBase(v, info)
		instrs = append(instrs, baseInstrs...)
	}

	// Create load instruction
	store := &StoreInstr{
		Src:    src,
		Base:   base,
		Offset: offset,
	}
	addUses(store)

	return append(instrs, store)
}

func storeGlobalToArm(src *Register, glob *ir.Register, info *functionInfo) []Instr {
	// Create temporary base register to hold page address
	base := nextTempReg(info.registers)

	// Get page aligned-address of global
	label := globals[glob.Name]
	adrp := &PageAddressInstr{
		Dst:   base,
		Label: label,
	}
	addUses(adrp)

	// Load global using page offset
	store := &StoreInstr{
		Src:        src,
		Base:       base,
		PageOffset: label,
	}
	addUses(store)

	return []Instr{adrp, store}
}

func getLoadStoreBase(reg *ir.Register, info *functionInfo) (instrs []Instr, base *Register, offset int) {
	switch v := reg.Def.(type) {
	case *ir.GepInstr:
		// Use GepInstr base as the base of this load/store
		switch gepBase := v.Base.(type) {
		case *ir.Literal:
			// If the GepInstr base is a literal, mov/load that value as the base
			instrs, base = movLoadImmediate(nil, gepBase, info)
		case *ir.Register:
			// Otherwise, use/load a register for the base
			instrs, base = useLoadRegister(gepBase.Name, info)
		}

		// Use GepInstr index to calculate offset
		offset = v.Index * dataSize

	case *ir.AllocInstr:
		// Only the read register will have an alloc as it's definition
		// In this case, use the stack variable directly for load/store
		sv := info.stackVars[v.Target.Name]
		base = sv.base
		offset = sv.offset
	}

	return
}

func callInstrToArm(call *ir.CallInstr, info *functionInfo) []Instr {
	numReg := util.Min(len(call.Arguments), maxRegisterParams)

	// Partition register arguments (x0-x7) vs. stack arguments
	regArgs := call.Arguments[:numReg]
	stackArgs := call.Arguments[numReg:]

	var instrs []Instr

	// Move/load argument registers (x0-x7)
	for i, arg := range regArgs {
		dst := findRegister(fmt.Sprintf("x%v", i), info)

		switch v := arg.(type) {
		case *ir.Register:
			// Handle scanf calls separately by adding an offset to an existing pointer
			if call.FnName == "scanf" {
				// Check if the load is from a global
				if v.Global {
					instrs = append(instrs, globalAddrToArm(dst, v, info)...)
					continue
				}

				baseInstrs, base, offset := getLoadStoreBase(v, info)
				instrs = append(instrs, baseInstrs...)

				add := &ArithInstr{
					Operator: AddOperator,
					Dst:      dst,
					Src1:     base,
					Src2: &Immediate{
						Value: strconv.Itoa(offset),
					},
				}
				addUses(add)
				instrs = append(instrs, add)

				// For all other calls, just load the appropriate argument register
			} else {
				var argInstrs []Instr
				argInstrs, _ = movLoadRegister(dst, v.Name, info)
				instrs = append(instrs, argInstrs...)
			}

		case *ir.Literal:
			var argInstrs []Instr
			argInstrs, _ = movLoadImmediate(dst, v, info)
			instrs = append(instrs, argInstrs...)
		}
	}

	// Store stack arguments (in reverse)
	stackInstrs, stackOffset := storeStackArgs(stackArgs, info)
	instrs = append(instrs, stackInstrs...)

	// Branch and link to subroutine
	bl := &BranchLinkInstr{call.FnName}
	addUses(bl)
	instrs = append(instrs, bl)

	// Restore stack pointer (if any stack arguments were pushed)
	if stackOffset != 0 {
		add := &ArithInstr{
			Operator: AddOperator,
			Dst:      findRegister("sp", info),
			Src1:     findRegister("sp", info),
			Src2: &Immediate{
				Value: strconv.Itoa(stackOffset),
			},
		}
		addUses(add)
		instrs = append(instrs, add)
	}

	// Store return value to call target (if needed)
	if call.Target != nil {
		src := findRegister("x0", info)

		targetInstrs := movStoreRegister(src, call.Target.Name, info)
		instrs = append(instrs, targetInstrs...)
	}

	return instrs
}

func storeStackArgs(args []ir.Value, info *functionInfo) (instrs []Instr, offset int) {
	numArgs := len(args)

	// Rectify number of pushed arguments to be even (sp must be a multiple of 16)
	if numArgs%2 == 1 {
		numArgs++
	}

	// Loop through all pairs of stack arguments (in reverse)
	for i := numArgs - 1; i >= 0; i -= 2 {
		var src1, src2 *Register
		var srcInstrs []Instr

		// Increment stack offset
		offset += 16

		// Get register for first argument from pair
		switch v := args[i-1].(type) {
		case *ir.Literal:
			srcInstrs, src1 = movLoadImmediate(nil, v, info)
		case *ir.Register:
			srcInstrs, src1 = useLoadRegister(v.Name, info)
		}
		instrs = append(instrs, srcInstrs...)

		// If a second argument exists in the pair, get a register for it
		if i <= numArgs-2 {
			switch v := args[i].(type) {
			case *ir.Literal:
				srcInstrs, src2 = movLoadImmediate(nil, v, info)
			case *ir.Register:
				srcInstrs, src2 = useLoadRegister(v.Name, info)
			}
			instrs = append(instrs, srcInstrs...)

			// Otherwise, just use the first argument as the second (as a buffer so the stack
			// pointer is still moved by a multiple of 16)
		} else {
			src2 = src1
		}

		// Decrement stack pointer and store argument pair
		store := &StorePairInstr{
			Src1:      src1,
			Src2:      src2,
			Base:      findRegister("sp", info),
			Offset:    -16,
			Increment: PreIncrement,
		}
		addUses(store)

		instrs = append(instrs, store)
	}

	return
}

func retInstrToArm(ret *ir.RetInstr, info *functionInfo) []Instr {
	// Ignore void returns
	if ret.Src == nil {
		return nil
	}

	// Create return register (x0)
	dst := findRegister("x0", info)

	var instrs []Instr

	// Move register or immediate into x0
	switch v := ret.Src.(type) {
	case *ir.Register:
		instrs, _ = movLoadRegister(dst, v.Name, info)
	case *ir.Literal:
		instrs, _ = movLoadImmediate(dst, v, info)
	}

	return instrs
}

func compInstrToArm(cmp *ir.CompInstr, info *functionInfo) []Instr {
	var forBranch bool
	uses := cmp.Target.Uses

	// Check if this compare target is used only as a branch condition
	if len(uses) == 1 {
		for k := range uses {
			if _, ok := k.(*ir.BranchInstr); ok {
				// If so, mark the comparison so that it doesn't require a conditional set
				forBranch = true
			}

			break
		}
	}

	// Create compare instruction
	return compToArm(cmp, forBranch, info)
}

func compToArm(cmp *ir.CompInstr, forBranch bool, info *functionInfo) []Instr {
	var instrs []Instr
	var op1 *Register
	var op2 Operand

	// Use/mov/load the first compare operand
	switch v := cmp.Op1.(type) {
	case *ir.Literal:
		instrs, op1 = movLoadImmediate(nil, v, info)
	case *ir.Register:
		instrs, op1 = useLoadRegister(v.Name, info)
	}

	// Use/mov/load the second compare operand
	var setInstrs []Instr
	switch v := cmp.Op2.(type) {
	case *ir.Literal:
		// Compare instructions can take a 12-bit immediate as the second operand
		setInstrs, op2 = arithImmediate(v, info)
	case *ir.Register:
		setInstrs, op2 = useLoadRegister(v.Name, info)
	}
	instrs = append(instrs, setInstrs...)

	// Create compare instruction
	comp := &CompInstr{
		Op1: op1,
		Op2: op2,
	}
	addUses(comp)
	instrs = append(instrs, comp)

	// Bail out if this compare will only set condition codes for a later branch
	if forBranch {
		return instrs
	}

	// Otherwise, add a conditional set instruction to save the comparison result
	cset := &ConditionalSetInstr{
		Dst:       findRegister(cmp.Target.Name, info),
		Condition: conditionToArm(cmp.Condition),
	}
	addUses(cset)

	return append(instrs, cset)
}

func branchInstrToArm(br *ir.BranchInstr, info *functionInfo) []Instr {
	var instrs []Instr
	var branchCond Condition

	// Check if this is an unconditional branch
	if br.Cond == nil {
		// If so, create a branch instruction without a condition
		branch := &BranchInstr{
			Block: info.blockMap[br.Next],
		}
		addUses(branch)

		return []Instr{branch}
	}

	// Otherwise, use a compare/test instruction along with a conditional branch
	switch cond := br.Cond.(type) {
	case *ir.Literal:
		// For literals, test the literal against itself
		var op1 *Register
		instrs, op1 = zeroMovLoadImmediate(nil, cond, info)

		test := &TestInstr{
			Op1: op1,
			Op2: boolImmediate(cond, info),
		}
		addUses(test)
		instrs = append(instrs, test)

		// Then use "not equal" for the branch condition (to check if the literal is not zero)
		branchCond = NotEqualCondition

	case *ir.Register:
		// For registers, check the condition register definition and insert a test/comp
		var regInstrs []Instr
		regInstrs, branchCond = branchRegisterToArm(cond, info)
		instrs = append(instrs, regInstrs...)
	}

	// Conditionally branch to the next block
	branchNext := &BranchInstr{
		Condition: branchCond,
		Block:     info.blockMap[br.Next],
	}
	addUses(branchNext)

	// Unconditionally branch to the else block
	branchElse := &BranchInstr{
		Block: info.blockMap[br.Els],
	}
	addUses(branchElse)

	return append(instrs, branchNext, branchElse)
}

func branchRegisterToArm(cond *ir.Register, info *functionInfo) (instrs []Instr, bc Condition) {
	// If the condition definition is a compare, use its condition
	if def, ok := cond.Def.(*ir.CompInstr); ok {
		bc = conditionToArm(def.Condition)
		return
	}

	var name string

	// Otherwise, get the register name from the ConvInstr or BinaryInstr target
	switch def := cond.Def.(type) {
	case *ir.ConvInstr:
		name = def.Src.(*ir.Register).Name
	case *ir.BinaryInstr:
		name = def.Target.Name
	}

	var op *Register
	instrs, op = useLoadRegister(name, info)

	// Then, test the condition register against itself
	test := &TestInstr{
		Op1: op,
		Op2: op,
	}
	addUses(test)
	instrs = append(instrs, test)

	// Use "not equal" for the branch condition
	bc = NotEqualCondition

	return
}

func binaryInstrToArm(bin *ir.BinaryInstr, info *functionInfo) []Instr {
	instrs := []Instr{}

	// Set destination register
	dst := findRegister(bin.Target.Name, info)

	op1 := bin.Op1
	op2 := bin.Op2

	// Flip operands to put an immediate second (for commutative operations)
	switch bin.Operator {
	case ir.AddOperator, ir.MulOperator, ir.AndOperator, ir.OrOperator, ir.XorOperator:
		if _, ok := op1.(*ir.Literal); ok {
			op1 = bin.Op2
			op2 = bin.Op1
		}
	}

	var setInstrs []Instr

	// Handle second operand
	var src2 Operand
	switch v := op2.(type) {
	case *ir.Register:
		setInstrs, src2 = useLoadRegister(v.Name, info)

		// Handle immediates on a per-operation basis
	case *ir.Literal:
		switch bin.Operator {
		case ir.AddOperator, ir.SubOperator:
			// Add/sub instructions can take an immediate
			setInstrs, src2 = arithImmediate(v, info)

		case ir.MulOperator, ir.DivOperator:
			// Mul/div instructions must use two registers
			setInstrs, src2 = movLoadImmediate(nil, v, info)

		case ir.AndOperator, ir.OrOperator, ir.XorOperator:
			// Logical instructions can take an immediate, but not "0" (so use xzr instead)
			src2 = boolImmediate(v, info)
			setInstrs = []Instr{}
		}
	}
	instrs = append(instrs, setInstrs...)

	// Handle first operand (always move or load immediates)
	var src1 *Register
	switch v := op1.(type) {
	case *ir.Register:
		setInstrs, src1 = useLoadRegister(v.Name, info)

	case *ir.Literal:
		useZero := true

		// Avoid using xzr for the first operand in certain instructions
		// when the second operand is an immediate
		switch bin.Operator {
		case ir.AddOperator, ir.SubOperator, ir.MulOperator, ir.DivOperator:
			if _, ok := src2.(*Immediate); ok {
				useZero = false
			}
		}

		if useZero {
			setInstrs, src1 = zeroMovLoadImmediate(nil, v, info)
		} else {
			setInstrs, src1 = movLoadImmediate(nil, v, info)
		}
	}
	instrs = append(instrs, setInstrs...)

	// Create arithmetic instruction
	arith := &ArithInstr{
		Operator: operatorToArm(bin.Operator),
		Dst:      dst,
		Src1:     src1,
		Src2:     src2,
	}
	addUses(arith)

	return append(instrs, arith)
}

func convInstrToArm(conv *ir.ConvInstr, info *functionInfo) []Instr {
	// Map target name to source register
	info.registers[conv.Target.Name] = findRegister(conv.Src.(*ir.Register).Name, info)
	return nil
}

func phiInstrToArm(phi *ir.PhiInstr, info *functionInfo) []Instr {
	// Create temporary register for phi value storage
	tmp := nextTempReg(info.registers)

	// Handle all phi values
	for _, val := range phi.Values {
		prev := info.blockMap[val.Block]

		var phiOutInstrs []Instr

		// Move/load phi value into temporary register
		switch v := val.Value.(type) {
		case *ir.Literal:
			phiOutInstrs, _ = movLoadImmediate(tmp, v, info)
		case *ir.Register:
			phiOutInstrs, _ = movLoadRegister(tmp, v.Name, info)
		}

		// Append move/load to previous block
		prev.PhiOuts = append(prev.PhiOuts, phiOutInstrs...)
	}

	return movStoreRegister(tmp, phi.Target.Name, info)
}

func useLoadRegister(name string, info *functionInfo) (instrs []Instr, reg *Register) {
	// Check for an existing stack variable
	if v, present := info.stackVars[name]; present {
		reg = nextTempReg(info.registers)

		// Load the stack variable into a new temporary register
		load := &LoadInstr{
			Dst:    reg,
			Base:   v.base,
			Offset: v.offset,
		}
		addUses(load)

		instrs = []Instr{load}

	} else {
		// Otherwise, use a register
		reg = findRegister(name, info)
	}

	return
}

func movLoadRegister(dst *Register, name string,
	info *functionInfo) (instrs []Instr, reg *Register) {

	// Use a new temporary register if a destination is not given
	if dst == nil {
		dst = nextTempReg(info.registers)
	}
	reg = dst

	// Check for an existing stack variable
	if v, ok := info.stackVars[name]; ok {
		// If one is found, load it into register
		load := &LoadInstr{
			Dst:    dst,
			Base:   v.base,
			Offset: v.offset,
		}
		addUses(load)
		instrs = []Instr{load}

		return
	}

	// Otherwise, move from src to dst
	mov := &MovInstr{
		Dst: dst,
		Src: findRegister(name, info),
	}
	addUses(mov)
	instrs = []Instr{mov}

	return
}

func movStoreRegister(src *Register, name string, info *functionInfo) []Instr {
	// Check for an existing stack variable
	if v, ok := info.stackVars[name]; ok {
		// If one is found, store it into register
		store := &StoreInstr{
			Src:    src,
			Base:   v.base,
			Offset: v.offset,
		}
		addUses(store)

		return []Instr{store}
	}

	dst := findRegister(name, info)

	// Otherwise, move from src to dst
	mov := &MovInstr{
		Dst: dst,
		Src: src,
	}
	addUses(mov)

	return []Instr{mov}
}

func boolImmediate(lit *ir.Literal, info *functionInfo) Operand {
	// Use zero register for false, or an immediate otherwise
	if lit.ToBool() {
		return &Immediate{
			Value:  "1",
			Global: lit.Global,
		}
	} else {
		return findRegister("xzr", info)
	}
}

func arithImmediate(lit *ir.Literal, info *functionInfo) (instrs []Instr, op Operand) {
	val, err := lit.ToInt()

	// Use zero register instead of zero immediate
	if err == nil && val == 0 {
		op = findRegister("xzr", info)
		return
	}

	// Check if an arithmetic instruction can fit the immediate (<= 12 bits, generally)
	if err == nil && val >= arithImmediateMin && val <= arithImmediateMax {
		op = &Immediate{
			Value:  lit.Value,
			Global: lit.Global,
		}
		return
	}

	instrs, op = movLoadImmediate(nil, lit, info)
	return
}

func zeroMovLoadImmediate(dst *Register, lit *ir.Literal, info *functionInfo) (instrs []Instr, reg *Register) {
	// Convert value to an integer
	val, err := lit.ToInt()

	// Use zero register instead of zero immediate
	if err == nil && val == 0 {
		reg = findRegister("xzr", info)
		return
	}

	instrs, reg = movLoadImmediate(nil, lit, info)
	return
}

func movLoadImmediate(dst *Register, lit *ir.Literal, info *functionInfo) (instrs []Instr, reg *Register) {
	// Use a new temporary register if a destination is not given
	if dst == nil {
		dst = nextTempReg(info.registers)
	}
	reg = dst

	strVal := lit.Value

	// Convert value to an integer
	val, err := lit.ToInt()

	// Always use "0" instead of "null"
	if err == nil && val == 0 {
		strVal = "0"
	}

	// Check if a mov instruction can fit the immediate (<= 16 bits, generally)
	if err == nil && val >= movImmediateMin && val <= movImmediateMax {
		mov := &MovInstr{
			Dst: dst,
			Src: &Immediate{
				Value:  strVal,
				Global: lit.Global,
			},
		}
		addUses(mov)
		instrs = []Instr{mov}

		return
	}

	// Otherwise, use a load immediate pseudo-instruction
	load := &LoadImmediateInstr{
		Dst: dst,
		Imm: &Immediate{
			Value:  strVal,
			Global: lit.Global,
		},
	}
	addUses(load)
	instrs = []Instr{load}

	return
}

func createPrologue(callees []*Register, info *functionInfo, offset int) []Instr {
	// Store frame pointer and link register for safe keeping
	store := &StorePairInstr{
		Src1:      findRegister("fp", info),
		Src2:      findRegister("lr", info),
		Base:      findRegister("sp", info),
		Offset:    -16,
		Increment: PreIncrement,
	}
	addUses(store)

	// Move frame pointer to new stack pointer
	mov := &MovInstr{
		Dst: findRegister("fp", info),
		Src: findRegister("sp", info),
	}
	addUses(mov)
	instrs := []Instr{store, mov}

	// Optionally, decrement stack pointer to make space for stack variables
	if offset != 0 {
		sub := &ArithInstr{
			Operator: SubOperator,
			Dst:      findRegister("sp", info),
			Src1:     findRegister("sp", info),
			Src2: &Immediate{
				Value: strconv.Itoa(offset),
			},
		}
		addUses(sub)

		instrs = append(instrs, sub)
	}

	// Store callee-saved registers
	return append(instrs, storeCallees(callees, info)...)
}

func storeCallees(callees []*Register, info *functionInfo) []Instr {
	var instrs []Instr
	offset := -16

	for i := 0; i < len(callees); i += 2 {
		// Get first stored source
		src1 := callees[i]
		var src2 *Register

		// Get second stored source (if it exists) or use the first source otherwise
		if i+1 < len(callees) {
			src2 = callees[i+1]
		} else {
			src2 = src1
		}

		// Store pair of callee-saved registers to the stack
		store := &StorePairInstr{
			Src1:   src1,
			Src2:   src2,
			Base:   findRegister("fp", info),
			Offset: offset,
		}
		addUses(store)
		instrs = append(instrs, store)

		offset -= 16
	}

	return instrs
}

func loadCallees(callees []*Register, info *functionInfo) []Instr {
	var instrs []Instr
	offset := -16

	for i := 0; i < len(callees); i += 2 {
		// Get first load destination
		dst1 := callees[i]

		var load Instr
		if i+1 < len(callees) {
			// If a second load destination exists, load pair of callee-saved registers
			load = &LoadPairInstr{
				Dst1:   dst1,
				Dst2:   callees[i+1],
				Base:   findRegister("fp", info),
				Offset: offset,
			}

		} else {
			// Otherwise, load the single callee-saved register
			load = &LoadInstr{
				Dst:    dst1,
				Base:   findRegister("fp", info),
				Offset: offset + dataSize,
			}
		}
		addUses(load)
		instrs = append(instrs, load)

		offset -= 16
	}

	return instrs
}

func createEpilogue(callees []*Register, info *functionInfo) []Instr {
	// Load callee-saved registers
	instrs := loadCallees(callees, info)

	// Restore stack pointer (so it mirrors the frame pointer)
	mov := &MovInstr{
		Dst: findRegister("sp", info),
		Src: findRegister("fp", info),
	}
	addUses(mov)

	// Restore frame pointer and link register from storage
	load := &LoadPairInstr{
		Dst1:      findRegister("fp", info),
		Dst2:      findRegister("lr", info),
		Base:      findRegister("sp", info),
		Offset:    16,
		Increment: PostIncrement,
	}
	addUses(load)

	// Use a return instruction to leave subroutine
	ret := &RetInstr{}

	return append(instrs, mov, load, ret)
}

func findRegister(name string, info *functionInfo) *Register {
	// Check for an existing register and return it
	if v, ok := info.registers[name]; ok {
		return v
	}

	// Create a new register if needed, and add it to the existing register table
	reg := &Register{
		Name:    name,
		Virtual: true,
	}
	info.registers[name] = reg

	return reg
}

func addUses(instr Instr) {
	// Append instruction to uses for each register source
	for _, v := range instr.getSrcs() {
		if reg, ok := v.(*Register); ok {
			reg.Uses = append(reg.Uses, instr)
		}
	}

	// Append instruction to uses for each destination
	for _, reg := range instr.getDsts() {
		reg.Uses = append(reg.Uses, instr)
	}
}

func nextTempReg(table map[string]*Register) *Register {
	tmpNum++

	reg := &Register{
		Name:    fmt.Sprintf("_tmp%v", tmpNum),
		Virtual: true,
	}

	table[reg.Name] = reg
	return reg
}

func conditionToArm(cond ir.Condition) Condition {
	switch cond {
	case ir.EqualCondition:
		return EqualCondition
	case ir.NotEqualCondition:
		return NotEqualCondition
	case ir.GreaterThanCondition:
		return GreaterThanCondition
	case ir.GreaterEqualCondition:
		return GreaterEqualCondition
	case ir.LessThanCondition:
		return LessThanCondition
	case ir.LessEqualCondition:
		return LessEqualCondition
	}

	panic("Unsupported condition")
}

func operatorToArm(op ir.Operator) Operator {
	switch op {
	case ir.AddOperator:
		return AddOperator
	case ir.SubOperator:
		return SubOperator
	case ir.MulOperator:
		return MulOperator
	case ir.DivOperator:
		return DivOperator
	case ir.AndOperator:
		return AndOperator
	case ir.OrOperator:
		return OrOperator
	case ir.XorOperator:
		return XorOperator
	}

	panic("Unsupported operator")
}

const genRegsCalleeEnd = 10
const genRegsCallerStart = 10

func populatePhysicalRegs(info *functionInfo) (genRegs []*Register, specRegs map[*Register]bool) {
	// Create general register list
	// The first registers in the list are more favored during register allocation
	genRegs = []*Register{
		// Callee save (x19-x28)
		{Name: "x28"},
		{Name: "x27"},
		{Name: "x26"},
		{Name: "x25"},
		{Name: "x24"},
		{Name: "x23"},
		{Name: "x22"},
		{Name: "x21"},
		{Name: "x20"},
		{Name: "x19"},

		// Caller save (x9-x15)
		{Name: "x15"},
		{Name: "x14"},
		{Name: "x13"},
		{Name: "x12"},
		{Name: "x11"},
		{Name: "x10"},
		{Name: "x9"},

		// Syscall number
		{Name: "x8"},

		// Parameters (x0-x7)
		{Name: "x7"},
		{Name: "x6"},
		{Name: "x5"},
		{Name: "x4"},
		{Name: "x3"},
		{Name: "x2"},
		{Name: "x1"},
		{Name: "x0"},
	}

	// Add general registers to register map
	for _, reg := range genRegs {
		info.registers[reg.Name] = reg
	}

	// Create special register map
	specRegs = map[*Register]bool{
		{Name: "fp"}:  true, // Frame pointer (x29)
		{Name: "lr"}:  true, // Link register (x30)
		{Name: "sp"}:  true, // Stack pointer (x31, depending on context)
		{Name: "xzr"}: true, // Zero register (x31, depending)
	}

	// Add special registers to register map
	for reg := range specRegs {
		info.registers[reg.Name] = reg
	}

	return
}

func findCallees(usedRegs map[*Register]*Register, genRegs []*Register) []*Register {
	// Create set of all callee registers
	callees := make(map[*Register]bool, genRegsCalleeEnd)
	for _, reg := range genRegs[:genRegsCalleeEnd] {
		callees[reg] = true
	}

	// Set all used callee registers to "false" in the map (but don't remove them)
	for _, reg := range usedRegs {
		if callees[reg] {
			callees[reg] = false
		}
	}

	// Collect all false (used) callee registers
	var ret []*Register
	for reg := range callees {
		if value, ok := callees[reg]; ok && !value {
			ret = append(ret, reg)
		}
	}

	return ret
}
