package parser // Mini

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/keen-cp/compiler-project-c/ast"
	"github.com/keen-cp/compiler-project-c/parser/mantlr"
)

// === Root ===

// MiniToAst is the entry function for creating AST from ANTLR parse tree
func MiniToAst(ctx mantlr.IProgramContext) *ast.Root {
	funcs := ctx.Functions().AllFunction()
	ch := make(chan *ast.Function)
	functions := make([]*ast.Function, 0, len(funcs))

	// Start Go Routines (1 per function)
	for _, v := range funcs {
		go functionToAst(v, ch)
	}

	types := typeDeclarationsToAst(ctx.Types().AllTypeDeclaration())
	declarations := declarationsToAst(ctx.Declarations().AllDeclaration())

	// Synchronize
	for range funcs {
		functions = append(functions, <-ch)
	}

	return &ast.Root{
		Types:        types,
		Declarations: declarations,
		Functions:    functions,
	}
}

// === Type declarations ===
func typeDeclarationsToAst(typeDecls []mantlr.ITypeDeclarationContext) []*ast.TypeDeclaration {
	ret := make([]*ast.TypeDeclaration, 0, len(typeDecls))

	for _, v := range typeDecls {
		ret = append(ret, &ast.TypeDeclaration{
			Position: constructPosition(v.ID().GetSymbol()),
			Id:       v.ID().GetText(),
			Fields:   gatherFieldDeclaration(v.NestedDecl().AllDecl()),
		})
	}

	return ret
}

func gatherFieldDeclaration(decls []mantlr.IDeclContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, 0, len(decls))

	for _, v := range decls {
		ret = append(ret, &ast.Declaration{
			Position: constructPosition(v.ID().GetSymbol()),
			Name:     v.ID().GetText(),
			Type:     typeToAst(v.Type_()),
		})
	}

	return ret
}

// === Types ===
func typeToAst(typ mantlr.ITypeContext) ast.Type {
	switch v := typ.(type) {
	case *mantlr.IntTypeContext:
		return &ast.IntType{}
	case *mantlr.BoolTypeContext:
		return &ast.BoolType{}
	case *mantlr.StructTypeContext:
		return &ast.StructType{Id: v.ID().GetText()}
	}

	return nil
}

// === Declarations ===
func declarationsToAst(declarations []mantlr.IDeclarationContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, 0, len(declarations))

	for _, v := range declarations {
		// Handles comma declarations
		for _, vv := range v.AllID() {
			decl := &ast.Declaration{
				Name:     vv.GetText(),
				Type:     typeToAst(v.Type_()),
				Position: constructPosition(vv.GetSymbol()),
			}
			ret = append(ret, decl)
		}
	}

	return ret
}

// === Functions ===
func functionToAst(fn mantlr.IFunctionContext, ch chan *ast.Function) {
	pos := constructPosition(fn.ID().GetSymbol())
	name := fn.ID().GetText()
	params := paramsToAst(fn.Parameters().AllDecl())
	returnType := returnTypeToAst(fn.ReturnType())
	locals := declarationsToAst(fn.Declarations().AllDeclaration())
	body := bodyToAst(fn.StatementList().AllStatement())

	ret := &ast.Function{
		Position:   pos,
		Name:       name,
		Parameters: params,
		ReturnType: returnType,
		Locals:     locals,
		Body:       body,
	}

	ch <- ret
}

func paramsToAst(params []mantlr.IDeclContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, 0, len(params))

	for _, v := range params {
		ret = append(ret, &ast.Declaration{
			Position: constructPosition(v.ID().GetSymbol()),
			Type:     typeToAst(v.Type_()),
			Name:     v.ID().GetText(),
		})
	}

	return ret
}

func returnTypeToAst(retType mantlr.IReturnTypeContext) ast.Type {
	switch v := retType.(type) {
	case *mantlr.ReturnTypeRealContext:
		switch vv := v.Type_().(type) {
		case *mantlr.IntTypeContext:
			return &ast.IntType{}
		case *mantlr.BoolTypeContext:
			return &ast.BoolType{}
		case *mantlr.StructTypeContext:
			return &ast.StructType{Id: vv.ID().GetText()}
		}
	case *mantlr.ReturnTypeVoidContext:
		return &ast.VoidType{}

	}

	return nil
}

func bodyToAst(body []mantlr.IStatementContext) []ast.Statement {
	ret := make([]ast.Statement, 0, len(body))

	for _, v := range body {
		ret = append(ret, statementToAst(v))
	}

	return ret
}

