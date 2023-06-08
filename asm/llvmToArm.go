package asm

import (
	"fmt"
	"strconv"

	"github.com/keen-cp/compiler-project-c/ir"
)

const (
	paramOffsetInit = 16
	localOffsetInit = 0

	arithImmediateMin = -4096
	arithImmediateMax = 4096
	movImmediateMin   = -65536
	movImmediateMax   = 65536
)

var regNum = 0

type functionInfo struct {
	epiBlock     *Block
	paramOffset  int
	localOffset  int
	calleeOffset int
	spillOffset  int
	blocks       map[string]*Block
	registers    map[string]*Register
}

func processFunction(fn *ir.Function, name string, ch chan *Function) {
	info := &functionInfo{
		paramOffset: paramOffsetInit,
		localOffset: localOffsetInit,
		registers:   map[string]*Register{},
	}

	visited := map[string]*Block{}
	blocks := processBlock(fn.Cfg, info, visited)

	proBlock := &Block{
		Label:  name,
		Instrs: createPrologue(info),
	}

	epiBlock := info.epiBlock
	epiBlock.Instrs = append(epiBlock.Instrs, createEpilogue(info)...)

	ret := &Function{
		Blocks: []*Block{proBlock},
	}
	ret.Blocks = append(ret.Blocks, blocks...)

	ch <- ret
}

func processBlock(b *ir.Block, info *functionInfo, visited map[string]*Block) []*Block {
	curr := &Block{
		Label: "." + b.Label(),
	}
	visited[curr.Label] = curr

	for _, v := range b.Instrs {
		curr.Instrs = append(curr.Instrs, instrToArm(v, curr, info)...)
	}

	ret := []*Block{curr}

	if b.Next != nil {
		if _, ok := visited["."+b.Next.Label()]; !ok {
			ret = append(ret, processBlock(b.Next, info, visited)...)
		}
	}

	if b.Els != nil {
		if _, ok := visited["."+b.Els.Label()]; !ok {
			ret = append(ret, processBlock(b.Els, info, visited)...)
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
	case *ir.GepInstr:
		return nil
	case *ir.CallInstr:
		return nil
	case *ir.RetInstr:
		info.epiBlock = curr
		return retInstrToArm(v, info)
	case *ir.CompInstr:
		return nil
	case *ir.BranchInstr:
		return nil
	case *ir.BinaryInstr:
		return nil
	case *ir.ConvInstr:
		return nil
	case *ir.PhiInstr:
		return nil
	}

	return nil
}

func retInstrToArm(ret *ir.RetInstr, info *functionInfo) []Instr {
	if ret.Src == nil {
		return nil
	}

	dst := createRegister("x0", false, info.registers)

	var instr Instr

	switch v := ret.Src.(type) {
	case *ir.Register:
		instr = &MovInstr{
			Dst: dst,
			Src: createRegister(v.Name, true, info.registers),
		}
		addUses(instr)

	case *ir.Literal:
		instr = movLoadImmediate(dst, v.Value, info)
	}

	return []Instr{instr}
}

// func valueToAsm(val ir.Value, dst *Register,
// 	immType ImmediateType, info *functionInfo) (instr Instr, op Operand) {

// 	switch v := val.(type) {
// 	case *ir.Register:
// 		op = createRegister(v.Name, true, info.registers)

// 	case *ir.Literal:
// 		switch immType {
// 		case BoolImmediateType:
// 			op = boolImmediate(v.Value, info)

// 		case ArithImmediateType:
// 			instr, op = arithImmediate(v.Value, info)

// 		case MovLoadImmediateType:
// 			instr = movLoadImmediate(v.Value, dst, info)
// 		}
// 	}

// 	return
// }

// type ImmediateType int

// const (
// 	BoolImmediateType ImmediateType = iota
// 	ArithImmediateType
// 	MovLoadImmediateType
// )

func boolImmediate(val string, info *functionInfo) Operand {
	if val == "0" {
		return createRegister("xzr", false, info.registers)
	} else {
		return &Immediate{val}
	}
}

func arithImmediate(val string, info *functionInfo) (instr Instr, imm *Immediate) {
	var v int
	var err error

	// Convert value to an integer
	if val == "null" {
		v = 0
	} else {
		v, err = strconv.Atoi(val)
	}

	// Check if an arithmetic instruction can fit the immediate (<= 12 bits, generally)
	if err == nil && v >= arithImmediateMin && v <= arithImmediateMax {
		imm = &Immediate{val}
		return
	}

	instr = movLoadImmediate(nil, val, info)
	return
}

func movLoadImmediate(dst *Register, val string, info *functionInfo) Instr {
	// Use a new temporary register if a destination is not given
	if dst == nil {
		dst = nextTempReg(info.registers)
	}

	v, err := strconv.Atoi(val)
	fmt.Println(err)

	// Check if a mov instruction can fit the immediate (<= 16 bits, generally)
	if err == nil && v >= movImmediateMin && v <= movImmediateMax {
		mov := &MovInstr{
			Dst: dst,
			Src: &Immediate{val},
		}
		addUses(mov)

		return mov
	}

	// Otherwise, use a load immediate pseudo-instruction
	load := &LoadImmediateInstr{
		Dst: dst,
		Imm: &Immediate{val},
	}
	addUses(load)

	return load
}

func createPrologue(info *functionInfo) []Instr {
	store := &StorePairInstr{
		Src1:      createRegister("fp", false, info.registers),
		Src2:      createRegister("lr", false, info.registers),
		Base:      createRegister("sp", false, info.registers),
		Offset:    Offset{Off: -16},
		Increment: PreIncrement,
	}
	addUses(store)

	mov := &MovInstr{
		Dst: createRegister("fp", false, info.registers),
		Src: createRegister("sp", false, info.registers),
	}
	addUses(mov)

	return []Instr{store, mov}
}

func createEpilogue(info *functionInfo) []Instr {
	mov := &MovInstr{
		Dst: createRegister("sp", false, info.registers),
		Src: createRegister("fp", false, info.registers),
	}
	addUses(mov)

	load := &LoadPairInstr{
		Dst1:      createRegister("fp", false, info.registers),
		Dst2:      createRegister("lr", false, info.registers),
		Base:      createRegister("sp", false, info.registers),
		Offset:    Offset{Off: 16},
		Increment: PostIncrement,
	}
	addUses(load)

	ret := &RetInstr{}

	return []Instr{load, mov, ret}
}

func createRegister(name string, virtual bool, table map[string]*Register) *Register {
	if v, ok := table[name]; ok {
		return v
	}

	reg := &Register{
		Name:    name,
		Virtual: virtual,
	}
	table[name] = reg

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
	regNum++

	reg := &Register{
		Name:    fmt.Sprintf("_tmp%v", regNum),
		Virtual: true,
	}

	table[reg.Name] = reg
	return reg
}
