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

type Declaration struct {
	*token.Token 
	ids				Expression
	declType 		types.Type
}

func NewDeclaration(token *token.Token, ids Expression, declType types.Type) *Declaration {
	return &Declaration{token, ids, declType}
}

func (d *Declaration) String() string {
	var out bytes.Buffer

	out.WriteString("var " + d.ids.String() + " ")

	if _, ok := d.declType.(*types.PointerTy); ok {
		out.WriteString("*")
	}

	out.WriteString(d.declType.String() + ";");

	return out.String()
}

func (d *Declaration) BuildSymbolTable(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	ids := d.ids.(*Ids)

	for _, id := range(ids.idsList) {
		if st, exists := tables.Structs.Contains(id); exists {
			msg := fmt.Sprintf("redeclaration of user defined type (%s) located at (%d,%d)", id, st.Line, st.Column)
			semError := context.NewCompilerError(ids.Line, ids.Column, msg, context.SEMANTICS)
			
			errors = append(errors, semError)
			continue
		}

		if gv, exists := tables.Globals.Contains(id); exists {
			msg := fmt.Sprintf("redeclaration of global variable (%s) located at (%d,%d)", id, gv.Line, gv.Column)
			semError := context.NewCompilerError(ids.Line, ids.Column, msg, context.SEMANTICS)
			
			errors = append(errors, semError)
			continue
		}

		varEntry := st.NewVarEntry(id, d.declType, st.GLOBAL, ids.Token)
		tables.Globals.Insert(id, varEntry)
	}

	return errors
}

func (d *Declaration) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	idsList := d.ids.(*Ids).idsList

	for _, id := range (idsList) {
		varEntry, _ := funcEntry.Variables.Contains(id)
		idRegister := llvm.NewLLVMRegister(varEntry.Name, varEntry)
		llvmProgram.AddEntryRegister(varEntry, idRegister)
	
		allocateInstn := llvm.NewAllocateInstn(idRegister, d.declType)
		currBlk.Instns = append(currBlk.Instns, allocateInstn)
	}
	
	return currBlk
}