package ast

import (
	"bytes"
	"golite/token"
)

type ConditionalStmt struct {
	*token.Token
	expr			Expression
	ifBlock			*Block
	elseBlock		*Block
}

func NewConditionalStmt(token *token.Token, expr Expression, ifBlock *Block, elseBlock *Block) *ConditionalStmt {
	return &ConditionalStmt{token, expr, ifBlock, elseBlock}
}

func (cs *ConditionalStmt) String() string {
	var out bytes.Buffer

	out.WriteString("if (" + cs.expr.String() + ") " + cs.ifBlock.String())

	if (cs.elseBlock != nil) {
		out.WriteString("else " + cs.elseBlock.String())
	}

	return out.String()
}