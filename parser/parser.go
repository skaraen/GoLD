package parser

import (
	"golite/ast"
	"golite/context"
	"golite/lexer"

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
	/*	if parser.antlrParser.Program() != nil {
			return true
		} else {
			return false
		}*/
	parser.antlrParser.Program()
	if context.HasErrors(parser.lexer.GetErrors()) || context.HasErrors(parser.GetErrors()) {
		return nil
	}
	// AST Construction
	return nil
}

func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	key := fmt.Sprintf("%d,%d", line, column)

	return line, column, key
}