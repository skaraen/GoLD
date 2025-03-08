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

type Ids struct {
	*token.Token
	idsList 	[]string
}

func NewIds(token *token.Token, idsList []string) *Ids {
	return &Ids{token, idsList}
}

func (ids *Ids) String() string {
	var out bytes.Buffer

	for i, id := range (ids.idsList) {
		out.WriteString(id)
		if (i < len(ids.idsList) - 1) {
			out.WriteString(", ")
		}
	}

	return out.String()
}

func (ids *Ids) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType("undefined")
}

func (ids *Ids) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	return ids.GetType(funcEntry, tables), errors
}

func (ids *Ids) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	return nil
}