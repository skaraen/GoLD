package llvm

import (
	"fmt"
	"golite/cfg"
	"strings"
)

type PrintInstn struct {
	fstrId		int
	args		[]LLVMOperand			
}

func NewPrintInstn(fstrId int, args	[]LLVMOperand) *PrintInstn {
	return &PrintInstn{fstrId, args}
}

func (p *PrintInstn) String() string {
	var argsStr []string

	for _, arg := range p.args {
		str := arg.GetType().String() + " " + arg.String()
		argsStr = append(argsStr, str)
	}

	printStr := fmt.Sprintf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr%d, i32 0, i32 0), %s)", p.fstrId, strings.Join(argsStr, ", "))

	return printStr
}

func (p *PrintInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, p.args...)
	return regs
}

func (p *PrintInstn) GetDef() *LLVMRegister {
	return nil
}

func (p *PrintInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}
