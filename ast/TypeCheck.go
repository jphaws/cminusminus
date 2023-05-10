package ast // Mini

import (
	"errors"
	"fmt"

	"github.com/keen-cp/compiler-project-c/color"

	om "github.com/wk8/go-ordered-map/v2"
)

// (struct) 'Id' -> [field 'Name' -> Type]
var structTable = make(map[string]*om.OrderedMap[string, Type])

// 'Name' -> Type
var symbolTable = make(map[string]Type)

var lines []string

type Tables struct {
	StructTable map[string]*om.OrderedMap[string, Type]
	SymbolTable map[string]Type
}

// === Root ===
func TypeCheck(root *Root, lns []string) (tab *Tables, err error) {
	funcErr := make(chan error)
	lines = lns

	// Check type definitions and add to structTable
	err = generateStructTable(root.Types)

	// Generate global symbol table and check for errors
	globErr := generateSymbolTable(root.Declarations, root.Functions)
	err = errors.Join(err, globErr)

	if err != nil {
		return
	}

	tab = &Tables{
		StructTable: structTable,
		SymbolTable: symbolTable,
	}

	// Type check function bodies in goroutines
	for _, v := range root.Functions {
		go typeCheckFunction(v, funcErr)
	}
	// Synchronize completed functions
	for range root.Functions {
		err = errors.Join(err, <-funcErr)
	}

	// Check for a fun main() int
	err = errors.Join(err, typeCheckMain())

	return
}

func typeCheckMain() (err error) {
	typ, err := lookupGlobal("main", &Position{}, "function")
	if err != nil {
		return
	}

	if !typ.canConvertTo(&FunctionType{ReturnType: &IntType{}}) {
		err = createError("program must include function 'fun main() int'", &Position{})
	}

	return
}

// === Struct table ===
func generateStructTable(types []*TypeDeclaration) (err error) {
	for _, v := range types {
		id := v.Id

		// Check for re-declaration
		if _, ok := structTable[id]; ok {
			e := createError("re-declaration of 'struct %s'", v.Position, id)
			err = errors.Join(err, e)
			continue
		}

		// Check validity of fields
		om, e := addFields(id, v.Fields)
		err = errors.Join(err, e)

		// Add to struct table
		structTable[id] = om
	}

	return
}

func addFields(structId string,
	fields []*Declaration) (omap *om.OrderedMap[string, Type], err error) {

	// Create fields map
	omap = om.New[string, Type]()

	// Loop through all fields
	for _, f := range fields {
		name := f.Name

		// Check that struct field names are valid
		if v, ok := f.Type.(*StructType); ok {
			id := v.Id

			if _, present := structTable[id]; !present && id != structId {
				e := createError("'%v' not declared (yet?)", f.Position, v)
				err = errors.Join(err, e)
				continue
			}
		}

		// Add to the field map (and check for re-declaration)
		if _, present := omap.Set(name, f.Type); present {
			e := createError("re-declaration of struct field '%s'", f.Position, name)
			err = errors.Join(err, e)
		}
	}

	return
}

// === Symbol table ===
func generateSymbolTable(decls []*Declaration, funcs []*Function) (err error) {
	// Add global definitions
	err = addGlobals(decls)

	// Add functions
	funcErr := addFunctions(funcs)
	err = errors.Join(err, funcErr)

	return
}

func addGlobals(decls []*Declaration) (err error) {
	for _, v := range decls {
		e := addDeclaration(v, "global variable", symbolTable)
		err = errors.Join(err, e)
	}

	return
}

func addDeclaration(decl *Declaration, desc string, table map[string]Type) (err error) {
	name := decl.Name
	typ := decl.Type

	// Check if struct is declared
	_, err = checkStructDeclared(typ, decl.Position)
	if err != nil {
		return
	}

	// Check for re-declaration
	if _, present := table[name]; present {
		return createError("re-declaration of %v '%v'", decl.Position, desc, name)
	}

	// Add to the symbol table
	table[name] = typ
	return
}

func checkStructDeclared(typ Type, pos *Position) (om *om.OrderedMap[string, Type], err error) {
	if v, ok := typ.(*StructType); ok {
		var present bool
		if om, present = structTable[v.Id]; !present {
			err = createError("'%v' not defined", pos, v)
		}
	}

	return
}

