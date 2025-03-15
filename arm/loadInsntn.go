package arm

import (
	"fmt"
)

type LoadInstn struct {
	dest		string
	op			string
}

func NewLoadInstn(dest string, op string) *LoadInstn {
	return &LoadInstn{dest, op}
}

func (l *LoadInstn) String() string {
	str := fmt.Sprintf("ldr %s, [%s]", l.dest, l.op)
	
	return str
}