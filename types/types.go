package types

type Type interface {
	String() string
}

type IntTy struct {}

func (intTy *IntTy) String() string {
	return "int"
}

type BoolTy struct {}

func (boolTy *BoolTy) String() string {
	return "bool"
}

type StructTy struct {}

func (structTy *StructTy) String() string {
	return "struct"
}

func StringToType(str string) Type {
	if str == "int" {
		return IntTySig
	} else if str == "bool" {
		return BoolTySig
	} else if str == "struct" {
		return StructTySig
	}
	panic("Type not found " + str)
}

var IntTySig *IntTy
var BoolTySig *BoolTy
var StructTySig *StructTy

func init() {
	IntTySig = &IntTy{}
	BoolTySig = &BoolTySig{}
	StructTySig = &StructTySig{}
}