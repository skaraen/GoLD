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

type Function struct {
	*token.Token
	id				string
	params			Expression
	returnType		types.Type
	declarations	Statement
	statements		Statement
}

func NewFunction(token *token.Token, id string, params Expression, returnType types.Type, declarations Statement, statements Statement) *Function {
	return &Function{token, id, params, returnType, declarations, statements}
}

func (f *Function) String() string {
	var out bytes.Buffer

	out.WriteString("func " + f.id + f.params.String())

	if f.returnType.String() != "nil" {
		out.WriteString(" " + f.returnType.String())
	}

	out.WriteString(" {\n")
	out.WriteString(f.declarations.String())
	out.WriteString(f.statements.String())

	out.WriteString("}\n")

	return out.String()
}

func (f *Function) BuildSymbolTable(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	if fn, exists := tables.Funcs.Contains(f.id); exists {
		msg := fmt.Sprintf("redeclaration of function (%s) located at (%d,%d)", f.id, fn.Line, fn.Column)
		semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
		errors = append(errors, semError)
	}
	
	funcEntry := st.NewFuncEntry(f.id, f.returnType, f.Token)
	paramsList := f.params.(*Params).paramsList

	for _, pExp := range (paramsList) {
		param := pExp.(*Decl)

		if p, exists := funcEntry.Variables.Contains(param.id); exists {
			msg := fmt.Sprintf("redeclaration of parameter (%s) located at (%d,%d)", param.id, p.Line, p.Column)
			semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
				
			errors = append(errors, semError)
			continue
		}

		varEntry := st.NewVarEntry(param.id, param.dType, st.LOCAL, param.Token)
		funcEntry.Signature = append(funcEntry.Signature, param.dType)
		funcEntry.Parameters.Insert(param.id, varEntry)
	}

	declrList := f.declarations.(*Declarations).declrList
	funcEntry.Variables.SetParent(funcEntry.Parameters)

	for _, declr := range (declrList) {
		ids := declr.ids.(*Ids)

		for _, id := range(ids.idsList) {
			if st, exists := tables.Structs.Contains(id); exists {
				msg := fmt.Sprintf("redeclaration of user defined type (%s) located at (%d,%d)", id, st.Line, st.Column)
				semError := context.NewCompilerError(ids.Line, ids.Column, msg, context.SEMANTICS)
				
				errors = append(errors, semError)
				continue
			}

			if v, exists := funcEntry.Variables.Contains(id); exists {
				msg := fmt.Sprintf("redeclaration of variable (%s) located at (%d,%d)", id, v.Line, v.Column)
				semError := context.NewCompilerError(ids.Line, ids.Column, msg, context.SEMANTICS)
				
				errors = append(errors, semError)
				continue
			}
			
			varEntry := st.NewVarEntry(id, declr.declType, st.LOCAL, ids.Token)
			funcEntry.Variables.Insert(id, varEntry)
		}
	}

	funcEntry.Parameters.SetParent(tables.Globals)
	tables.Funcs.Insert(f.id, funcEntry)
	return errors
}

func (f *Function) TypeCheck(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError { 
	if funcEntry, exists := tables.Funcs.Contains(f.id); exists {
		errors = f.statements.TypeCheck(errors, funcEntry, tables)
	}
	
	return errors
}

func (f *Function) TranslateToLLVMStack(llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *llvm.LLVMProgram {
	labels := llvmProgram.GenerateLabelsBatch(f.id, 2)
	entryBlk, exitBlk := cfg.NewBlock(labels[0], labels[1])
	llvmProgram.Funcs[f.id] = entryBlk

	currBlock := entryBlk

	var retReg *llvm.LLVMRegister
	if funcEntry.RetTy != types.NilTySig {
		retVarEntry := st.NewVarEntry("returnVal", funcEntry.RetTy, st.LOCAL, funcEntry.Token)
		retReg = llvm.NewLLVMRegister("_ret_val", retVarEntry, true)
		llvmProgram.RetRegisters[f.id] = retReg
		llvmProgram.EntryRegMap[retVarEntry] = retReg
		// currBlock.Instns = append(currBlock.Instns, llvm.NewAllocateInstn(retReg, funcEntry.RetTy))
		currBlock.AddInstn(llvm.NewAllocateInstn(retReg, funcEntry.RetTy))
		if funcEntry.Name == "main" {
			// currBlock.Instns = append(currBlock.Instns, llvm.NewStoreInstn(retReg, retReg.GetType(), llvm.NewLLVMImmediate(0, types.IntTySig), types.IntTySig))	
			currBlock.AddInstn(llvm.NewStoreInstn(retReg, retReg.GetType(), llvm.NewLLVMImmediate(0, types.IntTySig), types.IntTySig))
		}
	}

	for _, paramEntry := range funcEntry.Parameters.GetOrderedEntries() {
		paramReg := llvm.NewLLVMRegister(paramEntry.Name, paramEntry, false)
		localRegName := fmt.Sprintf("_P_%s", paramEntry.Name)
		localParamReg := llvm.NewLLVMRegister(localRegName, paramEntry, true)
		llvmProgram.ParamRegisters[f.id] = append(llvmProgram.ParamRegisters[f.id], paramReg)
		llvmProgram.AddEntryRegister(paramEntry, localParamReg)
		// currBlock.Instns = append(currBlock.Instns, llvm.NewAllocateInstn(localParamReg, param.Ty))
		currBlock.AddInstn(llvm.NewAllocateInstn(localParamReg, paramEntry.Ty))
		// currBlock.Instns = append(currBlock.Instns, llvm.NewStoreInstn(localParamReg, localParamReg.GetType(), paramReg, paramReg.GetType()))
		currBlock.AddInstn(llvm.NewStoreInstn(localParamReg, localParamReg.GetType(), paramReg, paramReg.GetType()))
	}

	currBlock = f.declarations.(*Declarations).TranslateToLLVMStack(currBlock, exitBlk, llvmProgram, funcEntry, tables)

	currBlock = f.statements.TranslateToLLVMStack(currBlock, exitBlk, llvmProgram, funcEntry, tables)
	// currBlock.Succs = append(currBlock.Succs, exitBlk)
	currBlock.AddSucc(exitBlk)
	exitBlk.AddPred(currBlock)
	// currBlock.Instns = append(currBlock.Instns, llvm.NewJumpInstn(exitBlk))
	currBlock.AddInstn(llvm.NewJumpInstn(exitBlk))

	if funcEntry.RetTy != types.NilTySig {
		tempReg := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry("tempRet", funcEntry.RetTy, st.LOCAL, funcEntry.Token), false)
		// exitBlk.Instns = append(exitBlk.Instns, llvm.NewLoadInstn(tempReg, tempReg.GetType(), retReg, retReg.GetType()))
		exitBlk.AddInstn(llvm.NewLoadInstn(tempReg, tempReg.GetType(), retReg, retReg.GetType()))
		// exitBlk.Instns = append(exitBlk.Instns, llvm.NewReturnInstn(tempReg))
		exitBlk.AddInstn(llvm.NewReturnInstn(tempReg))
	} else {
		// exitBlk.Instns = append(exitBlk.Instns, llvm.NewReturnInstn(nil))
		exitBlk.AddInstn(llvm.NewReturnInstn(nil))
	}
	
	return llvmProgram
}