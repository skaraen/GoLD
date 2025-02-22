package ast

import (
	"golite/context"
	st "golite/symboltable"
)

type Statement interface {
	String() string
	TypeCheck([]*context.CompilerError, *st.FuncEntry,  *st.SymbolTables) []*context.CompilerError
}