func addFunctions(funcs []*Function) (err error) {
	for _, v := range funcs {
		name := v.Name

		// Check for function re-declaration
		if _, present := symbolTable[name]; present {
			e := createError("re-declaration of 'fun %v'", v.Position, name)
			err = errors.Join(err, e)
			continue
		}

		// Check parameters
		params := make([]Type, len(v.Parameters))

		for i, vv := range v.Parameters {
			_, e := checkStructDeclared(vv.Type, vv.Position)
			err = errors.Join(err, e)

			params[i] = vv.Type
		}

		// Add corresponding function type to symbol table
		typ := &FunctionType{
			Parameters: params,
			ReturnType: v.ReturnType,
		}
		symbolTable[name] = typ
	}

	return
}

// === Functions ===
func typeCheckFunction(fn *Function, errCh chan error) {
	// Generate local table
	localTable, err := generateLocalTable(fn)
	if err != nil {
		errCh <- err
		return
	}

	// Type check statements
	retEq := false
	for _, v := range fn.Body {
		stmtRetEq, e := typeCheckStatement(v, localTable, fn.Name)
		err = errors.Join(err, e)
		retEq = retEq || stmtRetEq
	}

	// Check for return equivalence
	if !fn.ReturnType.canConvertTo(&VoidType{}) && !retEq {
		e := createError("function '%v' missing return (in at least one path)",
			fn.Position, fn.Name)
		err = errors.Join(err, e)
	}

	errCh <- err
	return
}

func generateLocalTable(fn *Function) (localTable map[string]Type, err error) {
	localTable = make(map[string]Type, len(fn.Parameters))

	// Add parameters to local table
	for _, v := range fn.Parameters {
		e := addDeclaration(v, "parameter", localTable)
		err = errors.Join(err, e)
	}

	// Add local declarations to local table
	for _, v := range fn.Locals {
		e := addDeclaration(v, "local variable", localTable)
		err = errors.Join(err, e)
	}

	return
}

// === Statements ===
func typeCheckStatement(stmt Statement, localTable map[string]Type,
	fnName string) (retEq bool, err error) {

	switch v := stmt.(type) {
	case *BlockStatement:
		retEq, err = typeCheckBlockStatement(v.Statements, localTable, fnName)
	case *AssignmentStatement:
		err = typeCheckAssignmentStatement(v, localTable)
	case *PrintStatement:
		err = typeCheckPrintStatement(v, localTable)
	case *IfStatement:
		retEq, err = typeCheckIfStatement(v, localTable, fnName)
	case *WhileStatement:
		err = typeCheckWhileStatement(v, localTable, fnName)
	case *DeleteStatement:
		err = typeCheckDeleteStatement(v, localTable)
	case *ReturnStatement:
		retEq, err = typeCheckReturnStatement(v, localTable, fnName)
	case *InvocationStatement:
		err = typeCheckInvocationStatement(v, localTable)
	}

	return
}

func typeCheckBlockStatement(stmts []Statement,
	localTable map[string]Type, fnName string) (retEq bool, err error) {

	for _, v := range stmts {
		stmtRetEq, e := typeCheckStatement(v, localTable, fnName)
		err = errors.Join(err, e)
		retEq = retEq || stmtRetEq
	}

	return
}

func typeCheckAssignmentStatement(asgn *AssignmentStatement,
	localTable map[string]Type) (err error) {

	// Get type of left side
	left, err := typeCheckLValue(asgn.Target, localTable)

	// Get type of right side
	var right Type
	var rErr error
	if asgn.Source != nil {
		right, rErr = typeCheckExpression(asgn.Source, localTable)
		err = errors.Join(err, rErr)
	} else {
		right = &IntType{}
	}

	// Left side type <- right side type
	if !right.canConvertTo(left) {
		e := createError("cannot assign to '%v': %v <- %v",
			asgn.Position, asgn.Target, left, right)
		err = errors.Join(err, e)
	}

	return
}

func lookupLocal(name string, pos *Position,
	desc string, localTable map[string]Type) (typ Type, err error) {

	if v, present := localTable[name]; present {
		typ = v
	} else {
		typ = &ErrorType{}
		err = createError("%v '%v' not declared", pos, desc, name)
	}

	return
}

