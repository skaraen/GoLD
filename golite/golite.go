package main

import (
	"flag"
	"fmt"
	"golite/lexer"
	"golite/parser"
	"golite/sa"
)

func main() {
	lFlag := flag.Bool("l", false, "Lexer Mode")
	astFlag := flag.Bool("ast", false, "Print AST")
	flag.Parse()

	inputSourcePath := flag.Arg(0)
	lexer := lexer.NewLexer(inputSourcePath)

	if (*lFlag) {
		lexer.PrintAllTokens()
	}

	parser := parser.NewParser(lexer)
	program := parser.Parse()

	if (*astFlag) {
		fmt.Println(parser.Parse())
	}

	sa.Execute(program)
}