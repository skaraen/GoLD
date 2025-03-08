package llvm

import (
	st "golite/symboltable"
	"golite/types"
)

type LLVMRegister struct {
	id				string
	entry 			*st.VarEntry
	isStackBased	bool
}

func NewLLVMRegister(id string, varEntry *st.VarEntry) *LLVMRegister {
	return &LLVMRegister{id, varEntry, false}
}

func (r *LLVMRegister) String() string {
	pred := "%"

	if r.entry.Scope == st.GLOBAL {
		pred = "@"
	}

	return pred + r.id
}

func (r *LLVMRegister) GetType() types.Type {
	return r.entry.Ty
}

func (r *LLVMRegister) GetName() string {
	return r.id
}

func (r *LLVMRegister) Mem2RegIsLocal() bool {
	return r.isStackBased
}