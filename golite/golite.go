package main

import (
	"flag"
	"golite/lexer"
	"golite/parser"
	"os"
)

func main() {
	lFlag := flag.Bool("l", false, "Lexer Mode")
	flag.Parse()

	inputSourcePath := os.Args[1]
	lexer := lexer.NewLexer(inputSourcePath)

	if (*lFlag) {
		// lexer.PrintAllTokens()
	} else {
		parser := parser.NewParser(lexer)
		parser.Parse()
	}
}