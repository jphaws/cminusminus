package ir

import (
	"fmt"
	"strconv"
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

type GepInstr struct {
	Target *Register
	Base   *Register
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
	target := ""
	if c.Target != nil {
		target = fmt.Sprintf("%v = ", c.Target)
	}

	// Fill argument strings list
	for _, v := range c.Arguments {
		args = append(args, fmt.Sprintf("%v %v", v.GetType(), v))
	}

	// Handle variadic argument types (if needed)
	vari := ""
	if c.Variadic > 0 {
		variTypes := make([]string, 0, c.Variadic)

		for i := 0; i < c.Variadic; i++ {
			variTypes = append(variTypes, fmt.Sprintf("%v", c.Arguments[i].GetType()))
		}

		vari = fmt.Sprintf(" (%v, ...)", strings.Join(variTypes, ", "))
	}

	// Create full string output
	return fmt.Sprintf("%vcall %v%v %v(%v)",
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

// === AST to LLVM (statements) ===
func statementToLlvm(stmt ast.Statement, locals map[string]*Register) []Instr {
	switch v := stmt.(type) {
	case *ast.AssignmentStatement:
		return assignmentStatementToLlvm(v, locals)
	case *ast.PrintStatement:
		return printStatementToLlvm(v, locals)
	case *ast.DeleteStatement:
		return deleteStatementToLlvm(v, locals)
	case *ast.InvocationStatement:
		ret, _ := invocationExpressionToLlvm(&v.Expression, locals, false)
		return ret
	}

	panic(fmt.Sprintf("Could not process statement of type %T", stmt))
}

func assignmentStatementToLlvm(asgn *ast.AssignmentStatement, locals map[string]*Register) []Instr {
	instrs, target := lValueToLlvm(asgn.Target, locals)

	// Read from stdin (if read assignment)
	if asgn.Source == nil {
		readInstrs := readToLlvm(target)
		instrs = append(instrs, readInstrs...)
		return instrs
	}

	// Otherwise, process expression
	expInstrs, val := expressionToLlvm(asgn.Source, locals, false)
	instrs = append(instrs, expInstrs...)

	return append(instrs, &StoreInstr{
		Mem: target,
		Reg: val,
	})
}

func lValueToLlvm(lval ast.LValue, locals map[string]*Register) (instrs []Instr, reg *Register) {
	switch v := lval.(type) {
	case *ast.NameLValue:
		// Base case
		reg = lookupSymbol(v.Name, locals)

	case *ast.DotLValue:
		// Recurse on left side
		var ptr *Register
		instrs, ptr = lValueToLlvm(v.Left, locals)

		// Create base register (pointer to start of struct)
		base := &Register{
			Name: getNextReg(),
			Type: ptr.GetType().(*PointerType).TargetType,
		}

		// Load base pointer
		instrs = append(instrs, &LoadInstr{
			Reg: base,
			Mem: ptr,
		})

		// Get field pointer
		var fieldInstrs []Instr
		fieldInstrs, reg = getFieldPointer(base, v.Name)
		instrs = append(instrs, fieldInstrs...)
	}

	return
}

func getFieldPointer(base *Register, fieldName string) (instrs []Instr, field *Register) {
	// Get struct ID
	id := base.GetType().(*PointerType).TargetType.(*StructType).Id

	// Get information for the relevant struct field
	f, _ := structTable[id].Fields.Get(fieldName)

	// Create field register (pointer to field in struct)
	field = &Register{
		Name: getNextReg(),
		Type: &PointerType{f.Type},
	}

	// Load base pointer and move pointer forward to field
	instrs = append(
		instrs,
		&GepInstr{
			Target: field,
			Base:   base,
			Index:  f.Index,
		},
	)

	return
}

func readToLlvm(target *Register) []Instr {
	format := "@" + scanStrName

	return []Instr{&CallInstr{
		FnName:     "@scanf",
		ReturnType: &IntType{32},
		Arguments: []Value{
			&Literal{
				Value: format,
				Type:  &PointerType{&IntType{8}},
			},
			target,
		},
		Variadic: 1,
	}}
}

func printStatementToLlvm(prnt *ast.PrintStatement, locals map[string]*Register) []Instr {
	// Process expression
	instrs, val := expressionToLlvm(prnt.Expression, locals, false)

	// Select relevant format string
	format := "@"
	if prnt.Newline {
		format += printlnStrName
	} else {
		format += printStrName
	}

	instrs = append(instrs, &CallInstr{
		FnName:     "@printf",
		ReturnType: &IntType{32},
		Arguments: []Value{
			&Literal{
				Value: format,
				Type:  &PointerType{&IntType{8}},
			},
			val,
		},
		Variadic: 1,
	})

	return instrs
}

func deleteStatementToLlvm(del *ast.DeleteStatement, locals map[string]*Register) []Instr {

	instrs, reg := expressionToLlvm(del.Expression, locals, false)

	instrs = append(instrs, &CallInstr{
		FnName:     "@free",
		ReturnType: &VoidType{},
		Arguments:  []Value{reg},
	})

	return instrs
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
	instrs, val := expressionToLlvm(ret.Expression, locals, false)

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
func expressionToLlvm(expr ast.Expression, locals map[string]*Register,
	isGuard bool) (instrs []Instr, val Value) {

	switch v := expr.(type) {
	case *ast.InvocationExpression:
		return invocationExpressionToLlvm(v, locals, true)
	case *ast.DotExpression:
		return dotExpressionToLlvm(v, locals)
	case *ast.UnaryExpression:
		return unaryExpressionToLlvm(v, locals, isGuard)
	case *ast.BinaryExpression:
		return binaryExpressionToLlvm(v, locals, isGuard)
	case *ast.IdentifierExpression:
		return identifierExpressionToLlvm(v, locals)
	case *ast.IntExpression:
		val = intExpressionToLlvm(v)
	case *ast.TrueExpression:
		val = trueExpressionToLlvm(isGuard)
	case *ast.FalseExpression:
		val = falseExpressionToLlvm(isGuard)
	case *ast.NewExpression:
		return newExpressionToLlvm(v)
	case *ast.NullExpression:
		val = nullExpressionToLlvm()
	}

	return
}

func invocationExpressionToLlvm(inv *ast.InvocationExpression,
	locals map[string]*Register, isExpr bool) (instrs []Instr, val Value) {

	args := make([]Value, 0, len(inv.Arguments))

	// Evaluate arguments
	for _, v := range inv.Arguments {
		argInstrs, argVal := expressionToLlvm(v, locals, false)
		instrs = append(instrs, argInstrs...)
		args = append(args, argVal)
	}

	// Get return type
	retType := functionTable[inv.Name].ReturnType

	// Create target register (if needed)
	var target *Register
	if isExpr {
		target = &Register{
			Name: getNextReg(),
			Type: retType,
		}
	}

	// Build call instruction
	instrs = append(instrs, &CallInstr{
		Target:     target,
		FnName:     "@" + inv.Name,
		ReturnType: retType,
		Arguments:  args,
	})

	val = target
	return
}

func dotExpressionToLlvm(dot *ast.DotExpression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	// Evaluate left expression
	instrs, lVal := expressionToLlvm(dot.Left, locals, false)

	// Get field pointer
	fieldInstrs, field := getFieldPointer(lVal.(*Register), dot.Field)
	instrs = append(instrs, fieldInstrs...)

	// Load value from field pointer
	val = &Register{
		Name: getNextReg(),
		Type: field.GetType().(*PointerType).TargetType,
	}

	instrs = append(instrs, &LoadInstr{
		Reg: val.(*Register),
		Mem: field,
	})

	return
}

func unaryExpressionToLlvm(una *ast.UnaryExpression,
	locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {

	switch una.Operator {
	case ast.NotOperator:
		return notOpToLlvm(una, locals, isGuard)
	case ast.MinusOperator:
		return minusOpToLlvm(una, locals)
	}

	return
}

func notOpToLlvm(not *ast.UnaryExpression,
	locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {

	// Select desired width
	width := 64
	if isGuard {
		width = 1
	}

	// Process operand expression
	instrs, oVal := expressionToLlvm(not.Operand, locals, isGuard)

	// Truncate expression (if needed)
	var convInstrs []Instr
	convInstrs, oVal = convertBoolWidth(oVal, isGuard)
	instrs = append(instrs, convInstrs...)

	// Create not instruction (1 ^ value)
	val = &Register{
		Name: getNextReg(),
		Type: &IntType{width},
	}

	instrs = append(instrs, &BinaryInstr{
		Target:   val.(*Register),
		Operator: XorOperator,
		Op1:      trueExpressionToLlvm(isGuard),
		Op2:      oVal,
	})

	return
}

func minusOpToLlvm(not *ast.UnaryExpression, locals map[string]*Register) (instrs []Instr, val Value) {
	// Process operand expression
	instrs, oVal := expressionToLlvm(not.Operand, locals, false)

	val = &Register{
		Name: getNextReg(),
		Type: &IntType{64},
	}

	// Create negation instruction (0 - value)
	instrs = append(instrs, &BinaryInstr{
		Target:   val.(*Register),
		Operator: SubOperator,
		Op1: &Literal{
			Value: "0",
			Type:  &IntType{64},
		},
		Op2: oVal,
	})

	return
}

func binaryExpressionToLlvm(bin *ast.BinaryExpression,
	locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {

	// Process left and right expressions
	instrs, lVal := expressionToLlvm(bin.Left, locals, isGuard)
	rInstrs, rVal := expressionToLlvm(bin.Right, locals, isGuard)

	instrs = append(instrs, rInstrs...)

	// Generate instructions depending on binary expression type
	condOp := operatorToLlvm(bin.Operator)

	switch v := condOp.(type) {
	case Operator:
		// Truncate expressions (if needed)
		switch v {
		case AndOperator:
			fallthrough

		case OrOperator:
			var convInstrs []Instr
			convInstrs, lVal = convertBoolWidth(lVal, isGuard)
			instrs = append(instrs, convInstrs...)

			convInstrs, rVal = convertBoolWidth(rVal, isGuard)
			instrs = append(instrs, convInstrs...)
		}

		// Use a binary instruction for arithmetic/boolean
		val = &Register{
			Name: getNextReg(),
			Type: lVal.GetType(),
		}

		instrs = append(instrs, &BinaryInstr{
			Target:   val.(*Register),
			Operator: v,
			Op1:      lVal,
			Op2:      rVal,
		})

	case Condition:
		// Otherwise, use a compare instruction
		reg := &Register{
			Name: getNextReg(),
			Type: &IntType{1},
		}

		instrs = append(instrs, &CompInstr{
			Target:    reg,
			Condition: v,
			Op1:       lVal,
			Op2:       rVal,
		})

		// Width-extend the bool if needed
		var convInstrs []Instr
		convInstrs, val = convertBoolWidth(reg, isGuard)
		instrs = append(instrs, convInstrs...)
	}

	return
}

func identifierExpressionToLlvm(ident *ast.IdentifierExpression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	reg := lookupSymbol(ident.Name, locals)

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

func trueExpressionToLlvm(isGuard bool) *Literal {
	width := 64
	if isGuard {
		width = 1
	}

	return &Literal{
		Value: "1",
		Type:  &IntType{width},
	}
}

func falseExpressionToLlvm(isGuard bool) *Literal {
	width := 64
	if isGuard {
		width = 1
	}

	return &Literal{
		Value: "0",
		Type:  &IntType{width},
	}
}

func newExpressionToLlvm(nw *ast.NewExpression) (instrs []Instr, val Value) {
	size := structTable[nw.Id].Size

	val = &Register{
		Name: getNextReg(),
		Type: &PointerType{&StructType{nw.Id}},
	}

	instrs = append(instrs, &CallInstr{
		Target:     val.(*Register),
		FnName:     "@malloc",
		ReturnType: val.GetType(),
		Arguments: []Value{
			&Literal{
				Value: strconv.Itoa(size),
				Type:  &IntType{64},
			},
		},
	})

	return
}

func nullExpressionToLlvm() Value {
	return &Literal{
		Value: "null",
		Type:  &PointerType{},
	}
}

func lookupSymbol(name string, locals map[string]*Register) *Register {
	var ret *Register

	if r, ok := locals[name]; ok {
		ret = r
	} else {
		ret = symbolTable[name]
	}

	return ret
}

func convertBoolWidth(src Value, isGuard bool) (instrs []Instr, val Value) {
	// Select desired width
	desiredWidth := 64
	if isGuard {
		desiredWidth = 1
	}

	// Transparently convert bool literals
	switch v := src.(type) {
	case *Literal:
		val = &Literal{
			Value: v.Value,
			Type:  &IntType{desiredWidth},
		}

	case *Register:
		width := src.GetType().(*IntType).Width

		// Truncate registers (if needed)
		if width > desiredWidth {
			var convInstr Instr
			convInstr, val = boolTruncReg(v)
			instrs = append(instrs, convInstr)

			// Or extend registers (if needed)
		} else if width < desiredWidth {
			var convInstr Instr
			convInstr, val = boolExtendReg(v)
			instrs = append(instrs, convInstr)

			// Otherwise, don't convert
		} else {
			val = src
		}
	}

	return
}

func boolExtendReg(src *Register) (instr Instr, reg *Register) {
	reg = &Register{
		Name: getNextReg(),
		Type: &IntType{64},
	}

	instr = &ConvInstr{
		Target:     reg,
		Conversion: ZeroExtConversion,
		Src:        src,
	}

	return
}

func boolTruncReg(src *Register) (instr Instr, reg *Register) {
	reg = &Register{
		Name: getNextReg(),
		Type: &IntType{1},
	}

	instr = &ConvInstr{
		Target:     reg,
		Conversion: TruncConversion,
		Src:        src,
	}

	return
}

func createGuardLlvm(guard ast.Expression,
	locals map[string]*Register) (instrs []Instr, val Value) {

	// Process guard expression
	instrs, guardVal := expressionToLlvm(guard, locals, true)

	// Truncate if needed
	convInstrs, val := convertBoolWidth(guardVal, true)
	instrs = append(instrs, convInstrs...)

	return
}

func functionInitLlvm(fn *ast.Function) (instrs []Instr,
	locals map[string]*Register, params []*Register) {

	// Allocate space for local variables and parameters
	instrs = make([]Instr, 0, 2*len(fn.Parameters)+len(fn.Locals))
	locals = make(map[string]*Register, len(fn.Parameters)+len(fn.Locals))
	params = make([]*Register, 0, len(fn.Parameters))

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

func createJump(next *Block) Instr {
	return &BranchInstr{
		Next: next,
	}
}

func createBranch(cond Value, next *Block, els *Block) Instr {
	return &BranchInstr{
		Cond: cond,
		Next: next,
		Els:  els,
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
	}

	panic("Unsupported type")
}

func operatorToLlvm(op ast.Operator) CondOp {
	switch op {
	case ast.TimesOperator:
		return MulOperator
	case ast.DivideOperator:
		return DivOperator
	case ast.PlusOperator:
		return AddOperator
	case ast.MinusOperator:
		return SubOperator
	case ast.LessThanOperator:
		return LessThanCondition
	case ast.GreaterThanOperator:
		return GreaterThanCondition
	case ast.LessEqualOperator:
		return LessEqualCondition
	case ast.GreaterEqualOperator:
		return GreaterEqualCondition
	case ast.EqualOperator:
		return EqualCondition
	case ast.NotEqualOperator:
		return NotEqualCondition
	case ast.AndOperator:
		return AndOperator
	case ast.OrOperator:
		return OrOperator
	}

	panic("Unsupported operator")
}

func getNextReg() string {
	regNum++
	return fmt.Sprintf("%%_r%v", regNum)
}
