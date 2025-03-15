package llvm

import (
	"bytes"
	"fmt"
	"golite/arm"
	"golite/cfg"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"sort"
	"strconv"
	"strings"
)

type LLVMProgram struct {
	SourceName			string
	TargetTriple		string
	Tables				*st.SymbolTables
	Structs				[]LLVMInstruction
	Globals				[]LLVMInstruction
	Funcs				map[string]*cfg.Block
	EntryRegMap			map[*st.VarEntry]*LLVMRegister
	ParamRegisters		map[string][]*LLVMRegister
	RetRegisters		map[string]*LLVMRegister
	PrintStrs			[]string
	RegCount			int
	LabelCount			map[string]int
	LineCount			map[string]int
	Defs				map[*cfg.Block]map[string]LLVMOperand
	ReadScratch			*LLVMRegister
	VariableTable		map[string]map[*LLVMRegister]*LiveRangeEntry
}

type LiveRangeEntry struct {
	Start			int
	End				int
}

func NewLLVMProgram(srcName string, targetTriple string, tables *st.SymbolTables) *LLVMProgram {
	return &LLVMProgram{srcName, targetTriple, tables, make([]LLVMInstruction, 0), make([]LLVMInstruction, 0), 
		make(map[string]*cfg.Block), make(map[*st.VarEntry]*LLVMRegister), make(map[string][]*LLVMRegister), make(map[string]*LLVMRegister), 
		make([]string, 0), 0, make(map[string]int, 0), make(map[string]int, 0), make(map[*cfg.Block]map[string]LLVMOperand), 
		NewLLVMRegister(".read_scratch", st.NewVarEntry(".read_scratch", types.IntTySig, st.GLOBAL, token.NewToken(0, 0)), false),
		make(map[string]map[*LLVMRegister]*LiveRangeEntry)}
}

func (p *LLVMProgram) String() string {
	var out bytes.Buffer

	out.WriteString("source_filename = " + "\"" + p.SourceName + "\"\n")
	out.WriteString("target triple = " + "\"" + p.TargetTriple + "\"\n")

	out.WriteString("\n")

	for _, structInstn := range p.Structs {
		out.WriteString(structInstn.String() + "\n")
	}

	out.WriteString("\n")

	for _, globalInstn := range p.Globals {
		out.WriteString(globalInstn.String() + "\n")
	}

	out.WriteString("\n")

	for fnName, entryBlk := range p.Funcs {
		fnEntry,_ := p.Tables.Funcs.Contains(fnName)

		var paramStr []string
		for _, paramEntry := range fnEntry.Parameters.GetOrderedEntries() {
			reg := p.EntryRegMap[paramEntry]
			str := fmt.Sprintf("%s %%%s", types.TypesToLLVMTypes(reg.GetType()), paramEntry.Name)
			paramStr = append(paramStr, str)
		}

		out.WriteString(fmt.Sprintf("define %s @%s(%s)\n", types.TypesToLLVMTypes(fnEntry.RetTy), fnName, strings.Join(paramStr, ", ")))
		out.WriteString("{\n")

		for _, block := range TopologicalSort(entryBlk) {
			out.WriteString(block.String() + "\n")
		}

		for i, fstr := range p.PrintStrs {
			str := fmt.Sprintf("@.fstr%d = private unnamed_addr constant [%d x i8] c\"%s\", align 1\n", i + 1, len(fstr), fstr)
			paramStr = append(paramStr, str)
		}
		out.WriteString("}\n")
	}

	out.WriteString("@.read = private unnamed_addr constant [4 x i8] c\"%ld\\00\", align 1\n");
	out.WriteString("declare i8* @malloc(i32)\n")
	out.WriteString("declare void @free(i8*)\n")
	out.WriteString("declare i32 @printf(i8*, ...)\n")
	out.WriteString("declare i32 @scanf(i8*, ...)\n")
	
	for i, fstr := range p.PrintStrs {
		llvmMod, goMod := p.ReplaceForLLVMFormat(fstr)
		uqStr, _ := strconv.Unquote(goMod)
		uqStrLen := len(uqStr)
		str := fmt.Sprintf("@.fstr%d = private unnamed_addr constant [%d x i8] c%s, align 1\n", i+1, uqStrLen, llvmMod)
		out.WriteString(str)
	}

	return out.String()
}

