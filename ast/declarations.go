package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
)

type Declarations struct {
	*token.Token
	declrList		[]*Declaration
}

func NewDeclarations(token *token.Token, declrList []*Declaration) *Declarations {
	return &Declarations{token, declrList}
}

func (d *Declarations) String() string {
	var out bytes.Buffer

	for _, declaration := range (d.declrList) {
		if (declaration == nil) {
			fmt.Println("Nil")
		}
		out.WriteString(declaration.String() + "\n")
	}

	return out.String()
}

func (d *Declarations) BuildSymbolTable(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	for _, declr := range (d.declrList) {
		errors = declr.BuildSymbolTable(errors, tables)
	}

	return errors
}

func (d *Declarations) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	return errors
}