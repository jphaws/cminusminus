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
	populateLocals(fn.Cfg.Allocs, info)

	// TODO: Remove me
	fmt.Printf("registers: %v\n", info.registers)
	fmt.Printf("stackVars: %v\n", info.stackVars)

	// Process all CFG blocks
	blocks := processBlock(fn.Cfg, info)

	// Create prologue block
	proBlock := &Block{
		Label:  name,
		Instrs: createPrologue(info),
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

func populateLocals(allocs []*ir.AllocInstr, info *functionInfo) {
	// Create name->offset mappings for stack locals
	for i, v := range allocs {
		info.stackVars[v.Target.Name] = StackOffset{
			Base:   findRegister("sp", false, info),
			Offset: localOffset + i*dataSize,
		}
	}

	// Set spill offset to top of local section
	info.spillOffset = len(allocs) * dataSize
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
	case *ir.AllocInstr:
		return nil
	case *ir.LoadInstr:
		return nil
	case *ir.StoreInstr:
		return nil
	case *ir.CallInstr:
		return nil
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
		instrs, _ = movLoadImmediate(dst, v.Value, info)
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

	// Handle first operand (always move or load immediates)
	var src1 *Register
	switch v := op1.(type) {
	case *ir.Register:
		src1 = findRegister(v.Name, true, info)

	case *ir.Literal:
		var movInstrs []Instr
		movInstrs, src1 = zeroMovLoadImmediate(nil, v.Value, info)
		instrs = append(instrs, movInstrs...)
	}

	// Handle second operand
	var src2 Operand
	switch v := op2.(type) {
	case *ir.Register:
		src2 = findRegister(v.Name, true, info)

		// Handle immediates on a per-operation basis
	case *ir.Literal:
		var setInstrs []Instr

		switch bin.Operator {
		case ir.AddOperator, ir.SubOperator:
			// Add/sub instructions can take an immediate
			setInstrs, src2 = arithImmediate(v.Value, info)

		case ir.MulOperator, ir.DivOperator:
			// Mul/div instructions must use two registers
			setInstrs, src2 = movLoadImmediate(nil, v.Value, info)

		case ir.AndOperator, ir.OrOperator, ir.XorOperator:
			// Logical instructions can take an immediate, but not "0" (so use xzr instead)
			src2 = boolImmediate(v.Value, info)
		}

		instrs = append(instrs, setInstrs...)
	}

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

func boolImmediate(val string, info *functionInfo) Operand {
	// Use zero register for "0", or an immediate otherwise
	if val == "0" {
		return findRegister("xzr", false, info)
	} else {
		return &Immediate{val}
	}
}

func arithImmediate(val string, info *functionInfo) (instrs []Instr, op Operand) {
	var v int
	var err error

	// Convert value to an integer
	if val == "null" {
		v = 0
	} else {
		v, err = strconv.Atoi(val)
	}

	// Use zero register instead of zero immediate
	if err == nil && v == 0 {
		op = findRegister("xzr", false, info)
		return
	}

	// Check if an arithmetic instruction can fit the immediate (<= 12 bits, generally)
	if err == nil && v >= arithImmediateMin && v <= arithImmediateMax {
		op = &Immediate{val}
		return
	}

	instrs, op = movLoadImmediate(nil, val, info)
	return
}

func zeroMovLoadImmediate(dst *Register, val string, info *functionInfo) (instrs []Instr, reg *Register) {
	var v int
	var err error

	// Convert value to an integer
	if val == "null" {
		v = 0
	} else {
		v, err = strconv.Atoi(val)
	}

	// Use zero register instead of zero immediate
	if err == nil && v == 0 {
		reg = findRegister("xzr", false, info)
		return
	}

	instrs, reg = movLoadImmediate(nil, val, info)
	return
}

func movLoadImmediate(dst *Register, val string, info *functionInfo) (instrs []Instr, reg *Register) {
	// Use a new temporary register if a destination is not given
	if dst == nil {
		dst = nextTempReg(info.registers)
	}
	reg = dst

	v, err := strconv.Atoi(val)

	// Check if a mov instruction can fit the immediate (<= 16 bits, generally)
	if err == nil && v >= movImmediateMin && v <= movImmediateMax {
		mov := &MovInstr{
			Dst: dst,
			Src: &Immediate{val},
		}
		addUses(mov)
		instrs = []Instr{mov}

		return
	}

	// Otherwise, use a load immediate pseudo-instruction
	load := &LoadImmediateInstr{
		Dst: dst,
		Imm: &Immediate{val},
	}
	addUses(load)
	instrs = []Instr{load}

	return
}

func createPrologue(info *functionInfo) []Instr {
	store := &StorePairInstr{
		Src1:      findRegister("fp", false, info),
		Src2:      findRegister("lr", false, info),
		Base:      findRegister("sp", false, info),
		Offset:    -16,
		Increment: PreIncrement,
	}
	addUses(store)

	mov := &MovInstr{
		Dst: findRegister("fp", false, info),
		Src: findRegister("sp", false, info),
	}
	addUses(mov)

	return []Instr{store, mov}
}

func createEpilogue(info *functionInfo) []Instr {
	mov := &MovInstr{
		Dst: findRegister("sp", false, info),
		Src: findRegister("fp", false, info),
	}
	addUses(mov)

	load := &LoadPairInstr{
		Dst1:      findRegister("fp", false, info),
		Dst2:      findRegister("lr", false, info),
		Base:      findRegister("sp", false, info),
		Offset:    16,
		Increment: PostIncrement,
	}
	addUses(load)

	ret := &RetInstr{}

	return []Instr{load, mov, ret}
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
