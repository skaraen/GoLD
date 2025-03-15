package llvm

import (
	"fmt"
	"golite/cfg"
	"golite/types"
)

type BitcastInstn struct {
	dest 	*LLVMRegister
	op		*LLVMRegister
}

func NewBitcastInstn(dest *LLVMRegister, op *LLVMRegister) *BitcastInstn {
	return &BitcastInstn{dest, op}
}

func (b *BitcastInstn) String() string {
	bitcastStr := fmt.Sprintf("%s = bitcast %s %s to %s", b.dest.String(), types.TypesToLLVMTypes(b.op.GetType()), b.op.String(), types.TypesToLLVMTypes(b.dest.GetType()))
	return bitcastStr
}

func (b *BitcastInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, b.op)
	return regs
}

func (b *BitcastInstn) GetDef() *LLVMRegister {
	return b.dest
}

func (b *BitcastInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[b.op.GetName()]; defEntry != nil {
		b.op = defEntry.(*LLVMRegister)
	}

	return true
}