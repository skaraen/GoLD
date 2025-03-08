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

type BinaryOp struct {
	*token.Token
	operator 		op.Operator
	left 			Expression
	right 			Expression
	leftInfType		types.Type
	rightInfType	types.Type
}

func NewBinaryOp(token *token.Token, operator op.Operator, left Expression, right Expression) *BinaryOp {
	return &BinaryOp{token, operator, left, right, types.StringToType("undefined"), types.StringToType("undefined")}
}

func (binOp *BinaryOp) String() string {
	var out bytes.Buffer

	out.WriteString(binOp.left.String())
	if (binOp.operator == op.DOT) {
		out.WriteString(op.OpToStr(binOp.operator))
	} else {
		out.WriteString(" " + op.OpToStr(binOp.operator) + " ")
	}
	out.WriteString(binOp.right.String())

	return out.String()
}

func (b *BinaryOp) GetType(funcEntry *st.FuncEntry, tables *st.SymbolTables) types.Type {
	if b.operator == op.DOT {
		var structEntry *st.StructEntry
		var fieldEntry *st.VarEntry
		var exists bool

		if structEntry, exists = tables.Structs.Contains(b.left.GetType(funcEntry, tables).String()); exists {
			if fieldEntry, exists = structEntry.Fields.Contains(b.right.String()); exists {
				return fieldEntry.Ty
			}
		}
	} else if op.IsComparisonOp(b.operator) || op.IsEqualityOp(b.operator) || op.IsLogicalOp(b.operator) {
		return types.StringToType("bool")
	} else {
		return types.StringToType("int")
	}

	return types.StringToType("undefined")
}

func (b *BinaryOp) TypeCheck(errors []*context.CompilerError, funcEntry *st.FuncEntry, tables *st.SymbolTables) (types.Type, []*context.CompilerError) {
	binTy := types.StringToType("undefined") 
	b.leftInfType, errors = b.left.TypeCheck(errors, funcEntry, tables)

	if b.operator != op.DOT {
		b.rightInfType, errors = b.right.TypeCheck(errors, funcEntry, tables)
	}

	if b.operator == op.DOT {
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
	} else if op.IsEqualityOp(b.operator) {
		if b.leftInfType != b.rightInfType {
			if _, ok := b.leftInfType.(*types.PointerTy); ok && b.rightInfType.String() == "nil" {
				binTy = types.StringToType("bool")
			} else if _, ok := b.rightInfType.(*types.PointerTy); ok && b.leftInfType.String() == "nil" {
				binTy = types.StringToType("bool")
			} else {
				msg := fmt.Sprintf("operands not of same type to apply %s operator, (%s) is of %s type and (%s) is of %s type", op.OpToStr(b.operator), b.left.String(), b.leftInfType, b.right.String(), b.rightInfType)
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
			msg := fmt.Sprintf("operands not of same type to apply %s operator, (%s) is of %s type and (%s) is of %s type", op.OpToStr(b.operator), b.left.String(), b.leftInfType, b.right.String(), b.rightInfType)
			semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
			errors = append(errors, semError)
		} else {
			if op.IsLogicalOp(b.operator) && b.leftInfType != types.StringToType("bool") {
				msg := fmt.Sprintf("operands need to be of bool type to apply %s operator, are of %s type instead", op.OpToStr(b.operator), b.leftInfType)
				semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
				errors = append(errors, semError)
			} else {
				binTy = types.StringToType("bool")
			} 
			
			if (op.IsArithmeticOp(b.operator) || op.IsComparisonOp(b.operator)) && b.leftInfType != types.StringToType("int") {
				msg := fmt.Sprintf("operands need to be of int type to apply %s operator, are of %s type instead", op.OpToStr(b.operator), b.leftInfType)
				semError := context.NewCompilerError(b.Line, b.Column, msg, context.SEMANTICS)
				errors = append(errors, semError)
			} else {
				if op.IsArithmeticOp(b.operator) {
					binTy = types.StringToType("int")
				} else {
					binTy = types.StringToType("bool")
				}
			}
		}
	}

	return binTy, errors
}

func (b *BinaryOp) TranslateToLLVMStack(funcEntry *st.FuncEntry, tables *st.SymbolTables, currBlk *cfg.Block, llvmProgram *llvm.LLVMProgram) llvm.LLVMOperand {
	var lOperand, rOperand llvm.LLVMOperand
	var instn llvm.LLVMInstruction
	lOperand = b.left.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)

	destEntry := st.NewVarEntry(b.String(), b.GetType(funcEntry, tables), st.LOCAL, b.Token)
	funcEntry.Expressions.Insert(destEntry.Name, destEntry)
	destRegister := llvm.NewLLVMRegister(llvmProgram.GenerateRegisterName(), destEntry)

	if (b.operator == op.DOT) {
		instn = llvm.NewGetElementInstn(destRegister, lOperand, b.right.String(), tables)
	} else {
		rOperand = b.right.TranslateToLLVMStack(funcEntry, tables, currBlk, llvmProgram)
		instn = llvm.NewBinOpInstn(destRegister, b.operator, lOperand, rOperand, b.leftInfType)
	}

	currBlk.Instns = append(currBlk.Instns, instn)
	return destRegister
}