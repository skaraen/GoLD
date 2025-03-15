package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type DeleteInstn struct {
	op			*LLVMRegister
}

func NewDeleteInstn(op *LLVMRegister) *DeleteInstn {
	return &DeleteInstn{op}
}

func (d *DeleteInstn) String() string {
	// bitcastStr := fmt.Sprintf("%s = bitcast %s %s to i8*", d.temp2.String(), types.TypesToLLVMTypes(d.temp1.entry.Ty), d.temp1.String())
	deleteStr := fmt.Sprintf("call void @free(%s %s)", types.TypesToLLVMTypes(d.op.GetType()), d.op.String())

	return deleteStr
}

func (d *DeleteInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, d.op)
	return regs
}

func (d *DeleteInstn) GetDef() *LLVMRegister {
	return nil
}

func (d *DeleteInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[d.op.GetName()]; defEntry != nil {
		d.op = defEntry.(*LLVMRegister)
	}

	return true
}

func (d *DeleteInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}