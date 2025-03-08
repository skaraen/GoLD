package ast

import (
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
)

type Statement interface {
	String() string
	TypeCheck([]*context.CompilerError, *st.FuncEntry,  *st.SymbolTables) []*context.CompilerError
	TranslateToLLVMStack(*cfg.Block, *cfg.Block, *llvm.LLVMProgram, *st.FuncEntry, *st.SymbolTables) *cfg.Block
}