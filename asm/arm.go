package asm

import (
	"fmt"
)

// === Block ===
type Block struct {
	Label  string
	Instrs []Instr
	Next   []*Block
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
	Dst       *Register
	Base      *Register
	Offset    Offset
	Increment Increment
}

func (l LoadInstr) getDsts() []*Register {
	return []*Register{l.Dst}
}

func (l LoadInstr) getSrcs() []Operand {
	return []Operand{l.Base}
}

func (l LoadInstr) String() string {
	offset := l.Offset.Calculate()

	var offStr, postStr string
	if offset != 0 {
		switch l.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", offset)
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
	return fmt.Sprintf("ldr %v, =%v", l.Dst, l.Imm.Value)
}

type LoadPairInstr struct {
	Dst1      *Register
	Dst2      *Register
	Base      *Register
	Offset    Offset
	Increment Increment
}

func (l LoadPairInstr) getDsts() []*Register {
	return []*Register{l.Dst1, l.Dst2}
}

func (l LoadPairInstr) getSrcs() []Operand {
	return []Operand{l.Base}
}

func (l LoadPairInstr) String() string {
	offset := l.Offset.Calculate()

	var offStr, postStr string
	if offset != 0 {
		switch l.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", offset)
		}
	}

	return fmt.Sprintf("ldp %v, %v, [%v%v]%v", l.Dst1, l.Dst2, l.Base, offStr, postStr)
}

type StoreInstr struct {
	Src       *Register
	Base      *Register
	Offset    Offset
	Increment Increment
}

func (s StoreInstr) getDsts() []*Register {
	return nil
}

func (s StoreInstr) getSrcs() []Operand {
	return []Operand{s.Src, s.Base}
}

func (s StoreInstr) String() string {
	offset := s.Offset.Calculate()

	var offStr, postStr string
	if offset != 0 {
		switch s.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", offset)
		}
	}

	return fmt.Sprintf("ldr %v, [%v%v]%v", s.Src, s.Base, offStr, postStr)
}

type StorePairInstr struct {
	Src1      *Register
	Src2      *Register
	Base      Operand
	Offset    Offset
	Increment Increment
}

func (s StorePairInstr) getDsts() []*Register {
	return nil
}

func (s StorePairInstr) getSrcs() []Operand {
	return []Operand{s.Src1, s.Src2, s.Base}
}

func (s StorePairInstr) String() string {
	offset := s.Offset.Calculate()

	var offStr, postStr string
	if offset != 0 {
		switch s.Increment {
		case PostIncrement:
			postStr = fmt.Sprintf(", %v", offset)
		case PreIncrement:
			postStr = "!"
			fallthrough
		case NoIncrement:
			offStr = fmt.Sprintf(", %v", offset)
		}
	}

	return fmt.Sprintf("stp %v, %v, [%v%v]%v", s.Src1, s.Src2, s.Base, offStr, postStr)
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
	Value string
}

func (i Immediate) operandFunc() {}

func (i Immediate) String() string {
	return i.Value
}

// === Offset ===
// Calculated offsets are *SectionOff + Off. The section offset can be to point to an unknown
// offset which will be finalized later.
// For example, the section offset could be the offset from fp to the spill frame section.
type Offset struct {
	SectionOff *int
	Off        int
}

func (o Offset) Calculate() int {
	frameOffset := 0
	if o.SectionOff != nil {
		frameOffset = *o.SectionOff
	}

	return frameOffset + o.Off
}

// === Increment ===
type Increment int

const (
	NoIncrement Increment = iota
	PreIncrement
	PostIncrement
)
