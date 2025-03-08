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

type ReturnStmt struct {
	*token.Token
	expr	Expression
}

func NewReturnStmt(token *token.Token, expr Expression) *ReturnStmt {
	return &ReturnStmt{token, expr}
}

func (rs *ReturnStmt) String() string {
	var out bytes.Buffer

	if rs.expr != nil {
		out.WriteString("return " + rs.expr.String() + ";")
	} else {
		out.WriteString("return;")
	}

	return out.String()
}

func (rs *ReturnStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var retTy types.Type
	if rs.expr != nil {
		retTy, errors = rs.expr.TypeCheck(errors, funcEntry, tables)
	} else {
		retTy = types.NilTySig
	}

	if retTy != funcEntry.RetTy {
		msg := fmt.Sprintf("return type of expression (%s) does not match return type of function (%s)", retTy, funcEntry.RetTy)
		semError := context.NewCompilerError(rs.Line, rs.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}

func (rs *ReturnStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	var temp1 llvm.LLVMOperand
	var temp1Ty types.Type
	if rs.expr != nil {
		temp1 = rs.expr.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
		temp1Ty = temp1.GetType()
	} else {
		temp1Ty = types.NilTySig
	}

	retReg := llvmProgram.RetRegisters[funcEntry.Name]
	currBlk.Instns = append(currBlk.Instns, llvm.NewStoreInstn(retReg, retReg.GetType(), temp1, temp1Ty))
	currBlk.Instns = append(currBlk.Instns, llvm.NewJumpInstn(exitBlk))
	return currBlk
}