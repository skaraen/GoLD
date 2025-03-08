package ast

import (
	"bytes"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
)

type Block struct {
	*token.Token
	statements 		Statement
}

func NewBlock(token *token.Token, statements Statement) *Block {
	return &Block{token, statements}
}

func (b *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{\n" + b.statements.String() + "}")

	return out.String()
}

func (b *Block) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	errors = b.statements.TypeCheck(errors, funcEntry, tables)

	return errors
}

func (b *Block) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	currBlk = b.statements.TranslateToLLVMStack(currBlk, exitBlk, llvmProgram, funcEntry, tables)

	return currBlk
}