package sa

import (
	"golite/ast"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
)

func Execute(program *ast.Program, sourceName string, targetTriple string) *llvm.LLVMProgram {
	tables := st.NewSymbolTables()

	errors := make([]*context.CompilerError, 0)
	errors = program.BuildSymbolTable(errors, tables)

	if !context.HasErrors(errors) {
		errors := make([]*context.CompilerError, 0)
		errors = program.TypeCheck(errors, tables)
		if !context.HasErrors(errors) {
			return program.TranslateToLLVM(sourceName, targetTriple, tables)
		}
	}
	return nil
}