func (p *LLVMProgram) GetRegister(funcId string, varId string) *LLVMRegister {
	funcEntry, _ := p.Tables.Funcs.Contains(funcId)
	if varEntry, exists := funcEntry.Variables.Contains(varId); exists {
		return p.EntryRegMap[varEntry]
	} else {
		fmt.Println("VarEntry not found on EntryRegMap: " + funcId + " " + varId)
		return nil
	}
}

func (p *LLVMProgram) ReplaceForLLVMFormat(target string) (string, string) {
	var llvmMod, goMod string

	llvmMod = strings.ReplaceAll(target, "%d", "%ld")
	goMod = strings.ReplaceAll(target, "%d", "%ld")

	llvmMod = strings.ReplaceAll(llvmMod, "\\n", "\\0A")
	goMod = strings.ReplaceAll(goMod, "\\n", "\\x0A")

	llvmMod = llvmMod[:len(llvmMod)-1] + "\\00" + llvmMod[len(llvmMod)-1:]
	goMod = goMod[:len(goMod)-1] + "\\x00" + goMod[len(goMod)-1:]

	return llvmMod, goMod
}

func (p *LLVMProgram) AddFormatStr(str string) int {
	p.PrintStrs = append(p.PrintStrs, str)
	return len(p.PrintStrs)
}

func (p *LLVMProgram) GenerateRegisterName() string {
	val := p.RegCount
	p.RegCount++
	return fmt.Sprintf("r%d", val)
}

func (p *LLVMProgram) GenerateLabel(id string) string {
	val := p.LabelCount[id]
	p.LabelCount[id]++
	return fmt.Sprintf("L%d", val)
}

func (p *LLVMProgram) GenerateLabelsBatch(id string, count int) []string {
	var labels []string

	for i := 0; i < count; i++ {
		labels = append(labels, p.GenerateLabel(id))
	}

	return labels
}

func (p *LLVMProgram) AddEntryRegister(varEntry *st.VarEntry, reg *LLVMRegister) {
	p.EntryRegMap[varEntry] = reg
}

func (p *LLVMProgram) UpdateVariableState(currBlk *cfg.Block, varName string, currOperand LLVMOperand) {
	p.Defs[currBlk][varName] = currOperand
}

func (p *LLVMProgram) GetVariableState(currBlk *cfg.Block, varName string) LLVMOperand {
	return p.Defs[currBlk][varName]
}

func (p *LLVMProgram) TranslateFromStackToReg() {
	for fnName, entryBlk := range p.Funcs {
		funcEntry, _ := p.Tables.Funcs.Contains(fnName)
		visitedMap := make(map[string]int)
		p.Mem2Reg(entryBlk, nil, funcEntry, visitedMap)
	}
}

func (p *LLVMProgram) CreateNewPhiInstn(varEntry *st.VarEntry, funcEntry *st.FuncEntry, currBlk *cfg.Block, prevBlk *cfg.Block) *PhiInstn {
	var stackReg *LLVMRegister
	if varEntry != nil {
		stackReg = p.EntryRegMap[varEntry]
	} else {
		stackReg = p.RetRegisters[funcEntry.Name]
		varEntry = stackReg.entry
	}

	prevBlkOp := p.Defs[prevBlk][stackReg.GetName()]
	if prevBlkOp == nil {
		fmt.Printf("Nil at %s %s %s\n", funcEntry.Name, currBlk.Label, stackReg.GetName())
	}
	destName := fmt.Sprintf("%s_%s", stackReg.GetName(), currBlk.Label[1:])
	destReg := NewLLVMRegister(destName, varEntry, false)
	valMap := make(map[*cfg.Block]LLVMOperand)
	valMap[prevBlk] = prevBlkOp

	phiInstn := NewPhiInstn(destReg, varEntry, valMap)
	p.Defs[currBlk][stackReg.GetName()] = destReg
	return phiInstn
}

