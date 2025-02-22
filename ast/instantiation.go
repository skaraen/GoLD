package ast

import (
	"bytes"
	"golite/token"
)

type Instantiation struct {
	*token.Token
	id string
}

func NewInstantation(token *token.Token, id string) *Instantiation {
	return &Instantiation{token, id}
}

func (ins *Instantiation) String() string {
	var out bytes.Buffer

	out.WriteString("new "+ ins.id)

	return out.String()
}