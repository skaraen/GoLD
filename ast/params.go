package ast

import (
	"bytes"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Params struct {
	*token.Token
	paramsList []Expression
}

func NewParams(token *token.Token, paramsList []Expression) *Params {
	return &Params{token, paramsList}
}

func (params *Params) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	for i, p := range (params.paramsList) {
		out.WriteString(p.String())
		if (i < len(params.paramsList) - 1) {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	return out.String()
}

func (params *Params) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("undefined")
}

func (params *Params) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return params.GetType(funcEntry, tables), errors
}