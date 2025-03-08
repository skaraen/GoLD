package ast

import (
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/types"
)

type Expression interface {
	String() string
	//Returns the type after the expression is fully typed checked.
	GetType(*st.FuncEntry, *st.SymbolTables) types.Type
	//Typechecks an expression.
	TypeCheck([]*context.CompilerError, *st.FuncEntry,  *st.SymbolTables) (types.Type, []*context.CompilerError)
	TranslateToLLVMStack(*st.FuncEntry, *st.SymbolTables, *cfg.Block, *llvm.LLVMProgram) llvm.LLVMOperand
}