// === Statements ===
func statementToAst(stmt mantlr.IStatementContext) ast.Statement {
	switch v := stmt.(type) {
	case *mantlr.AssignmentContext:
		return assignmentStatementToAst(v)
	case *mantlr.NestedBlockContext:
		return blockStatementToAst(v.Block())
	case *mantlr.PrintContext:
		return printStatementToAst(v)
	case *mantlr.PrintLnContext:
		return printLnStatementToAst(v)
	case *mantlr.ConditionalContext:
		return conditionalStatementToAst(v)
	case *mantlr.WhileContext:
		return whileStatementToAst(v)
	case *mantlr.DeleteContext:
		return deleteStatementToAst(v)
	case *mantlr.ReturnContext:
		return returnStatementToAst(v)
	case *mantlr.InvocationContext:
		return invocationStatementToAst(v)
	}

	return nil
}

func blockStatementToAst(block mantlr.IBlockContext) *ast.BlockStatement {
	pos := constructPosition(block.GetStart())
	blockStmts := block.StatementList().AllStatement()
	stmts := make([]ast.Statement, 0, len(blockStmts))

	for _, v := range blockStmts {
		stmts = append(stmts, statementToAst(v))
	}

	return &ast.BlockStatement{
		Position:   pos,
		Statements: stmts,
	}
}

func assignmentStatementToAst(asgn *mantlr.AssignmentContext) *ast.AssignmentStatement {
	return &ast.AssignmentStatement{
		Position: constructPosition(asgn.GetStart()),
		Target:   lValueToAst(asgn.Lvalue()),
		Source:   expressionToAst(asgn.Expression()),
	}
}

func printStatementToAst(prnt *mantlr.PrintContext) *ast.PrintStatement {
	return &ast.PrintStatement{
		Position:   constructPosition(prnt.GetStart()),
		Expression: expressionToAst(prnt.Expression()),
		Newline:    false,
	}
}

func printLnStatementToAst(prnt *mantlr.PrintLnContext) *ast.PrintStatement {
	return &ast.PrintStatement{
		Position:   constructPosition(prnt.GetStart()),
		Expression: expressionToAst(prnt.Expression()),
		Newline:    true,
	}
}

func conditionalStatementToAst(fi *mantlr.ConditionalContext) *ast.IfStatement {
	if fi.GetElseBlock() != nil {
		return &ast.IfStatement{
			Position: constructPosition(fi.GetStart()),
			Guard:    expressionToAst(fi.Expression()),
			Then:     blockStatementToAst(fi.GetThenBlock()),
			Else:     blockStatementToAst(fi.GetElseBlock()),
		}
	}

	return &ast.IfStatement{
		Position: constructPosition(fi.GetStart()),
		Guard:    expressionToAst(fi.Expression()),
		Then:     blockStatementToAst(fi.GetThenBlock()),
		Else:     nil,
	}
}

func whileStatementToAst(whl *mantlr.WhileContext) *ast.WhileStatement {
	return &ast.WhileStatement{
		Position: constructPosition(whl.GetStart()),
		Guard:    expressionToAst(whl.Expression()),
		Body:     blockStatementToAst(whl.Block()),
	}
}

func deleteStatementToAst(del *mantlr.DeleteContext) *ast.DeleteStatement {
	return &ast.DeleteStatement{
		Position:   constructPosition(del.GetStart()),
		Expression: expressionToAst(del.Expression()),
	}
}

func returnStatementToAst(ret *mantlr.ReturnContext) *ast.ReturnStatement {
	return &ast.ReturnStatement{
		Position:   constructPosition(ret.GetStart()),
		Expression: expressionToAst(ret.Expression()),
	}
}

func invocationStatementToAst(inv *mantlr.InvocationContext) *ast.InvocationStatement {
	pos := constructPosition(inv.ID().GetSymbol())
	name := inv.ID().GetText()
	args := make([]ast.Expression, 0, len(inv.Arguments().AllExpression()))

	for _, v := range inv.Arguments().AllExpression() {
		args = append(args, expressionToAst(v))
	}

	return &ast.InvocationStatement{
		Expression: ast.InvocationExpression{
			Position:  pos,
			Name:      name,
			Arguments: args,
		},
	}
}

// === LValues ===
func lValueToAst(lval mantlr.ILvalueContext) ast.LValue {
	switch v := lval.(type) {
	// Base case
	case *mantlr.LvalueIdContext:
		return &ast.NameLValue{
			Position: constructPosition(v.GetStart()),
			Name:     v.GetText(),
		}
	case *mantlr.LvalueDotContext:
		return &ast.DotLValue{
			Position: constructPosition(v.GetStart()),
			Name:     v.ID().GetText(),
			Left:     lValueToAst(v.Lvalue()),
		}
	}

	return nil
}

