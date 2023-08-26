package ir // Mini

import (
	"fmt"
	"strconv"

	"github.com/jphaws/cminusminus/ast"
)

const (
	RetValName  = "_retval"
	RetPtrName  = "_ret"
	ReadPtrName = "_read"
)

var regNum = 0

// === Statements ===
func statementToLlvm(stmt ast.Statement, curr *Block, locals map[string]*Register) []Instr {
	switch v := stmt.(type) {
	case *ast.AssignmentStatement:
		if stackLlvm {
			return assignmentStatementToLlvmStack(v, curr, locals)
		} else {
			return assignmentStatementToLlvmReg(v, curr, locals)
		}
	case *ast.PrintStatement:
		return printStatementToLlvm(v, curr, locals)
	case *ast.DeleteStatement:
		return deleteStatementToLlvm(v, curr, locals)
	case *ast.InvocationStatement:
		ret, _ := invocationExpressionToLlvm(&v.Expression, curr, locals, false)
		return ret
	}

	panic(fmt.Sprintf("Could not process statement of type %T", stmt))
}

func printStatementToLlvm(prnt *ast.PrintStatement,
	curr *Block, locals map[string]*Register) []Instr {

	// Process expression
	instrs, val := expressionToLlvm(prnt.Expression, curr, locals, false)

	// Select relevant format string
	var format string
	if prnt.Newline {
		format = PrintlnStrName
	} else {
		format = PrintStrName
	}

	call := &CallInstr{
		FnName:     "printf",
		ReturnType: &IntType{32},
		Arguments: []Value{
			&Literal{
				Value:  format,
				Global: true,
				Type:   &PointerType{&IntType{8}},
			},
			val,
		},
		Variadic: 1,
	}
	addDefUse(call)

	return append(instrs, call)
}

func deleteStatementToLlvm(del *ast.DeleteStatement,
	curr *Block, locals map[string]*Register) []Instr {

	instrs, reg := expressionToLlvm(del.Expression, curr, locals, false)

	call := &CallInstr{
		FnName:     "free",
		ReturnType: &VoidType{},
		Arguments:  []Value{reg},
	}
	addDefUse(call)

	return append(instrs, call)
}

func getFieldPointer(base Value, fieldName string,
	locals map[string]*Register) (gep *GepInstr, field *Register) {

	// Get struct ID
	id := base.GetType().(*PointerType).TargetType.(*StructType).Id

	// Get information for the relevant struct field
	f, _ := structTable[id].Fields.Get(fieldName)

	// Create field register (pointer to field in struct)
	name := nextRegName()
	field = &Register{
		Name: name,
		Type: &PointerType{f.Type},
	}
	locals[name] = field

	// Load base pointer and move pointer forward to field
	gep = &GepInstr{
		Target: field,
		Base:   base,
		Index:  f.Index,
	}
	addDefUse(gep)

	return
}

func readToLlvm(target Value) []Instr {
	call := &CallInstr{
		FnName:     "scanf",
		ReturnType: &IntType{32},
		Arguments: []Value{
			&Literal{
				Value:  ScanStrName,
				Global: true,
				Type:   &PointerType{&IntType{8}},
			},
			target,
		},
		Variadic: 1,
	}
	addDefUse(call)

	return []Instr{call}
}

// === Expressions ===
func expressionToLlvm(expr ast.Expression, curr *Block, locals map[string]*Register,
	isGuard bool) (instrs []Instr, val Value) {

	switch v := expr.(type) {
	case *ast.InvocationExpression:
		return invocationExpressionToLlvm(v, curr, locals, true)
	case *ast.DotExpression:
		return dotExpressionToLlvm(v, curr, locals)
	case *ast.UnaryExpression:
		return unaryExpressionToLlvm(v, curr, locals, isGuard)
	case *ast.BinaryExpression:
		return binaryExpressionToLlvm(v, curr, locals, isGuard)
	case *ast.IdentifierExpression:
		if stackLlvm {
			return identifierExpressionToLlvmStack(v, locals)
		} else {
			return identifierExpressionToLlvmReg(v, curr, locals)
		}
	case *ast.IntExpression:
		val = intExpressionToLlvm(v)
	case *ast.TrueExpression:
		val = trueExpressionToLlvm(isGuard)
	case *ast.FalseExpression:
		val = falseExpressionToLlvm(isGuard)
	case *ast.NewExpression:
		return newExpressionToLlvm(v, locals)
	case *ast.NullExpression:
		val = nullExpressionToLlvm()
	}

	return
}

func invocationExpressionToLlvm(inv *ast.InvocationExpression, curr *Block,
	locals map[string]*Register, isExpr bool) (instrs []Instr, val Value) {

	args := make([]Value, 0, len(inv.Arguments))

	// Evaluate arguments
	for _, v := range inv.Arguments {
		argInstrs, argVal := expressionToLlvm(v, curr, locals, false)
		instrs = append(instrs, argInstrs...)
		args = append(args, argVal)
	}

	// Get return type
	retType := functionTable[inv.Name].ReturnType

	// Create target register (if needed)
	var target *Register
	if isExpr {
		name := nextRegName()
		target = &Register{
			Name: name,
			Type: retType,
		}
		locals[name] = target
	}

	// Build call instruction
	call := &CallInstr{
		Target:     target,
		FnName:     inv.Name,
		ReturnType: retType,
		Arguments:  args,
	}
	addDefUse(call)

	instrs = append(instrs, call)
	val = target
	return
}

