package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type TruncInstn struct {
	dest		*LLVMRegister
	destTy		types.Type
	op			LLVMOperand
	opTy		types.Type
}

func NewTruncInstn(dest *LLVMRegister, destTy types.Type, op LLVMOperand, opTy types.Type) *TruncInstn {
	return &TruncInstn{dest, destTy, op, opTy}
}

func (t *TruncInstn) String() string {
	string := fmt.Sprintf("%s = trunc %s %s to %s", t.dest.String(), types.TypesToLLVMTypes(t.opTy), t.op.String(), types.TypesToLLVMTypes(t.destTy))
	return string
}

func (t *TruncInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, t.op)
	return regs
}

func (t *TruncInstn) GetDef() *LLVMRegister {
	return t.dest
}

func (t *TruncInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[t.op.GetName()]; defEntry != nil {
		t.op = defEntry
	}
	return true
}

func (t *TruncInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}