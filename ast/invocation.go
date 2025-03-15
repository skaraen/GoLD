package ast

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/context"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
)

type Invocation struct {
	*token.Token
	function 	string
	args 		Expression
}

func NewInvocation(token *token.Token, function string, args Expression) *Invocation {
	return &Invocation{token, function, args}
}

func (inv *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.function + inv.args.String() + ";")

	return out.String()
}

func (inv *Invocation) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) []*context.CompilerError {
	// var fEntry *st.FuncEntry

	if invEntry, exists := tables.Funcs.Contains(inv.function); !exists {
		msg := fmt.Sprintf("method (%s) trying to be invoked does not exist", inv.function)
		semError := context.NewCompilerError(inv.Line, inv.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)

		return errors
	} else {
		_, errors = inv.args.TypeCheck(errors, funcEntry, tables)

		if len(invEntry.Signature) != len(inv.args.(*Args).argsList) {
			msg := fmt.Sprintf("function %s need %d arguments, found %d instead", invEntry.Name, len(invEntry.Signature), len(inv.args.(*Args).argsList))
			semError := context.NewCompilerError(inv.Line, inv.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		} else {
			for i, arg := range (inv.args.(*Args).argsList) {
				argTy := arg.GetType(funcEntry, tables)
				if invEntry.Signature[i] != argTy {
					msg := fmt.Sprintf("need argument of type %s, argument (%s) is of type %s", invEntry.Signature[i].String(), arg.String(), argTy.String())
					semError := context.NewCompilerError(inv.Line, inv.Column, msg, context.SEMANTICS)
					errors = append(errors, semError)
				}
			}
		}
	}

	return errors
}

func (inv *Invocation) TranslateToLLVMStack(currBlk *cfg.Block, exitBlk *cfg.Block, llvmProgram *llvm.LLVMProgram, funcEntry *st.FuncEntry, tables *st.SymbolTables) *cfg.Block {
	invEntry, _ := tables.Funcs.Contains(inv.function)

	var argList []llvm.LLVMOperand
	for _, arg := range inv.args.(*Args).argsList {
		reg := arg.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
		argList = append(argList, reg)
	}

	invInstn := llvm.NewInvInstn(invEntry, nil, argList)
	// currBlk.Instns = append(currBlk.Instns, invInstn)
	currBlk.AddInstn(invInstn)
	
	return currBlk
}