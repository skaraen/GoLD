package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type ReturnStmt struct {
	*token.Token
	expr	Expression
}

func NewReturnStmt(token *token.Token, expr Expression) *ReturnStmt {
	return &ReturnStmt{token, expr}
}

func (rs *ReturnStmt) String() string {
	var out bytes.Buffer

	out.WriteString("return " + rs.expr.String() + ";")

	return out.String()
}

func (rs *ReturnStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var retTy types.Type
	
	retTy, errors = rs.expr.TypeCheck(errors, funcEntry, tables)

	if retTy != funcEntry.RetTy {
		msg := fmt.Sprintf("return type of expression (%s) does not match return type of function (%s)", retTy, funcEntry.RetTy)
		semError := context.NewCompilerError(rs.Line, rs.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}