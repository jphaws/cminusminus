package parser // Mini

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/keep-cp/compiler-project-c/ast"
)

// MiniToAst is the entry function for creating AST from ANTLR parse tree
func MiniToAst(ctx IProgramContext) ast.Root {
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

	return ast.Root{
		Types:        types,
		Declarations: declarations,
		Functions:    functions,
	}
}

func constructPosition(tok antlr.Token) *ast.Position {
	return &ast.Position{
		Line:   tok.GetLine(),
		Column: tok.GetColumn(),
	}
}

func typeDeclarationsToAst(typeDecls []ITypeDeclarationContext) []*ast.TypeDeclaration {
	ret := make([]*ast.TypeDeclaration, len(typeDecls))

	for i, v := range typeDecls {
		ret[i] = &ast.TypeDeclaration{
			Id:       v.ID().GetText(),
			Fields:   gatherFieldDeclaration(v.NestedDecl().AllDecl()),
			Position: constructPosition(v.ID().GetSymbol()),
		}
	}

	return ret
}

func gatherFieldDeclaration(decls []IDeclContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, len(decls))

	for i, v := range decls {
		ret[i] = &ast.Declaration{
			Name: v.ID().GetText(),
			Type: typeToAst(v.Type_()),
		}
	}

	return ret
}

func typeToAst(typ ITypeContext) ast.Type {
	switch v := typ.(type) {
	case *IntTypeContext:
		return ast.IntType{}
	case *BoolTypeContext:
		return ast.BoolType{}
	case *StructTypeContext:
		return ast.StructType{Id: v.ID().GetText()}
	}

	return nil
}

func declarationsToAst(declarations []IDeclarationContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, 0, len(declarations))

	for _, v := range declarations {
		// Handles comma declarations
		for _, vv := range v.AllID() {
			decl := ast.Declaration{
				Name:     vv.GetText(),
				Type:     typeToAst(v.Type_()),
				Position: constructPosition(vv.GetSymbol()),
			}
			ret = append(ret, &decl)
		}
	}

	return ret
}

func functionToAst(f IFunctionContext, ch chan *ast.Function) {
	name := f.ID().GetText()
	params := paramsToAst(f.Parameters().AllDecl())
	locals := declarationsToAst(f.Declarations().AllDeclaration())
	body := bodyToAst(f.StatementList().AllStatement())
	returnT := returnTypeToAst(f.ReturnType())
	pos := constructPosition(f.GetStart())

	ret := ast.Function{
		Name:       name,
		Parameters: params,
		ReturnType: returnT,
		Locals:     locals,
		Body:       body,
		Position:   pos,
	}

	ch <- &ret
}

func paramsToAst(params []IDeclContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, len(params))

	for i, v := range params {
		ret[i] = &ast.Declaration{
			Type:     typeToAst(v.Type_()),
			Name:     v.ID().GetText(),
			Position: constructPosition(v.GetStart()),
		}
	}

	return ret
}

func returnTypeToAst(retType IReturnTypeContext) ast.Type {
	switch v := retType.(type) {
	case *ReturnTypeRealContext:
		switch vv := v.Type_().(type) {
		case *IntTypeContext:
			return ast.IntType{}
		case *BoolTypeContext:
			return ast.BoolType{}
		case *StructTypeContext:
			return ast.StructType{Id: vv.ID().GetText()}
		}
	case *ReturnTypeVoidContext:
		return ast.VoidType{}

	}

	return nil
}

func bodyToAst(body []IStatementContext) []ast.Statement {
	ret := make([]ast.Statement, len(body))

	for i, v := range body {
		ret[i] = statementToAst(v)
	}

	return ret
}

