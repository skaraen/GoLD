package parser

import (
	"fmt"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
	"golite/types"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Parser interface {
	Parse() *ast.Program
	GetErrors() []*context.CompilerError
}

type parserWrapper struct {
	*antlr.DefaultErrorListener
	*BaseGoliteParserListener
	antlrParser *GoliteParser
	lexer       lexer.Lexer
	errors      []*context.CompilerError
	nodes 		map[string]interface{}
}

func NewParser(lexer lexer.Lexer) Parser {
	parser := &parserWrapper{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{}, nil, lexer, nil, make(map[string]interface{})}
	antlrParser := NewGoliteParser(lexer.GetTokenStream())
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(parser)
	parser.antlrParser = antlrParser
	return parser
}
func (parser *parserWrapper) GetErrors() []*context.CompilerError {
	return parser.errors
}
func (parser *parserWrapper) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	parser.errors = append(parser.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.PARSER,
	})
}
func (parser *parserWrapper) Parse() *ast.Program {
	ctx := parser.antlrParser.Program()
	if context.HasErrors(parser.lexer.GetErrors()) || context.HasErrors(parser.GetErrors()) {
		return nil
	}
	antlr.ParseTreeWalkerDefault.Walk(parser, ctx)
	_, _, key := GetTokenInfo(ctx)
	return parser.nodes[key].(*ast.Program)
}

func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	id := ctx.GetRuleIndex()
	key := fmt.Sprintf("%d,%d,%d", line, column,id)

	return line, column, key
}

func (parser *parserWrapper) ExitFactor(c *FactorContext) {
	line, column, key := GetTokenInfo(c)
	factorToken := token.NewToken(line, column)

	if intFactor := c.NUM(); intFactor != nil {
		intVal,_ := strconv.Atoi(intFactor.GetText())
		parser.nodes[key] = ast.NewIntLit(factorToken, int64(intVal))
	} else if boolFactor := c.TRUE(); boolFactor != nil {
		boolVal,_ := strconv.ParseBool(boolFactor.GetText())
		parser.nodes[key] = ast.NewBoolLit(factorToken, boolVal)
	} else if boolFactor := c.FALSE(); boolFactor != nil {
		boolVal,_ := strconv.ParseBool(boolFactor.GetText())
		parser.nodes[key] = ast.NewBoolLit(factorToken, boolVal)
	} else if expr := c.Expression(); expr != nil {
		_, _, exprKey := GetTokenInfo(expr)
		parser.nodes[key] = parser.nodes[exprKey].(ast.Expression)
	} else if nilFactor := c.NIL(); nilFactor != nil {
		parser.nodes[key] = ast.NewNilLit(factorToken)	
	} else if idFactor := c.ID(); idFactor != nil {
		if args := c.Args(); args != nil {
			_, _, argsKey := GetTokenInfo(args)
			parser.nodes[key] = ast.NewInlineInvocation(factorToken, idFactor.GetText(), parser.nodes[argsKey].(ast.Expression), false)
		} else if newKw := c.NEW(); newKw != nil {
			parser.nodes[key] = ast.NewInstantation(factorToken, idFactor.GetText())
		} else {
			parser.nodes[key] = ast.NewVariable(factorToken, idFactor.GetText())
		}
	}

	// fmt.Printf("Leaving factor: %s\n", c.GetText())
}

func (parser *parserWrapper) ExitArgs(c *ArgsContext) {
	line, column, key := GetTokenInfo(c)

	exprCtx := c.Expression()
	argsPrimes := c.AllArgsPrime()
	argsList := make([]ast.Expression, 0)

	if exprCtx != nil && argsPrimes != nil {
		_, _, exprKey := GetTokenInfo(exprCtx)
		argsList = append(argsList, parser.nodes[exprKey].(ast.Expression))

		for _, argsPrimeCtx := range (argsPrimes) {
			pExprCtx := argsPrimeCtx.GetExpr()
			_, _, pExprKey := GetTokenInfo(pExprCtx)
			argsList = append(argsList, parser.nodes[pExprKey].(ast.Expression))
		}
	}

	args := ast.NewArgs(token.NewToken(line, column), argsList)
	parser.nodes[key] = args
}

