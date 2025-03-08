package llvm

import (
	"bytes"
	"fmt"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type InstantInstn struct {
	structEntry		*st.StructEntry
	temp			*LLVMRegister
	dest			*LLVMRegister
}

func NewInstantInstn(structEntry *st.StructEntry, temp *LLVMRegister, dest *LLVMRegister) *InstantInstn {
	return &InstantInstn{structEntry, temp, dest}
}

func (i *InstantInstn) String() string {
	var out bytes.Buffer

	mallocStr := fmt.Sprintf("%s = call i8* @malloc(i32 %d)", i.temp.String(), len(i.structEntry.Fields.GetTable()) * 8)
	bitcastStr := fmt.Sprintf("%s = bitcast i8* %s to %s", i.dest.String(), i.temp.String(), types.TypesToLLVMTypes(types.StringToType(i.structEntry.Name)))

	out.WriteString(mallocStr + "\n")
	out.WriteString(bitcastStr)

	return out.String()
}

func (i *InstantInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, i.temp)
	return regs
}

func (i *InstantInstn) GetDef() *LLVMRegister {
	return i.dest
}

func (i *InstantInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}