func statementToAst(stmt IStatementContext) ast.Statement {
	switch v := stmt.(type) {
	case *AssignmentContext:
		return assignmentStatementToAst(v)
	case *NestedBlockContext:
		return blockStatementToAst(v.Block())
	case *PrintContext:
		return printStatementToAst(v)
	case *PrintLnContext:
		return printLnStatementToAst(v)
	case *ConditionalContext:
		return conditionalStatementToAst(v)
	case *WhileContext:
		return whileStatementToAst(v)
	case *DeleteContext:
		return deleteStatementToAst(v)
	case *ReturnContext:
		return returnStatementToAst(v)
	case *InvocationContext:
		return invocationStatementToAst(v)
	}

	return nil
}

func blockStatementToAst(block IBlockContext) *ast.BlockStatement {
	blockStmts := block.StatementList().AllStatement()
	stmts := make([]ast.Statement, len(blockStmts))
	pos := constructPosition(block.GetStart())

	for i, v := range blockStmts {
		stmts[i] = statementToAst(v)
	}

	return &ast.BlockStatement{
		Position:   pos,
		Statements: stmts,
	}
}

func invocationStatementToAst(inv *InvocationContext) *ast.InvocationStatement {
	name := inv.ID().GetText()
	args := make([]ast.Expression, len(inv.Arguments().AllExpression()))
	pos := constructPosition(inv.GetStart())

	for i, v := range inv.Arguments().AllExpression() {
		args[i] = expressionToAst(v)
	}

	return &ast.InvocationStatement{
		Expression: ast.InvocationExpression{
			Position:  pos,
			Name:      name,
			Arguments: args,
		},
	}
}

func conditionalStatementToAst(fi *ConditionalContext) *ast.IfStatement {
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

func whileStatementToAst(whl *WhileContext) *ast.WhileStatement {
	return &ast.WhileStatement{
		Position: constructPosition(whl.GetStart()),
		Guard:    expressionToAst(whl.Expression()),
		Body:     blockStatementToAst(whl.Block()),
	}
}

func returnStatementToAst(ret *ReturnContext) *ast.ReturnStatement {
	return &ast.ReturnStatement{
		Position:   constructPosition(ret.GetStart()),
		Expression: expressionToAst(ret.Expression()),
	}
}

func deleteStatementToAst(del *DeleteContext) *ast.DeleteStatement {
	return &ast.DeleteStatement{
		Position:   constructPosition(del.GetStart()),
		Expression: expressionToAst(del.Expression()),
	}
}

func printStatementToAst(prnt *PrintContext) *ast.PrintStatement {
	return &ast.PrintStatement{
		Position:   constructPosition(prnt.GetStart()),
		Expression: expressionToAst(prnt.Expression()),
		Newline:    false,
	}
}

func printLnStatementToAst(prnt *PrintLnContext) *ast.PrintStatement {
	return &ast.PrintStatement{
		Position:   constructPosition(prnt.GetStart()),
		Expression: expressionToAst(prnt.Expression()),
		Newline:    true,
	}
}

func assignmentStatementToAst(asgn *AssignmentContext) *ast.AssignmentStatement {
	if asgn.Expression() == nil {
		return &ast.AssignmentStatement{
			Position: constructPosition(asgn.GetStart()),
			Target:   lValueToAst(asgn.Lvalue()),
			Source: ast.ReadExpression{
				Position: constructPosition(asgn.GetStart()),
			},
		}
	}

	return &ast.AssignmentStatement{
		Position: constructPosition(asgn.GetStart()),
		Target:   lValueToAst(asgn.Lvalue()),
		Source:   expressionToAst(asgn.Expression()),
	}
}

func lValueToAst(lval ILvalueContext) ast.LValue {
	switch v := lval.(type) {
	// Base case
	case *LvalueIdContext:
		return &ast.NameLValue{
			Position: constructPosition(v.GetStart()),
			Name:     v.GetText(),
		}
	case *LvalueDotContext:
		return &ast.DotLValue{
			Position: constructPosition(v.GetStart()),
			Name:     v.ID().GetText(),
			Left:     lValueToAst(v.Lvalue()),
		}
	}

	return nil
}

// TODO
func expressionToAst(expr IExpressionContext) ast.Expression {
	return nil
}
