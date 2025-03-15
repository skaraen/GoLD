package ast

import (
	"bytes"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Program struct {
	*token.Token
	userTypes		*UserTypes
	declarations 	*Declarations
	functions		*Functions
}

func NewProgram(token *token.Token, userTypes *UserTypes, declarations *Declarations, functions *Functions) *Program {
	return &Program{token, userTypes, declarations, functions}
}

func (p *Program) String() string {
	var out bytes.Buffer

	out.WriteString(p.userTypes.String() + "\n")
	out.WriteString(p.declarations.String() + "\n")
	out.WriteString(p.functions.String() + "\n")

	return out.String()
}

func (p *Program) BuildSymbolTable(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	errors = p.userTypes.BuildSymbolTable(errors, tables.Structs)
	errors = p.declarations.BuildSymbolTable(errors, tables)
	errors = p.functions.BuildSymbolTable(errors, tables)
	
	return errors
}

func (p *Program) TypeCheck(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	errors = p.functions.TypeCheck(errors, tables)
	
	return errors
}

func (p *Program) TranslateToLLVM(srcName string, targetTriple string, tables *st.SymbolTables) *llvm.LLVMProgram {
	llvmProgram := llvm.NewLLVMProgram(srcName, targetTriple, tables)

	llvmProgram = p.userTypes.TranslateToLLVMStack(llvmProgram, tables)

	declrList := p.declarations.declrList
	for _, declr := range (declrList) {
		idsList := declr.ids.(*Ids).idsList

		for _, id := range (idsList) {
			varEntry,_ := tables.Globals.Contains(id) 
			reg := llvm.NewLLVMRegister(varEntry.Name, varEntry, false)
			llvmProgram.AddEntryRegister(varEntry, reg)
			llvmProgram.Globals = append(llvmProgram.Globals, llvm.NewGlobalDecl(id, declr.declType))
		}
	}
	llvmProgram.Globals = append(llvmProgram.Globals, llvm.NewGlobalDecl(".read_scratch", types.IntTySig))

	llvmProgram = p.functions.TranslateToLLVMStack(llvmProgram, tables)

	return llvmProgram
}