package ast

import (
	"bytes"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
)

type Program struct {
	*token.Token
	userTypes		*UserTypes
	declarations 	*Declarations
	functions		*Functions
}

func NewProgram(token *token.Token, userTypes *UserTypes, declarations *Declarations, functions *Functions) *Program {
	return &Program{token, userTypes, declarations, functions}
}

func (p *Program) String() string {
	var out bytes.Buffer

	out.WriteString(p.userTypes.String() + "\n")
	out.WriteString(p.declarations.String() + "\n")
	out.WriteString(p.functions.String() + "\n")

	return out.String()
}

func (p *Program) BuildSymbolTable(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	errors = p.userTypes.BuildSymbolTable(errors, tables.Structs)
	errors = p.declarations.BuildSymbolTable(errors, tables)
	errors = p.functions.BuildSymbolTable(errors, tables)
	
	return errors
}

func (p *Program) TypeCheck(errors []*context.CompilerError, tables *st.SymbolTables) []*context.CompilerError {
	errors = p.functions.TypeCheck(errors, tables)
	
	return errors
}