func lookupGlobal(name string, pos *Position, desc string) (typ Type, err error) {
	if v, present := symbolTable[name]; present {
		typ = v
	} else {
		typ = &ErrorType{}
		err = createError("%v '%v' not declared", pos, desc, name)
	}

	return
}

func lookupType(name string, pos *Position,
	desc string, localTable map[string]Type) (typ Type, err error) {

	lType, lErr := lookupLocal(name, pos, desc, localTable)
	if lErr == nil {
		typ = lType
		return
	}

	typ, err = lookupGlobal(name, pos, desc)
	return
}

func typeCheckLValue(target LValue, localTable map[string]Type) (typ Type, err error) {
	switch v := target.(type) {
	case *NameLValue:
		// Check if variable is defined
		typ, err = lookupType(v.Name, v.Position, "variable", localTable)
		if err != nil {
			typ = &ErrorType{}
			return
		}

	case *DotLValue:
		// Recurse and check left side
		var left Type
		left, err = typeCheckLValue(v.Left, localTable)
		if err != nil {
			typ = &ErrorType{}
			return
		}

		// Check if left side is a struct
		var struc *StructType
		var ok bool
		if struc, ok = left.(*StructType); !ok {
			err = createError("dotted value is not type struct", v.Position)
			typ = &ErrorType{}
			return
		}

		// Check if struct is declared
		var om *om.OrderedMap[string, Type]
		om, err = checkStructDeclared(left, v.Position)
		if err != nil {
			typ = &ErrorType{}
			return
		}

		// Check if field exists
		field, present := om.Get(v.Name)
		if !present {
			err = createError("'struct %v' has no field '%v'",
				v.Position, struc.Id, v.Name)
			typ = &ErrorType{}
			return
		}

		typ = field
		return

	}

	return
}

func typeCheckPrintStatement(prnt *PrintStatement, localTable map[string]Type) (err error) {
	// Check expression
	typ, err := typeCheckExpression(prnt.Expression, localTable)

	// Check that print expression is an int
	if !typ.canConvertTo(&IntType{}) {
		e := createError("cannot print expression (type %v)", prnt.Position, typ)
		err = errors.Join(err, e)
	}

	return
}

func typeCheckIfStatement(fi *IfStatement, localTable map[string]Type,
	fnName string) (retEq bool, err error) {

	// Check guard expression
	typ, err := typeCheckExpression(fi.Guard, localTable)

	// Check that guard expression is a bool
	if !typ.canConvertTo(&BoolType{}) {
		e := createError("if guard is not type bool", fi.Position)
		err = errors.Join(err, e)
	}

	// Check then expression
	thenRetEq, thenErr := typeCheckBlockStatement(fi.Then.Statements, localTable, fnName)
	err = errors.Join(err, thenErr)

	// Check else expression
	elseRetEq := false
	if fi.Else != nil {
		var elseErr error
		elseRetEq, elseErr = typeCheckBlockStatement(fi.Else.Statements, localTable, fnName)
		err = errors.Join(err, elseErr)
	}
	retEq = thenRetEq && elseRetEq

	return
}

func typeCheckWhileStatement(whl *WhileStatement,
	localTable map[string]Type, fnName string) (err error) {

	// Check guard expression
	typ, err := typeCheckExpression(whl.Guard, localTable)

	// Check that guard expression is a bool
	if !typ.canConvertTo(&BoolType{}) {
		e := createError("while guard is not type bool", whl.Position)
		err = errors.Join(err, e)
	}

	// Check body
	_, thenErr := typeCheckBlockStatement(whl.Body.Statements, localTable, fnName)
	err = errors.Join(err, thenErr)

	return
}

func typeCheckDeleteStatement(del *DeleteStatement, localTable map[string]Type) (err error) {
	// Check expression
	typ, err := typeCheckExpression(del.Expression, localTable)

	// Check that expression is a struct
	if !typ.canConvertTo(&StructType{}) {
		e := createError("cannot delete expression (type %v)", del.Position, typ)
		err = errors.Join(err, e)
	}

	return
}

