package arm

import (
	"fmt"
	"golite/operators"
)

type BinOpInstn struct {
	dest		string
	op1			string
	op2			string
	operator 	operators.Operator
}

func NewBinOpInstn(dest string, op1 string, op2 string, operator operators.Operator) *BinOpInstn {
	return &BinOpInstn{dest, op1, op2, operator}
}

func (b *BinOpInstn) String() string {
	str := ""

	if operators.IsArithmeticOp(b.operator) || operators.IsLogicalOp(b.operator) {
		str = fmt.Sprintf("%s %s, %s, %s", operators.OpToAssembly(b.operator), b.dest, b.op1, b.op2)
	} else if operators.IsComparisonOp(b.operator) || operators.IsEqualityOp(b.operator) {
		str = fmt.Sprintf("cmp %s, %s", b.op1, b.op2)
	}
	
	return str
}