func dotExpressionToLlvm(dot *ast.DotExpression, curr *Block,
	locals map[string]*Register) (instrs []Instr, val Value) {

	// Evaluate left expression
	instrs, lVal := expressionToLlvm(dot.Left, curr, locals, false)

	// Get field pointer
	gepInstr, field := getFieldPointer(lVal, dot.Field, locals)
	instrs = append(instrs, gepInstr)

	// Load value from field pointer
	name := nextRegName()
	reg := &Register{
		Name: name,
		Type: field.GetType().(*PointerType).TargetType,
	}
	locals[name] = reg

	load := &LoadInstr{
		Reg: reg,
		Mem: field,
	}
	addDefUse(load)

	instrs = append(instrs, load)
	val = reg

	return
}

func unaryExpressionToLlvm(una *ast.UnaryExpression, curr *Block,
	locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {

	switch una.Operator {
	case ast.NotOperator:
		return notOpToLlvm(una, curr, locals, isGuard)
	case ast.MinusOperator:
		return minusOpToLlvm(una, curr, locals)
	}

	return
}

func notOpToLlvm(not *ast.UnaryExpression, curr *Block,
	locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {

	// Select desired width
	width := 64
	if isGuard {
		width = 1
	}

	// Process operand expression
	instrs, oVal := expressionToLlvm(not.Operand, curr, locals, isGuard)

	// Truncate expression (if needed)
	var convInstrs []Instr
	convInstrs, oVal = convertBoolWidth(oVal, locals, isGuard)
	instrs = append(instrs, convInstrs...)

	// Create not instruction (1 ^ value)
	name := nextRegName()
	reg := &Register{
		Name: name,
		Type: &IntType{width},
	}
	locals[name] = reg

	bin := &BinaryInstr{
		Target:   reg,
		Operator: XorOperator,
		Op1:      trueExpressionToLlvm(isGuard),
		Op2:      oVal,
	}
	addDefUse(bin)

	instrs = append(instrs, bin)
	val = reg

	return
}

func minusOpToLlvm(not *ast.UnaryExpression, curr *Block,
	locals map[string]*Register) (instrs []Instr, val Value) {

	// Process operand expression
	instrs, oVal := expressionToLlvm(not.Operand, curr, locals, false)

	name := nextRegName()
	reg := &Register{
		Name: name,
		Type: &IntType{64},
	}
	locals[name] = reg

	// Create negation instruction (0 - value)
	bin := &BinaryInstr{
		Target:   reg,
		Operator: SubOperator,
		Op1: &Literal{
			Value: "0",
			Type:  &IntType{64},
		},
		Op2: oVal,
	}
	addDefUse(bin)

	instrs = append(instrs, bin)
	val = reg

	return
}

func binaryExpressionToLlvm(bin *ast.BinaryExpression, curr *Block,
	locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {

	// Process left and right expressions
	instrs, lVal := expressionToLlvm(bin.Left, curr, locals, isGuard)
	rInstrs, rVal := expressionToLlvm(bin.Right, curr, locals, isGuard)

	instrs = append(instrs, rInstrs...)

	// Generate instructions depending on binary expression type
	condOp := operatorToLlvm(bin.Operator)

	switch v := condOp.(type) {
	case Operator:
		// Truncate expressions (if needed)
		switch v {
		case AndOperator, OrOperator:
			var convInstrs []Instr
			convInstrs, lVal = convertBoolWidth(lVal, locals, isGuard)
			instrs = append(instrs, convInstrs...)

			convInstrs, rVal = convertBoolWidth(rVal, locals, isGuard)
			instrs = append(instrs, convInstrs...)
		}

		// Use a binary instruction for arithmetic/boolean
		name := nextRegName()
		reg := &Register{
			Name: name,
			Type: lVal.GetType(),
		}
		locals[name] = reg

		bin := &BinaryInstr{
			Target:   reg,
			Operator: v,
			Op1:      lVal,
			Op2:      rVal,
		}
		instrs = append(instrs, bin)
		addDefUse(bin)
		val = reg

	case Condition:
		// Otherwise, use a compare instruction
		name := nextRegName()
		reg := &Register{
			Name: name,
			Type: &IntType{1},
		}
		locals[name] = reg

		cmp := &CompInstr{
			Target:    reg,
			Condition: v,
			Op1:       lVal,
			Op2:       rVal,
			IsGuard:   isGuard,
		}
		addDefUse(cmp)

		instrs = append(instrs, cmp)

		// Width-extend the bool if needed
		var convInstrs []Instr
		convInstrs, val = convertBoolWidth(reg, locals, isGuard)
		instrs = append(instrs, convInstrs...)
	}

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

func newExpressionToLlvm(nw *ast.NewExpression, locals map[string]*Register) (instrs []Instr, val Value) {
	size := structTable[nw.Id].Size

	name := nextRegName()
	reg := &Register{
		Name: name,
		Type: &PointerType{&StructType{nw.Id}},
	}
	locals[name] = reg

	malloc := &CallInstr{
		Target:     reg,
		FnName:     "malloc",
		ReturnType: reg.GetType(),
		Arguments: []Value{
			&Literal{
				Value: strconv.Itoa(size),
				Type:  &IntType{64},
			},
		},
	}
	addDefUse(malloc)
	instrs = append(instrs, malloc)

	val = reg
	return
}

func nullExpressionToLlvm() Value {
	return &Literal{
		Value: "null",
		Type:  &PointerType{},
	}
}

func convertBoolWidth(src Value, locals map[string]*Register, isGuard bool) (instrs []Instr, val Value) {
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
			convInstr, val = boolTruncReg(v, locals)
			instrs = append(instrs, convInstr)

			// Or extend registers (if needed)
		} else if width < desiredWidth {
			var convInstr Instr
			convInstr, val = boolExtendReg(v, locals)
			instrs = append(instrs, convInstr)

			// Otherwise, don't convert
		} else {
			val = src
		}
	}

	return
}

func boolExtendReg(src *Register, locals map[string]*Register) (instr Instr, reg *Register) {
	name := nextRegName()
	reg = &Register{
		Name: name,
		Type: &IntType{64},
	}

	locals[name] = reg

	instr = &ConvInstr{
		Target:     reg,
		Conversion: ZeroExtConversion,
		Src:        src,
	}
	addDefUse(instr)

	return
}

func boolTruncReg(src *Register, locals map[string]*Register) (instr Instr, reg *Register) {
	name := nextRegName()
	reg = &Register{
		Name: name,
		Type: &IntType{1},
	}

	locals[name] = reg

	instr = &ConvInstr{
		Target:     reg,
		Conversion: TruncConversion,
		Src:        src,
	}
	addDefUse(instr)

	return
}

// === Helpers ===
func createGuardLlvm(guard ast.Expression, curr *Block,
	locals map[string]*Register) (instrs []Instr, val Value) {

	// Process guard expression
	instrs, guardVal := expressionToLlvm(guard, curr, locals, true)

	// Truncate if needed
	convInstrs, val := convertBoolWidth(guardVal, locals, true)
	instrs = append(instrs, convInstrs...)

	return
}

func createJump(next *Block) Instr {
	return &BranchInstr{
		Next: next,
	}
}

func createBranch(cond Value, next *Block, els *Block) Instr {
	br := &BranchInstr{
		Cond: cond,
		Next: next,
		Els:  els,
	}
	addDefUse(br)

	return br
}

func typeToLlvm(typ ast.Type) Type {
	switch v := typ.(type) {
	case *ast.IntType, *ast.BoolType:
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

func nextRegName() string {
	regNum++
	return fmt.Sprintf("_r%v", regNum)
}

func addDefUse(instr Instr) {
	switch v := instr.(type) {
	case *AllocInstr:
		v.Target.Def = v

	case *LoadInstr:
		v.Reg.Def = v
		if reg, ok := v.Mem.(*Register); ok {
			reg.AddUse(v)
		}

	case *StoreInstr:
		if reg, ok := v.Mem.(*Register); ok {
			reg.AddUse(v)
		}
		if reg, ok := v.Reg.(*Register); ok {
			reg.AddUse(v)
		}

	case *GepInstr:
		v.Target.Def = v
		if reg, ok := v.Base.(*Register); ok {
			reg.AddUse(v)
		}

	case *CallInstr:
		if v.Target != nil {
			v.Target.Def = v
		}

		for _, arg := range v.Arguments {
			if reg, ok := arg.(*Register); ok {
				reg.AddUse(v)
			}
		}

	case *RetInstr:
		if reg, ok := v.Src.(*Register); ok {
			reg.AddUse(v)
		}

	case *CompInstr:
		v.Target.Def = v
		if reg, ok := v.Op1.(*Register); ok {
			reg.AddUse(v)
		}
		if reg, ok := v.Op2.(*Register); ok {
			reg.AddUse(v)
		}

	case *BranchInstr:
		if v.Cond != nil {
			if reg, ok := v.Cond.(*Register); ok {
				reg.AddUse(v)
			}
		}

	case *BinaryInstr:
		v.Target.Def = v
		if reg, ok := v.Op1.(*Register); ok {
			reg.AddUse(v)
		}
		if reg, ok := v.Op2.(*Register); ok {
			reg.AddUse(v)
		}

	case *ConvInstr:
		v.Target.Def = v
		if reg, ok := v.Src.(*Register); ok {
			reg.AddUse(v)
		}

	case *PhiInstr:
		v.Target.Def = v
		for _, phiVal := range v.Values {
			if reg, ok := phiVal.Value.(*Register); ok {
				reg.AddUse(v)
			}
		}
	}
}
