package llvm

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/types"
)

type DeleteInstn struct {
	temp1			*LLVMRegister
	temp2			*LLVMRegister
}

func NewDeleteInstn(temp1 *LLVMRegister, temp2 *LLVMRegister) *DeleteInstn {
	return &DeleteInstn{temp1, temp2}
}

func (d *DeleteInstn) String() string {
	var out bytes.Buffer

	bitcastStr := fmt.Sprintf("%s = bitcast %s %s to i8*", d.temp2.String(), types.TypesToLLVMTypes(d.temp1.entry.Ty), d.temp1.String())
	deleteStr := fmt.Sprintf("call void @free(i8* %s)", d.temp2.String())

	out.WriteString(bitcastStr + "\n")
	out.WriteString(deleteStr)

	return out.String()
}

func (d *DeleteInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, d.temp1)
	regs = append(regs, d.temp2)
	return regs
}

func (d *DeleteInstn) GetDef() *LLVMRegister {
	return nil
}

func (d *DeleteInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}