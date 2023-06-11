package asm

import (
	"fmt"

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
	case *ir.LoadInstr:
		return loadInstrToArm(v, info)
	case *ir.StoreInstr:
		return storeInstrToArm(v, info)
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
	adrp := &PageAddressInstr{
		Dst:   base,
		Label: "$" + glob.Name,
	}
	addUses(adrp)

	// Load global using page offset
	load := &LoadInstr{
		Dst:        dst,
		Base:       base,
		PageOffset: glob.Name,
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
	adrp := &PageAddressInstr{
		Dst:   base,
		Label: "$" + glob.Name,
	}
	addUses(adrp)

	// Load global using page offset
	store := &StoreInstr{
		Src:        src,
		Base:       base,
		PageOffset: glob.Name,
	}
	addUses(store)

	return []Instr{adrp, store}
}

func getLoadStoreBase(reg *ir.Register, info *functionInfo) (instrs []Instr, base *Register, offset int) {
	gep := reg.Def.(*ir.GepInstr)

	// Use GepInstr base as the base of this load/store
	switch v := gep.Base.(type) {
	case *ir.Literal:
		// If the GepInstr base is a literal, mov/load that value as the base
		instrs, base = movLoadImmediate(nil, v, info)

	case *ir.Register:
		// Otherwise, use/load a register for the base
		instrs, base = useLoadRegister(v.Name, info)
	}

	// Use GepInstr index to calculate offset
	offset = gep.Index * dataSize

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

	// Handle first operand (always move or load immediates)
	var src1 *Register
	switch v := op1.(type) {
	case *ir.Register:
		src1 = findRegister(v.Name, true, info)

	case *ir.Literal:
		var movInstrs []Instr
		movInstrs, src1 = zeroMovLoadImmediate(nil, v, info)
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
			setInstrs, src2 = arithImmediate(v, info)

		case ir.MulOperator, ir.DivOperator:
			// Mul/div instructions must use two registers
			setInstrs, src2 = movLoadImmediate(nil, v, info)

		case ir.AndOperator, ir.OrOperator, ir.XorOperator:
			// Logical instructions can take an immediate, but not "0" (so use xzr instead)
			src2 = boolImmediate(v, info)
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

func boolImmediate(lit *ir.Literal, info *functionInfo) Operand {
	// Use zero register for false, or an immediate otherwise
	if lit.ToBool() {
		return &Immediate{"1"}
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
		op = &Immediate{lit.Value}
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
			Src: &Immediate{strVal},
		}
		addUses(mov)
		instrs = []Instr{mov}

		return
	}

	// Otherwise, use a load immediate pseudo-instruction
	load := &LoadImmediateInstr{
		Dst: dst,
		Imm: &Immediate{strVal},
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
