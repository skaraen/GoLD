package arm

import (
	"fmt"
)

type LoadPairInstn struct {
	dest1		string
	dest2		string
	op			string
}

func NewLoadPairInstn(dest1 string, dest2 string, op string) *LoadPairInstn {
	return &LoadPairInstn{dest1, dest2, op}
}

func (l *LoadPairInstn) String() string {
	str := fmt.Sprintf("ldp %s, %s, [%s]", l.dest1, l.dest2, l.op)
	
	return str
}