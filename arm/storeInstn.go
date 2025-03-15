package arm

import (
	"fmt"
)

type StoreInstn struct {
	dest		string
	op			string
}

func NewStoreInstn(dest string, op string) *StoreInstn {
	return &StoreInstn{dest, op}
}

func (s *StoreInstn) String() string {
	str := fmt.Sprintf("str %s, [%s]", s.op, s.dest)
	
	return str
}