package ast

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Instantiation struct {
	*token.Token
	id string
}

func NewInstantation(token *token.Token, id string) *Instantiation {
	return &Instantiation{token, id}
}

func (ins *Instantiation) String() string {
	var out bytes.Buffer

	out.WriteString("new "+ ins.id)

	return out.String()
}

func (ins *Instantiation) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return types.StringToType(ins.id)
}

func (ins *Instantiation) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	if _, exists := tables.Structs.Contains(ins.id); !exists {
		msg := fmt.Sprintf("trying to instantiate an undefined user defined type (%s)", ins.id);
		semError := context.NewCompilerError(ins.Line, ins.Column, msg, context.SEMANTICS)

		errors = append(errors, semError)
	}

	return types.StringToType(ins.id), errors
}

func (ins *Instantiation) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	mallocPtr := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(ins.String(), types.Int8PtrTySig, st.LOCAL, ins.Token), false)
	dest := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(ins.String(), ins.GetType(funcEntry, tables), st.LOCAL, ins.Token), false)

	structEntry, _ := tables.Structs.Contains(ins.id)
	// currBlk.Instns = append(currBlk.Instns, llvm.NewMallocInstn(structEntry, mallocPtr))
	currBlk.AddInstn(llvm.NewMallocInstn(structEntry, mallocPtr))
	// currBlk.Instns = append(currBlk.Instns, llvm.NewBitcastInstn(dest, mallocPtr))
	currBlk.AddInstn(llvm.NewBitcastInstn(dest, mallocPtr))
	
	return dest
}