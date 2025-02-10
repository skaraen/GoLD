package ast

import (
	"bytes"
	"golite/token"
)

type AssignStmt struct {
	*token.Token
	lValue 			string
	rExpression		Expression
}

func NewAssignStmt(token *token.Token, lValue string, rExpression Expression) *AssignStmt {
	return &AssignStmt{token, lValue, rExpression}
}

func (as *AssignStmt) String() string {
	var out bytes.Buffer

	out.WriteString(as.lValue + " = " + as.rExpression.String() + ";")

	return out.String()
}