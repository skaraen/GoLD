package ast

import (
	"bytes"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
)

type UserTypes struct {
	*token.Token
	typesDeclList		[]*TypeDecl
}

func NewUserTypes(token *token.Token, typesDeclList []*TypeDecl) *UserTypes {
	return &UserTypes{token, typesDeclList}
}

func (u *UserTypes) String() string {
	var out bytes.Buffer

	for _, td := range (u.typesDeclList) {
		out.WriteString(td.String() + "\n")
	}

	return out.String()
}

func (u *UserTypes) BuildSymbolTable(errors []*context.CompilerError, table *st.SymbolTable[*st.StructEntry]) []*context.CompilerError {
	for _, typeDecl := range (u.typesDeclList) {
		errors = typeDecl.BuildSymbolTable(errors, table)
	}

	return errors
}

func (u *UserTypes) TranslateToLLVMStack(llvmProgram *llvm.LLVMProgram, tables *st.SymbolTables) *llvm.LLVMProgram {
	for _, typeDecl := range (u.typesDeclList) {
		structEntry, _ := tables.Structs.Contains(typeDecl.id) 
		llvmProgram = typeDecl.TranslateToLLVMStack(llvmProgram, structEntry, tables)
	}

	return llvmProgram
}