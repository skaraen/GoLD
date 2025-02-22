package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type ReadStmt struct {
	*token.Token
	lValue	Expression
}

func NewReadStmt(token *token.Token, lValue Expression) *ReadStmt {
	return &ReadStmt{token, lValue}
}

func (ds *ReadStmt) String() string {
	var out bytes.Buffer

	out.WriteString("scan " + ds.lValue.String() + ";")

	return out.String()
}

func (rs *ReadStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var lTy types.Type
	lTy, errors = rs.lValue.TypeCheck(errors, funcEntry, tables)
	if (lTy.String() != "int") {
		msg := fmt.Sprintf("trying to read non integer value (%s)", rs.lValue.String())
		semError := context.NewCompilerError(rs.Line, rs.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}