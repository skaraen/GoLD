package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type TypeDecl struct {
	*token.Token
	id				string
	fields			Expression
}

func NewTypeDecl(token *token.Token, id string, fields Expression) *TypeDecl {
	return &TypeDecl{token, id, fields}
}

func (td *TypeDecl) String() string {
	var out bytes.Buffer

	out.WriteString("type " + td.id + " struct {\n" + td.fields.String() + "};")

	return out.String()
}

func (t *TypeDecl) BuildSymbolTable(errors []*context.CompilerError, structTable *st.SymbolTable[*st.StructEntry]) []*context.CompilerError {
	if st, ok := structTable.Contains(t.id); ok {
		msg := fmt.Sprintf("redeclaration of user defined type (%s) located at (%d,%d)", t.id, st.Line, st.Column)
		semError := context.NewCompilerError(t.Line, t.Column, msg, context.SEMANTICS)
			
		errors = append(errors, semError)
	}

	structEntry := st.NewStructEntry(t.id, t.Token)
	structFields := t.fields.(*Fields)

	for _, field := range(structFields.fieldsList) {
		f := field.(*Decl)

		if fld, ok := structEntry.Fields.Contains(f.id); ok {
			msg := fmt.Sprintf("redeclaration of field (%s) located at (%d,%d)", f.id, fld.Line, fld.Column)
			semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
			errors = append(errors, semError)
			continue
		}

		varEntry := st.NewVarEntry(f.id, f.dType, st.STRUCT, f.Token)
		structEntry.Fields.Insert(f.id, varEntry)
	}
	
	structTable.Insert(t.id, structEntry)

	return errors
}

func (t *TypeDecl) TranslateToLLVMStack(llvmProgram *llvm.LLVMProgram, structEntry *st.StructEntry, tables *st.SymbolTables) *llvm.LLVMProgram {
	var tyList []types.Type

	for _, fieldEntry := range structEntry.Fields.GetOrderedEntries() {
		tyList = append(tyList, fieldEntry.Ty)
	}

	llvmProgram.Structs = append(llvmProgram.Structs, llvm.NewStructDecl(t.id, tyList))

	return llvmProgram
}