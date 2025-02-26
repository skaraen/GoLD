package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

type Functions struct {
	*token.Token
	funcList		[]*Function
}

func NewFunctions(token *token.Token, funcList []*Function) *Functions {
	return &Functions{token, funcList}
}

func (f *Functions) String() string {
	var out bytes.Buffer

	for _, function := range (f.funcList) {
		out.WriteString(function.String() + "\n")
	}

	return out.String()
}

func (f *Functions) BuildSymbolTable(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	for _, fn := range (f.funcList) {
		errors = fn.BuildSymbolTable(errors, tables)
	}

	if mainFn, exists := tables.Funcs.Contains("main"); !exists {
		msg := "main function absent in program"
		semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
		errors = append(errors, semError)
	} else {
		if mainFn.RetTy.String() != "nil" {
			msg := fmt.Sprintf("main function has non void return type (%s)", mainFn.RetTy.String())
			semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
			errors = append(errors, semError)
		} 

		if len(mainFn.Signature) > 0 {
			var str []string
			for _, sigTy := range (mainFn.Signature) {
				str = append(str, sigTy.String())
			}
			msg := fmt.Sprintf("main function has arguments (%s)", strings.Join(str, ", "))
			semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
			errors = append(errors, semError)
		}
	}

	return errors
}

func (f *Functions) TypeCheck(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	isMain := false 
	for _, fn := range (f.funcList) {
		if fn.id == "main" {
			isMain = true
			if fn.returnType != types.StringToType("nil") {
				msg := "main function has non void return type"
				semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
				errors = append(errors, semError)
			}
		}
		errors = fn.TypeCheck(errors, tables)
	}
	
	if !isMain {
		msg := "main function absent in program"
		semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
		errors = append(errors, semError)
	}
	
	return errors
}