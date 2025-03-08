package ast

import (
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type NilLit struct {
	*token.Token
	litType 	types.Type
}

func NewNilLit(token *token.Token) *NilLit {
	return &NilLit{token, types.NilTySig}
}

func (nilLit *NilLit) String() string {
	return "nil"
}

func (nilLit *NilLit) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("nil")
}

func (nilLit *NilLit) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return nilLit.GetType(funcEntry, tables), errors
}

func (nilLit *NilLit) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	return nil
}