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
	"strings"
)

type AssignStmt struct {
	*token.Token
	lValue 			Expression
	rExpression		Expression
}

func NewAssignStmt(token *token.Token, lValue Expression, rExpression Expression) *AssignStmt {
	return &AssignStmt{token, lValue, rExpression}
}

func (as *AssignStmt) String() string {
	var out bytes.Buffer

	out.WriteString(as.lValue.String() + " = " + as.rExpression.String() + ";")

	return out.String()
}

func (as *AssignStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var leftInfType, rightInfType types.Type
	
	leftInfType, errors = as.lValue.TypeCheck(errors, funcEntry, tables)
	rightInfType, errors = as.rExpression.TypeCheck(errors, funcEntry, tables)

	if leftInfType == types.StringToType("bool") || leftInfType == types.StringToType("int") || leftInfType == types.StringToType("string") {
		if leftInfType != rightInfType {
			msg := fmt.Sprintf("operands not of same type to assign, (%s) is of %s type and (%s) is of %s type", as.lValue.String(), leftInfType, as.rExpression.String(), rightInfType)
			semError := context.NewCompilerError(as.Line, as.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		}
	} else if leftInfType != types.StringToType("nil") && rightInfType != types.StringToType("nil") && leftInfType != rightInfType {
		msg := fmt.Sprintf("operands not of same reference type to assign, (%s) is of %s type and (%s) is of %s type", as.lValue.String(), leftInfType, as.rExpression.String(), rightInfType)
		semError := context.NewCompilerError(as.Line, as.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}

func (as *AssignStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	var lOperand llvm.LLVMOperand
	if strings.Contains(as.lValue.String(), ".") {
		as.lValue.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
		loadInstn := currBlk.Instns[len(currBlk.Instns)-1].(*llvm.LoadInstn)
		lOperand = loadInstn.GetUses()[0]
		currBlk.Instns = currBlk.Instns[:len(currBlk.Instns)-1]
		llvmProgram.RegCount--
	} else {
		lOperand = llvmProgram.GetRegister(funcEntry.Name, as.lValue.String())
	}
	
	rOperand := as.rExpression.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
	if rOperand.GetType() == types.NilTySig {
		fmt.Printf("ROP: %s, %s\n", rOperand.String(), types.TypesToLLVMTypes(rOperand.GetType()))
		rOperand.SetType(lOperand.GetType())
		fmt.Printf("ROP: %s, %s\n", rOperand.String(), types.TypesToLLVMTypes(rOperand.GetType()))
	}

	// currBlk.Instns = append(currBlk.Instns, llvm.NewStoreInstn(lOperand.(*llvm.LLVMRegister), lOperand.GetType(), rOperand, rOperand.GetType()))
	currBlk.AddInstn(llvm.NewStoreInstn(lOperand.(*llvm.LLVMRegister), lOperand.GetType(), rOperand, rOperand.GetType()))
	
	return currBlk
}