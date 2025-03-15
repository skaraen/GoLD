package arm

import (
	"fmt"
)

type AdrpInsntn struct {
	dest		string
	op			string
}

func NewAdrpInsntn(dest string, op string) *AdrpInsntn {
	return &AdrpInsntn{dest, op}
}

func (a *AdrpInsntn) String() string {
	str := fmt.Sprintf("adrp %s, %s", a.dest, a.op)
	
	return str
}