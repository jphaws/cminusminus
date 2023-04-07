// Code generated from Mini.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Mini
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

/* package declaration here */

// A complete Visitor for a parse tree produced by MiniParser.
type MiniVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MiniParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MiniParser#types.
	VisitTypes(ctx *TypesContext) interface{}

	// Visit a parse tree produced by MiniParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by MiniParser#nestedDecl.
	VisitNestedDecl(ctx *NestedDeclContext) interface{}

	// Visit a parse tree produced by MiniParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by MiniParser#IntType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by MiniParser#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by MiniParser#StructType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by MiniParser#declarations.
	VisitDeclarations(ctx *DeclarationsContext) interface{}

	// Visit a parse tree produced by MiniParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by MiniParser#functions.
	VisitFunctions(ctx *FunctionsContext) interface{}

	// Visit a parse tree produced by MiniParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by MiniParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by MiniParser#ReturnTypeReal.
	VisitReturnTypeReal(ctx *ReturnTypeRealContext) interface{}

	// Visit a parse tree produced by MiniParser#ReturnTypeVoid.
	VisitReturnTypeVoid(ctx *ReturnTypeVoidContext) interface{}

	// Visit a parse tree produced by MiniParser#NestedBlock.
	VisitNestedBlock(ctx *NestedBlockContext) interface{}

	// Visit a parse tree produced by MiniParser#Assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by MiniParser#Print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by MiniParser#PrintLn.
	VisitPrintLn(ctx *PrintLnContext) interface{}

	// Visit a parse tree produced by MiniParser#Conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

	// Visit a parse tree produced by MiniParser#While.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by MiniParser#Delete.
	VisitDelete(ctx *DeleteContext) interface{}

	// Visit a parse tree produced by MiniParser#Return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by MiniParser#Invocation.
	VisitInvocation(ctx *InvocationContext) interface{}

	// Visit a parse tree produced by MiniParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MiniParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by MiniParser#LvalueId.
	VisitLvalueId(ctx *LvalueIdContext) interface{}

	// Visit a parse tree produced by MiniParser#LvalueDot.
	VisitLvalueDot(ctx *LvalueDotContext) interface{}

	// Visit a parse tree produced by MiniParser#IntegerExpr.
	VisitIntegerExpr(ctx *IntegerExprContext) interface{}

	// Visit a parse tree produced by MiniParser#TrueExpr.
	VisitTrueExpr(ctx *TrueExprContext) interface{}

	// Visit a parse tree produced by MiniParser#IdentifierExpr.
	VisitIdentifierExpr(ctx *IdentifierExprContext) interface{}

	// Visit a parse tree produced by MiniParser#BinaryExpr.
	VisitBinaryExpr(ctx *BinaryExprContext) interface{}

	// Visit a parse tree produced by MiniParser#NewExpr.
	VisitNewExpr(ctx *NewExprContext) interface{}

	// Visit a parse tree produced by MiniParser#NestedExpr.
	VisitNestedExpr(ctx *NestedExprContext) interface{}

	// Visit a parse tree produced by MiniParser#DotExpr.
	VisitDotExpr(ctx *DotExprContext) interface{}

	// Visit a parse tree produced by MiniParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by MiniParser#InvocationExpr.
	VisitInvocationExpr(ctx *InvocationExprContext) interface{}

	// Visit a parse tree produced by MiniParser#FalseExpr.
	VisitFalseExpr(ctx *FalseExprContext) interface{}

	// Visit a parse tree produced by MiniParser#NullExpr.
	VisitNullExpr(ctx *NullExprContext) interface{}

	// Visit a parse tree produced by MiniParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}
}
