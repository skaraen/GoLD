package ast

import (
	"bytes"
	"golite/token"
)

type Block struct {
	*token.Token
	statements 		[]Statement
}

func NewBlock(token *token.Token, statement []Statement) *Block {
	return &Block{token, statement}
}

func (b *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{\n")

	for _, stmt := range (b.statements) {
		out.WriteString(stmt.String() + "\n")
	}
	out.WriteString("}\n")

	return out.String()
}