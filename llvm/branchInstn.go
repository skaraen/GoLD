package llvm

import (
	"fmt"
	"golite/cfg"
)

type BranchInstn struct {
	indicator		*LLVMRegister
	trueBlk			*cfg.Block
	falseBlk		*cfg.Block
}

func NewBranchInstn(indicator *LLVMRegister, trueBlk *cfg.Block, falseBlk *cfg.Block) *BranchInstn {
	return &BranchInstn{indicator, trueBlk, falseBlk}
}

func (b *BranchInstn) String() string {
	brStr := fmt.Sprintf("br i1 %s, label %%%s, label %%%s", b.indicator.String(), b.trueBlk.Label, b.falseBlk.Label)
	return brStr
}

func (b *BranchInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, b.indicator)
	return regs
}

func (b *BranchInstn) GetDef() *LLVMRegister {
	return nil
}

func (b *BranchInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}