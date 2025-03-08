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
	stackFlag := flag.Bool("llvm-stack", false, "LLVM Stack")
	flag.Parse()

	inputSourcePath := flag.Arg(0)
	fileNameWithExt := filepath.Base(inputSourcePath)
	sourceName := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))
	llvmFile := sourceName + ".ll"
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
	if (*stackFlag) {
		os.WriteFile(llvmFile, []byte(llvmProgram.String()), 0644)
	}

	fmt.Print(llvmProgram)
}