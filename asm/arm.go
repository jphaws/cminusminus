package asm

import (
	"fmt"
)

// === Block ===
type Block struct {
	Label     string
	Instrs    []Instr
	EndInstrs []Instr
	Next      []*Block
}

// === Instructions ===
type Instr interface {
	getDsts() []*Register
	getSrcs() []Operand
}

type MovInstr struct {
	Dst *Register
	Src Operand
}

func (m MovInstr) getDsts() []*Register {
	return []*Register{m.Dst}
}

func (m MovInstr) getSrcs() []Operand {
	return []Operand{m.Src}
}

func (m MovInstr) String() string {
	return fmt.Sprintf("mov %v, %v", m.Dst, m.Src)
}

type LoadInstr struct {
	Dst        *Register
	Base       *Register
	Offset     int
	PageOffset string
	Increment  Increment
}

func (l LoadInstr) getDsts() []*Register {
	return []*Register{l.Dst}
}

func (l LoadInstr) getSrcs() []Operand {
	return []Operand{l.Base}
}

func (l LoadInstr) String() string {
	var offStr, postStr string
	if l.PageOffset != "" {
		offStr = fmt.Sprintf(", #:lo12:%v", l.PageOffset)

	} else if l.Offset != 0 {
		switch l.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", l.Offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", l.Offset)
		}
	}

	return fmt.Sprintf("ldr %v, [%v%v]%v", l.Dst, l.Base, offStr, postStr)
}

type LoadImmediateInstr struct {
	Dst *Register
	Imm *Immediate
}

func (l LoadImmediateInstr) getDsts() []*Register {
	return []*Register{l.Dst}
}

func (l LoadImmediateInstr) getSrcs() []Operand {
	return nil
}

func (l LoadImmediateInstr) String() string {
	return fmt.Sprintf("ldr %v, =%v", l.Dst, l.Imm)
}

type LoadPairInstr struct {
	Dst1      *Register
	Dst2      *Register
	Base      *Register
	Offset    int
	Increment Increment
}

func (l LoadPairInstr) getDsts() []*Register {
	return []*Register{l.Dst1, l.Dst2}
}

func (l LoadPairInstr) getSrcs() []Operand {
	return []Operand{l.Base}
}

func (l LoadPairInstr) String() string {
	var offStr, postStr string
	if l.Offset != 0 {
		switch l.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", l.Offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", l.Offset)
		}
	}

	return fmt.Sprintf("ldp %v, %v, [%v%v]%v", l.Dst1, l.Dst2, l.Base, offStr, postStr)
}

type StoreInstr struct {
	Src        *Register
	Base       *Register
	Offset     int
	PageOffset string
	Increment  Increment
}

func (s StoreInstr) getDsts() []*Register {
	return nil
}

func (s StoreInstr) getSrcs() []Operand {
	return []Operand{s.Src, s.Base}
}

func (s StoreInstr) String() string {
	var offStr, postStr string
	if s.PageOffset != "" {
		offStr = fmt.Sprintf(", #:lo12:%v", s.PageOffset)

	} else if s.Offset != 0 {
		switch s.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", s.Offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", s.Offset)
		}
	}

	return fmt.Sprintf("str %v, [%v%v]%v", s.Src, s.Base, offStr, postStr)
}

type StorePairInstr struct {
	Src1      *Register
	Src2      *Register
	Base      Operand
	Offset    int
	Increment Increment
}

func (s StorePairInstr) getDsts() []*Register {
	return nil
}

func (s StorePairInstr) getSrcs() []Operand {
	return []Operand{s.Src1, s.Src2, s.Base}
}

func (s StorePairInstr) String() string {
	var offStr, postStr string
	if s.Offset != 0 {
		switch s.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", s.Offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", s.Offset)
		}
	}

	return fmt.Sprintf("stp %v, %v, [%v%v]%v", s.Src1, s.Src2, s.Base, offStr, postStr)
}

type PageAddressInstr struct {
	Dst   *Register
	Label string
}

func (p PageAddressInstr) getDsts() []*Register {
	return []*Register{p.Dst}
}

func (p PageAddressInstr) getSrcs() []Operand {
	return nil
}

