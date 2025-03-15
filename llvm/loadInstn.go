package llvm

import (
	"bytes"
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type LoadInstn struct {
	desTy	types.Type	 
	opTy	types.Type
	dest	*LLVMRegister
	op		*LLVMRegister
}

func NewLoadInstn(dest *LLVMRegister, destTy types.Type, op *LLVMRegister, opTy types.Type) *LoadInstn {
	return &LoadInstn{destTy, opTy, dest, op}
}

func (l *LoadInstn) String() string {
	var out bytes.Buffer

	str := fmt.Sprintf("%s = load %s, %s* %s", l.dest.String(), types.TypesToLLVMTypes(l.desTy), types.TypesToLLVMTypes(l.opTy), l.op.String())

	out.WriteString(str)
	return out.String()
}

func (l *LoadInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, l.op)
	return regs
}

func (l *LoadInstn) GetDef() *LLVMRegister {
	return l.dest
}

func (l *LoadInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if l.op.isStackBased {
		defs[l.dest.id] = defs[l.op.id]
		return false
	}
	return true
}

func (l *LoadInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}