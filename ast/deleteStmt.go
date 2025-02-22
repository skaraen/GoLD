package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type DeleteStmt struct {
	*token.Token
	expr	Expression
}

func NewDeleteStmt(token *token.Token, expr Expression) *DeleteStmt {
	return &DeleteStmt{token, expr}
}

func (ds *DeleteStmt) String() string {
	var out bytes.Buffer

	out.WriteString("delete " + ds.expr.String() + ";")

	return out.String()
}

func (ds *DeleteStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var eTy types.Type
	eTy, errors = ds.expr.TypeCheck(errors, funcEntry, tables)
	if _, ok := eTy.(*types.PointerTy); !ok {
		msg := fmt.Sprintf("trying to delete a non reference value (%s)", ds.expr.String())
		semError := context.NewCompilerError(ds.Line, ds.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}