package symboltable

import (
	"bytes"
	"fmt"
	"golite/token"
	"golite/types"
)

type VarScope int

const (
	GLOBAL VarScope = iota
	LOCAL
	STRUCT
)

type VarEntry struct {
	*token.Token
	Name  		string
	Ty    		types.Type
	Scope 		VarScope
}

func NewVarEntry(name string, ty types.Type, scope VarScope, token *token.Token) *VarEntry {
	return &VarEntry{token, name, ty, scope}
}

type FuncEntry struct {
	*token.Token
	Name       	string
	RetTy      	types.Type
	Signature	[]types.Type
	Parameters 	*SymbolTable[*VarEntry]
	Variables  	*SymbolTable[*VarEntry]
	Expressions *SymbolTable[*VarEntry]
}

func NewFuncEntry(name string, returnTy types.Type, token *token.Token) *FuncEntry {
	return &FuncEntry{token, name, returnTy, make([]types.Type, 0), NewSymbolTable[*VarEntry](nil), NewSymbolTable[*VarEntry](nil), NewSymbolTable[*VarEntry](nil)}
}

type StructEntry struct {
	*token.Token
	Name       	string
	Fields 		*SymbolTable[*VarEntry]
}

func NewStructEntry(name string, token *token.Token) *StructEntry {
	return &StructEntry{token, name, NewSymbolTable[*VarEntry](nil)}
}

func (entry *VarEntry) String() string {
	return fmt.Sprintf("VarEntry{Name: %s, Type: %s, Scope: %d}", entry.Name, entry.Ty.String(), entry.Scope)
}

func (entry *FuncEntry) String() string {
	return fmt.Sprintf("FuncEntry{Name: %s, ReturnType: %s, Variables: %d}", entry.Name, entry.RetTy.String(), len(entry.Variables.table))
}

func (entry *StructEntry) String() string {
	return fmt.Sprintf("StructEntry{Name: %s, Fields: %d}", entry.Name, len(entry.Fields.table))
}

type SymbolTables struct {
	Funcs   	*SymbolTable[*FuncEntry]
	Globals 	*SymbolTable[*VarEntry]
	Structs 	*SymbolTable[*StructEntry]
}

func NewSymbolTables() *SymbolTables {
	st := &SymbolTables{NewSymbolTable[*FuncEntry](nil), NewSymbolTable[*VarEntry](nil), NewSymbolTable[*StructEntry](nil)}
	return st
}

func (tables *SymbolTables) String() string {
	var out bytes.Buffer
	for _, structEntry := range (tables.Structs.table) {
		out.WriteString(structEntry.String() + "\n")
	}

	for _, globalEntry := range (tables.Globals.table) {
		out.WriteString(globalEntry.String() + "\n")
	}

	for _, funcEntry := range (tables.Funcs.table) {
		out.WriteString(funcEntry.String() + "\n")
	}

	return out.String()
}

type SymbolTable[T fmt.Stringer] struct {
	parent 	*SymbolTable[T]
	table  	map[string]T
	order	[]string 	
}

func NewSymbolTable[T fmt.Stringer](parent *SymbolTable[T]) *SymbolTable[T] {
	return &SymbolTable[T]{parent, make(map[string]T), make([]string, 0)}
}

func (st *SymbolTable[T]) Insert(id string, entry T) {
	st.table[id] = entry
	st.order = append(st.order, id)
}

func (st *SymbolTable[T]) Contains(id string) (T, bool) {
	table  := st
	
	for table != nil {
		if entry, exists := table.table[id]; exists {
			return entry, true
		}

		table = table.parent
	}

	var dummyEntry T
	return dummyEntry, false
}

func (st *SymbolTable[T]) SetParent(parent *SymbolTable[T]) {
	st.parent = parent
}

func (st *SymbolTable[T]) GetTable() map[string]T {
	return st.table
}

func (st *SymbolTable[T]) GetIndex(target string) (int, bool) {
	for idx, id := range st.order {
		if id == target {
			return idx, true
		}
	}
	return -1, false
}

func (st *SymbolTable[T]) GetOrderedEntries() []T {
	var orderedList []T

	for _, id := range st.order {
		entry := st.table[id]
		orderedList = append(orderedList, entry)
	}

	return orderedList
}