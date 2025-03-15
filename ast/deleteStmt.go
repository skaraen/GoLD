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

type DeleteStmt struct {
	*token.Token
	expr	Expression
}

func NewDeleteStmt(token *token.Token, expr Expression) *DeleteStmt {
	return &DeleteStmt{token, expr}
}

func (ds *DeleteStmt) String() string {
	var out bytes.Buffer

	out.WriteString("delete " + ds.expr.String() + ";")

	return out.String()
}

func (ds *DeleteStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var eTy types.Type
	eTy, errors = ds.expr.TypeCheck(errors, funcEntry, tables)
	if _, ok := eTy.(*types.PointerTy); !ok {
		msg := fmt.Sprintf("trying to delete a non reference value (%s)", ds.expr.String())
		semError := context.NewCompilerError(ds.Line, ds.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}

func (ds *DeleteStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	op := ds.expr.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram).(*llvm.LLVMRegister)
	bitcastPtr := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(ds.String(), types.Int8PtrTySig, st.LOCAL, ds.Token), false)

	// currBlk.Instns = append(currBlk.Instns, llvm.NewBitcastInstn(bitcastPtr, op))
	currBlk.AddInstn(llvm.NewBitcastInstn(bitcastPtr, op))
	// currBlk.Instns = append(currBlk.Instns, llvm.NewDeleteInstn(bitcastPtr))
	currBlk.AddInstn(llvm.NewDeleteInstn(bitcastPtr))
	return currBlk
}