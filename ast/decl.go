package ast

import (
	"bytes"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Decl struct {
	*token.Token
	id 			string
	dType		types.Type
}

func NewDecl(token *token.Token, id string, dType types.Type) *Decl {
	return &Decl{token, id, dType}
}

func (d *Decl) String() string {
	var out bytes.Buffer

	out.WriteString(d.id + " ")
	if _, ok := d.dType.(*types.PointerTy); ok {
		out.WriteString("*")
	}
	out.WriteString(d.dType.String())

	return out.String()
}

func (d *Decl) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("undefined")
}

func (d *Decl) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return d.GetType(funcEntry, tables), errors
}

func (d *Decl) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	return nil
}