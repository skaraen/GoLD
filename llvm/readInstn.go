package llvm

import (
	"fmt"
	"golite/cfg"
	st "golite/symboltable"
)

type ReadInstn struct {
	destEntry		*st.VarEntry			
}

func NewReadInstn(destEntry *st.VarEntry) *ReadInstn {
	return &ReadInstn{destEntry}
}

func (r *ReadInstn) String() string {
	pred := "%"
	if r.destEntry.Scope == st.GLOBAL {
		pred = "@"
	} 
	readStr := fmt.Sprintf("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %s%s", pred, r.destEntry.Name)

	return readStr
}

func (r *ReadInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (r *ReadInstn) GetDef() *LLVMRegister {
	return nil
}

func (r *ReadInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}