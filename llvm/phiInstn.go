package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
	"strings"
)

type PhiInstn struct {
	dest 		*LLVMRegister
	varEntry	*st.VarEntry
	valMap		map[*cfg.Block]LLVMOperand			
}

func NewPhiInstn(dest *LLVMRegister, varEntry *st.VarEntry, valMap map[*cfg.Block]LLVMOperand) *PhiInstn {
	return &PhiInstn{dest, varEntry, valMap}
}

func (p *PhiInstn) String() string {
	valStr := make([]string, 0)
	
	for block, operand := range p.valMap {
		str := fmt.Sprintf("[%s, %%%s]", operand.String(), block.Label)
		valStr = append(valStr, str)
	}

	instn := fmt.Sprintf("%s = phi %s %s", p.dest.String(), types.TypesToLLVMTypes(p.dest.GetType()), strings.Join(valStr, ", "))

	return instn
}

func (p *PhiInstn) AddVal(block *cfg.Block, operand LLVMOperand) {
	p.valMap[block] = operand
}

func (p *PhiInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (p *PhiInstn) GetDef() *LLVMRegister {
	return p.dest
}

func (p *PhiInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return true
}

func (p *PhiInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}

