package parser

import (
	"fmt"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
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
	key := fmt.Sprintf("%d,%d", line, column)

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
	} else if args := c.Args(); args != nil {
		_, _, argsKey := GetTokenInfo(args)
		parser.nodes[key] = ast.NewInvocation(factorToken, "Dummy", parser.nodes[argsKey].(ast.Expression), false)
	}
	// NEW ID, NIL
}

func (parser *parserWrapper) ExitArgs(c *ArgsContext) {
	line, column, key := GetTokenInfo(c)

	exprCtx := c.Expression()
	argsPrimes := c.AllArgsPrime()
	
	_, _, exprKey := GetTokenInfo(exprCtx)
	argsList := make([]ast.Expression, 0)
	argsList = append(argsList, parser.nodes[exprKey].(ast.Expression))

	for _, argsPrimeCtx := range (argsPrimes) {
		_, _, argsPrimeKey := GetTokenInfo(argsPrimeCtx)
		argsList = append(argsList, parser.nodes[argsPrimeKey].(ast.Expression))
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
		binOp := ast.NewBinaryOp(token, ast.StrToOp("."), expr, ast.NewVariable(token, lValuePrimeCtx.GetText()))
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

	var expr ast.Expression

	expr = parser.nodes[fKey].(ast.Expression)

	for _, lValuePrimeCtx := range (lValuePrimes) {
		binOp := ast.NewBinaryOp(token, ast.StrToOp("."), expr, ast.NewVariable(token, lValuePrimeCtx.GetText()))
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

	parser.nodes[key] = ast.NewInvocation(token, id.GetText(), parser.nodes[argsKey].(ast.Expression), true)
}

// func (parser *parserWrapper) ExitBlock(c *BlockContext) {
// 	line, column, key := GetTokenInfo(c)
// 	token := token.NewToken(line, column)

// 	stmts := c.AllStatement()

// 	for _, stmt := range (stmts) {
// 		if (stmt.GetA() != nil) {

// 		} else if (stmt.GetP() != nil) {

// 		} else if (stmt.GetP() != nil) {

// 		} else if (stmt.GetP() != nil) {

// 		}
// 	}
// }

func (parser *parserWrapper) ExitStatement(c *StatementContext) {
	line, column, key := GetTokenInfo(c)
	token := token.NewToken(line, column)

	var stmtNode ast.Statement

	if (c.Assignment() != nil) {
		as := c.Assignment()
		lv := as.GetLv()
		expr := as.GetExpr()
		_,_,exprKey := GetTokenInfo(expr)

		stmtNode = ast.NewAssignStmt(token, lv.GetText(), parser.nodes[exprKey].(ast.Expression))
	} else if (c.Print_() != nil) {
		// ps := c.Print_()
	} else if (c.Read() != nil) {
		rs := c.Read()
		lv := rs.GetLv() 
		_,_,lvKey := GetTokenInfo(lv)

		stmtNode = ast.NewReadStmt(token, parser.nodes[lvKey].(ast.Expression).String())
	} else if (c.Delete_() != nil) {
		ds := c.Delete_()
		expr := ds.GetExpr()
		_,_,exprKey := GetTokenInfo(expr)

		stmtNode = ast.NewDeleteStmt(token, parser.nodes[exprKey].(ast.Expression))
	} else if (c.Conditional() != nil) {
		// cs := c.Conditional()
		// expr := cs.GetExpr()
		// _,_,exprKey := GetTokenInfo(expr)
		// ifBlock := cs.GetIf_()
		// _,_,ifKey := GetTokenInfo(ifBlock)

		// var elseBlock ast.Statement
		// elseBlock = nil

		// if (cs.GetElse_() != nil) {
		// 	elseBlock := cs.GetIf_()
		// 	_,_,elseKey := GetTokenInfo(elseBlock)
		// 	elseBlock = parser.nodes[elseKey].(ast.Statement)
		// }
		// expr := ds.GetExpr()
		// _,_,exprKey := GetTokenInfo(expr)
	} else if (c.Loop() != nil) {

	} else if (c.Invocation() != nil) {

	}

	parser.nodes[key] = stmtNode
}
