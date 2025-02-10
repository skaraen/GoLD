package ast

import (
	"bytes"
	"golite/token"
)

type PrintStmt struct {
	*token.Token
	str		 		string
	args			Expression
}

func NewPrintStmt(token *token.Token, str string, args Expression) *PrintStmt {
	return &PrintStmt{token, str, args}
}

func (ps *PrintStmt) String() string {
	var out bytes.Buffer

	out.WriteString("printf(\"" + ps.str + "\"" + ps.args.String())

	if (ps.args != nil) {
		out.WriteString(ps.args.String())
	}
	out.WriteString(");")

	return out.String()
}