func (p PageAddressInstr) String() string {
	return fmt.Sprintf("adrp %v, %v", p.Dst, p.Label)
}

type ArithInstr struct {
	Operator Operator
	Dst      *Register
	Src1     *Register
	Src2     Operand
}

func (a ArithInstr) getDsts() []*Register {
	return []*Register{a.Dst}
}

func (a ArithInstr) getSrcs() []Operand {
	return []Operand{a.Src1, a.Src2}
}

func (a ArithInstr) String() string {
	return fmt.Sprintf("%v %v, %v, %v", a.Operator, a.Dst, a.Src1, a.Src2)
}

type CompInstr struct {
	Op1 *Register
	Op2 Operand
}

func (c CompInstr) getDsts() []*Register {
	return nil
}

func (c CompInstr) getSrcs() []Operand {
	return []Operand{c.Op1, c.Op2}
}

func (c CompInstr) String() string {
	return fmt.Sprintf("cmp %v, %v", c.Op1, c.Op2)
}

type TestInstr struct {
	Op1 *Register
	Op2 Operand
}

func (t TestInstr) getDsts() []*Register {
	return nil
}

func (t TestInstr) getSrcs() []Operand {
	return []Operand{t.Op1, t.Op2}
}

func (t TestInstr) String() string {
	return fmt.Sprintf("tst %v, %v", t.Op1, t.Op2)
}

type ConditionalSetInstr struct {
	Dst       *Register
	Condition Condition
}

func (c ConditionalSetInstr) getDsts() []*Register {
	return []*Register{c.Dst}
}

func (c ConditionalSetInstr) getSrcs() []Operand {
	return nil
}

func (c ConditionalSetInstr) String() string {
	return fmt.Sprintf("cset %v, %v", c.Dst, c.Condition)
}

type BranchInstr struct {
	Condition Condition
	Block     *Block
}

func (b BranchInstr) getDsts() []*Register {
	return nil
}

func (b BranchInstr) getSrcs() []Operand {
	return nil
}

func (b BranchInstr) String() string {
	var cond string
	if b.Condition != NoCondition {
		cond = fmt.Sprintf(".%v", b.Condition)
	}

	return fmt.Sprintf("b%v %v", cond, b.Block.Label)
}

type BranchLinkInstr struct {
	Label string
}

func (b BranchLinkInstr) getDsts() []*Register {
	return nil
}

func (b BranchLinkInstr) getSrcs() []Operand {
	return nil
}

func (b BranchLinkInstr) String() string {
	return fmt.Sprintf("bl %v", b.Label)
}

type RetInstr struct{}

func (m RetInstr) getDsts() []*Register {
	return nil
}

func (m RetInstr) getSrcs() []Operand {
	return nil
}

func (m RetInstr) String() string {
	return "ret"
}

// === Conditions ===
type Condition string

const (
	NoCondition           Condition = ""
	EqualCondition        Condition = "eq"
	NotEqualCondition     Condition = "ne"
	GreaterThanCondition  Condition = "gt"
	GreaterEqualCondition Condition = "ge"
	LessThanCondition     Condition = "lt"
	LessEqualCondition    Condition = "le"
)

func (c Condition) String() string {
	return string(c)
}

// === Operators ===
type Operator string

const (
	AddOperator Operator = "add"
	SubOperator Operator = "sub"
	MulOperator Operator = "mul"
	DivOperator Operator = "sdiv"
	AndOperator Operator = "and"
	OrOperator  Operator = "orr"
	XorOperator Operator = "eor"
)

func (o Operator) String() string {
	return string(o)
}

// === Operand ===
type Operand interface {
	operandFunc()
}

type Register struct {
	Name    string
	Virtual bool
	Uses    []Instr
}

func (r Register) operandFunc() {}

func (r Register) String() string {
	return r.Name
}

type Immediate struct {
	Value  string
	Global bool
}

func (i Immediate) operandFunc() {}

func (i Immediate) String() string {
	if i.Global {
		return "$" + i.Value
	} else {
		return i.Value
	}
}

// === Increment ===
type Increment int

const (
	NoIncrement Increment = iota
	PreIncrement
	PostIncrement
)
