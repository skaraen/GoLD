package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	"golite/operators"
	st "golite/symboltable"
)

type ReadInstn struct {
	destEntry		*LLVMRegister			
}

func NewReadInstn(destEntry *LLVMRegister) *ReadInstn {
	return &ReadInstn{destEntry}
}

func (r *ReadInstn) String() string {
	readStr := fmt.Sprintf("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %s)", r.destEntry.String())

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
	return true
}

func (r *ReadInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction
	readReg := "x11"

	armInstns = append(armInstns, arm.NewAdrpInsntn(readReg, ".READ"))
	armInstns = append(armInstns, arm.NewBinOpInstn(readReg, readReg, ":lo12:.READ", operators.PLUS))
	armInstns = append(armInstns, arm.NewMovInstn("x0", readReg))
	armInstns = append(armInstns, arm.NewMovInstn("x1", r.destEntry.id))

	return armInstns
}