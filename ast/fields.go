package ast

import (
	"bytes"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Fields struct {
	*token.Token
	fieldsList []Expression
}

func NewFields(token *token.Token, fieldsList []Expression) *Fields {
	return &Fields{token, fieldsList}
}

func (f *Fields) String() string {
	var out bytes.Buffer

	for _, field := range (f.fieldsList) {
		out.WriteString(field.String() + ";\n")
	}

	return out.String()
}

func (f *Fields) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("undefined")
}

func (f *Fields) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return f.GetType(funcEntry, tables), errors
}