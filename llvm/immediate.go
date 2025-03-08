package llvm

import (
	"golite/types"
	"strconv"
)

type LLVMImmediate struct {
	val 	int64
	iTy 	types.Type
}

func NewLLVMImmediate(val int64, iTy types.Type) *LLVMImmediate {
	return &LLVMImmediate{val, iTy}
}

func (i *LLVMImmediate) GetType() types.Type {
	return i.iTy
}

func (i *LLVMImmediate) String() string {
	return strconv.FormatInt(i.val, 10)
}

func (i *LLVMImmediate) GetName() string {
	return ""
}

func (i *LLVMImmediate) Mem2RegIsLocal() bool {
	return false
}