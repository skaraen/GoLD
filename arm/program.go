package arm

import (
	"bytes"
	"fmt"
	"golite/cfg"
	"golite/operators"
	st "golite/symboltable"
)

type ARMProgram struct {
	Arch			string
	Tables			*st.SymbolTables
	Funcs			map[string][]*cfg.Block
	BlockInstns		map[*cfg.Block][]Instruction
}

func NewARMProgram (arch string, tables *st.SymbolTables, funcs map[string][]*cfg.Block, blockInstns map[*cfg.Block][]Instruction) *ARMProgram {
	return &ARMProgram{arch, tables, funcs, blockInstns}
}

func (p *ARMProgram) String() string {
	var out bytes.Buffer

	// Architecture
	out.WriteString(fmt.Sprintf("\t\t.arch %s\n", p.Arch))

	// Globals
	for globalName, _ := range p.Tables.Globals.GetTable() {
		out.WriteString(fmt.Sprintf("\t\t.comm %s, 8, 8\n", globalName))
	}

	out.WriteString("\t\t.text\n\n");

	for fnName, _ := range p.Tables.Funcs.GetTable() {
		out.WriteString(fmt.Sprintf("\t\t.type %s, %%function\n", fnName));
		out.WriteString(fmt.Sprintf("\t\t.global %s\n", fnName));
		out.WriteString(fmt.Sprintf("\t\t.p2align %d\n", 2))

		out.WriteString(fnName + ":\n")
		out.WriteString("\t\t\\\\Function prologue\n")
		out.WriteString(fmt.Sprintf("\t\t%s\n", NewBinOpInstn("sp", "sp", "32", operators.MINUS).String()))
		out.WriteString(fmt.Sprintf("\t\t%s\n",NewStorePairInstn("sp", "x29", "x30")))
		out.WriteString(fmt.Sprintf("\t\t%s\n", NewMovInstn("x29", "sp").String()))

		for _, block := range p.Funcs[fnName] {
			out.WriteString(fmt.Sprintf(".%s:\n", block.Label))
			for _, instn := range p.BlockInstns[block] {
				out.WriteString("\t\t" + instn.String() + "\n")
			}
		}

		out.WriteString(fmt.Sprintf("\t\t.size %s,(.- %s)\n", fnName, fnName));
	}

	return out.String()
}