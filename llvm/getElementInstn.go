package llvm

import (
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
)

type GetElementInstn struct {
	dest		*LLVMRegister
	operand		LLVMOperand
	field 		string
	tables		*st.SymbolTables
}

func NewGetElementInstn(dest *LLVMRegister, operand LLVMOperand, field string, tables *st.SymbolTables) *GetElementInstn {
	return &GetElementInstn{dest, operand, field, tables}
}

func (g *GetElementInstn) String() string {
	str := ""

	structTy := g.operand.GetType()
	structEntry, _ := g.tables.Structs.Contains(structTy.String())
	idx, _ := structEntry.Fields.GetIndex(g.field)
	str = fmt.Sprintf("%s = getelementptr %%struct.%s, %s %s, i32 0, i32 %d", g.dest.String(), structTy.String(), types.TypesToLLVMTypes(structTy), g.operand.String(), idx)

	return str
}

func (g *GetElementInstn) GetUses() []LLVMOperand {
	var regs []LLVMOperand
	regs = append(regs, g.operand)
	return regs
}

func (g *GetElementInstn) GetDef() *LLVMRegister {
	return g.dest
}

func (g *GetElementInstn) Mem2Reg(defs map[string]LLVMOperand, predLbl string, currBlock *cfg.Block) bool {
	if defEntry := defs[g.operand.GetName()]; defEntry != nil {
		g.operand = defEntry
	}

	return true
}

func (g *GetElementInstn) TranslateToAssembly(tables *st.SymbolTables) []arm.Instruction {
	var armInstns []arm.Instruction

	return armInstns
}