func (p *LLVMProgram) Mem2Reg(currBlk *cfg.Block, prevBlk *cfg.Block, funcEntry *st.FuncEntry, visitedMap map[string]int) {
	visitedMap[currBlk.Label]++
	var defs map[string]LLVMOperand
	var prevLabel string

	p.Defs[currBlk] = make(map[string]LLVMOperand)
	if prevBlk != nil {
    	for name, operand := range p.Defs[prevBlk] {
        	p.Defs[currBlk][name] = operand
    	}
		prevLabel = prevBlk.Label	
	}

	newInstns := make([]cfg.Instruction, 0)

	if len(currBlk.Preds) > 1 {
		if visitedMap[currBlk.Label] == 1 {
			for _, paramEntry := range funcEntry.Parameters.GetTable() {
				phiInstn := p.CreateNewPhiInstn(paramEntry, funcEntry, currBlk, prevBlk)
				newInstns = append(newInstns, phiInstn)
			}

			for _, varEntry := range funcEntry.Variables.GetTable() {
				phiInstn := p.CreateNewPhiInstn(varEntry, funcEntry, currBlk, prevBlk)
				newInstns = append(newInstns, phiInstn)
			}

			// For return value
			phiInstn := p.CreateNewPhiInstn(nil, funcEntry, currBlk, prevBlk)
			newInstns = append(newInstns, phiInstn)
		} else {
			for _, instn := range currBlk.Instns {
				if phiInstn, ok := instn.(*PhiInstn); ok {
					stackReg := p.EntryRegMap[phiInstn.varEntry]
					prevBlkOp := p.Defs[prevBlk][stackReg.GetName()]

					phiInstn.AddVal(prevBlk, prevBlkOp)
				} else {
					return;
				}
			}
		}
	}

	defs = p.Defs[currBlk]
	for _, instn := range currBlk.Instns {
		llvmInstn := instn.(LLVMInstruction)
		if llvmInstn.Mem2Reg(defs, prevLabel, currBlk) {
			newInstns = append(newInstns, llvmInstn)
		}
	}
	currBlk.Instns = newInstns

	for _, succBlk := range currBlk.Succs {
		if visitedMap[succBlk.Label] < len(succBlk.Preds) {
			p.Mem2Reg(succBlk, currBlk, funcEntry, visitedMap)
		}
	}
}

func TopologicalSort(entryBlock *cfg.Block) []*cfg.Block {
	result := make([]*cfg.Block, 0)
	visited := make(map[*cfg.Block]bool)
	temp := make(map[*cfg.Block]bool)
	
	var visit func(block *cfg.Block)
	visit = func(block *cfg.Block) {
		if visited[block] {
			return
		}
		
		if temp[block] {
			return
		}
		
		temp[block] = true
	
		for i := len(block.Succs) - 1; i >= 0; i-- {
			visit(block.Succs[i])
		}
		
		if block.Exit != nil {
			visit(block.Exit)
		}
		
		temp[block] = false
		visited[block] = true
		result = append(result, block)
	}
	
	visit(entryBlock)
	
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	return result
}

func (p *LLVMProgram) LScanRegAlloc() {
	p.CalculateLiveRanges()
	p.AssignPhysicalRegs()
}

