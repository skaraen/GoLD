package llvm

import (
	"bytes"
	"fmt"
	"golite/cfg"
	st "golite/symboltable"
	"golite/types"
	"golite/utils"
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
	RetRegisters		map[string]*LLVMRegister
	PrintStrs			[]string
	RCount				int
	LCount				map[string]int
}

func NewLLVMProgram(srcName string, targetTriple string, tables *st.SymbolTables) *LLVMProgram {
	return &LLVMProgram{srcName, targetTriple, tables, make([]LLVMInstruction, 0), make([]LLVMInstruction, 0), 
		make(map[string]*cfg.Block), make(map[*st.VarEntry]*LLVMRegister), 
		make(map[string]*LLVMRegister), make([]string, 0), 0, make(map[string]int, 0)}
}

func (p *LLVMProgram) String() string {
	var out bytes.Buffer
	p.Tables.Funcs.GetTable()["main"].RetTy = types.IntTySig

	out.WriteString("source_name = " + "\"" + p.SourceName + "\"\n")
	out.WriteString("target_triple = " + "\"" + p.TargetTriple + "\"\n")

	out.WriteString("\n")

	for structName, structEntry := range p.Tables.Structs.GetTable() {
		var fieldTypes []string

		for _, fieldEntry := range (structEntry.Fields.GetTable()) {
			fieldTypes = append(fieldTypes, types.TypesToLLVMTypes(fieldEntry.Ty))
		}

		fields := strings.Join(fieldTypes, ", ")
		str := fmt.Sprintf("%%struct.%s = type {%s}\n", structName, fields)

		out.WriteString(str) 
	}

	out.WriteString("\n")

	for varName, varEntry := range p.Tables.Globals.GetTable() {
		str := fmt.Sprintf("@%s = common global %s %s\n", varName, types.TypesToLLVMTypes(varEntry.Ty), types.GetLLVMDefaultValue(varEntry.Ty))
		out.WriteString(str)
	}

	out.WriteString("\n")

	for fnName, entryBlk := range p.Funcs {
		fnEntry,_ := p.Tables.Funcs.Contains(fnName)

		var paramStr []string
		for _, param := range fnEntry.Parameters.GetTable() {
			reg := p.EntryRegMap[param]
			str := fmt.Sprintf("%s %s", types.TypesToLLVMTypes(reg.GetType()), reg.String())
			paramStr = append(paramStr, str)
		}

		out.WriteString(fmt.Sprintf("define %s @%s(%s)\n", types.TypesToLLVMTypes(fnEntry.RetTy), fnName, strings.Join(paramStr, ", ")))
		out.WriteString("{\n")

		blockQueue := utils.NewQueue[*cfg.Block]()
		discoveredMap := make(map[string]*cfg.Block)
		blockQueue.Enqueue(entryBlk)

		for {
			block, ok := blockQueue.Dequeue()
			if !ok {
				break 
			}

			for _, succBlock := range block.Succs {
				if _, exists := discoveredMap[succBlock.Label]; exists {
					continue
				}
				blockQueue.Enqueue(succBlock)
				discoveredMap[succBlock.Label] = succBlock
			}

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

func (p *LLVMProgram) AddFormatStr(str string) int {
	p.PrintStrs = append(p.PrintStrs, str)
	return len(p.PrintStrs)
}

func (p *LLVMProgram) GenerateRegisterName() string {
	val := p.RCount
	p.RCount++
	return fmt.Sprintf("r%d", val)
}

func (p *LLVMProgram) GenerateLabel(id string) string {
	val := p.LCount[id]
	p.LCount[id]++
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