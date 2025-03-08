package ast

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
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

func (ls *LoopStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	labels := llvmProgram.GenerateLabelsBatch(funcEntry.Name, 3)
	entryForBlk, bodyBlk, exitForBlk := cfg.NewForBlock(currBlk, exitBlk, labels[0], labels[1], labels[2])
	condReg := ls.expr.TranslateToLLVMStack(funcEntry, tables, entryForBlk, llvmProgram).(*llvm.LLVMRegister)
	
	entryForBlk.Instns = append(entryForBlk.Instns, llvm.NewBranchInstn(condReg, bodyBlk, exitForBlk))
	
	currBlk = ls.loopBlock.TranslateToLLVMStack(bodyBlk, exitBlk, llvmProgram, funcEntry, tables)
	currBlk.Instns = append(bodyBlk.Instns, llvm.NewJumpInstn(entryForBlk))
	currBlk.Succs = append(currBlk.Succs, entryForBlk)
	entryForBlk.Preds = append(entryForBlk.Preds, currBlk)

	return exitForBlk
}