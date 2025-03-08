package llvm

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/types"
	"strings"
)

type StructDecl struct {
	id				string
	fieldTyList		[]types.Type
}

func NewStructDecl(id string, fieldTyList []types.Type) *StructDecl {
	return &StructDecl{id, fieldTyList}
}

func (s *StructDecl) String() string {
	var out bytes.Buffer

	var fieldTyStr []string
	
	for _, fieldTy := range(s.fieldTyList) {
		fieldTyStr = append(fieldTyStr, types.TypesToLLVMTypes(fieldTy))
	}

	str := fmt.Sprintf("%%struct.%s = type {%s}", s.id, strings.Join(fieldTyStr, ", "))

	out.WriteString(str)
	return out.String()
}

func (s *StructDecl) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (s *StructDecl) GetDef() *LLVMRegister {
	return nil
}

func (s *StructDecl) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return false
}