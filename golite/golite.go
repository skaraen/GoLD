package main

import (
	"flag"
	"fmt"
	"golite/lexer"
	"golite/parser"
	"golite/sa"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	lFlag := flag.Bool("l", false, "Lexer Mode")
	astFlag := flag.Bool("ast", false, "Print AST")
	llvmIRFlag := flag.String("llvm-ir", "", "Specify LLVM IR phase: 'stack' or 'reg'")
	assemblyFlag := flag.Bool("S", false, "Assembly code file")
	flag.Parse()

	inputSourcePath := flag.Arg(0)
	fileNameWithExt := filepath.Base(inputSourcePath)
	sourceName := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))
	llvmStackFile := sourceName + "_stack.ll"
	llvmRegFile := sourceName + "_reg.ll"
	armMachineFile := sourceName + "_arm.s"
	targetTriple := "x86_64-linux-gnu"

	lexer := lexer.NewLexer(inputSourcePath)

	if (*lFlag) {
		lexer.PrintAllTokens()
	}

	parser := parser.NewParser(lexer)
	program := parser.Parse()

	if (*astFlag) {
		fmt.Println(parser.Parse())
	}

	llvmProgram := sa.Execute(program, sourceName, targetTriple)
	if *llvmIRFlag == "stack" {
		os.WriteFile(llvmStackFile, []byte(llvmProgram.String()), 0644)
		return
	}

	llvmProgram.TranslateFromStackToReg()
	if *llvmIRFlag == "reg" {
		os.WriteFile(llvmRegFile, []byte(llvmProgram.String()), 0644)
		return
	}

	llvmProgram.LScanRegAlloc()
	armProgram := llvmProgram.TranslateToAssembly()
	if *assemblyFlag {
		os.WriteFile(armMachineFile, []byte(armProgram.String()), 0644)
	} 
}