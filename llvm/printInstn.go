package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	"golite/operators"
	st "golite/symboltable"
	"golite/types"
	"strconv"
	"strings"
)

type PrintInstn struct {
	fstrId		int
	strLit		string
	args		[]LLVMOperand			
}

func NewPrintInstn(fstrId int, strLit string, args []LLVMOperand) *PrintInstn {
	return &PrintInstn{fstrId, strLit, args}
}

func (p *PrintInstn) String() string {
	var argsStr []string

	for _, arg := range p.args {
		str := types.TypesToLLVMTypes(arg.GetType()) + " " + arg.String()
		argsStr = append(argsStr, str)
	}
	unQuotedStr,_ := strconv.Unquote(p.strLit)
	unQuotedLen := len(unQuotedStr)

	argsAppend := ""
	if len(p.args) > 0 {
		argsAppend = ", " + strings.Join(argsStr, ", ")
	}

	printStr := fmt.Sprintf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([%d x i8], [%d x i8]* @.fstr%d, i32 0, i32 0)%s)", unQuotedLen, unQuotedLen, p.fstrId, argsAppend)

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
	for idx, argOp := range p.args {
		if defEntry := defs[argOp.GetName()]; defEntry != nil {
			p.args[idx] = defEntry
		}
	}

	return true
}

func (p *PrintInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	
	fmtStrId := fmt.Sprintf(".FMTSTR%d", p.fstrId)
	printReg := "x10"

	armInstns = append(armInstns, arm.NewAdrpInsntn(printReg, fmtStrId))
	armInstns = append(armInstns, arm.NewBinOpInstn(printReg, printReg, ":lo12:" + fmtStrId, operators.PLUS))
	armInstns = append(armInstns, arm.NewMovInstn("x0", printReg))

	if len(p.args) > 0 {
		if _, ok := p.args[0].(*LLVMRegister); ok {
			armInstns = append(armInstns, arm.NewMovInstn("x1", p.args[0].(*LLVMRegister).pReg.id))
		} else {
			armInstns = append(armInstns, arm.NewMovInstn("x1", fmt.Sprintf("#%s", p.args[0].String())))
		}
	}

	return armInstns
}
