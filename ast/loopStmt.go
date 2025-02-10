package ast

import (
	"bytes"
	"golite/token"
)

type LoopStmt struct {
	*token.Token
	expr			Expression
	loopBlock			*Block
}

func NewLoopStmt(token *token.Token, expr Expression, block *Block) *LoopStmt {
	return &LoopStmt{token, expr, block}
}

func (ls *LoopStmt) String() string {
	var out bytes.Buffer

	out.WriteString("for (" + ls.expr.String() + ") " + ls.loopBlock.String())

	return out.String()
}