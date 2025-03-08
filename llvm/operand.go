package llvm

import (
	"golite/cfg"
	"golite/types"
)

type LLVMOperand interface {
	cfg.Operand
	GetType() types.Type
	GetName() string
    Mem2RegIsLocal() bool
}