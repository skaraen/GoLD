package ast

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	op "golite/operators"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type UnaryOp struct {
	*token.Token
	operator 					op.Operator
	right 						Expression
	rightInfType				types.Type
}

func NewUnaryOp(token *token.Token, operator op.Operator, right Expression) *UnaryOp {
	return &UnaryOp{token, operator, right, types.StringToType("undefined")}
}

func (unrOp *UnaryOp) String() string {
	var out bytes.Buffer

	out.WriteString(op.OpToStr(unrOp.operator))
	out.WriteString(unrOp.right.String())

	return out.String()
}

func (u *UnaryOp) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	if u.operator == op.NOT && u.rightInfType.String() == "bool" {
		return u.rightInfType
	} else if u.operator == op.MINUS && u.rightInfType.String() == "int" {
		return u.rightInfType
	}

	return types.StringToType("undefined")
}

func (u *UnaryOp) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	u.rightInfType, errors = u.right.TypeCheck(errors, funcEntry, tables)

	if (u.operator == op.NOT && u.rightInfType.String() != "bool") || (u.operator == op.MINUS && u.rightInfType.String() != "int") {
		msg := fmt.Sprintf("cannot apply a (%s) operator to an expression of type (%s) located at (%d,%d)", op.OpToStr(u.operator), u.rightInfType.String(), u.Line, u.Column)
		semError := context.NewCompilerError(u.Line, u.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return u.GetType(funcEntry, tables), errors
}

func (u *UnaryOp) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	rOperand := u.right.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)

	destEntry := st.NewVarEntry(u.String(), u.GetType(funcEntry, tables), st.LOCAL, u.Token)
	funcEntry.Expressions.Insert(destEntry.Name, destEntry)
	destRegister := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), destEntry)

	instn := llvm.NewUnOpInstn(destRegister, u.operator, rOperand, u.rightInfType)

	currBlk.Instns = append(currBlk.Instns, instn)
	return destRegister
}