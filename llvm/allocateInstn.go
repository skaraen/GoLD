package llvm

import (
	"bytes"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type AllocateInstn struct {
	dest		*LLVMRegister
	ty			types.Type
}

func NewAllocateInstn(dest *LLVMRegister, ty types.Type) *AllocateInstn {
	dest.isStackBased = true
	return &AllocateInstn{dest, ty}
}

func (a *AllocateInstn) String() string {
	var out bytes.Buffer

	out.WriteString(a.dest.String() + " = alloca " + types.TypesToLLVMTypes(a.ty))

	return out.String()
}

func (a *AllocateInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (a *AllocateInstn) GetDef() *LLVMRegister {
	return a.dest
}

func (a *AllocateInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	var defaultOperand LLVMOperand
	// fmt.Printf("Inside alloca memReg: %s\n", a.String())
	if _, ok := a.ty.(*types.PointerTy); ok {
		defaultOperand = NewLLVMImmediate(0, types.NilTySig)
	} else {
		defaultOperand = NewLLVMImmediate(0, a.ty)
	}

	defs[a.dest.id] = defaultOperand
	return false
}

func (a *AllocateInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}
