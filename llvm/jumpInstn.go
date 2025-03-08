package llvm

import (
	"fmt"
	"golite/cfg"
)

type JumpInstn struct {
	targetBlk		*cfg.Block
}

func NewJumpInstn(targetBlk	*cfg.Block) *JumpInstn {
	return &JumpInstn{targetBlk}
}

func (j *JumpInstn) String() string {
	jumpStr := fmt.Sprintf("br label %%%s", j.targetBlk.Label)
	return jumpStr
}

func (j *JumpInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (j *JumpInstn) GetDef() *LLVMRegister {
	return nil
}

func (j *JumpInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}