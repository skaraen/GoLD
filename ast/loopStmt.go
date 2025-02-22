package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type LoopStmt struct {
	*token.Token
	expr			Expression
	loopBlock		*Block
}

func NewLoopStmt(token *token.Token, expr Expression, block *Block) *LoopStmt {
	return &LoopStmt{token, expr, block}
}

func (ls *LoopStmt) String() string {
	var out bytes.Buffer

	out.WriteString("for (" + ls.expr.String() + ") " + ls.loopBlock.String())

	return out.String()
}

func (ls *LoopStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var eTy types.Type
	eTy, errors = ls.expr.TypeCheck(errors, funcEntry, tables)
	if eTy.String() != "bool" {
		msg := fmt.Sprintf("loop condition (%s) is not a boolean expression", ls.expr.String())
		semError := context.NewCompilerError(ls.Line, ls.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	errors = ls.loopBlock.TypeCheck(errors, funcEntry, tables)

	return errors
}