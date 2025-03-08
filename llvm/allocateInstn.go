package llvm

import (
	"bytes"
	"golite/cfg"
	"golite/types"
)

type AllocateInstn struct {
	dest		*LLVMRegister
	ty			types.Type
}

func NewAllocateInstn(dest *LLVMRegister, ty types.Type) *AllocateInstn {
	dest.isStackBased = true
	return &AllocateInstn{dest, ty}
}

func (a *AllocateInstn) String() string {
	var out bytes.Buffer

	out.WriteString(a.dest.String() + " = alloca " + types.TypesToLLVMTypes(a.ty))

	return out.String()
}

func (a *AllocateInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (a *AllocateInstn) GetDef() *LLVMRegister {
	return a.dest
}

func (a *AllocateInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}