package ast

import (
	"golite/token"
	"golite/types"
	"strconv"
)

type IntLit struct {
	*token.Token
	litType 	types.Type
	value 		int64
}

func NewIntLit(token *token.Token, value int64) *IntLit {
	return &IntLit{token, types.StringToType("int"), value}
}

func (intLit *IntLit) String() string {
	return strconv.FormatInt(intLit.value, 10)
}

type BoolLit struct {
	*token.Token
	litType 	types.Type
	value 		bool
}

func NewBoolLit(token *token.Token, value bool) *BoolLit {
	return &BoolLit{token, types.StringToType("bool"), value}
}

func (boolLit *BoolLit) String() string {
	return strconv.FormatBool(boolLit.value)
}

type StrLit struct {
	*token.Token
	litType 	types.Type
	value 		string
}

func NewStrLit(token *token.Token, value string) *StrLit {
	return &StrLit{token, types.StringToType("string"), value}
}

func (strLit *StrLit) String() string {
	return strLit.value
}