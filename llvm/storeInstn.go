package llvm

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/types"
)

type StoreInstn struct {
	valTy	types.Type	 
	destTy	types.Type
	val		LLVMOperand
	dest	*LLVMRegister
}

func NewStoreInstn(dest *LLVMRegister, destTy types.Type, val LLVMOperand, valTy types.Type) *StoreInstn {
	return &StoreInstn{valTy, destTy, val, dest}
}

func (s *StoreInstn) String() string {
	var out bytes.Buffer

	str := fmt.Sprintf("store %s %s, %s* %s", types.TypesToLLVMTypes(s.valTy), s.val.String(), types.TypesToLLVMTypes(s.destTy), s.dest.String())

	out.WriteString(str)
	return out.String()
}

func (s *StoreInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, s.val)
	return regs
}

func (s *StoreInstn) GetDef() *LLVMRegister {
	return s.dest
}

func (l *StoreInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}