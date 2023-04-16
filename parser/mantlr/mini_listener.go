// Code generated from Mini.g4 by ANTLR 4.12.0. DO NOT EDIT.

package mantlr // Mini
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// MiniListener is a complete listener for a parse tree produced by MiniParser.
type MiniListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterNestedDecl is called when entering the nestedDecl production.
	EnterNestedDecl(c *NestedDeclContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterIntType is called when entering the IntType production.
	EnterIntType(c *IntTypeContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterStructType is called when entering the StructType production.
	EnterStructType(c *StructTypeContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterReturnTypeReal is called when entering the ReturnTypeReal production.
	EnterReturnTypeReal(c *ReturnTypeRealContext)

	// EnterReturnTypeVoid is called when entering the ReturnTypeVoid production.
	EnterReturnTypeVoid(c *ReturnTypeVoidContext)

	// EnterNestedBlock is called when entering the NestedBlock production.
	EnterNestedBlock(c *NestedBlockContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterPrintLn is called when entering the PrintLn production.
	EnterPrintLn(c *PrintLnContext)

	// EnterConditional is called when entering the Conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterWhile is called when entering the While production.
	EnterWhile(c *WhileContext)

	// EnterDelete is called when entering the Delete production.
	EnterDelete(c *DeleteContext)

	// EnterReturn is called when entering the Return production.
	EnterReturn(c *ReturnContext)

	// EnterInvocation is called when entering the Invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterLvalueId is called when entering the LvalueId production.
	EnterLvalueId(c *LvalueIdContext)

	// EnterLvalueDot is called when entering the LvalueDot production.
	EnterLvalueDot(c *LvalueDotContext)

	// EnterIntegerExpr is called when entering the IntegerExpr production.
	EnterIntegerExpr(c *IntegerExprContext)

	// EnterTrueExpr is called when entering the TrueExpr production.
	EnterTrueExpr(c *TrueExprContext)

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterBinaryExpr is called when entering the BinaryExpr production.
	EnterBinaryExpr(c *BinaryExprContext)

	// EnterNewExpr is called when entering the NewExpr production.
	EnterNewExpr(c *NewExprContext)

	// EnterNestedExpr is called when entering the NestedExpr production.
	EnterNestedExpr(c *NestedExprContext)

	// EnterDotExpr is called when entering the DotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterInvocationExpr is called when entering the InvocationExpr production.
	EnterInvocationExpr(c *InvocationExprContext)

	// EnterFalseExpr is called when entering the FalseExpr production.
	EnterFalseExpr(c *FalseExprContext)

	// EnterNullExpr is called when entering the NullExpr production.
	EnterNullExpr(c *NullExprContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitNestedDecl is called when exiting the nestedDecl production.
	ExitNestedDecl(c *NestedDeclContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitIntType is called when exiting the IntType production.
	ExitIntType(c *IntTypeContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitStructType is called when exiting the StructType production.
	ExitStructType(c *StructTypeContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitReturnTypeReal is called when exiting the ReturnTypeReal production.
	ExitReturnTypeReal(c *ReturnTypeRealContext)

	// ExitReturnTypeVoid is called when exiting the ReturnTypeVoid production.
	ExitReturnTypeVoid(c *ReturnTypeVoidContext)

	// ExitNestedBlock is called when exiting the NestedBlock production.
	ExitNestedBlock(c *NestedBlockContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitPrintLn is called when exiting the PrintLn production.
	ExitPrintLn(c *PrintLnContext)

	// ExitConditional is called when exiting the Conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitWhile is called when exiting the While production.
	ExitWhile(c *WhileContext)

	// ExitDelete is called when exiting the Delete production.
	ExitDelete(c *DeleteContext)

	// ExitReturn is called when exiting the Return production.
	ExitReturn(c *ReturnContext)

	// ExitInvocation is called when exiting the Invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitLvalueId is called when exiting the LvalueId production.
	ExitLvalueId(c *LvalueIdContext)

	// ExitLvalueDot is called when exiting the LvalueDot production.
	ExitLvalueDot(c *LvalueDotContext)

	// ExitIntegerExpr is called when exiting the IntegerExpr production.
	ExitIntegerExpr(c *IntegerExprContext)

	// ExitTrueExpr is called when exiting the TrueExpr production.
	ExitTrueExpr(c *TrueExprContext)

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitBinaryExpr is called when exiting the BinaryExpr production.
	ExitBinaryExpr(c *BinaryExprContext)

	// ExitNewExpr is called when exiting the NewExpr production.
	ExitNewExpr(c *NewExprContext)

	// ExitNestedExpr is called when exiting the NestedExpr production.
	ExitNestedExpr(c *NestedExprContext)

	// ExitDotExpr is called when exiting the DotExpr production.
	ExitDotExpr(c *DotExprContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitInvocationExpr is called when exiting the InvocationExpr production.
	ExitInvocationExpr(c *InvocationExprContext)

	// ExitFalseExpr is called when exiting the FalseExpr production.
	ExitFalseExpr(c *FalseExprContext)

	// ExitNullExpr is called when exiting the NullExpr production.
	ExitNullExpr(c *NullExprContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)
}
