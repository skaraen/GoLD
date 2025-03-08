package ast

import (
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strconv"
)

type IntLit struct {
	*token.Token
	litType 	types.Type
	value 		int64
}

func NewIntLit(token *token.Token, value int64) *IntLit {
	return &IntLit{token, types.StringToType("int"), value}
}

func (intLit *IntLit) String() string {
	return strconv.FormatInt(intLit.value, 10)
}

func (intLit *IntLit) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("int")
}

func (intLit *IntLit) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return intLit.GetType(funcEntry, tables), errors
}

func (intLit *IntLit) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	return llvm.NewLLVMImmediate(intLit.value, intLit.litType)
}