package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
)

type MovInstn struct {
	dest		*LLVMRegister
	op			LLVMOperand
}

func NewMovInstn(dest *LLVMRegister, op LLVMOperand) *MovInstn {
	return &MovInstn{dest, op}
}

func (m *MovInstn) String() string {
	string := fmt.Sprintf("mov %s, %s\n", m.op.String(),  m.dest.String())

	return string
}

func (m *MovInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, m.op)
	return regs
}

func (m *MovInstn) GetDef() *LLVMRegister {
	return m.dest
}

func (m *MovInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return true
}

func (m *MovInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}