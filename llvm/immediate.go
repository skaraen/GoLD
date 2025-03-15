package llvm

import (
	"fmt"
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

func (i *LLVMImmediate) SetType(ty types.Type) {
	i.iTy = ty
}

func (i *LLVMImmediate) String() string {
	if _, ok := i.iTy.(*types.PointerTy); ok {
		return "null"
	}
	return strconv.FormatInt(i.val, 10)
}

func (i *LLVMImmediate) GetName() string {
	return ""
}

func (i *LLVMImmediate) Mem2RegIsLocal() bool {
	return false
}

func (i *LLVMImmediate) GetAssemblyId() string {
	return fmt.Sprintf("#%d", i.val)
}