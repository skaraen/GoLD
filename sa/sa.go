package sa

import (
	"golite/ast"
	"golite/context"
	st "golite/symboltable"
)

func Execute(program *ast.Program) *st.SymbolTables {
	tables := st.NewSymbolTables()

	errors := make([]*context.CompilerError, 0)
	errors = program.BuildSymbolTable(errors, tables)

	if !context.HasErrors(errors) {
		errors := make([]*context.CompilerError, 0)
		errors = program.TypeCheck(errors, tables)
		if !context.HasErrors(errors) {
			return tables
		}
	}
	return nil
}