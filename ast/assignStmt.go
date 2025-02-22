package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
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