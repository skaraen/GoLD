package ast

type Program struct {
	stmts []*Statement
}

func NewProgram( /*..*/ ) *Program {
	return &Program{nil}
}