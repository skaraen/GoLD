package ast

import (
	"bytes"
	"fmt"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type BinaryOp struct {
	*token.Token
	operator 		Operator
	left 			Expression
	right 			Expression
	leftInfType		types.Type
	rightInfType	types.Type
}

func NewBinaryOp(token *token.Token, operator Operator, left Expression, right Expression) *BinaryOp {
	return &BinaryOp{token, operator, left, right, types.StringToType("undefined"), types.StringToType("undefined")}
}

func (binOp *BinaryOp) String() string {
	var out bytes.Buffer

	out.WriteString(binOp.left.String())
	if (binOp.operator == DOT) {
		out.WriteString(OpToStr(binOp.operator))
	} else {
		out.WriteString(" " + OpToStr(binOp.operator) + " ")
	}
	out.WriteString(binOp.right.String())

	return out.String()
}

func (b *BinaryOp) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	fmt.Printf("BinOp getType (%s), op: %s\n", b.String(), OpToStr(b.operator))
	if b.operator == DOT {
		var structEntry *st.StructEntry
		var fieldEntry *st.VarEntry
		var exists bool

		if structEntry, exists = tables.Structs.Contains(b.left.GetType(funcEntry, tables).String()); exists {
			if fieldEntry, exists = structEntry.Fields.Contains(b.right.String()); exists {
				return fieldEntry.Ty
			}
		}
	} else if isComparisonOp(b.operator) || isEqualityOp(b.operator) || isLogicalOp(b.operator) {
		return types.StringToType("bool")
	} else {
		return types.StringToType("int")
	}

	return types.StringToType("undefined")
}

func (b *BinaryOp) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	binTy := types.StringToType("undefined") 
	b.leftInfType, errors = b.left.TypeCheck(errors, funcEntry, tables)

	if b.operator != DOT {
		b.rightInfType, errors = b.right.TypeCheck(errors, funcEntry, tables)
	}

	if b.operator == DOT {
		if structEntry, exists := tables.Structs.Contains(b.left.GetType(funcEntry, tables).String()); !exists {
			msg := fmt.Sprintf("undefined type of %s located at (%d,%d)", b.leftInfType.String(), b.Line, b.Column)
			semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		} else {
			if fieldEntry, exists := structEntry.Fields.Contains(b.right.String()); !exists {
				msg := fmt.Sprintf("undefined field %s for user defined type %s, located at (%d,%d)", b.right.String(), b.leftInfType.String(), b.Line, b.Column)
				semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
				errors = append(errors, semError)
			} else {
				binTy = fieldEntry.Ty
			}
		}
	} else if isEqualityOp(b.operator) {
		if b.leftInfType != b.rightInfType {
			if _, ok := b.leftInfType.(*types.PointerTy); ok && b.rightInfType.String() == "nil" {
				binTy = types.StringToType("bool")
			} else if _, ok := b.rightInfType.(*types.PointerTy); ok && b.leftInfType.String() == "nil" {
				binTy = types.StringToType("bool")
			} else {
				msg := fmt.Sprintf("operands not of same type to apply %s operator, (%s) is of %s type and (%s) is of %s type", OpToStr(b.operator), b.left.String(), b.leftInfType, b.right.String(), b.rightInfType)
				semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
				errors = append(errors, semError)
			}
		} else {
			binTy = types.StringToType("bool")
		}

		// if b.leftInfType == types.StringToType("bool") || b.leftInfType == types.StringToType("int") || b.leftInfType == types.StringToType("string") {
		// 	if b.leftInfType != b.rightInfType {
		// 		msg := fmt.Sprintf("operands not of same type to apply %s operator, (%s) is of %s type and (%s) is of %s type", OpToStr(b.operator), b.left.String(), b.leftInfType, b.right.String(), b.rightInfType)
		// 		semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
		// 		errors = append(errors, semError)
		// 	} else {
		// 		binTy = types.StringToType("bool")
		// 	}
		// } else {
		// 	if b.leftInfType != types.StringToType("nil") && b.rightInfType != types.StringToType("nil") && b.leftInfType != b.rightInfType {
		// 		msg := fmt.Sprintf("operands not of same reference type to apply %s operator, (%s) is of %s type and (%s) is of %s type", OpToStr(b.operator), b.left.String(), b.leftInfType, b.right.String(), b.rightInfType)
		// 		semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
		// 		errors = append(errors, semError)
		// 	}
		// }		
	} else {
		if b.leftInfType != b.rightInfType {
			msg := fmt.Sprintf("operands not of same type to apply %s operator, (%s) is of %s type and (%s) is of %s type", OpToStr(b.operator), b.left.String(), b.leftInfType, b.right.String(), b.rightInfType)
			semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		} else {
			if isLogicalOp(b.operator) && b.leftInfType != types.StringToType("bool") {
				msg := fmt.Sprintf("operands need to be of bool type to apply %s operator, are of %s type instead", OpToStr(b.operator), b.leftInfType)
				semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
				errors = append(errors, semError)
			} else {
				binTy = types.StringToType("bool")
			} 
			
			if (isArithmeticOp(b.operator) || isComparisonOp(b.operator)) && b.leftInfType != types.StringToType("int") {
				msg := fmt.Sprintf("operands need to be of int type to apply %s operator, are of %s type instead", OpToStr(b.operator), b.leftInfType)
				semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
				errors = append(errors, semError)
			} else {
				if isArithmeticOp(b.operator) {
					binTy = types.StringToType("int")
				} else {
					binTy = types.StringToType("bool")
				}
			}
		}
	}

	return binTy, errors
}