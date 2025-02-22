package ast

import (
	"bytes"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
)

type Statements struct {
	*token.Token
	stmtsList			[]Statement
}

func NewStatements(token *token.Token, stmtsList []Statement) *Statements {
	return &Statements{token, stmtsList}
}

func (stmts *Statements) String() string {
	var out bytes.Buffer

	for _, stmt := range (stmts.stmtsList) {
		out.WriteString(stmt.String() + "\n")
	}

	return out.String()
}

func (stmts *Statements) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	for _, stmt := range (stmts.stmtsList) {
		errors = stmt.TypeCheck(errors, funcEntry, tables)
	}

	return errors
}