package cfg

import (
	"bytes"
	"fmt"
)

type Block struct {
	Label 		string
	Instns 		[]Instruction
	Preds 		[]*Block
	Succs 		[]*Block
	Exit 		*Block
	IsReturning bool
}

func (b *Block) FindBlock(blockList []*Block, target *Block) int {
	for idx, block := range blockList {
		if block == target {
			return idx
		}
	}
	return -1
}

func (b *Block) FindInstn(target Instruction) int {
	for idx, instn := range b.Instns {
		if instn == target {
			return idx
		}
	}
	return -1
}

func (b *Block) AddPred(pred *Block) {
	if b.FindBlock(b.Preds, pred) == -1 {
		b.Preds = append(b.Preds, pred)
	}
}

func (b *Block) AddSucc(succ *Block) {
	if b.FindBlock(b.Succs, succ) == -1 {
		b.Succs = append(b.Succs, succ)
	}
}

func (b *Block) RemovePred(pred *Block) {
	idx := b.FindBlock(b.Preds, pred)
	if idx == -1 {
		fmt.Printf("Cannot find pred block %s for %s \n", pred.Label, b.Label)
		return
	}

	b.Preds = append(b.Preds[:idx], b.Preds[idx+1:]...)
}

func (b *Block) RemoveInstn(instn Instruction) {
	idx := b.FindInstn(instn)
	if idx == -1 {
		fmt.Printf("Cannot find instn %s for %s \n", instn.String(), b.Label)
		return
	}

	b.Instns = append(b.Instns[:idx], b.Instns[idx+1:]...)
}

func (b *Block) ClearSuccs() {
	b.Succs = b.Succs[:0]
}

func (b *Block) AddInstn(instn Instruction) {
	if b.IsReturning {
		return
	}
	b.Instns = append(b.Instns, instn)
}

func NewBlock(entryLabel string, exitLabel string) (*Block, *Block) {
	exit :=  &Block{exitLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), nil, false}
	entry := &Block{entryLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}

	return entry, exit
}

func NewIfBlock(entry *Block, exit *Block, trueLabel string, falseLabel string, exitLabel string) (*Block, *Block, *Block) {
	trueBlk :=  &Block{trueLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}
	falseBlk := &Block{falseLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}
	exitIfBlk := &Block{exitLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}

	entry.AddSucc(trueBlk)
	entry.AddSucc(falseBlk)
	trueBlk.AddPred(entry)
	falseBlk.AddPred(entry)

	return trueBlk, falseBlk, exitIfBlk
}

func NewForBlock(prev *Block, exit *Block, entryLabel string, bodyLabel string, exitLabel string) (*Block, *Block, *Block) {
	entryForBlk :=  &Block{entryLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}
	exitForBlk :=  &Block{exitLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}
	bodyBlk :=	&Block{bodyLabel, make([]Instruction, 0), make([]*Block, 0), make([]*Block, 0), exit, false}

	prev.AddSucc(entryForBlk)
	entryForBlk.AddPred(prev)
	entryForBlk.AddSucc(bodyBlk)
	entryForBlk.AddSucc(exitForBlk)
	bodyBlk.AddPred(entryForBlk)
	exitForBlk.AddPred(entryForBlk)

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