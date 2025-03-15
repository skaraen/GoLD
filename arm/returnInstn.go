package arm

type ReturnInstn struct {
}

func NewReturnInstn() *ReturnInstn {
	return &ReturnInstn{}
}

func (r *ReturnInstn) String() string {	
	return "ret"
}