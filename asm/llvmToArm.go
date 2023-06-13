package asm

import (
	"fmt"
	"strconv"

	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/util"
)

const (
	dataSize = 8

	paramOffset  = 16 // Offset from fp
	calleeOffset = 0  // Offset from fp
	localOffset  = 0  // Offset from sp

	maxRegisterParams = 8

	arithImmediateMin = -4096
	arithImmediateMax = 4096
	movImmediateMin   = -65536
	movImmediateMax   = 65536
)

var tmpNum = 0

type functionInfo struct {
	epiBlock    *Block
	blocks      map[string]*Block
	registers   map[string]*Register
	stackVars   map[string]StackOffset
	spillOffset int
}

type StackOffset struct {
	Base   *Register
	Offset int
}

func processFunction(fn *ir.Function, name string, ch chan *Function) {
	info := &functionInfo{
		blocks:    map[string]*Block{},
		registers: map[string]*Register{},
		stackVars: map[string]StackOffset{},
	}

	// Pre-populate register and stack maps with parameters
	populateParams(fn.Parameters, info)

	// Pre-populate stack map with local stack variables
	stackPointerOffset := populateLocals(fn.Cfg.Allocs, info)

	// Set spill offset to top of local section
	info.spillOffset = stackPointerOffset

	// TODO: Remove me
	fmt.Printf("registers: %v\n", info.registers)
	fmt.Printf("stackVars: %v\n", info.stackVars)

	// Process all CFG blocks
	blocks := processBlock(fn.Cfg, info)

	// Rectify stack pointer to a 16-byte boundary
	if stackPointerOffset%16 != 0 {
		stackPointerOffset += 8
	}

	// Create prologue block
	proBlock := &Block{
		Label:  name,
		Instrs: createPrologue(info, stackPointerOffset),
	}

	// Add epilogue instructions to return block
	epiBlock := info.epiBlock
	epiBlock.Instrs = append(epiBlock.Instrs, createEpilogue(info)...)

	// Create function wrapper
	ret := &Function{
		Blocks: []*Block{proBlock},
	}
	ret.Blocks = append(ret.Blocks, blocks...)

	ch <- ret
}

func populateParams(params []*ir.Register, info *functionInfo) {
	numReg := util.Min(len(params), maxRegisterParams)

	// Partition register parameters (x0-x7) vs. stack parameters
	regParams := params[:numReg]
	stackParams := params[numReg:]

	// Create name->register mappings for register parameters
	for i, v := range regParams {
		info.registers[v.Name] = &Register{
			Name:    fmt.Sprintf("x%v", i),
			Virtual: false,
		}
	}

	// Create name->offset mappings for stack parameters
	for i, v := range stackParams {
		info.stackVars[v.Name] = StackOffset{
			Base:   findRegister("fp", false, info),
			Offset: paramOffset + i*dataSize,
		}
	}
}

func populateLocals(allocs []*ir.AllocInstr, info *functionInfo) int {
	// Create name->offset mappings for stack locals
	for i, v := range allocs {
		info.stackVars[v.Target.Name] = StackOffset{
			Base:   findRegister("sp", false, info),
			Offset: localOffset + i*dataSize,
		}
	}

	return len(allocs) * dataSize
}

func processBlock(b *ir.Block, info *functionInfo) []*Block {
	// Create new ASM block and add to the block list
	curr := &Block{
		Label: "." + b.Label(),
	}
	info.blocks[curr.Label] = curr

	// Translate all LLVM instructions to ARM
	for _, v := range b.Instrs {
		curr.Instrs = append(curr.Instrs, instrToArm(v, curr, info)...)
	}

	ret := []*Block{curr}

	// Process next block (if it exists)
	if b.Next != nil {
		if _, ok := info.blocks["."+b.Next.Label()]; !ok {
			ret = append(ret, processBlock(b.Next, info)...)
		}
	}

	// Process else block (if it exists)
	if b.Els != nil {
		if _, ok := info.blocks["."+b.Els.Label()]; !ok {
			ret = append(ret, processBlock(b.Els, info)...)
		}
	}

	return ret
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
	case *ir.BranchInstr:
		return nil
	case *ir.BinaryInstr:
		return binaryInstrToArm(v, info)
	case *ir.PhiInstr:
		return nil
	}

	return nil
}

