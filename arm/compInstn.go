package arm

import (
	"fmt"
)

type CompInstn struct {
	op1		string
	op2		string
}

func NewCompInstn(op1 string, op2 string) *CompInstn {
	return &CompInstn{op1, op2}
}

func (c *CompInstn) String() string {
	str := fmt.Sprintf("cmp %s, %s", c.op1, c.op2)
	
	return str
}