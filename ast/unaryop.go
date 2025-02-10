package ast

import (
	"bytes"
	"golite/token"
)

type UnaryOp struct {
	*token.Token
	operator 	Operator
	right 		Expression
}

func NewUnaryOp(token *token.Token, operator Operator, right Expression) *UnaryOp {
	return &UnaryOp{token, operator, right}
}

func (unrOp *UnaryOp) String() string {
	var out bytes.Buffer

	out.WriteString(OpToStr(unrOp.operator))
	out.WriteString(unrOp.right.String())

	return out.String()
}