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

type BoolLit struct {
	*token.Token
	litType 	types.Type
	value 		bool
}

func NewBoolLit(token *token.Token, value bool) *BoolLit {
	return &BoolLit{token, types.StringToType("bool"), value}
}

func (boolLit *BoolLit) String() string {
	return strconv.FormatBool(boolLit.value)
}

func (boolLit *BoolLit) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("bool")
}

func (boolLit *BoolLit) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return boolLit.GetType(funcEntry, tables), errors
}

func (boolLit *BoolLit) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	val := 0
	if boolLit.value {
		val = 1
	}
	return llvm.NewLLVMImmediate(int64(val), boolLit.litType)
}