// === Expressions ===
func expressionToAst(expr mantlr.IExpressionContext) ast.Expression {
	switch v := expr.(type) {
	case *mantlr.InvocationExprContext:
		return invocationExpressionToAst(v)
	case *mantlr.DotExprContext:
		return dotExpressionToAst(v)
	case *mantlr.UnaryExprContext:
		return unaryExpressionToAst(v)
	case *mantlr.BinaryExprContext:
		return binaryExpressionToAst(v)
	case *mantlr.IdentifierExprContext:
		return identifierExpressionToAst(v)
	case *mantlr.IntegerExprContext:
		return integerExpressionToAst(v)
	case *mantlr.TrueExprContext:
		return trueExpressionToAst(v)
	case *mantlr.FalseExprContext:
		return falseExpressionToAst(v)
	case *mantlr.NewExprContext:
		return newExpressionToAst(v)
	case *mantlr.NullExprContext:
		return nullExpressionToAst(v)
	case *mantlr.NestedExprContext:
		return nestedExpressionToAst(v)
	}

	return nil
}

func invocationExpressionToAst(inv *mantlr.InvocationExprContext) *ast.InvocationExpression {
	return &ast.InvocationExpression{
		Position:  constructPosition(inv.ID().GetSymbol()),
		Name:      inv.ID().GetText(),
		Arguments: argumentsToAst(inv.Arguments().AllExpression()),
	}
}

func dotExpressionToAst(dot *mantlr.DotExprContext) *ast.DotExpression {
	return &ast.DotExpression{
		Position: constructPosition(dot.ID().GetSymbol()),
		Left:     expressionToAst(dot.Expression()),
		Field:    dot.ID().GetText(),
	}
}

func unaryExpressionToAst(unary *mantlr.UnaryExprContext) *ast.UnaryExpression {
	return &ast.UnaryExpression{
		Position: constructPosition(unary.GetStart()),
		Operator: ast.Operator(unary.GetOp().GetText()),
		Operand:  expressionToAst(unary.Expression()),
	}
}

func binaryExpressionToAst(bin *mantlr.BinaryExprContext) *ast.BinaryExpression {
	return &ast.BinaryExpression{
		Position: constructPosition(bin.GetOp()),
		Left:     expressionToAst(bin.GetLft()),
		Operator: ast.Operator(bin.GetOp().GetText()),
		Right:    expressionToAst(bin.GetRht()),
	}
}

func identifierExpressionToAst(ident *mantlr.IdentifierExprContext) *ast.IdentifierExpression {
	return &ast.IdentifierExpression{
		Position: constructPosition(ident.GetStart()),
		Name:     ident.ID().GetText(),
	}
}

func integerExpressionToAst(intgr *mantlr.IntegerExprContext) *ast.IntExpression {
	return &ast.IntExpression{
		Position: constructPosition(intgr.GetStart()),
		Value:    intgr.INTEGER().GetText(),
	}
}

func trueExpressionToAst(tr *mantlr.TrueExprContext) *ast.TrueExpression {
	return &ast.TrueExpression{
		Position: constructPosition(tr.GetStart()),
	}
}

func falseExpressionToAst(fls *mantlr.FalseExprContext) *ast.FalseExpression {
	return &ast.FalseExpression{
		Position: constructPosition(fls.GetStart()),
	}
}

func newExpressionToAst(nw *mantlr.NewExprContext) *ast.NewExpression {
	return &ast.NewExpression{
		Position: constructPosition(nw.GetStart()),
		Id:       nw.ID().GetText(),
	}
}

func nullExpressionToAst(null *mantlr.NullExprContext) *ast.NullExpression {
	return &ast.NullExpression{
		Position: constructPosition(null.GetStart()),
	}
}

func nestedExpressionToAst(nest *mantlr.NestedExprContext) ast.Expression {
	return expressionToAst(nest.Expression())
}

func argumentsToAst(args []mantlr.IExpressionContext) []ast.Expression {
	ret := make([]ast.Expression, 0, len(args))

	for _, v := range args {
		ret = append(ret, expressionToAst(v))
	}

	return ret
}

// === Position ===
func constructPosition(tok antlr.Token) *ast.Position {
	return &ast.Position{
		Line:   tok.GetLine(),
		Column: tok.GetColumn() + 1,
	}
}