func (parser *parserWrapper) ExitLValue(c *LValueContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	idCtx := c.ID()
	lValuePrimes := c.AllLValuePrime()

	var expr ast.Expression

	expr = ast.NewVariable(token, idCtx.GetText())

	for _, lValuePrimeCtx := range (lValuePrimes) {
		binOp := ast.NewBinaryOp(token, ast.StrToOp("."), expr, ast.NewVariable(token, lValuePrimeCtx.GetPid().GetText()))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitSelectorTerm(c *SelectorTermContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	factorCtx := c.Factor()
	_, _, fKey := GetTokenInfo(factorCtx)
	lValuePrimes := c.AllLValuePrime()

	expr := parser.nodes[fKey].(ast.Expression)

	for _, lValuePrimeCtx := range (lValuePrimes) {
		binOp := ast.NewBinaryOp(token, ast.StrToOp("."), expr, ast.NewVariable(token, lValuePrimeCtx.GetPid().GetText()))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitUnaryTerm(c *UnaryTermContext) {
	line, column, key := GetTokenInfo(c)
	var op ast.Operator

	selectorCtx := c.SelectorTerm()
	_, _, sKey := GetTokenInfo(selectorCtx)
	selectorTerm := parser.nodes[sKey]

	if opCtx := c.MINUS(); opCtx != nil {
		op = ast.StrToOp(opCtx.GetText())
		unOp := ast.NewUnaryOp(token.NewToken(line, column), op, selectorTerm.(ast.Expression))
		parser.nodes[key] = unOp
	} else if opCtx := c.NOT(); opCtx != nil {
		op = ast.StrToOp(opCtx.GetText())
		unOp := ast.NewUnaryOp(token.NewToken(line, column), op, selectorTerm.(ast.Expression))
		parser.nodes[key] = unOp
	} else {
		parser.nodes[key] = selectorTerm
	}
}

func (parser *parserWrapper) ExitTerm(c *TermContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	unaryCtx := c.UnaryTerm()
	_, _, uKey := GetTokenInfo(unaryCtx)
	termPrimes := c.AllTermPrime()

	var expr ast.Expression

	expr = parser.nodes[uKey].(ast.Expression)

	for _, termPrimeCtx := range (termPrimes) {
		op := ast.StrToOp(termPrimeCtx.GetOp().GetText())
		ut := termPrimeCtx.GetUt()
		_, _, utKey := GetTokenInfo(ut)

		binOp := ast.NewBinaryOp(token, op, expr, parser.nodes[utKey].(ast.Expression))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitSimpleTerm(c *SimpleTermContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	termCtx := c.Term()
	_, _, tKey := GetTokenInfo(termCtx)
	sTermPrimes := c.AllSimpleTermPrime()

	var expr ast.Expression

	expr = parser.nodes[tKey].(ast.Expression)

	for _, sTermPrimeCtx := range (sTermPrimes) {
		op := ast.StrToOp(sTermPrimeCtx.GetOp().GetText())
		t := sTermPrimeCtx.GetT()
		_, _, tKey := GetTokenInfo(t)

		binOp := ast.NewBinaryOp(token, op, expr, parser.nodes[tKey].(ast.Expression))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitRelTerm(c *RelTermContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	sTermCtx := c.SimpleTerm()
	_, _, sTermKey := GetTokenInfo(sTermCtx)
	rTermPrimes := c.AllRelTermPrime()

	var expr ast.Expression

	expr = parser.nodes[sTermKey].(ast.Expression)

	for _, rTermPrimeCtx := range (rTermPrimes) {
		op := ast.StrToOp(rTermPrimeCtx.GetOp().GetText())
		st := rTermPrimeCtx.GetSt()
		_, _, stKey := GetTokenInfo(st)

		binOp := ast.NewBinaryOp(token, op, expr, parser.nodes[stKey].(ast.Expression))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitEqualTerm(c *EqualTermContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	rTermCtx := c.RelTerm()
	_, _, rTermKey := GetTokenInfo(rTermCtx)
	eTermPrimes := c.AllEqualTermPrime()

	var expr ast.Expression

	expr = parser.nodes[rTermKey].(ast.Expression)

	for _, eTermPrimeCtx := range (eTermPrimes) {
		op := ast.StrToOp(eTermPrimeCtx.GetOp().GetText())
		rt := eTermPrimeCtx.GetRt()
		_, _, rtKey := GetTokenInfo(rt)

		binOp := ast.NewBinaryOp(token, op, expr, parser.nodes[rtKey].(ast.Expression))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitBoolTerm(c *BoolTermContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	eTermCtx := c.EqualTerm()
	_, _, eTermKey := GetTokenInfo(eTermCtx)
	bTermPrimes := c.AllBoolTermPrime()

	var expr ast.Expression

	expr = parser.nodes[eTermKey].(ast.Expression)

	for _, bTermPrimeCtx := range (bTermPrimes) {
		et := bTermPrimeCtx.GetEt()
		_, _, etKey := GetTokenInfo(et)
		binOp := ast.NewBinaryOp(token, ast.StrToOp("&&"), expr, parser.nodes[etKey].(ast.Expression))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitExpression(c *ExpressionContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	bTermCtx := c.BoolTerm()
	_, _, bTermKey := GetTokenInfo(bTermCtx)
	exprPrimes := c.AllExprPrime()

	var expr ast.Expression

	expr = parser.nodes[bTermKey].(ast.Expression)

	for _, exprPrimeCtx := range (exprPrimes) {
		bt := exprPrimeCtx.GetBt()
		_, _, btKey := GetTokenInfo(bt)

		binOp := ast.NewBinaryOp(token, ast.StrToOp("||"), expr, parser.nodes[btKey].(ast.Expression))
		expr = binOp
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitInvocation(c *InvocationContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	id := c.ID()
	argsCtx := c.Args()
	_, _, argsKey := GetTokenInfo(argsCtx)

	parser.nodes[key] = ast.NewInvocation(token, id.GetText(), parser.nodes[argsKey].(ast.Expression))
}

func (parser *parserWrapper) ExitBlock(c *BlockContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	stmts := c.Statements()
	_, _, stmtsKey := GetTokenInfo(stmts)

	parser.nodes[key] = ast.NewBlock(token, parser.nodes[stmtsKey].(ast.Statement))
}

func (parser *parserWrapper) ExitDelete(c *DeleteContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	exprCtx := c.Expression()
	_, _, exprKey := GetTokenInfo(exprCtx)

	parser.nodes[key] = ast.NewDeleteStmt(token, parser.nodes[exprKey].(ast.Expression))
}

func (parser *parserWrapper) ExitRead(c *ReadContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	lvCtx := c.LValue()
	_, _, lvKey := GetTokenInfo(lvCtx)

	parser.nodes[key] = ast.NewReadStmt(token, parser.nodes[lvKey].(ast.Expression))
}

func (parser *parserWrapper) ExitAssignment(c *AssignmentContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	lvCtx := c.LValue()
	exprCtx := c.Expression()
	_, _, lvKey := GetTokenInfo(lvCtx)
	_, _, exprKey := GetTokenInfo(exprCtx)

	parser.nodes[key] = ast.NewAssignStmt(token, parser.nodes[lvKey].(ast.Expression), parser.nodes[exprKey].(ast.Expression))
}

func (parser *parserWrapper) ExitPrint(c *PrintContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	argsPCtxList := c.AllArgsPrime()
	argsList := make([]ast.Expression, 0)

	for _, argsPCtx := range (argsPCtxList) {
		pExprCtx := argsPCtx.GetExpr()
		_, _, pExprKey := GetTokenInfo(pExprCtx)
		argsList = append(argsList, parser.nodes[pExprKey].(ast.Expression))
	}

	parser.nodes[key] = ast.NewPrintStmt(token, c.STRING().GetText(), argsList)
}

func (parser *parserWrapper) ExitConditional(c *ConditionalContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	exprCtx := c.Expression()
	_, _, exprKey := GetTokenInfo(exprCtx)

	var ifCtx, elseCtx IBlockContext

	blockCtxList := c.AllBlock()
	for i, blockCtx := range (blockCtxList) {
		if i == 0 {
			ifCtx = blockCtx
		} else {
			elseCtx = blockCtx
		}
	}

	var ifBlock, elseBlock *ast.Block
	_, _, ifKey := GetTokenInfo(ifCtx)
	ifBlock = parser.nodes[ifKey].(*ast.Block)

	if elseCtx != nil {
		_, _, elseKey := GetTokenInfo(elseCtx)
		elseBlock = parser.nodes[elseKey].(*ast.Block)
	}

	parser.nodes[key] = ast.NewConditionalStmt(token, parser.nodes[exprKey].(ast.Expression), ifBlock, elseBlock)
}

func (parser *parserWrapper) ExitLoop(c *LoopContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	exprCtx := c.Expression()
	blockCtx := c.Block()
	_, _, exprKey := GetTokenInfo(exprCtx)
	_, _, blockKey := GetTokenInfo(blockCtx)

	parser.nodes[key] = ast.NewLoopStmt(token, parser.nodes[exprKey].(ast.Expression), parser.nodes[blockKey].(*ast.Block))
}

func (parser *parserWrapper) ExitReturn(c *ReturnContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	exprCtx := c.Expression()
	_, _, exprKey := GetTokenInfo(exprCtx)

	parser.nodes[key] = ast.NewReturnStmt(token, parser.nodes[exprKey].(ast.Expression))
}

func (parser *parserWrapper) ExitStatement(c *StatementContext) {
	_, _, key := GetTokenInfo(c)

	if (c.Assignment() != nil) {
		as := c.Assignment()
		_, _, asKey := GetTokenInfo(as)

		parser.nodes[key] = parser.nodes[asKey]
	} else if (c.Print_() != nil) {
		ps := c.Print_()
		_, _, psKey := GetTokenInfo(ps)

		parser.nodes[key] = parser.nodes[psKey]
	} else if (c.Read() != nil) {
		rs := c.Read()
		_, _, rsKey := GetTokenInfo(rs)

		parser.nodes[key] = parser.nodes[rsKey]
	} else if (c.Delete_() != nil) {
		ds := c.Delete_()
		_, _, dsKey := GetTokenInfo(ds)

		parser.nodes[key] = parser.nodes[dsKey]
	} else if (c.Conditional() != nil) {
		cs := c.Conditional()
		_, _, csKey := GetTokenInfo(cs)

		parser.nodes[key] = parser.nodes[csKey]
	} else if (c.Loop() != nil) {
		ls := c.Loop()
		_, _, lsKey := GetTokenInfo(ls)

		parser.nodes[key] = parser.nodes[lsKey]
	} else if (c.Invocation() != nil) {
		is := c.Invocation()
		_, _, isKey := GetTokenInfo(is)

		parser.nodes[key] = parser.nodes[isKey]
	} else if (c.Return_() != nil) {
		rs := c.Return_()
		_, _, rsKey := GetTokenInfo(rs)

		parser.nodes[key] = parser.nodes[rsKey]
	}
}

func (parser *parserWrapper) ExitIds(c *IdsContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	idsPrimes := c.AllIdsPrime()
	idsList := make([]string, 0)
	idsList = append(idsList, c.ID().GetText())

	for _, idsPrime := range (idsPrimes) {
		idsList = append(idsList, idsPrime.GetId().GetText())
	}

	parser.nodes[key] = ast.NewIds(token, idsList)
}

func (parser *parserWrapper) ExitDeclaration(c *DeclarationContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	idsCtx := c.Ids()
	_, _, idsKey := GetTokenInfo(idsCtx)

	var dType types.Type

	if c.Type_().GetPointer() != nil {
		dType = types.StringToType(c.Type_().GetPointer().GetName().GetText())
	} else {
		dType = types.StringToType(c.Type_().GetText())
	}
	
	parser.nodes[key] = ast.NewDeclaration(token, parser.nodes[idsKey].(ast.Expression), dType)
}

func (parser *parserWrapper) ExitDeclarations(c *DeclarationsContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	declarations :=  make([]*ast.Declaration, 0)

	for _, declrCtx := range (c.AllDeclaration()) {
		_, _, declrKey := GetTokenInfo(declrCtx)
		declarations = append(declarations, parser.nodes[declrKey].(*ast.Declaration))
	}

	parser.nodes[key] = ast.NewDeclarations(token, declarations)
}

func (parser *parserWrapper) ExitParams(c *ParamsContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	declCtx := c.Decl()
	paramsPrimes := c.AllParamsPrime()
	paramsList := make([]ast.Expression, 0)

	if declCtx != nil && paramsPrimes != nil {
		_, _, declKey := GetTokenInfo(declCtx)
		paramsList = append(paramsList, parser.nodes[declKey].(ast.Expression))

		for _, paramsPrime := range (paramsPrimes) {
			pDeclCtx := paramsPrime.GetDeclr()
			_, _, pDeclKey := GetTokenInfo(pDeclCtx)
			paramsList = append(paramsList, parser.nodes[pDeclKey].(ast.Expression))
		}
	}

	parser.nodes[key] = ast.NewParams(token, paramsList)
}

func (parser *parserWrapper) ExitStatements(c *StatementsContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	statements :=  make([]ast.Statement, 0)

	for _, sCtx := range (c.AllStatement()) {
		_, _, sKey := GetTokenInfo(sCtx)
		statements = append(statements, parser.nodes[sKey].(ast.Statement))
	}

	parser.nodes[key] = ast.NewStatements(token, statements)
}

func (parser *parserWrapper) ExitFunction(c *FunctionContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	id := c.ID().GetText()
	paramsCtx := c.Params()
	_, _, paramsKey := GetTokenInfo(paramsCtx)
	dCtx := c.Declarations()
	_, _, dKey := GetTokenInfo(dCtx)
	sCtx := c.Statements()
	_, _, sKey := GetTokenInfo(sCtx)

	returnType := types.StringToType("nil") 

	if rtCtx := c.ReturnType(); rtCtx != nil {
		rt := rtCtx.GetRt()
		if intType := rt.GetInteger(); intType != nil {
			returnType = types.StringToType("int")
		} else if boolType := rt.GetBoolean(); boolType != nil {
			returnType = types.StringToType("bool")
		} else if structType := rt.GetPointer(); structType != nil {
			returnType = types.StringToType(structType.GetName().GetText())
		}
	}

	parser.nodes[key] = ast.NewFunction(token, id, 
		parser.nodes[paramsKey].(ast.Expression), 
		returnType, 
		parser.nodes[dKey].(ast.Statement), 
		parser.nodes[sKey].(ast.Statement))
}

func (parser *parserWrapper) ExitFunctions(c *FunctionsContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	functions :=  make([]*ast.Function, 0)

	for _, fCtx := range (c.AllFunction()) {
		_, _, fKey := GetTokenInfo(fCtx)
		functions = append(functions, parser.nodes[fKey].(*ast.Function))
	}

	parser.nodes[key] = ast.NewFunctions(token, functions)
}

func (parser *parserWrapper) ExitDecl(c *DeclContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	var dType types.Type
	dType = nil

	if intType := c.Type_().GetInteger(); intType != nil {
		dType = types.StringToType("int")
	} else if boolType := c.Type_().GetBoolean(); boolType != nil {
		dType = types.StringToType("bool")
	} else if structType := c.Type_().GetPointer(); structType != nil {
		dType = types.StringToType(structType.GetName().GetText())
	}

	parser.nodes[key] = ast.NewDecl(token, c.ID().GetText(), dType)
}

func (parser *parserWrapper) ExitFields(c *FieldsContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	fieldsPrimes := c.AllFieldsPrime()
	fieldsList := make([]ast.Expression, 0)

	for _, fpCtx := range (fieldsPrimes) {
		declrCtx := fpCtx.GetDeclr()
		_, _, declrKey := GetTokenInfo(declrCtx)

		fieldsList = append(fieldsList, parser.nodes[declrKey].(ast.Expression))
	}

	parser.nodes[key] = ast.NewFields(token, fieldsList)
}

func (parser *parserWrapper) ExitTypeDecl(c *TypeDeclContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	id := c.ID().GetText()
	fieldsCtx := c.Fields()
	_, _, fieldsKey := GetTokenInfo(fieldsCtx)

	types.AddNewUserType(id)

	parser.nodes[key] = ast.NewTypeDecl(token, id, parser.nodes[fieldsKey].(ast.Expression))
}

func (parser *parserWrapper) ExitUserTypes(c *UserTypesContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	typeDecls :=  make([]*ast.TypeDecl, 0)

	for _, tdCtx := range (c.AllTypeDecl()) {
		_, _, tdKey := GetTokenInfo(tdCtx)
		typeDecls = append(typeDecls, parser.nodes[tdKey].(*ast.TypeDecl))
	}

	parser.nodes[key] = ast.NewUserTypes(token, typeDecls)
}

func (parser *parserWrapper) ExitProgram(c *ProgramContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	uCtx := c.UserTypes()
	dCtx := c.Declarations()
	fCtx := c.Functions()

	_, _, uKey := GetTokenInfo(uCtx)
	_, _, dKey := GetTokenInfo(dCtx)
	_, _, fKey := GetTokenInfo(fCtx)

	parser.nodes[key] = ast.NewProgram(token, 
							parser.nodes[uKey].(*ast.UserTypes),
							parser.nodes[dKey].(*ast.Declarations),
							parser.nodes[fKey].(*ast.Functions))
}	