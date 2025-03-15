package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	"golite/operators"
	st "golite/symboltable"
	"golite/types"
)

type ReturnInstn struct {
	op		LLVMOperand
}

func NewReturnInstn(op LLVMOperand) *ReturnInstn {
	return &ReturnInstn{op}
}

func (r *ReturnInstn) String() string {
	str := "ret void"
	if r.op.GetType() != types.NilTySig {
		str = fmt.Sprintf("ret %s %s", types.TypesToLLVMTypes(r.op.GetType()), r.op.String())
	}
	return str
}

func (r *ReturnInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, r.op)
	return regs
}

func (r *ReturnInstn) GetDef() *LLVMRegister {
	return nil
}

func (r *ReturnInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[r.op.GetName()]; defEntry != nil {
		r.op = defEntry
	}
	return true
}

func (r *ReturnInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	armInstns = append(armInstns, arm.NewLoadPairInstn("x29", "x30", "sp"))
	armInstns = append(armInstns, arm.NewBinOpInstn("sp", "sp", "32", operators.PLUS))
	armInstns = append(armInstns,  arm.NewMovInstn("x29", "sp"))
	armInstns = append(armInstns, arm.NewLoadInstn("x0", r.op.GetAssemblyId()))
	armInstns = append(armInstns, arm.NewReturnInstn())

	return armInstns
}