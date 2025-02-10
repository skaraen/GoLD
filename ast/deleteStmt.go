package ast

import (
	"bytes"
	"golite/token"
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