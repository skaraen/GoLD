package types

import "fmt"

type Type interface {
	String() string
}

type IntTy struct {}

func (intTy *IntTy) String() string {
	return "int"
}

type Int1Ty struct {}

func (int1Ty *Int1Ty) String() string {
	return "int1"
}

type Int8PtrTy struct {}

func (int8Ty *Int8PtrTy) String() string {
	return "*int8"
}

type BoolTy struct {}

func (boolTy *BoolTy) String() string {
	return "bool"
}

type PointerTy struct {
	id		string
}

func (pTy *PointerTy) String() string {
	return pTy.id
}

type StrTy struct {}

func (structTy *StrTy) String() string {
	return "string"
}

type NilTy struct {}

func (nilTy *NilTy) String() string {
	return "nil"
}

type UndefTy struct {}

func (undefTy *UndefTy) String() string {
	return "undefined"
}

func StringToType(str string) Type {
	if str == "int" {
		return IntTySig
	} else if str == "int1" {
		return Int1TySig
	} else if str == "*int8" {
		return Int8PtrTySig
	} else if str == "bool" {
		return BoolTySig
	} else if str == "string" {
		return StrTySig
	} else if str == "nil" {
		return NilTySig
	} else if str == "undefined" {
		return UndefTySig
	} else if userTypes[str] != nil {
		return userTypes[str]
	} else {
		AddNewUserType(str)
		return userTypes[str]
	}
}

var IntTySig *IntTy
var Int1TySig *Int1Ty
var Int8PtrTySig *Int8PtrTy
var BoolTySig *BoolTy
var StrTySig *StrTy
var NilTySig *NilTy
var UndefTySig *UndefTy
var userTypes map[string]Type

func init() {
	IntTySig = &IntTy{}
	Int1TySig = &Int1Ty{}
	Int8PtrTySig = &Int8PtrTy{}
	BoolTySig = &BoolTy{}
	StrTySig = &StrTy{}
	NilTySig = &NilTy{}
	UndefTySig = &UndefTy{}
	userTypes = make(map[string]Type)
}

func AddNewUserType(id string) {
	if (userTypes[id] != nil) {
		fmt.Println("Type declared twice")
		return
	}
	userTypes[id] = &PointerTy{id: id}
}

func TypesToLLVMTypes(ty Type) string {
	str := ""

	if ty == IntTySig || ty == BoolTySig {
		str = "i64"
	} else if ty == Int1TySig {
		str = "i1"
	} else if ty == Int8PtrTySig {
		str = "i8*"
	} else if pTy, ok := ty.(*PointerTy); ok {
		str = fmt.Sprintf("%%struct.%s*", pTy.id)
	} else if ty == NilTySig {
		str = "void"
	}

	return str
}

func GetLLVMDefaultValue(ty Type) string {
	str := ""

	if ty == IntTySig || ty == BoolTySig {
		str = "0"
	} else if _, ok := ty.(*PointerTy); ok {
		str = "null"
	}

	return str
}