func loadInstrToArm(ld *ir.LoadInstr, info *functionInfo) []Instr {
	dst := findRegister(ld.Reg.Name, true, info)

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
		base = sv.Base
		offset = sv.Offset
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
		dst := findRegister(fmt.Sprintf("x%v", i), false, info)

		switch v := arg.(type) {
		case *ir.Register:
			// Handle scanf calls separately by adding an offset to an existing pointer
			if call.FnName == "scanf" {
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
				argInstrs, _ = movLoadRegister(dst, v.Name, false, info)
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
			Dst:      findRegister("sp", false, info),
			Src1:     findRegister("sp", false, info),
			Src2: &Immediate{
				Value: strconv.Itoa(stackOffset),
			},
		}
		addUses(add)
		instrs = append(instrs, add)
	}

	// Store return value to call target (if needed)
	if call.Target != nil {
		src := findRegister("x0", false, info)

		targetInstrs, _ := movStoreRegister(src, call.Target.Name, true, info)
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
			Base:      findRegister("sp", false, info),
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
	dst := findRegister("x0", false, info)

	var instrs []Instr

	// Move register or immediate into x0
	switch v := ret.Src.(type) {
	case *ir.Register:
		instrs, _ = movLoadRegister(dst, v.Name, true, info)

	case *ir.Literal:
		instrs, _ = movLoadImmediate(dst, v, info)
	}

	return instrs
}

func binaryInstrToArm(bin *ir.BinaryInstr, info *functionInfo) []Instr {
	instrs := []Instr{}

	// Set destination register
	dst := findRegister(bin.Target.Name, true, info)

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

	// Handle first operand (always move or load immediates)
	var src1 *Register
	switch v := op1.(type) {
	case *ir.Register:
		setInstrs, src1 = useLoadRegister(v.Name, info)

	case *ir.Literal:
		setInstrs, src1 = zeroMovLoadImmediate(nil, v, info)
	}
	instrs = append(instrs, setInstrs...)

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
	instrs = append(instrs, arith)

	return instrs
}

func useLoadRegister(name string, info *functionInfo) (instrs []Instr, reg *Register) {
	// Check for an existing stack variable
	if v, present := info.stackVars[name]; present {
		reg = nextTempReg(info.registers)

		// Load the stack variable into a new temporary register
		load := &LoadInstr{
			Dst:    reg,
			Base:   v.Base,
			Offset: v.Offset,
		}
		addUses(load)

		instrs = []Instr{load}

	} else {
		// Otherwise, use a register
		reg = findRegister(name, true, info)
	}

	return
}

func movLoadRegister(dst *Register, name string,
	virtual bool, info *functionInfo) (instrs []Instr, reg *Register) {

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
			Base:   v.Base,
			Offset: v.Offset,
		}
		addUses(load)
		instrs = []Instr{load}

		return
	}

	// Otherwise, move from src to dst
	mov := &MovInstr{
		Dst: dst,
		Src: findRegister(name, virtual, info),
	}
	addUses(mov)
	instrs = []Instr{mov}

	return
}

func movStoreRegister(src *Register, name string,
	virtual bool, info *functionInfo) (instrs []Instr, reg *Register) {

	// Check for an existing stack variable
	if v, ok := info.stackVars[name]; ok {
		// If one is found, store it into register
		store := &StoreInstr{
			Src:    src,
			Base:   v.Base,
			Offset: v.Offset,
		}
		addUses(store)
		instrs = []Instr{store}

		return
	}

	dst := findRegister(name, virtual, info)

	// Otherwise, move from src to dst
	mov := &MovInstr{
		Dst: dst,
		Src: src,
	}
	addUses(mov)
	instrs = []Instr{mov}

	return
}

func boolImmediate(lit *ir.Literal, info *functionInfo) Operand {
	// Use zero register for false, or an immediate otherwise
	if lit.ToBool() {
		return &Immediate{
			Value:  "1",
			Global: lit.Global,
		}
	} else {
		return findRegister("xzr", false, info)
	}
}

func arithImmediate(lit *ir.Literal, info *functionInfo) (instrs []Instr, op Operand) {
	val, err := lit.ToInt()

	// Use zero register instead of zero immediate
	if err == nil && val == 0 {
		op = findRegister("xzr", false, info)
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
		reg = findRegister("xzr", false, info)
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

func createPrologue(info *functionInfo, offset int) []Instr {
	// Store frame pointer and link register for safe keeping
	store := &StorePairInstr{
		Src1:      findRegister("fp", false, info),
		Src2:      findRegister("lr", false, info),
		Base:      findRegister("sp", false, info),
		Offset:    -16,
		Increment: PreIncrement,
	}
	addUses(store)

	// Move frame pointer to new stack pointer
	mov := &MovInstr{
		Dst: findRegister("fp", false, info),
		Src: findRegister("sp", false, info),
	}
	addUses(mov)

	instrs := []Instr{store, mov}

	// Optionally, decrement stack pointer to make space for stack variables
	if offset != 0 {
		sub := &ArithInstr{
			Operator: SubOperator,
			Dst:      findRegister("sp", false, info),
			Src1:     findRegister("sp", false, info),
			Src2: &Immediate{
				Value: strconv.Itoa(offset),
			},
		}
		addUses(sub)

		instrs = append(instrs, sub)
	}

	return instrs
}

func createEpilogue(info *functionInfo) []Instr {
	// Restore stack pointer (so it mirrors the frame pointer)
	mov := &MovInstr{
		Dst: findRegister("sp", false, info),
		Src: findRegister("fp", false, info),
	}
	addUses(mov)

	// Restore frame pointer and link register from storage
	load := &LoadPairInstr{
		Dst1:      findRegister("fp", false, info),
		Dst2:      findRegister("lr", false, info),
		Base:      findRegister("sp", false, info),
		Offset:    16,
		Increment: PostIncrement,
	}
	addUses(load)

	// Use a return instruction to leave subroutine
	ret := &RetInstr{}

	return []Instr{mov, load, ret}
}

func findRegister(name string, virtual bool, info *functionInfo) *Register {
	// Check for an existing register and return it
	if v, ok := info.registers[name]; ok {
		return v
	}

	// Create a new register if needed, and add it to the existing register table
	reg := &Register{
		Name:    name,
		Virtual: virtual,
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
