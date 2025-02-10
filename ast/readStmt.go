package ast

import (
	"bytes"
	"golite/token"
)

type ReadStmt struct {
	*token.Token
	lValue	string
}

func NewReadStmt(token *token.Token, lValue string) *ReadStmt {
	return &ReadStmt{token, lValue}
}

func (ds *ReadStmt) String() string {
	var out bytes.Buffer

	out.WriteString("scan " + ds.lValue + ";")

	return out.String()
}