package llvm

import (
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
)

type LLVMInstruction interface {
	cfg.Instruction
	GetUses() []LLVMOperand
    GetDef() *LLVMRegister
    Mem2Reg(map[string]LLVMOperand, string, *cfg.Block) bool
	TranslateToAssembly(tables * st.SymbolTables) []arm.Instruction
}