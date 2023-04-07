// Code generated from Mini.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Mini
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseMiniListener is a complete listener for a parse tree produced by MiniParser.
type BaseMiniListener struct{}

var _ MiniListener = &BaseMiniListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMiniListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMiniListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMiniListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMiniListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMiniListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMiniListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseMiniListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseMiniListener) ExitTypes(ctx *TypesContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseMiniListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseMiniListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterNestedDecl is called when production nestedDecl is entered.
func (s *BaseMiniListener) EnterNestedDecl(ctx *NestedDeclContext) {}

// ExitNestedDecl is called when production nestedDecl is exited.
func (s *BaseMiniListener) ExitNestedDecl(ctx *NestedDeclContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseMiniListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseMiniListener) ExitDecl(ctx *DeclContext) {}

// EnterIntType is called when production IntType is entered.
func (s *BaseMiniListener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production IntType is exited.
func (s *BaseMiniListener) ExitIntType(ctx *IntTypeContext) {}

// EnterBoolType is called when production BoolType is entered.
func (s *BaseMiniListener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production BoolType is exited.
func (s *BaseMiniListener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterStructType is called when production StructType is entered.
func (s *BaseMiniListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production StructType is exited.
func (s *BaseMiniListener) ExitStructType(ctx *StructTypeContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseMiniListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseMiniListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseMiniListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseMiniListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseMiniListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseMiniListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseMiniListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseMiniListener) ExitFunction(ctx *FunctionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseMiniListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseMiniListener) ExitParameters(ctx *ParametersContext) {}

// EnterReturnTypeReal is called when production ReturnTypeReal is entered.
func (s *BaseMiniListener) EnterReturnTypeReal(ctx *ReturnTypeRealContext) {}

// ExitReturnTypeReal is called when production ReturnTypeReal is exited.
func (s *BaseMiniListener) ExitReturnTypeReal(ctx *ReturnTypeRealContext) {}

// EnterReturnTypeVoid is called when production ReturnTypeVoid is entered.
func (s *BaseMiniListener) EnterReturnTypeVoid(ctx *ReturnTypeVoidContext) {}

// ExitReturnTypeVoid is called when production ReturnTypeVoid is exited.
func (s *BaseMiniListener) ExitReturnTypeVoid(ctx *ReturnTypeVoidContext) {}

// EnterNestedBlock is called when production NestedBlock is entered.
func (s *BaseMiniListener) EnterNestedBlock(ctx *NestedBlockContext) {}

// ExitNestedBlock is called when production NestedBlock is exited.
func (s *BaseMiniListener) ExitNestedBlock(ctx *NestedBlockContext) {}

// EnterAssignment is called when production Assignment is entered.
func (s *BaseMiniListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production Assignment is exited.
func (s *BaseMiniListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrint is called when production Print is entered.
func (s *BaseMiniListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BaseMiniListener) ExitPrint(ctx *PrintContext) {}

// EnterPrintLn is called when production PrintLn is entered.
func (s *BaseMiniListener) EnterPrintLn(ctx *PrintLnContext) {}

// ExitPrintLn is called when production PrintLn is exited.
func (s *BaseMiniListener) ExitPrintLn(ctx *PrintLnContext) {}

// EnterConditional is called when production Conditional is entered.
func (s *BaseMiniListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production Conditional is exited.
func (s *BaseMiniListener) ExitConditional(ctx *ConditionalContext) {}

// EnterWhile is called when production While is entered.
func (s *BaseMiniListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production While is exited.
func (s *BaseMiniListener) ExitWhile(ctx *WhileContext) {}

// EnterDelete is called when production Delete is entered.
func (s *BaseMiniListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production Delete is exited.
func (s *BaseMiniListener) ExitDelete(ctx *DeleteContext) {}

// EnterReturn is called when production Return is entered.
func (s *BaseMiniListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production Return is exited.
func (s *BaseMiniListener) ExitReturn(ctx *ReturnContext) {}

// EnterInvocation is called when production Invocation is entered.
func (s *BaseMiniListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production Invocation is exited.
func (s *BaseMiniListener) ExitInvocation(ctx *InvocationContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMiniListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMiniListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseMiniListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseMiniListener) ExitStatementList(ctx *StatementListContext) {}

// EnterLvalueId is called when production LvalueId is entered.
func (s *BaseMiniListener) EnterLvalueId(ctx *LvalueIdContext) {}

// ExitLvalueId is called when production LvalueId is exited.
func (s *BaseMiniListener) ExitLvalueId(ctx *LvalueIdContext) {}

// EnterLvalueDot is called when production LvalueDot is entered.
func (s *BaseMiniListener) EnterLvalueDot(ctx *LvalueDotContext) {}

// ExitLvalueDot is called when production LvalueDot is exited.
func (s *BaseMiniListener) ExitLvalueDot(ctx *LvalueDotContext) {}

// EnterIntegerExpr is called when production IntegerExpr is entered.
func (s *BaseMiniListener) EnterIntegerExpr(ctx *IntegerExprContext) {}

// ExitIntegerExpr is called when production IntegerExpr is exited.
func (s *BaseMiniListener) ExitIntegerExpr(ctx *IntegerExprContext) {}

// EnterTrueExpr is called when production TrueExpr is entered.
func (s *BaseMiniListener) EnterTrueExpr(ctx *TrueExprContext) {}

// ExitTrueExpr is called when production TrueExpr is exited.
func (s *BaseMiniListener) ExitTrueExpr(ctx *TrueExprContext) {}

// EnterIdentifierExpr is called when production IdentifierExpr is entered.
func (s *BaseMiniListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production IdentifierExpr is exited.
func (s *BaseMiniListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterBinaryExpr is called when production BinaryExpr is entered.
func (s *BaseMiniListener) EnterBinaryExpr(ctx *BinaryExprContext) {}

// ExitBinaryExpr is called when production BinaryExpr is exited.
func (s *BaseMiniListener) ExitBinaryExpr(ctx *BinaryExprContext) {}

// EnterNewExpr is called when production NewExpr is entered.
func (s *BaseMiniListener) EnterNewExpr(ctx *NewExprContext) {}

// ExitNewExpr is called when production NewExpr is exited.
func (s *BaseMiniListener) ExitNewExpr(ctx *NewExprContext) {}

// EnterNestedExpr is called when production NestedExpr is entered.
func (s *BaseMiniListener) EnterNestedExpr(ctx *NestedExprContext) {}

// ExitNestedExpr is called when production NestedExpr is exited.
func (s *BaseMiniListener) ExitNestedExpr(ctx *NestedExprContext) {}

// EnterDotExpr is called when production DotExpr is entered.
func (s *BaseMiniListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production DotExpr is exited.
func (s *BaseMiniListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterUnaryExpr is called when production UnaryExpr is entered.
func (s *BaseMiniListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production UnaryExpr is exited.
func (s *BaseMiniListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterInvocationExpr is called when production InvocationExpr is entered.
func (s *BaseMiniListener) EnterInvocationExpr(ctx *InvocationExprContext) {}

// ExitInvocationExpr is called when production InvocationExpr is exited.
func (s *BaseMiniListener) ExitInvocationExpr(ctx *InvocationExprContext) {}

// EnterFalseExpr is called when production FalseExpr is entered.
func (s *BaseMiniListener) EnterFalseExpr(ctx *FalseExprContext) {}

// ExitFalseExpr is called when production FalseExpr is exited.
func (s *BaseMiniListener) ExitFalseExpr(ctx *FalseExprContext) {}

// EnterNullExpr is called when production NullExpr is entered.
func (s *BaseMiniListener) EnterNullExpr(ctx *NullExprContext) {}

// ExitNullExpr is called when production NullExpr is exited.
func (s *BaseMiniListener) ExitNullExpr(ctx *NullExprContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseMiniListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseMiniListener) ExitArguments(ctx *ArgumentsContext) {}
