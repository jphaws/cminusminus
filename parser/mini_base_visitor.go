// Code generated from Mini.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Mini
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseMiniVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMiniVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitTypes(ctx *TypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitNestedDecl(ctx *NestedDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitDeclarations(ctx *DeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitFunctions(ctx *FunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitReturnTypeReal(ctx *ReturnTypeRealContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitReturnTypeVoid(ctx *ReturnTypeVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitNestedBlock(ctx *NestedBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitPrintLn(ctx *PrintLnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitConditional(ctx *ConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitReturn(ctx *ReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitInvocation(ctx *InvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitLvalueId(ctx *LvalueIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitLvalueDot(ctx *LvalueDotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitIntegerExpr(ctx *IntegerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitTrueExpr(ctx *TrueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitIdentifierExpr(ctx *IdentifierExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitBinaryExpr(ctx *BinaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitNewExpr(ctx *NewExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitNestedExpr(ctx *NestedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitDotExpr(ctx *DotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitInvocationExpr(ctx *InvocationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitFalseExpr(ctx *FalseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitNullExpr(ctx *NullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
