package arm

import (
	"fmt"
)

type StorePairInstn struct {
	dest		string
	op1			string
	op2			string
}

func NewStorePairInstn(dest string, op1 string, op2 string) *StorePairInstn {
	return &StorePairInstn{dest, op1, op2}
}

func (s *StorePairInstn) String() string {
	str := fmt.Sprintf("stp %s, %s, [%s]", s.op1, s.op2, s.dest)
	
	return str
}