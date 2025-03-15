package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	op "golite/operators"
	st "golite/symboltable"
	"golite/types"
)

type BinOpInstn struct {
	dest		*LLVMRegister
	operator	op.Operator
	op1			LLVMOperand
	op2 		LLVMOperand
	opTy 		types.Type
}

func NewBinOpInstn(dest *LLVMRegister, operator op.Operator, op1 LLVMOperand, op2 LLVMOperand, opTy types.Type) *BinOpInstn {
	return &BinOpInstn{dest, operator, op1, op2, opTy}
}

func (b *BinOpInstn) String() string {
	str := ""

	if op.IsArithmeticOp(b.operator) || op.IsLogicalOp(b.operator) {
		str = fmt.Sprintf("%s = %s %s %s, %s", b.dest.String(), op.OpToLLVM(b.operator), types.TypesToLLVMTypes(b.op1.GetType()), b.op1.String(), b.op2.String())
	} else if op.IsComparisonOp(b.operator) || op.IsEqualityOp(b.operator) {
		str = fmt.Sprintf("%s = icmp %s %s %s, %s", b.dest.String(), op.OpToLLVM(b.operator), types.TypesToLLVMTypes(b.op1.GetType()), b.op1.String(), b.op2.String())
	}

	return str
}

func (b *BinOpInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, b.op1)
	regs = append(regs, b.op2)
	return regs
}

func (b *BinOpInstn) GetDef() *LLVMRegister {
	return b.dest
}

func (b *BinOpInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[b.op1.GetName()]; defEntry != nil {
		b.op1 = defEntry
	}

	if defEntry := defs[b.op2.GetName()]; defEntry != nil {
		b.op2 = defEntry
	}

	return true
}

func (b *BinOpInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}