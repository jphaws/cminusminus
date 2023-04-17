package ast // Mini

import "fmt"

// === Root ===
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
type Declaration struct {
	Position *Position
	Type     Type
	Name     string
}

// === Types ===
type Type interface {
	typeFunc()
}

type IntType struct{}

func (i IntType) GoString() string {
	return fmt.Sprintf("int")
}

func (i IntType) typeFunc() {}

type BoolType struct{}

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

type VoidType struct{}

func (v VoidType) GoString() string {
	return fmt.Sprintf("void")
}

func (v VoidType) typeFunc() {}

// === Functions ===
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

type NameLValue struct {
	Position *Position
	Name     string
}

func (n NameLValue) lValueFunc() {}

type DotLValue struct {
	Position *Position
	Left     LValue
	Name     string
}

func (d DotLValue) lValueFunc() {}

// === Expressions ===
type Expression interface {
	expressionFunc()
}

type InvocationExpression struct {
	Position  *Position
	Name      string
	Arguments []Expression
}

func (i InvocationExpression) expressionFunc() {}

type DotExpression struct {
	Position *Position
	Left     Expression
	Field    string
}

func (d DotExpression) expressionFunc() {}

type UnaryExpression struct {
	Position *Position
	Operator Operator
	Operand  Expression
}

func (u UnaryExpression) expressionFunc() {}

type BinaryExpression struct {
	Position *Position
	Left     Expression
	Operator Operator
	Right    Expression
}

func (b BinaryExpression) expressionFunc() {}

type IdentifierExpression struct {
	Position *Position
	Name     string
}

func (i IdentifierExpression) expressionFunc() {}

type IntExpression struct {
	Position *Position
	Value    string
}

func (i IntExpression) expressionFunc() {}

type ReadExpression struct {
	Position *Position
}

func (r ReadExpression) expressionFunc() {}

type TrueExpression struct {
	Position *Position
}

func (t TrueExpression) expressionFunc() {}

type FalseExpression struct {
	Position *Position
}

func (f FalseExpression) expressionFunc() {}

type NewExpression struct {
	Position *Position
	Id       string
}

func (n NewExpression) expressionFunc() {}

type NullExpression struct {
	Position *Position
}

func (n NullExpression) expressionFunc() {}

// === Operators ===
type Operator string

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

func (o Operator) GoString() string {
	return string(o)
}

// === Position ===
type Position struct {
	Line   int
	Column int
}

func (p Position) GoString() string {
	return fmt.Sprintf("Line: %v, Col: %v", p.Line, p.Column)
}
