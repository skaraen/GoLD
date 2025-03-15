package llvm

import (
	"fmt"
	st "golite/symboltable"
	"golite/types"
)

type LLVMRegister struct {
	id				string
	entry 			*st.VarEntry
	isStackBased	bool
	pReg			*LLVMRegister
}

func NewLLVMRegister(id string, varEntry *st.VarEntry, isStackBased bool) *LLVMRegister {
	return &LLVMRegister{id, varEntry, isStackBased, nil}
}

func (r *LLVMRegister) String() string {
	pred := "%"


	if r.entry == nil {
		fmt.Printf("Nil reg varentry: %s\n", r.id)
	}

	if r.entry.Scope == st.GLOBAL {
		pred = "@"
	}

	return pred + r.id
}

func (r *LLVMRegister) GetType() types.Type {
	return r.entry.Ty
}

func (r *LLVMRegister) SetType(ty types.Type) {
	r.entry.Ty = ty
}

func (r *LLVMRegister) GetName() string {
	return r.id
}

func (r *LLVMRegister) Mem2RegIsLocal() bool {
	return r.isStackBased
}

func (r *LLVMRegister) GetAssemblyId() string {
	if r.pReg != nil {
		return r.pReg.id
	}

	return ""
}