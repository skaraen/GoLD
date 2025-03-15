package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	op "golite/operators"
	st "golite/symboltable"
	"golite/types"
)

type UnOpInstn struct {
	dest		*LLVMRegister
	operator	op.Operator
	operand 	LLVMOperand
	opTy 		types.Type
}

func NewUnOpInstn(dest *LLVMRegister, operator op.Operator, operand LLVMOperand, opTy types.Type) *UnOpInstn {
	return &UnOpInstn{dest, operator, operand, opTy}
}

func (u *UnOpInstn) String() string {
	str := ""

	if u.operator == op.MINUS {
		str = fmt.Sprintf("%s = sub %s 0, %s", u.dest.String(), types.TypesToLLVMTypes(u.operand.GetType()), u.operand.String())
	} else if u.operator == op.NOT {
		str = fmt.Sprintf("%s = xor %s %s, 1", u.dest.String(), types.TypesToLLVMTypes(u.operand.GetType()), u.operand.String())
	}

	return str
}

func (u *UnOpInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, u.operand)
	return regs
}

func (u *UnOpInstn) GetDef() *LLVMRegister {
	return u.dest
}

func (u *UnOpInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[u.operand.GetName()]; defEntry != nil {
		u.operand = defEntry
	}
	return true
}

func (u *UnOpInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}