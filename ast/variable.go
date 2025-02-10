package ast

import (
	"golite/token"
)

type Variable struct {
	*token.Token
	identifier	string
	// varType 	types.Type
}

func NewVariable(token *token.Token, identifier string) *Variable {
	return &Variable{token, identifier}
}

func (v *Variable) String() string {
	return v.identifier
}