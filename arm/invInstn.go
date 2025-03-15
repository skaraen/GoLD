package arm

import (
	"fmt"
)

type InvInstn struct {
	fnName		string
}

func NewInvInstn(fnName string) *InvInstn {
	return &InvInstn{fnName}
}

func (i *InvInstn) String() string {
	str := fmt.Sprintf("bl %s", i.fnName)
	
	return str
}