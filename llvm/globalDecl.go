package llvm

import (
	"bytes"
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type GlobalDecl struct {
	id				string
	ty				types.Type
}

func NewGlobalDecl(id string, ty types.Type) *GlobalDecl {
	return &GlobalDecl{id, ty}
}

func (g *GlobalDecl) String() string {
	var out bytes.Buffer

	str := fmt.Sprintf("@%s = common global %s %s", g.id, types.TypesToLLVMTypes(g.ty), types.GetLLVMDefaultValue(g.ty))

	out.WriteString(str)
	return out.String()
}

func (g *GlobalDecl) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	return regs
}

func (g *GlobalDecl) GetDef() *LLVMRegister {
	return nil
}

func (g *GlobalDecl) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	return true
}

func (g *GlobalDecl) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}