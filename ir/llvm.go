package ir

import (
	"fmt"
	"strconv"
	"strings"
)

// === Instructions ===
type Instr interface {
	instrFunc()
}

type AllocInstr struct {
	Target *Register
}

func (a AllocInstr) instrFunc() {}

func (a AllocInstr) String() string {
	typ := a.Target.GetType().(*PointerType)
	return fmt.Sprintf("%v = alloca %v", a.Target, typ.TargetType)
}

type LoadInstr struct {
	Reg *Register
	Mem Value
}

func (l LoadInstr) instrFunc() {}

func (l LoadInstr) String() string {
	memType := l.Mem.GetType()
	return fmt.Sprintf("%v = load %v, %v %v", l.Reg, l.Reg.GetType(), memType, l.Mem)
}

type StoreInstr struct {
	Mem Value
	Reg Value
}

func (s StoreInstr) instrFunc() {}

func (s StoreInstr) String() string {
	memType := s.Mem.GetType()
	return fmt.Sprintf("store %v %v, %v %v", s.Reg.GetType(), s.Reg, memType, s.Mem)
}

type GepInstr struct {
	Target *Register
	Base   Value
	Index  int
}

func (g GepInstr) instrFunc() {}

func (g GepInstr) String() string {
	targetType := g.Base.GetType().(*PointerType).TargetType
	return fmt.Sprintf("%v = getelementptr %%struct.%v, ptr %v, i32 0, i32 %v",
		g.Target, targetType, g.Base, g.Index)
}

type CallInstr struct {
	Target     *Register
	FnName     string
	ReturnType Type
	Arguments  []Value
	Variadic   int
}

func (c CallInstr) instrFunc() {}

func (c CallInstr) String() string {
	args := make([]string, 0, len(c.Arguments))

	// Only print a target if it exists
	var target string
	if c.Target != nil {
		target = fmt.Sprintf("%v = ", c.Target)
	}

	// Fill argument strings list
	for _, v := range c.Arguments {
		args = append(args, fmt.Sprintf("%v %v", v.GetType(), v))
	}

	// Handle variadic argument types (if needed)
	var vari string
	if c.Variadic > 0 {
		variTypes := make([]string, 0, c.Variadic)

		for i := 0; i < c.Variadic; i++ {
			variTypes = append(variTypes, fmt.Sprintf("%v", c.Arguments[i].GetType()))
		}

		vari = fmt.Sprintf(" (%v, ...)", strings.Join(variTypes, ", "))
	}

	// Create full string output
	return fmt.Sprintf("%vcall %v%v @%v(%v)",
		target, c.ReturnType, vari, c.FnName, strings.Join(args, ", "))
}

type RetInstr struct {
	Src Value
}

func (r RetInstr) instrFunc() {}

func (r RetInstr) String() string {
	if r.Src == nil {
		return "ret void"
	} else {
		return fmt.Sprintf("ret %v %v", r.Src.GetType(), r.Src)
	}
}

type CompInstr struct {
	Target    *Register
	Condition Condition
	Op1       Value
	Op2       Value
	IsGuard   bool
}

func (c CompInstr) instrFunc() {}

func (c CompInstr) String() string {
	return fmt.Sprintf("%v = icmp %v %v %v, %v",
		c.Target, c.Condition, c.Op1.GetType(), c.Op1, c.Op2)
}

type BranchInstr struct {
	Cond Value
	Next *Block
	Els  *Block
}

func (b BranchInstr) instrFunc() {}

func (b BranchInstr) String() string {
	if b.Cond == nil {
		return fmt.Sprintf("br label %%%v", b.Next.Label())
	} else {
		return fmt.Sprintf("br i1 %v, label %%%v, label %%%v", b.Cond, b.Next.Label(), b.Els.Label())
	}
}

type BinaryInstr struct {
	Target   *Register
	Operator Operator
	Op1      Value
	Op2      Value
}

func (b BinaryInstr) instrFunc() {}

func (b BinaryInstr) String() string {
	return fmt.Sprintf("%v = %v %v %v, %v", b.Target, b.Operator, b.Op1.GetType(), b.Op1, b.Op2)
}

type ConvInstr struct {
	Target     *Register
	Conversion Conversion
	Src        Value
}

func (c ConvInstr) instrFunc() {}

func (c ConvInstr) String() string {
	return fmt.Sprintf("%v = %v %v %v to %v",
		c.Target, c.Conversion, c.Src.GetType(), c.Src, c.Target.GetType())
}

type PhiInstr struct {
	Target *Register
	Values []*PhiVal
}

type PhiVal struct {
	Value Value
	Block *Block
}

func (p PhiInstr) instrFunc() {}

func (p PhiInstr) String() string {
	values := make([]string, 0, len(p.Values))

	for _, v := range p.Values {
		values = append(values, fmt.Sprintf("[%v, %%%v]", v.Value, v.Block.Label()))
	}

	return fmt.Sprintf("%v = phi %v %v", p.Target, p.Target.GetType(), strings.Join(values, ", "))
}

// === Conversions ===
type Conversion string

const (
	ZeroExtConversion Conversion = "zext"
	SignExtConversion Conversion = "sext"
	TruncConversion   Conversion = "trunc"
)

func (c Conversion) String() string {
	return string(c)
}

// === Conditions/operators ===
type CondOp interface {
	condOpFunc()
}

type Condition string

const (
	EqualCondition        Condition = "eq"
	NotEqualCondition     Condition = "ne"
	GreaterThanCondition  Condition = "sgt"
	GreaterEqualCondition Condition = "sge"
	LessThanCondition     Condition = "slt"
	LessEqualCondition    Condition = "sle"
)

