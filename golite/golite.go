package main

import (
	"golite/lexer"
	"golite/parser"
	"os"
)

func main() {
	inputSourcePath := os.Args[1]
	lexer := lexer.NewLexer(inputSourcePath)
	parser := parser.NewParser(lexer)
	parser.Parse()
}