func (p *LLVMProgram) CalculateLiveRanges() {
	for fnName, entryBlock := range p.Funcs {
		p.VariableTable[fnName] = make(map[*LLVMRegister]*LiveRangeEntry)
		fnVarTable := p.VariableTable[fnName]
		sortedBlocks := TopologicalSort(entryBlock)

		lineCnt := 0
		for _, paramReg := range p.ParamRegisters[fnName] {
			fnVarTable[paramReg] = &LiveRangeEntry{lineCnt, -1}
		}

		blockLineMap := make(map[*cfg.Block]int)

		for _, block := range sortedBlocks {
			for _, instn := range block.Instns {
				lineCnt++
				llvmInstn := instn.(LLVMInstruction)
				regDef := llvmInstn.GetDef()

				if regDef != nil {
					if fnVarTable[regDef] == nil {
						fnVarTable[regDef] = &LiveRangeEntry{}
						fnVarTable[regDef].End = -1
					}
					p.VariableTable[fnName][regDef].Start = lineCnt
				}

				for _, regUsed := range llvmInstn.GetUses() {
					if _, ok := regUsed.(*LLVMRegister); ok {
						if  _, ok := llvmInstn.(*PhiInstn); ok {
							maxEndLine := -1

							for block := range llvmInstn.(*PhiInstn).valMap {
								if blockLineMap[block] > maxEndLine {
									maxEndLine = blockLineMap[block]
								}
							}

							fnVarTable[regUsed.(*LLVMRegister)].End = maxEndLine - 1
						} else {
							if regUsed != p.ReadScratch {
								fnVarTable[regUsed.(*LLVMRegister)].End = lineCnt - 1
							}
						}
					}
				}
			}
			blockLineMap[block] = lineCnt
		}
	}
}

func (p *LLVMProgram) AssignPhysicalRegs() {
	for fnName := range p.Funcs {
		fnVarTable := p.VariableTable[fnName]

		type TablePair struct {
			reg		*LLVMRegister
			entry	*LiveRangeEntry
		}

		var paramPairList []*TablePair
		paramRegSet  := make(map[*LLVMRegister]*LiveRangeEntry)

		for _, paramReg := range p.ParamRegisters[fnName] {
			paramPairList = append(paramPairList, &TablePair{paramReg, fnVarTable[paramReg]})
			paramRegSet[paramReg] = fnVarTable[paramReg]
		}

		var pairList []*TablePair
		for reg, entry := range fnVarTable {
			if _, exists := paramRegSet[reg]; !exists {
				pairList = append(pairList, &TablePair{reg, entry})
			}
		}

		sort.Slice(pairList, func(i, j int) bool {
			return pairList[i].entry.Start < pairList[j].entry.Start
		})

		pairList = append(paramPairList, pairList...)

		ra := NewRegisterAlloc()
		for _, pair := range pairList {
			ra.CheckExpiry(pair.entry.Start, fnVarTable)
			if pReg := ra.GetRegister(pair.reg); pReg != nil {
				pair.reg.pReg = pReg
			}
		}
	}
}

func (p *LLVMProgram) TranslateToQuasi() {
	for _, entryBlock := range p.Funcs { 
		sortedBlocks := TopologicalSort(entryBlock)
		for _, block := range sortedBlocks {

			i := 0
			for _, instn := range block.Instns {
				if phiInstn, ok := instn.(*PhiInstn); ok {
					for targetBlk, operand := range phiInstn.valMap {
						if reg, ok := operand.(*LLVMRegister); ok {
							if phiInstn.dest == reg {
								continue
							}
						}
						instns := targetBlk.Instns
						movInstn := NewMovInstn(phiInstn.dest, operand)

						if len(instns) == 0 {
							targetBlk.Instns = append(instns, movInstn)
						} else {
							targetBlk.Instns = append(instns[:len(instns)-1], movInstn, instns[len(instns)-1])
						}
					}
				} else {
					block.Instns[i] = instn
					i++
				}
			}
			block.Instns = block.Instns[:i]
		}
	}
}

func (p *LLVMProgram) TranslateToAssembly() *arm.ARMProgram {
	p.TranslateToQuasi()

	sortedFuncBlocks := make(map[string][]*cfg.Block)
	blockInstns :=  make(map[*cfg.Block][]arm.Instruction)
	for fnName, entryBlock := range p.Funcs {
		sortedFuncBlocks[fnName] = TopologicalSort(entryBlock) 

		for _, block := range sortedFuncBlocks[fnName] {
			armInstns := make([]arm.Instruction, 0)

			for _, instn := range block.Instns {
				armInstns = append(armInstns, instn.(LLVMInstruction).TranslateToAssembly(p.Tables)...)
			}
			blockInstns[block] = armInstns
		}
	} 

	armProgram := arm.NewARMProgram("armv8-a", p.Tables, sortedFuncBlocks, blockInstns)

	return armProgram
}