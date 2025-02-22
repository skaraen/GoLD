package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
)

type Invocation struct {
	*token.Token
	function 	string
	args 		Expression
}

func NewInvocation(token *token.Token, function string, args Expression) *Invocation {
	return &Invocation{token, function, args}
}

func (inv *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.function + inv.args.String() + ";")

	return out.String()
}

func (inv *Invocation) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	// var fEntry *st.FuncEntry

	if _, exists := tables.Funcs.Contains(inv.function); !exists {
		msg := fmt.Sprintf("method (%s) trying to be invoked does not exist", inv.function)
		semError := context.NewCompilerError(inv.Line, inv.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)

		return errors
	}

	// for _, arg := range (inv.args.(*Args).argsList) {
	// 	argTy := arg.GetType(funcEntry, tables)

	// 	// Match function signatures
	// }

	return errors
}