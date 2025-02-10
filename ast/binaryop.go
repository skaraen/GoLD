package ast

import (
	"bytes"
	"golite/token"
)

type BinaryOp struct {
	*token.Token
	operator 	Operator
	left 		Expression
	right 		Expression
}

func NewBinaryOp(token *token.Token, operator Operator, left Expression, right Expression) *BinaryOp {
	return &BinaryOp{token, operator, left, right}
}

func (binOp *BinaryOp) String() string {
	var out bytes.Buffer

	out.WriteString(binOp.left.String())
	out.WriteString(" " + OpToStr(binOp.operator) + " ")
	out.WriteString(binOp.right.String())

	return out.String()
}