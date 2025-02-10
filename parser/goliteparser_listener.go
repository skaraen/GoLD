// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// GoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type GoliteParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterFieldsPrime is called when entering the fieldsPrime production.
	EnterFieldsPrime(c *FieldsPrimeContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterIds is called when entering the ids production.
	EnterIds(c *IdsContext)

	// EnterIdsPrime is called when entering the idsPrime production.
	EnterIdsPrime(c *IdsPrimeContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParamsPrime is called when entering the paramsPrime production.
	EnterParamsPrime(c *ParamsPrimeContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterArgsPrime is called when entering the argsPrime production.
	EnterArgsPrime(c *ArgsPrimeContext)

	// EnterLValue is called when entering the lValue production.
	EnterLValue(c *LValueContext)

	// EnterLValuePrime is called when entering the lValuePrime production.
	EnterLValuePrime(c *LValuePrimeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExprPrime is called when entering the exprPrime production.
	EnterExprPrime(c *ExprPrimeContext)

	// EnterBoolTerm is called when entering the boolTerm production.
	EnterBoolTerm(c *BoolTermContext)

	// EnterBoolTermPrime is called when entering the boolTermPrime production.
	EnterBoolTermPrime(c *BoolTermPrimeContext)

	// EnterEqualTermPrime is called when entering the equalTermPrime production.
	EnterEqualTermPrime(c *EqualTermPrimeContext)

	// EnterEqualTerm is called when entering the equalTerm production.
	EnterEqualTerm(c *EqualTermContext)

	// EnterRelTermPrime is called when entering the relTermPrime production.
	EnterRelTermPrime(c *RelTermPrimeContext)

	// EnterRelTerm is called when entering the relTerm production.
	EnterRelTerm(c *RelTermContext)

	// EnterSimpleTermPrime is called when entering the simpleTermPrime production.
	EnterSimpleTermPrime(c *SimpleTermPrimeContext)

	// EnterSimpleTerm is called when entering the simpleTerm production.
	EnterSimpleTerm(c *SimpleTermContext)

	// EnterTermPrime is called when entering the termPrime production.
	EnterTermPrime(c *TermPrimeContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterUnaryTerm is called when entering the unaryTerm production.
	EnterUnaryTerm(c *UnaryTermContext)

	// EnterSelectorTerm is called when entering the selectorTerm production.
	EnterSelectorTerm(c *SelectorTermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitFieldsPrime is called when exiting the fieldsPrime production.
	ExitFieldsPrime(c *FieldsPrimeContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitIds is called when exiting the ids production.
	ExitIds(c *IdsContext)

	// ExitIdsPrime is called when exiting the idsPrime production.
	ExitIdsPrime(c *IdsPrimeContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParamsPrime is called when exiting the paramsPrime production.
	ExitParamsPrime(c *ParamsPrimeContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitArgsPrime is called when exiting the argsPrime production.
	ExitArgsPrime(c *ArgsPrimeContext)

	// ExitLValue is called when exiting the lValue production.
	ExitLValue(c *LValueContext)

	// ExitLValuePrime is called when exiting the lValuePrime production.
	ExitLValuePrime(c *LValuePrimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExprPrime is called when exiting the exprPrime production.
	ExitExprPrime(c *ExprPrimeContext)

	// ExitBoolTerm is called when exiting the boolTerm production.
	ExitBoolTerm(c *BoolTermContext)

	// ExitBoolTermPrime is called when exiting the boolTermPrime production.
	ExitBoolTermPrime(c *BoolTermPrimeContext)

	// ExitEqualTermPrime is called when exiting the equalTermPrime production.
	ExitEqualTermPrime(c *EqualTermPrimeContext)

	// ExitEqualTerm is called when exiting the equalTerm production.
	ExitEqualTerm(c *EqualTermContext)

	// ExitRelTermPrime is called when exiting the relTermPrime production.
	ExitRelTermPrime(c *RelTermPrimeContext)

	// ExitRelTerm is called when exiting the relTerm production.
	ExitRelTerm(c *RelTermContext)

	// ExitSimpleTermPrime is called when exiting the simpleTermPrime production.
	ExitSimpleTermPrime(c *SimpleTermPrimeContext)

	// ExitSimpleTerm is called when exiting the simpleTerm production.
	ExitSimpleTerm(c *SimpleTermContext)

	// ExitTermPrime is called when exiting the termPrime production.
	ExitTermPrime(c *TermPrimeContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitUnaryTerm is called when exiting the unaryTerm production.
	ExitUnaryTerm(c *UnaryTermContext)

	// ExitSelectorTerm is called when exiting the selectorTerm production.
	ExitSelectorTerm(c *SelectorTermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
