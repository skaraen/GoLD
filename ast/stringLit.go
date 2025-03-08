package ast

import (
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type StrLit struct {
	*token.Token
	litType 	types.Type
	value 		string
}

func NewStrLit(token *token.Token, value string) *StrLit {
	return &StrLit{token, types.StringToType("string"), value}
}

func (strLit *StrLit) String() string {
	return strLit.value
}

func (strLit *StrLit) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("string")
}

func (strLit *StrLit) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return strLit.GetType(funcEntry, tables), errors
}

func (strLit *StrLit) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	return nil
}