package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
)

type MallocInstn struct {
	structEntry		*st.StructEntry
	dest			*LLVMRegister
}

func NewMallocInstn(structEntry *st.StructEntry, dest *LLVMRegister) *MallocInstn {
	return &MallocInstn{structEntry, dest}
}

func (i *MallocInstn) String() string {
	mallocStr := fmt.Sprintf("%s = call i8* @malloc(i32 %d)", i.dest.String(), len(i.structEntry.Fields.GetTable()) * 8)
	// bitcastStr := fmt.Sprintf("%s = bitcast i8* %s to %s", i.dest.String(), i.temp.String(), types.TypesToLLVMTypes(types.StringToType(i.structEntry.Name)))
	return mallocStr
}

func (i *MallocInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (i *MallocInstn) GetDef() *LLVMRegister {
	return i.dest
}

func (i *MallocInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return true
}

func (i *MallocInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}