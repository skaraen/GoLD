package llvm

import (
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type RegisterAlloc struct {
	regs		[]*LLVMRegister
	activeSet	map[*LLVMRegister]*LLVMRegister
}

func NewRegisterAlloc() *RegisterAlloc {
	regNames := []string{	"x2", "x3", "x4", "x5", "x6", "x7",
							"x9", "x10", "x11", "x12", "x13", "x14", "x15",
							"x19", "x20", "x21", "x22", "x23", "x24", "x25",
							"x26", "x27", "x28", "spill_1", "spill_2"}
	
	var regs []*LLVMRegister
	for _, name := range regNames {
		reg := NewLLVMRegister(name, st.NewVarEntry(name, types.IntTySig, st.LOCAL, &token.Token{}), false)
		regs = append(regs, reg)
	}

	activeSet := make(map[*LLVMRegister]*LLVMRegister)

	return &RegisterAlloc{regs, activeSet}
}

func (r *RegisterAlloc) CheckExpiry(ts int, varTable map[*LLVMRegister]*LiveRangeEntry) {
	for pReg, vReg := range r.activeSet {
		if varTable[vReg].End < ts {
			delete(r.activeSet, pReg)
		}
	}
}

func (r *RegisterAlloc) GetRegister(vReg *LLVMRegister) *LLVMRegister {
	for _, pReg := range r.regs {
		if r.activeSet[pReg] == nil {
			r.activeSet[pReg] = vReg
			return pReg
		}
	}

	return nil
}

