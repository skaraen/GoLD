package llvm

import (
	"fmt"
	"golite/arm"
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
		str := types.TypesToLLVMTypes(arg.GetType()) + " " + arg.String()
		argsStr = append(argsStr, str)
	}

	var fnCallStr string
	if i.dest == nil {
		fnCallStr = fmt.Sprintf("call %s @%s(%s)", types.TypesToLLVMTypes(i.invEntry.RetTy), i.invEntry.Name, strings.Join(argsStr, ", "))
	} else {
		fnCallStr = fmt.Sprintf("%s = call %s @%s(%s)", i.dest.String(), types.TypesToLLVMTypes(i.invEntry.RetTy), i.invEntry.Name, strings.Join(argsStr, ", "))
	}

	return fnCallStr
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
	for idx, argOp := range i.args {
		if defEntry := defs[argOp.GetName()]; defEntry != nil {
			i.args[idx] = defs[argOp.GetName()]
		}
	} 

	return true
}

func (i *InvInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}