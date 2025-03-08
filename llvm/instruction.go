package llvm

import "golite/cfg"

type LLVMInstruction interface {
	cfg.Instruction
	GetUses() []LLVMOperand
    GetDef() *LLVMRegister
    Mem2Reg(map[string]LLVMOperand, string, *cfg.Block) bool
}