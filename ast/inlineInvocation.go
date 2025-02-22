package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type InlineInvocation struct {
	*token.Token
	function 	string
	args 		Expression
}

func NewInlineInvocation(token *token.Token, function string, args Expression, isInv bool) *InlineInvocation {
	return &InlineInvocation{token, function, args}
}

func (inv *InlineInvocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.function + inv.args.String())

	return out.String()
}

func (inv *InlineInvocation) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	if fEntry, exists := tables.Funcs.Contains(inv.function); exists {
		return fEntry.RetTy
	}

	return types.StringToType("undefined")
}

func (inv *InlineInvocation) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	if _, exists := tables.Funcs.Contains(inv.function); !exists {
		msg := fmt.Sprintf("method (%s) trying to be invoked does not exist", inv.function)
		semError := context.NewCompilerError(inv.Line, inv.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	// for _, arg := range (inv.args.(*Args).argsList) {
	// 	argTy := arg.GetType(funcEntry, tables)

	// 	// Match function signatures
	// }

	return inv.GetType(funcEntry, tables), errors
}