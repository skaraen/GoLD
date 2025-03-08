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

type ReadStmt struct {
	*token.Token
	lValue	Expression
}

func NewReadStmt(token *token.Token, lValue Expression) *ReadStmt {
	return &ReadStmt{token, lValue}
}

func (ds *ReadStmt) String() string {
	var out bytes.Buffer

	out.WriteString("scan " + ds.lValue.String() + ";")

	return out.String()
}

func (rs *ReadStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var lTy types.Type
	lTy, errors = rs.lValue.TypeCheck(errors, funcEntry, tables)
	if (lTy.String() != "int") {
		msg := fmt.Sprintf("trying to read non integer value (%s)", rs.lValue.String())
		semError := context.NewCompilerError(rs.Line, rs.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}

func (rs *ReadStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	valEntry, exists := funcEntry.Variables.Contains(rs.lValue.String())

	if !exists {
		valEntry = st.NewVarEntry("", types.IntTySig, st.GLOBAL, rs.Token)
	}
	currBlk.Instns = append(currBlk.Instns, llvm.NewReadInstn(valEntry))
	
	return currBlk
}