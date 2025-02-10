package ast

import (
	"bytes"
	"golite/token"
)

type Args struct {
	*token.Token
	argsList []Expression
}

func NewArgs(token *token.Token, argsList []Expression) *Args {
	return &Args{token, argsList}
}

func (args *Args) String() string {
	var out bytes.Buffer

	for i, arg := range (args.argsList) {
		out.WriteString(arg.String())
		if (i < len(args.argsList) - 1) {
			out.WriteString(", ")
		}
	}

	return out.String()
}