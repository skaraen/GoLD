package llvm

import (
	"golite/cfg"
	"golite/types"
)

type LLVMOperand interface {
	cfg.Operand
	GetType() types.Type
	SetType(types.Type)
	GetName() string
    Mem2RegIsLocal() bool
	GetAssemblyId() string
}