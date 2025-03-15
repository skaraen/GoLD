package ast

import (
	"fmt"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Variable struct {
	*token.Token
	identifier		string
	iType			types.Type
}

func NewVariable(token *token.Token, identifier string) *Variable {
	return &Variable{token, identifier, types.StringToType("undefined")}
}

func (v *Variable) String() string {
	return v.identifier
}

func (v *Variable) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	return v.iType
}

func (v *Variable) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	var semError *context.CompilerError
	var msg string 

	if varEntry, exists := funcEntry.Variables.Contains(v.identifier); !exists {
		msg = fmt.Sprintf("undefined variable (%s) located at (%d,%d)", v.identifier, v.Line, v.Column)
		semError = context.NewCompilerError(v.Line, v.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	} else {
		v.iType = varEntry.Ty
	}

	return v.GetType(funcEntry, tables), errors
}

func (v *Variable) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	varEntry,_ := funcEntry.Variables.Contains(v.identifier)
	register := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(varEntry.Name, varEntry.Ty, st.LOCAL, v.Token), false)
	varRegister := llvmProgram.EntryRegMap[varEntry]
	// fmt.Println(varRegister)
	loadInstn := llvm.NewLoadInstn(register, v.iType, varRegister, v.iType)
	// currBlk.Instns = append(currBlk.Instns, loadInstn)
	currBlk.AddInstn(loadInstn)

	if v.iType == types.BoolTySig {
		newRegister := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(varEntry.Name, types.Int1TySig, st.LOCAL, v.Token), false)
		truncInstn := llvm.NewTruncInstn(newRegister, newRegister.GetType(), register, register.GetType())
		// currBlk.Instns = append(currBlk.Instns, truncInstn)
		currBlk.AddInstn(truncInstn)

		return newRegister
	}

	return register
}