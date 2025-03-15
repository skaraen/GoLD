package arm

import (
	"fmt"
	op "golite/operators"
)

type BranchInstn struct {
	operator		op.Operator
	targetLabel		string
}

func NewBranchInstn(operator op.Operator, target string) *BranchInstn {
	return &BranchInstn{operator, target}
}

func (b *BranchInstn) String() string {
	str := fmt.Sprintf("%s .%s", op.OpToAssembly(b.operator), b.targetLabel)
	
	return str
}