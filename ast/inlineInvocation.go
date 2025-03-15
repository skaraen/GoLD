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

type InlineInvocation struct {
	*token.Token
	function 	string
	args 		Expression
}

func NewInlineInvocation(token *token.Token, function string, args Expression, isInv bool) *InlineInvocation {
	return &InlineInvocation{token, function, args}
}

func (inv *InlineInvocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.function + inv.args.String())

	return out.String()
}

func (inv *InlineInvocation) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	if fEntry, exists := tables.Funcs.Contains(inv.function); exists {
		return fEntry.RetTy
	}

	return types.StringToType("undefined")
}

func (inv *InlineInvocation) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	if invEntry, exists := tables.Funcs.Contains(inv.function); !exists {
		msg := fmt.Sprintf("method (%s) trying to be invoked does not exist", inv.function)
		semError := context.NewCompilerError(inv.Line, inv.Column, msg, context.SEMANTICS)
		errors = append(errors, semError)
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

	return inv.GetType(funcEntry, tables), errors
}

func (inv *InlineInvocation) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	invEntry, _ := tables.Funcs.Contains(inv.function)

	var opList []llvm.LLVMOperand
	for _, arg := range inv.args.(*Args).argsList {
		op := arg.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
		opList = append(opList, op)
	}

	dest := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(inv.String(), inv.GetType(funcEntry, tables), st.LOCAL, inv.Token), false)
	invInstn := llvm.NewInvInstn(invEntry, dest, opList)
	// currBlk.Instns = append(currBlk.Instns, invInstn)
	currBlk.AddInstn(invInstn)

	if invEntry.RetTy == types.BoolTySig {
		newDest := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), st.NewVarEntry(inv.String(), types.Int1TySig, st.LOCAL, inv.Token), false)
		truncInstn := llvm.NewTruncInstn(newDest, newDest.GetType(), dest, dest.GetType())
		// currBlk.Instns = append(currBlk.Instns, truncInstn)
		currBlk.AddInstn(truncInstn)

		return newDest
	}
	
	return dest
}