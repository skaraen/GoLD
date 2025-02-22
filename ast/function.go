package ast

import (
	"bytes"
	"fmt"
	"golite/context"
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
		msg := fmt.Sprintf("redefinition of function (%s) located at (%d,%d)", f.id, fn.Line, fn.Column)
		semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
			
		errors = append(errors, semError)
	}
	
	funcEntry := st.NewFuncEntry(f.id, f.returnType, f.Token)
	paramsList := f.params.(*Params).paramsList

	for _, pExp := range (paramsList) {
		param := pExp.(*Decl)

		if p, exists := funcEntry.Variables.Contains(param.id); exists {
			msg := fmt.Sprintf("redefinition of parameter (%s) located at (%d,%d)", param.id, p.Line, p.Column)
			semError := context.NewCompilerError(f.Line, f.Column, msg, context.SEMANTICS)
				
			errors = append(errors, semError)
			continue
		}

		varEntry := st.NewVarEntry(param.id, param.dType, st.LOCAL, param.Token)
		funcEntry.Variables.Insert(param.id, varEntry)
	}

	declrList := f.declarations.(*Declarations).declrList

	for _, declr := range (declrList) {
		ids := declr.ids.(*Ids)

		for _, id := range(ids.idsList) {
			if st, exists := tables.Structs.Contains(id); exists {
				msg := fmt.Sprintf("redefinition of user defined type (%s) located at (%d,%d)", id, st.Line, st.Column)
				semError := context.NewCompilerError(ids.Line, ids.Column, msg, context.SEMANTICS)
				
				errors = append(errors, semError)
				continue
			}

			if v, exists := funcEntry.Variables.Contains(id); exists {
				msg := fmt.Sprintf("redefinition of variable (%s) located at (%d,%d)", id, v.Line, v.Column)
				semError := context.NewCompilerError(ids.Line, ids.Column, msg, context.SEMANTICS)
				
				errors = append(errors, semError)
				continue
			}
			
			varEntry := st.NewVarEntry(id, declr.declType, st.LOCAL, ids.Token)
			funcEntry.Variables.Insert(id, varEntry)
		}
	}

	funcEntry.Variables.SetParent(tables.Globals)
	tables.Funcs.Insert(f.id, funcEntry)
	return errors
}

func (f *Function) TypeCheck(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError { 
	if funcEntry, exists := tables.Funcs.Contains(f.id); exists {
		errors = f.statements.TypeCheck(errors, funcEntry, tables)
	}
	
	return errors
}