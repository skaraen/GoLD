package llvm

import (
	"fmt"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
	"strings"
)

type InvInstn struct {
	invEntry		*st.FuncEntry
	dest			*LLVMRegister
	args			[]LLVMOperand
}

func NewInvInstn(invEntry *st.FuncEntry, dest *LLVMRegister, args []LLVMOperand) *InvInstn {
	return &InvInstn{invEntry, dest, args}
}

func (i *InvInstn) String() string {
	var argsStr []string

	for _, arg := range i.args {
		str := arg.GetType().String() + " " + arg.String()
		argsStr = append(argsStr, str)
	}

	fnCalStr := fmt.Sprintf("%s = call %s @%s(%s)", i.dest.String(), types.TypesToLLVMTypes(i.invEntry.RetTy), i.invEntry.Name, strings.Join(argsStr, ", "))

	return fnCalStr
}

func (i *InvInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, i.args...)
	return regs
}

func (i *InvInstn) GetDef() *LLVMRegister {
	return i.dest
}

func (i *InvInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}