func (c Condition) condOpFunc() {}

func (c Condition) String() string {
	return string(c)
}

type Operator string

const (
	AddOperator Operator = "add"
	SubOperator Operator = "sub"
	MulOperator Operator = "mul"
	DivOperator Operator = "sdiv"
	AndOperator Operator = "and"
	OrOperator  Operator = "or"
	XorOperator Operator = "xor"
)

func (o Operator) condOpFunc() {}

func (o Operator) String() string {
	return string(o)
}

// === Value ====
type Value interface {
	GetType() Type
	IsEqual(v Value) bool
}

type Register struct {
	Name   string
	Type   Type
	Global bool
	Def    Instr
	Uses   map[Instr]int
}

func (r *Register) AddUse(instr Instr) {
	if r.Global {
		return
	}

	if r.Uses == nil {
		r.Uses = map[Instr]int{}
	}

	r.Uses[instr]++
}

func (r *Register) DeleteUse(instr Instr) {
	if r.Global {
		return
	}

	r.Uses[instr]--

	if r.Uses[instr] == 0 {
		delete(r.Uses, instr)
	}
}

func (r Register) GetType() Type {
	return r.Type
}

func (r *Register) IsEqual(v Value) bool {
	if reg, ok := v.(*Register); ok {
		return r == reg
	}

	return false
}

func (r Register) String() string {
	if r.Global {
		return "@" + r.Name
	} else {
		return "%" + r.Name
	}
}

type Literal struct {
	Value  string
	Global bool
	Type   Type
}

func (l Literal) GetType() Type {
	return l.Type
}

// Check if a literal is equal to some other value
// Assumes that the type (.GetType()) for both l and v is the same
func (l *Literal) IsEqual(v Value) bool {
	var ol *Literal
	var ok bool
	if ol, ok = v.(*Literal); !ok {
		return false
	}

	// Assumption: literals with the same value are always represented by the same string
	return l.Value == ol.Value
}

func (l *Literal) DoCondition(ol *Literal, cond Condition, isGuard bool) (ret *Literal, err error) {
	// Check for literal "null" equality (or non-equality)
	// The only literal pointers in Mini are nulls, so we can get away with a simple EqualCondition
	// check here. Uses default value for ret if cond is NotEqualCondition
	if _, ok := l.GetType().(*PointerType); ok {
		ret = createCondLiteral(cond == EqualCondition, isGuard)
		return
	}

	// Convert integer literals
	var thisVal, thatVal int
	thisVal, err = strconv.Atoi(l.Value)
	if err != nil {
		return
	}

	thatVal, err = strconv.Atoi(ol.Value)
	if err != nil {
		return
	}

	// Perform comparison
	var result bool
	switch cond {
	case EqualCondition:
		result = thisVal == thatVal
	case NotEqualCondition:
		result = thisVal != thatVal
	case GreaterThanCondition:
		result = thisVal > thatVal
	case GreaterEqualCondition:
		result = thisVal >= thatVal
	case LessThanCondition:
		result = thisVal < thatVal
	case LessEqualCondition:
		result = thisVal <= thatVal
	default:
		panic("Unsupported condition")
	}

	ret = createCondLiteral(result, isGuard)
	return
}

func createCondLiteral(b bool, isGuard bool) *Literal {
	val := "0"
	if b {
		val = "1"
	}

	width := 64
	if isGuard {
		width = 1
	}

	return &Literal{
		Value: val,
		Type:  &IntType{width},
	}
}

func (l *Literal) DoOperation(ol *Literal, op Operator) (ret *Literal, err error) {
	var thisVal, thatVal, result int
	thisVal, err = strconv.Atoi(l.Value)
	if err != nil {
		return
	}

	thatVal, err = strconv.Atoi(ol.Value)
	if err != nil {
		return
	}

	switch op {
	case AddOperator:
		result = thisVal + thatVal
	case SubOperator:
		result = thisVal - thatVal
	case MulOperator:
		result = thisVal * thatVal
	case DivOperator:
		if thatVal == 0 {
			err = fmt.Errorf("Division by zero")
		} else {
			result = thisVal / thatVal
		}
	case AndOperator:
		result = thisVal & thatVal
	case OrOperator:
		result = thisVal | thatVal
	case XorOperator:
		result = thisVal ^ thatVal
	default:
		panic("Unsupported operator")
	}

	ret = &Literal{
		Value: strconv.Itoa(result),
		Type:  l.GetType(),
	}

	return
}

// Check the bool value of a literal
// Assumes that l is a boolean (has value "0" or "1")
func (l Literal) ToBool() bool {
	return l.Value == "1"
}

func (l Literal) ToInt() (val int, err error) {
	if l.Value == "null" {
		val = 0
		return
	}

	return strconv.Atoi(l.Value)
}

func (l Literal) String() string {
	if l.Global {
		return "@" + l.Value
	} else {
		return l.Value
	}
}

// === Types ===
type Type interface {
	typeFunc()
}

type IntType struct {
	Width int
}

func (i IntType) typeFunc() {}

func (i IntType) String() string {
	return fmt.Sprintf("i%v", i.Width)
}

type StructType struct {
	Id string
}

func (s StructType) typeFunc() {}

func (s StructType) String() string {
	return s.Id
}

type PointerType struct {
	TargetType Type
}

func (p PointerType) typeFunc() {}

func (p PointerType) String() string {
	return "ptr"
}

type VoidType struct{}

func (v VoidType) typeFunc() {}

func (v VoidType) String() string {
	return "void"
}
