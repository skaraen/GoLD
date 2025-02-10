package lexer

import (
	"fmt"
	"golite/context"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Lexer interface {
	GetTokenStream() *antlr.CommonTokenStream
	GetErrors() []*context.CompilerError
}

type lexerWrapper struct {
	*antlr.DefaultErrorListener
	antrlLexer *GoliteLexer
	stream     *antlr.CommonTokenStream
	errors     []*context.CompilerError
}

func (lex *lexerWrapper) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	lex.errors = append(lex.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.LEXER,
	})
}

func (lexer *lexerWrapper) GetTokenStream() *antlr.CommonTokenStream {
	return lexer.stream
}
func (lexer *lexerWrapper) GetErrors() []*context.CompilerError {
	return lexer.errors
}

func NewLexer(inputSourcePath string) Lexer {
	input, _ := antlr.NewFileStream(inputSourcePath)
	lexer := &lexerWrapper{antlr.NewDefaultErrorListener(), nil, nil, nil}
	antlrLexer := NewGoliteLexer(input)
	antlrLexer.RemoveErrorListeners()
	antlrLexer.AddErrorListener(lexer)
	tokenStream := antlr.NewCommonTokenStream(antlrLexer, 0)
	lexer.antrlLexer = antlrLexer
	lexer.stream = tokenStream
	return lexer
}

func (lexer *lexerWrapper) PrintTokens() {
    lexer.stream.Fill()
    for _, token := range lexer.stream.GetAllTokens() {
        if token.GetTokenType() != antlr.TokenEOF {
            fmt.Printf("TOKEN.%v(%v,%v)\n", lexer.antrlLexer.SymbolicNames[token.GetTokenType()], token.GetStart(), token.GetText())
        }
    }
}