func typeCheckReturnStatement(ret *ReturnStatement,
	localTable map[string]Type, fnName string) (retEq bool, err error) {

	// Check expression
	var typ Type
	if ret.Expression != nil {
		typ, err = typeCheckExpression(ret.Expression, localTable)
	} else {
		typ = &VoidType{}
	}

	// Get expected return type
	fn, _ := lookupGlobal(fnName, ret.Position, "- COMPILER ERROR -")

	// Check that given return type is as expected
	retType := fn.(*FunctionType).ReturnType

	if !typ.canConvertTo(retType) {
		e := createError("mismatched return type: %v <- %v", ret.Position, retType, typ)
		err = errors.Join(err, e)
	}

	retEq = true
	return
}

func typeCheckInvocationStatement(inv *InvocationStatement, localTable map[string]Type) (err error) {
	_, err = typeCheckInvocationExpression(&inv.Expression, localTable)
	return
}

// === Expressions ===
func typeCheckExpression(expr Expression, localTable map[string]Type) (typ Type, err error) {
	switch v := expr.(type) {
	case *InvocationExpression:
		typ, err = typeCheckInvocationExpression(v, localTable)
	case *DotExpression:
		typ, err = typeCheckDotExpression(v, localTable)
	case *UnaryExpression:
		typ, err = typeCheckUnaryExpression(v, localTable)
	case *BinaryExpression:
		typ, err = typeCheckBinaryExpression(v, localTable)
	case *IdentifierExpression:
		typ, err = typeCheckIdentifierExpression(v, localTable)
	case *IntExpression:
		typ = &IntType{}
	case *TrueExpression:
		typ = &BoolType{}
	case *FalseExpression:
		typ = &BoolType{}
	case *NewExpression:
		typ = &StructType{Id: v.Id}
	case *NullExpression:
		typ = &NullType{}
	}
	return
}

func typeCheckInvocationExpression(inv *InvocationExpression,
	localTable map[string]Type) (typ Type, err error) {

	name := inv.Name
	pos := inv.Position

	// Get attempted invocation type
	t, err := lookupType(name, pos, "invoked function", localTable)
	if err != nil {
		typ = &ErrorType{}
		return
	}

	// Check that attempted invocation variable is type function
	fn, ok := t.(*FunctionType)
	if !ok {
		err = createError("cannot invoke non-function '%v' (type %v)", pos, name, t)
		typ = &ErrorType{}
		return
	}

	typ = fn.ReturnType

	// Check correct number of arguments
	argLen := len(inv.Arguments)
	paramLen := len(fn.Parameters)

	if argLen != paramLen {
		paramLenPlural := "s"
		if paramLen == 1 {
			paramLenPlural = ""
		}

		err = createError("invalid invocation of function '%v': "+
			"expected %v argument%s, got %v", pos, name, paramLen, paramLenPlural, argLen)
		return
	}

	// Type check the arguments
	for i, v := range inv.Arguments {
		arg, e := typeCheckExpression(v, localTable)
		err = errors.Join(err, e)
		if !arg.canConvertTo(fn.Parameters[i]) {
			e := createError("argument %v to function '%v' is of type %v (expected %v)",
				pos, i+1, name, arg, fn.Parameters[i])
			err = errors.Join(err, e)
		}

	}

	return
}

func typeCheckDotExpression(dot *DotExpression, localTable map[string]Type) (typ Type, err error) {
	// Type check left expression
	left, err := typeCheckExpression(dot.Left, localTable)
	if err != nil {
		typ = &ErrorType{}
		return
	}

	// Check that left expression is a struct
	var struc *StructType
	var ok bool
	if struc, ok = left.(*StructType); !ok {
		err = createError("dotted expression is not type struct", dot.Position)
		typ = &ErrorType{}
		return
	}
	// Check if struct is declared
	var om *om.OrderedMap[string, Type]
	om, err = checkStructDeclared(left, dot.Position)
	if err != nil {
		typ = &ErrorType{}
		return
	}

	// Check if field exists
	field, present := om.Get(dot.Field)
	if !present {
		err = createError("'struct %v' has no field '%v'",
			dot.Position, struc.Id, dot.Field)
		typ = &ErrorType{}
		return
	}

	typ = field
	return
}

