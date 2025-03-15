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

func (cs *ConditionalStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	labels := llvmProgram.GenerateLabelsBatch(funcEntry.Name, 3)
	trueBlk, falseBlk, exitIfBlk := cfg.NewIfBlock(currBlk, exitBlk, labels[0], labels[1], labels[2])
	condReg := cs.expr.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram).(*llvm.LLVMRegister)
	
	// currBlk.Instns = append(currBlk.Instns, llvm.NewBranchInstn(condReg, trueBlk, falseBlk))
	currBlk.AddInstn(llvm.NewBranchInstn(condReg, trueBlk, falseBlk))
	currTrueBlk := cs.ifBlock.TranslateToLLVMStack(trueBlk, exitBlk, llvmProgram, funcEntry, tables)
	// currTrueBlk.Instns = append(currTrueBlk.Instns, llvm.NewJumpInstn(exitIfBlk))
	currTrueBlk.AddInstn(llvm.NewJumpInstn(exitIfBlk))
	
	if !currTrueBlk.IsReturning {
		currTrueBlk.AddSucc(exitIfBlk)
		exitIfBlk.AddPred(currTrueBlk)
	}
	
	currFalseBlk := falseBlk
	if cs.elseBlock != nil {
		currFalseBlk = cs.elseBlock.TranslateToLLVMStack(falseBlk, exitBlk, llvmProgram, funcEntry, tables)
	}
	// currFalseBlk.Instns = append(currFalseBlk.Instns, llvm.NewJumpInstn(exitIfBlk))
	currFalseBlk.AddInstn(llvm.NewJumpInstn(exitIfBlk))
	
	if !currFalseBlk.IsReturning {
		currFalseBlk.AddSucc(exitIfBlk)
		exitIfBlk.AddPred(currFalseBlk)
	}

	return exitIfBlk
}