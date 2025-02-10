package ast

import (
	"bytes"
	"golite/token"
)

type Invocation struct {
	*token.Token
	function 	string
	args 		Expression
	isInv		bool
}

func NewInvocation(token *token.Token, function string, args Expression, isInv bool) *Invocation {
	return &Invocation{token, function, args, isInv}
}

func (inv *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.function + "(" + inv.args.String() + ")")

	if (inv.isInv) {
		out.WriteString(";")
	}

	return out.String()
}