func typeCheckUnaryExpression(una *UnaryExpression,
	localTable map[string]Type) (typ Type, err error) {

	op := una.Operator
	pos := una.Position

	typ, err = typeCheckExpression(una.Operand, localTable)

	switch op {
	case MinusOperator:
		if !typ.canConvertTo(&IntType{}) {
			e := createError("cannot apply unary operator '%v' on non-integer expression "+
				"(of type %v)", pos, op, typ)
			err = errors.Join(err, e)
		}
		typ = &IntType{}

	case NotOperator:
		if !typ.canConvertTo(&BoolType{}) {
			e := createError("cannot apply unary operator '%v' on non-boolean expression "+
				"(of type %v)", pos, op, typ)
			err = errors.Join(err, e)
		}
		typ = &BoolType{}
	}

	return
}

func typeCheckBinaryExpression(bin *BinaryExpression,
	localTable map[string]Type) (typ Type, err error) {

	op := bin.Operator
	pos := bin.Position

	lType, err := typeCheckExpression(bin.Left, localTable)
	rType, rErr := typeCheckExpression(bin.Right, localTable)
	err = errors.Join(err, rErr)

	switch op {
	// Arithmetic operators (take int, return int)
	case TimesOperator:
		fallthrough
	case DivideOperator:
		fallthrough
	case PlusOperator:
		fallthrough
	case MinusOperator:
		if !lType.canConvertTo(&IntType{}) {
			e := createError("cannot apply arithmetic operator '%v' on non-integer expression "+
				"(of type %v)", pos, op, lType)
			err = errors.Join(err, e)
		}

		if !rType.canConvertTo(&IntType{}) {
			e := createError("cannot apply arithmetic operator '%v' on non-integer expression "+
				"(of type %v)", pos, op, rType)
			err = errors.Join(err, e)
		}
		typ = &IntType{}

	// Relational (take int, return bool)
	case LessThanOperator:
		fallthrough
	case GreaterThanOperator:
		fallthrough
	case LessEqualOperator:
		fallthrough
	case GreaterEqualOperator:
		if !lType.canConvertTo(&IntType{}) {
			e := createError("cannot apply relational operator '%v' on non-integer expression "+
				"(of type %v)", pos, op, lType)
			err = errors.Join(err, e)
		}

		if !rType.canConvertTo(&IntType{}) {
			e := createError("cannot apply relational operator '%v' on non-integer expression "+
				"(of type %v)", pos, op, rType)
			err = errors.Join(err, e)
		}
		typ = &BoolType{}

	// Equality operators (take int/struct, return bool)
	case EqualOperator:
		fallthrough
	case NotEqualOperator:
		if lType.canConvertTo(&IntType{}) {
			if !rType.canConvertTo(&IntType{}) {
				e := createError("cannot apply relational operator '%v' to expressions "+
					"(of types %v, %v)", pos, op, lType, rType)
				err = errors.Join(err, e)
			}

		} else if lType.canConvertTo(&StructType{}) {
			if !rType.canConvertTo(lType) {
				e := createError("cannot apply relational operator '%v' to expressions "+
					"(of types %v, %v)", pos, op, lType, rType)
				err = errors.Join(err, e)
			}
			typ = lType

		} else {
			e := createError("cannot apply relational operator '%v' on expression "+
				"(of type %v)", pos, op, lType)
			err = errors.Join(err, e)
		}
		typ = &BoolType{}

	// Boolean operators (take bool, return bool)
	case AndOperator:
		fallthrough
	case OrOperator:
		if !lType.canConvertTo(&BoolType{}) {
			e := createError("cannot apply boolean operator '%v' on non-boolean expression "+
				"(of type %v)", pos, op, lType)
			err = errors.Join(err, e)
		}

		if !rType.canConvertTo(&BoolType{}) {
			e := createError("cannot apply boolean operator '%v' on non-boolean expression "+
				"(of type %v)", pos, op, rType)
			err = errors.Join(err, e)
		}
		typ = &BoolType{}
	}

	return
}

func typeCheckIdentifierExpression(ident *IdentifierExpression,
	localTable map[string]Type) (typ Type, err error) {

	return lookupType(ident.Name, ident.Position, "variable", localTable)
}

func createError(format string, pos *Position, a ...any) error {
	s := fmt.Sprintf("%s: error: ", pos)
	s += fmt.Sprintf(format, a...)

	if pos.Line != 0 {
		s += fmt.Sprintf("\n %4v | %s%s%s%s\n      |",
			pos.Line, color.Red, color.Bright, lines[pos.Line], color.Reset)
	}

	return fmt.Errorf(s)
}
