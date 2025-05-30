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
	"regexp"
)

type PrintStmt struct {
	*token.Token
	str		 		string
	args 			[]Expression
}

func NewPrintStmt(token *token.Token, str string, args []Expression) *PrintStmt {
	return &PrintStmt{token, str, args}
}

func (ps *PrintStmt) String() string {
	var out bytes.Buffer

	out.WriteString("printf(" + ps.str)

	if ps.args != nil {
		for _, arg := range (ps.args) {
			out.WriteString(", " + arg.String());
		}
	}

	out.WriteString(");")

	return out.String()
}

func countPlaceholders(format string) int {
	re := regexp.MustCompile(`%[+#0-9.]*[vTtbcdoqxXUeEfFgGsp]`)
	matches := re.FindAllString(format, -1)
	return len(matches)
}

func (ps *PrintStmt) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	var argTy types.Type
	for _, arg := range (ps.args) {
		argTy, errors = arg.TypeCheck(errors, funcEntry, tables)
		if (argTy.String() != "int") {
			msg := fmt.Sprintf("trying to print non integer value (%s)", arg.String())
			semError := context.NewCompilerError(ps.Line, ps.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		}
	}

	if countPlaceholders(ps.str) != len(ps.args) {
		msg := fmt.Sprintf("number of placeholders (%d) does not match number of arguments (%d) to print", countPlaceholders(ps.str), len(ps.args))
		semError := context.NewCompilerError(ps.Line, ps.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
	}

	return errors
}

func (ps *PrintStmt) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	_, goMod := llvmProgram.ReplaceForLLVMFormat(ps.str)
	id := llvmProgram.AddFormatStr(ps.str)

	var argList []llvm.LLVMOperand
	for _, arg := range ps.args {
		reg := arg.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
		argList = append(argList, reg)
	}

	// currBlk.Instns = append(currBlk.Instns, llvm.NewPrintInstn(id, goMod, argList))
	currBlk.AddInstn(llvm.NewPrintInstn(id, goMod, argList))
	
	return currBlk
}