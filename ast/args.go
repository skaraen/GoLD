package ast

import (
	"bytes"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Args struct {
	*token.Token
	argsList []Expression
}

func NewArgs(token *token.Token, argsList []Expression) *Args {
	return &Args{token, argsList}
}

func (args *Args) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	for i, arg := range (args.argsList) {
		out.WriteString(arg.String())
		if (i < len(args.argsList) - 1) {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	return out.String()
}

func (args *Args) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("undefined")
}

func (args *Args) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	for _, arg := range args.argsList {
		_, errors =  arg.TypeCheck(errors, funcEntry, tables)
	}
	
	return args.GetType(funcEntry, tables), errors
}