// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseGoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type BaseGoliteParserListener struct{}

var _ GoliteParserListener = &BaseGoliteParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoliteParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoliteParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoliteParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoliteParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoliteParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoliteParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseGoliteParserListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseGoliteParserListener) ExitTypes(ctx *TypesContext) {}

// EnterTypedecl is called when production typedecl is entered.
func (s *BaseGoliteParserListener) EnterTypedecl(ctx *TypedeclContext) {}

// ExitTypedecl is called when production typedecl is exited.
func (s *BaseGoliteParserListener) ExitTypedecl(ctx *TypedeclContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseGoliteParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseGoliteParserListener) ExitFields(ctx *FieldsContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseGoliteParserListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseGoliteParserListener) ExitDecl(ctx *DeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoliteParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoliteParserListener) ExitType(ctx *TypeContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseGoliteParserListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseGoliteParserListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseGoliteParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseGoliteParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterIds is called when production ids is entered.
func (s *BaseGoliteParserListener) EnterIds(ctx *IdsContext) {}

// ExitIds is called when production ids is exited.
func (s *BaseGoliteParserListener) ExitIds(ctx *IdsContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseGoliteParserListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseGoliteParserListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseGoliteParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseGoliteParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGoliteParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGoliteParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterReturntype is called when production returntype is entered.
func (s *BaseGoliteParserListener) EnterReturntype(ctx *ReturntypeContext) {}

// ExitReturntype is called when production returntype is exited.
func (s *BaseGoliteParserListener) ExitReturntype(ctx *ReturntypeContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseGoliteParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseGoliteParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoliteParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoliteParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoliteParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoliteParserListener) ExitBlock(ctx *BlockContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseGoliteParserListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseGoliteParserListener) ExitDelete(ctx *DeleteContext) {}

// EnterRead is called when production read is entered.
func (s *BaseGoliteParserListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseGoliteParserListener) ExitRead(ctx *ReadContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGoliteParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGoliteParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseGoliteParserListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseGoliteParserListener) ExitPrint(ctx *PrintContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseGoliteParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseGoliteParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseGoliteParserListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseGoliteParserListener) ExitLoop(ctx *LoopContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseGoliteParserListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseGoliteParserListener) ExitReturn(ctx *ReturnContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BaseGoliteParserListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BaseGoliteParserListener) ExitInvocation(ctx *InvocationContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGoliteParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGoliteParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseGoliteParserListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseGoliteParserListener) ExitLvalue(ctx *LvalueContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoliteParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoliteParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBoolterm is called when production boolterm is entered.
func (s *BaseGoliteParserListener) EnterBoolterm(ctx *BooltermContext) {}

// ExitBoolterm is called when production boolterm is exited.
func (s *BaseGoliteParserListener) ExitBoolterm(ctx *BooltermContext) {}

// EnterEqualterm is called when production equalterm is entered.
func (s *BaseGoliteParserListener) EnterEqualterm(ctx *EqualtermContext) {}

// ExitEqualterm is called when production equalterm is exited.
func (s *BaseGoliteParserListener) ExitEqualterm(ctx *EqualtermContext) {}

// EnterRelationterm is called when production relationterm is entered.
func (s *BaseGoliteParserListener) EnterRelationterm(ctx *RelationtermContext) {}

// ExitRelationterm is called when production relationterm is exited.
func (s *BaseGoliteParserListener) ExitRelationterm(ctx *RelationtermContext) {}

// EnterSimpleterm is called when production simpleterm is entered.
func (s *BaseGoliteParserListener) EnterSimpleterm(ctx *SimpletermContext) {}

// ExitSimpleterm is called when production simpleterm is exited.
func (s *BaseGoliteParserListener) ExitSimpleterm(ctx *SimpletermContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseGoliteParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseGoliteParserListener) ExitTerm(ctx *TermContext) {}

// EnterUnaryterm is called when production unaryterm is entered.
func (s *BaseGoliteParserListener) EnterUnaryterm(ctx *UnarytermContext) {}

// ExitUnaryterm is called when production unaryterm is exited.
func (s *BaseGoliteParserListener) ExitUnaryterm(ctx *UnarytermContext) {}

// EnterSelectorterm is called when production selectorterm is entered.
func (s *BaseGoliteParserListener) EnterSelectorterm(ctx *SelectortermContext) {}

// ExitSelectorterm is called when production selectorterm is exited.
func (s *BaseGoliteParserListener) ExitSelectorterm(ctx *SelectortermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseGoliteParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseGoliteParserListener) ExitFactor(ctx *FactorContext) {}
