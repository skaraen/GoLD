package arm

import (
	"fmt"
)

type JumpInstn struct {
	label		string
}

func NewJumpInstn(label string) *JumpInstn {
	return &JumpInstn{label}
}

func (j *JumpInstn) String() string {
	str := fmt.Sprintf("b .%s", j.label)
	
	return str
}