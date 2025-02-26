package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
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

func (cs *ConditionalStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var eTy types.Type
	eTy, errors = cs.expr.TypeCheck(errors, funcEntry, tables)
	if eTy.String() != "bool" {
		msg := fmt.Sprintf("Expected boolean expression in if conditional, found %s expression (%s) instead", eTy.String(), cs.expr.String())
		semError := context.NewCompilerError(cs.Line, cs.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	errors = cs.ifBlock.TypeCheck(errors, funcEntry, tables)

	if (cs.elseBlock != nil) {
		errors = cs.elseBlock.TypeCheck(errors, funcEntry, tables)
	}

	return errors
}