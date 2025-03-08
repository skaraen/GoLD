package llvm

import (
	"fmt"
	"golite/cfg"
	"golite/types"
)

type ReturnInstn struct {
	retReg		*LLVMRegister
}

func NewReturnInstn(retReg *LLVMRegister) *ReturnInstn {
	return &ReturnInstn{retReg}
}

func (r *ReturnInstn) String() string {
	str := "ret void"
	if r.retReg != nil {
		str = fmt.Sprintf("ret %s %s", types.TypesToLLVMTypes(r.retReg.GetType()), r.retReg.String())
	}
	return str
}

func (r *ReturnInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, r.retReg)
	return regs
}

func (r *ReturnInstn) GetDef() *LLVMRegister {
	return nil
}

func (r *ReturnInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}