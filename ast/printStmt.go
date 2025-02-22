package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"regexp"
)

type PrintStmt struct {
	*token.Token
	str		 		string
	args 			[]Expression
}

func NewPrintStmt(token *token.Token, str string, args []Expression) *PrintStmt {
	return &PrintStmt{token, str, args}
}

func (ps *PrintStmt) String() string {
	var out bytes.Buffer

	out.WriteString("printf(" + ps.str)

	if ps.args != nil {
		for _, arg := range (ps.args) {
			out.WriteString(", " + arg.String());
		}
	}

	out.WriteString(");")

	return out.String()
}

func countPlaceholders(format string) int {
	re := regexp.MustCompile(`%[+#0-9.]*[vTtbcdoqxXUeEfFgGsp]`)
	matches := re.FindAllString(format, -1)
	return len(matches)
}

func (ps *PrintStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var argTy types.Type
	for _, arg := range (ps.args) {
		argTy, errors = arg.TypeCheck(errors, funcEntry, tables)
		if (argTy.String() != "int") {
			msg := fmt.Sprintf("trying to print non integer value (%s)", arg.String())
			semError := context.NewCompilerError(ps.Line, ps.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		}
	}

	if countPlaceholders(ps.str) != len(ps.args) {
		msg := fmt.Sprintf("number of placeholders (%d) does not match number of arguments (%d) to print", countPlaceholders(ps.str), len(ps.args))
		semError := context.NewCompilerError(ps.Line, ps.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}