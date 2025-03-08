package cfg

import (
	"bytes"
)

type Block struct {
	Label 	string
	Instns 	[]Instruction
	Preds 	[]*Block
	Succs 	[]*Block
	Exit 	*Block
}

func NewBlock(entryLabel string, exitLabel string) (*Block, *Block) {
	exit :=  &Block{exitLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), nil}
	entry := &Block{entryLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}

	return entry, exit
}

func NewIfBlock(entry *Block, exit *Block, trueLabel string, falseLabel string, exitLabel string) (*Block, *Block, *Block) {
	trueBlk :=  &Block{trueLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}
	falseBlk := &Block{falseLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}
	exitIfBlk := &Block{exitLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}

	trueBlk.Preds = append(trueBlk.Preds, entry)
	falseBlk.Preds = append(falseBlk.Preds, entry)

	entry.Succs = append(entry.Succs, trueBlk)
	entry.Succs = append(entry.Succs, falseBlk)

	return trueBlk, falseBlk, exitIfBlk
}

func NewForBlock(prev *Block, exit *Block, entryLabel string, bodyLabel string, exitLabel string) (*Block, *Block, *Block) {
	entryForBlk :=  &Block{entryLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}
	exitForBlk :=  &Block{exitLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}
	bodyBlk :=	&Block{bodyLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit}

	prev.Succs = append(prev.Succs, entryForBlk)

	entryForBlk.Succs = append(entryForBlk.Succs, bodyBlk)
	entryForBlk.Succs = append(entryForBlk.Succs, exitForBlk)
	entryForBlk.Preds = append(entryForBlk.Preds, prev)
	// entryForBlk.Preds = append(entryForBlk.Preds, bodyBlk)
	
	bodyBlk.Preds = append(bodyBlk.Preds, entryForBlk)
	exitForBlk.Preds = append(exitForBlk.Preds, entryForBlk)

	return entryForBlk, bodyBlk, exitForBlk
}

func (b *Block) String() string {
	var out bytes.Buffer

	out.WriteString(b.Label + ":\n")

	for _, instn := range b.Instns {
		out.WriteString("\t" + instn.String() + "\n")
	}
	
	return out.String()
}