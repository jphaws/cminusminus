package ir

import (
	"fmt"
	"strings"

	"github.com/keen-cp/compiler-project-c/ast"
)

const (
	retValName = "%_retval"
	retPtrName = "%_ret"
)

var regNum = 0

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
	Mem *Register
}

func (l LoadInstr) instrFunc() {}

func (l LoadInstr) String() string {
	memType := l.Mem.GetType().(*PointerType)
	return fmt.Sprintf("%v = load %v, %v %v", l.Reg, l.Reg.GetType(), memType, l.Mem)
}

type StoreInstr struct {
	Mem *Register
	Reg Value
}

func (s StoreInstr) instrFunc() {}

func (s StoreInstr) String() string {
	memType := s.Mem.GetType().(*PointerType)
	return fmt.Sprintf("store %v %v, %v %v", s.Reg.GetType(), s.Reg, memType, s.Mem)
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
	args := make([]string, len(c.Arguments))

	for i, v := range c.Arguments {
		args[i] = fmt.Sprintf("%v %v", v.GetType(), v)
	}

	vari := ""
	if c.Variadic != 0 {
		vari = fmt.Sprintf(" (%v, ...)", strings.Join(args[:c.Variadic], ", "))
	}

	return fmt.Sprintf("%v = call %v%v %v(%v)",
		c.Target, c.ReturnType, vari, c.FnName, strings.Join(args, ", "))
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

type BranchInstr struct {
	Cond *Register
	Next *Block
	Els  *Block
}

func (b BranchInstr) instrFunc() {}

func (b BranchInstr) String() string {
	if b.Cond == nil {
		return fmt.Sprintf("br label %v", b.Next.Label())
	} else {
		return fmt.Sprintf("br i1 %v, label %v, label %v", b.Cond, b.Next.Label(), b.Els.Label())
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

// === Operators ===
type Operator string

const (
	AddOperator = "add"
	SubOperator = "sub"
	MulOperator = "mul"
	DivOperator = "sdiv"
	AndOperator = "and"
	OrOperator  = "or"
)

func (o Operator) String() string {
	return string(o)
}

// === Value ====
type Value interface {
	GetType() Type
}

type Register struct {
	Name string
	Type Type
}

func (r Register) GetType() Type {
	return r.Type
}

func (r Register) String() string {
	return r.Name
}

type Literal struct {
	Value string
	Type  Type
}

func (l Literal) GetType() Type {
	return l.Type
}

func (l Literal) String() string {
	return l.Value
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

type FunctionType struct {
	ReturnType Type
}

func (f FunctionType) typeFunc() {}

func (f FunctionType) String() string {
	return "fun"
}

// === AST to LLVM (statements) ===
func statementToLlvm(stmt ast.Statement) []Instr {
	switch stmt.(type) {
	case *ast.AssignmentStatement:
		return nil
	case *ast.PrintStatement:
		return nil
	case *ast.DeleteStatement:
		return nil
	case *ast.ReturnStatement:
		return nil
	case *ast.InvocationStatement:
		return nil
	}

	panic("Trying to process incorrect statement")
}

func returnStatementToLlvm(ret *ast.ReturnStatement, funcExit *Block, locals map[string]*Register) []Instr {
	jump := &BranchInstr{
		Next: funcExit,
	}

	// Check for a void return
	if ret.Expression == nil {
		return []Instr{jump}
	}

	// Process expression
	instrs, val := expressionToLlvm(ret.Expression, locals)

	// Store expression value and jump to exit block
	return append(
		instrs,
		&StoreInstr{
			Mem: locals[retPtrName],
			Reg: val,
		},
		jump,
	)
}

// === AST to LLVM (expressions) ===
func expressionToLlvm(expr ast.Expression, locals map[string]*Register) (instrs []Instr, val Value) {
	// TODO: ?
	instrs = make([]Instr, 0)

	switch v := expr.(type) {
	case *ast.InvocationExpression:
		return invocationExpressionToLlvm(v, locals)
	case *ast.DotExpression:
		return
	case *ast.UnaryExpression:
		return
	case *ast.BinaryExpression:
		return binaryExpressionToLlvm(v, locals)
	case *ast.IdentifierExpression:
		return identifierExpressionToLlvm(v, locals)
	case *ast.IntExpression:
		val = intExpressionToLlvm(v)
	case *ast.TrueExpression:
		val = trueExpressionToLlvm()
	case *ast.FalseExpression:
		val = falseExpressionToLlvm()
	case *ast.NewExpression:
		return
	case *ast.NullExpression:
		return
	}

	return
}

func invocationExpressionToLlvm(inv *ast.InvocationExpression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	instrs = make([]Instr, 0)
	args := make([]Value, len(inv.Arguments))

	// Evaluate arguments
	for i, v := range inv.Arguments {
		argInstrs, argVal := expressionToLlvm(v, locals)
		instrs = append(instrs, argInstrs...)
		args[i] = argVal
	}

	// Get return type
	retType := symbolTable[inv.Name].GetType().(*PointerType).
		TargetType.(*FunctionType).ReturnType // TODO Wow!

	// Build call instruction
	val = &Register{
		Name: getNextReg(),
		Type: retType,
	}

	instrs = append(instrs, &CallInstr{
		Target:     val.(*Register),
		FnName:     "@" + inv.Name,
		ReturnType: retType,
		Arguments:  args,
	})

	return
}

func binaryExpressionToLlvm(bin *ast.BinaryExpression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	instrs, lVal := expressionToLlvm(bin.Left, locals)
	rInstrs, rVal := expressionToLlvm(bin.Right, locals)

	instrs = append(instrs, rInstrs...)

	// TODO OPTIMAZING!
	// _, lOk := lVal.(*IntType)

	val = &Register{
		Name: getNextReg(),
		Type: lVal.GetType(),
	}

	instrs = append(instrs, &BinaryInstr{
		Target:   val.(*Register),
		Operator: operatorToLlvm(bin.Operator),
		Op1:      lVal,
		Op2:      rVal,
	})

	return
}

/*
const (
	// Unary
	NotOperator Operator = "!"

	// Binary
	TimesOperator        = "*"
	DivideOperator       = "/"
	PlusOperator         = "+"
	MinusOperator        = "-" // Also unary
	LessThanOperator     = "<"
	GreaterThanOperator  = ">"
	LessEqualOperator    = "<="
	GreaterEqualOperator = ">="
	EqualOperator        = "=="
	NotEqualOperator     = "!="
	AndOperator          = "&&"
	OrOperator           = "||"
)
*/

func identifierExpressionToLlvm(ident *ast.IdentifierExpression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	var reg *Register
	if r, ok := locals[ident.Name]; ok {
		reg = r
	} else {
		reg = symbolTable[ident.Name]
	}

	val = &Register{
		Name: getNextReg(),
		Type: reg.GetType().(*PointerType).TargetType,
	}
	instrs = []Instr{&LoadInstr{
		Reg: val.(*Register),
		Mem: reg,
	}}

	return
}

func intExpressionToLlvm(tin *ast.IntExpression) *Literal {
	return &Literal{
		Value: tin.Value,
		Type:  &IntType{64},
	}
}

func trueExpressionToLlvm() *Literal {
	return &Literal{
		Value: "1",
		Type:  &IntType{64},
	}
}

func falseExpressionToLlvm() *Literal {
	return &Literal{
		Value: "0",
		Type:  &IntType{64},
	}
}

/*
func newExpressionToLlvm(nw *ast.NewExpression) (instrs []Instr, val Value) {
}
*/

func functionInitLlvm(fn *ast.Function) (instrs []Instr,
	locals map[string]*Register, params []*Register) {

	// Allocate space for local variables and parameters
	instrs = make([]Instr, 0)
	locals = make(map[string]*Register)
	params = make([]*Register, 0)

	for _, v := range fn.Parameters {
		pReg := &Register{
			Name: "%_p" + v.Name,
			Type: typeToLlvm(v.Type),
		}

		reg := &Register{
			Name: "%" + v.Name,
			Type: &PointerType{typeToLlvm(v.Type)},
		}

		instrs = append(
			instrs,
			&AllocInstr{reg},
			&StoreInstr{
				Mem: reg,
				Reg: pReg,
			},
		)

		locals[v.Name] = reg
		params = append(params, pReg)
	}

	for _, v := range fn.Locals {
		reg := &Register{
			Name: "%" + v.Name,
			Type: &PointerType{typeToLlvm(v.Type)},
		}
		instrs = append(instrs, &AllocInstr{reg})
		locals[v.Name] = reg
	}

	// Create return register (if needed)
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		return
	}

	retPtr := &Register{
		Name: retPtrName,
		Type: &PointerType{typeToLlvm(fn.ReturnType)},
	}

	instrs = append(instrs, &AllocInstr{
		Target: retPtr,
	})

	locals[retPtrName] = retPtr

	return
}

func functionFiniLlvm(fn *ast.Function, locals map[string]*Register) []Instr {
	// Store the return value in a register (if needed)
	if _, ok := fn.ReturnType.(*ast.VoidType); ok {
		return []Instr{&RetInstr{}}
	}

	retVal := &Register{
		Name: retValName,
		Type: typeToLlvm(fn.ReturnType),
	}

	// Create load and return instructions
	return []Instr{
		&LoadInstr{
			Reg: retVal,
			Mem: locals[retPtrName],
		},
		&RetInstr{retVal},
	}
}

func typeToLlvm(typ ast.Type) Type {
	switch v := typ.(type) {
	case *ast.IntType:
		return &IntType{64}
	case *ast.BoolType:
		return &IntType{64}
	case *ast.StructType:
		return &PointerType{&StructType{v.Id}}
	case *ast.VoidType:
		return &VoidType{}
	case *ast.NullType:
		return &PointerType{}
	case *ast.FunctionType:
		return &FunctionType{typeToLlvm(v.ReturnType)}
	}

	panic("Unsupported type")
}

func operatorToLlvm(op ast.Operator) Operator {
	switch op {
	case ast.NotOperator:
		panic("A")
	case ast.TimesOperator:
		return MulOperator
	case ast.DivideOperator:
		return DivOperator
	case ast.PlusOperator:
		return AddOperator
	case ast.MinusOperator:
		return SubOperator
	case ast.LessThanOperator:
		panic("A")
	case ast.GreaterThanOperator:
		panic("A")
	case ast.LessEqualOperator:
		panic("A")
	case ast.GreaterEqualOperator:
		panic("A")
	case ast.EqualOperator:
		panic("A")
	case ast.NotEqualOperator:
		panic("A")
	case ast.AndOperator:
		return AndOperator
	case ast.OrOperator:
		return OrOperator
	}

	panic("Unsupported operator")
}

/*
	AddOperator = "add"
	SubOperator = "sub"
	MulOperator = "mul"
	DivOperator = "sdiv"
	AndOperator = "and"
	OrOperator = "or"
*/

func getNextReg() string {
	regNum++
	return fmt.Sprintf("%%_r%v", regNum)
}
