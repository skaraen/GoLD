package ast

import (
	"bytes"
	"golite/token"
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