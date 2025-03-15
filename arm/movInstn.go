package arm

import (
	"fmt"
)

type MovInstn struct {
	dest		string
	op			string
}

func NewMovInstn(dest string, op string) *MovInstn {
	return &MovInstn{dest, op}
}

func (m *MovInstn) String() string {
	str := fmt.Sprintf("mov %s, %s", m.dest, m.op)
	
	return str
}