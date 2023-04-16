package ast // Mini

import "fmt"

type Root struct {
	Types        []*TypeDeclaration
	Declarations []*Declaration
	Functions    []*Function
}

// === Type declarations ===
type TypeDeclaration struct {
	Position *Position
	Id       string
	Fields   []*Declaration
}

// === Declarations ===
// TODO: Add comma support
type Declaration struct {
	Position *Position
	Type     Type
	Name     string
}

// === Types ===
type Type interface {
	typeFunc()
}

type IntType struct {
}

func (i IntType) GoString() string {
	return fmt.Sprintf("int")
}

func (i IntType) typeFunc() {}

type BoolType struct {
}

func (b BoolType) GoString() string {
	return fmt.Sprintf("bool")
}

func (b BoolType) typeFunc() {}

type StructType struct {
	Id string
}

func (s StructType) GoString() string {
	return fmt.Sprintf("struct %v", s.Id)
}

func (s StructType) typeFunc() {}

type VoidType struct {
}

func (v VoidType) GoString() string {
	return fmt.Sprintf("void")
}

func (v VoidType) typeFunc() {}

// === Function ===
type Function struct {
	Position   *Position
	Name       string
	Parameters []*Declaration
	ReturnType Type
	Locals     []*Declaration
	Body       []Statement
}

// === Statements ===
type Statement interface {
	statementFunc()
}

type BlockStatement struct {
	Position   *Position
	Statements []Statement
}

func (b BlockStatement) statementFunc() {}

type AssignmentStatement struct {
	Position *Position
	Target   LValue
	Source   Expression
}

func (a AssignmentStatement) statementFunc() {}

type PrintStatement struct {
	Position   *Position
	Expression Expression
	Newline    bool
}

func (p PrintStatement) statementFunc() {}

type IfStatement struct {
	Position *Position
	Guard    Expression
	Then     *BlockStatement
	Else     *BlockStatement
}

func (i IfStatement) statementFunc() {}

type WhileStatement struct {
	Position *Position
	Guard    Expression
	Body     *BlockStatement
}

func (w WhileStatement) statementFunc() {}

type DeleteStatement struct {
	Position   *Position
	Expression Expression
}

func (d DeleteStatement) statementFunc() {}

type ReturnStatement struct {
	Position   *Position
	Expression Expression
}

func (r ReturnStatement) statementFunc() {}

type InvocationStatement struct {
	Expression InvocationExpression
}

func (i InvocationStatement) statementFunc() {}

// === LValues ===
type LValue interface {
	lValueFunc()
}

type DotLValue struct {
	Position *Position
	Left     LValue
	Name     string
}

func (d DotLValue) lValueFunc() {}

type NameLValue struct {
	Position *Position
	Name     string
}

func (n NameLValue) lValueFunc() {}

// === Expressions ===
type Expression interface {
	expressionFunc()
}

type DotLeftExpression interface {
	dotLeftExpressionFunc()
}

type BinaryLeftExpression interface {
	binaryLeftExpressionFunc()
}

type DotExpression struct {
	Position *Position
	Left     DotLeftExpression
	Right    Expression
}

func (d DotExpression) expressionFunc()           {}
func (d DotExpression) binaryLeftExpressionFunc() {}

type BinaryExpression struct {
	Left     BinaryLeftExpression
	Operator string
	Right    Expression
}

func (b BinaryExpression) expressionFunc()        {}
func (b BinaryExpression) dotLeftExpressionFunc() {}

type InvocationExpression struct {
	Position  *Position
	Name      string
	Arguments []Expression
}

func (i InvocationExpression) expressionFunc()           {}
func (i InvocationExpression) dotLeftExpressionFunc()    {}
func (i InvocationExpression) binaryLeftExpressionFunc() {}

// TODO: Unary expression
// TODO: Binary expression precedence
// TODO: Nested expressions

type ReadExpression struct {
	Position *Position
}

func (r ReadExpression) expressionFunc()           {}
func (r ReadExpression) dotLeftExpressionFunc()    {}
func (r ReadExpression) binaryLeftExpressionFunc() {}

type IntExpression struct {
	Position *Position
	Value    int
}

func (i IntExpression) expressionFunc()           {}
func (i IntExpression) dotLeftExpressionFunc()    {}
func (i IntExpression) binaryLeftExpressionFunc() {}

type TrueExpression struct {
	Position *Position
}

func (t TrueExpression) expressionFunc()           {}
func (t TrueExpression) dotLeftExpressionFunc()    {}
func (t TrueExpression) binaryLeftExpressionFunc() {}

type FalseExpression struct {
	Position *Position
}

func (f FalseExpression) expressionFunc()           {}
func (f FalseExpression) dotLeftExpressionFunc()    {}
func (f FalseExpression) binaryLeftExpressionFunc() {}

type NewExpression struct {
	Position *Position
	Id       string
}

func (n NewExpression) expressionFunc()           {}
func (n NewExpression) dotLeftExpressionFunc()    {}
func (n NewExpression) binaryLeftExpressionFunc() {}

type NullExpression struct {
	Position *Position
}

func (n NullExpression) expressionFunc()           {}
func (n NullExpression) dotLeftExpressionFunc()    {}
func (n NullExpression) binaryLeftExpressionFunc() {}

type IdentifierExpression struct {
	Position *Position
	Name     string
}

func (i IdentifierExpression) expressionFunc()           {}
func (i IdentifierExpression) dotLeftExpressionFunc()    {}
func (i IdentifierExpression) binaryLeftExpressionFunc() {}

type Position struct {
	Line   int
	Column int
}

func (p Position) GoString() string {
	return fmt.Sprintf("Line: %v, Col: %v", p.Line, p.Column)
}
