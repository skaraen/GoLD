// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lexer

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoliteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var golitelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golitelexerLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'type'", "'var'", "'int'", "'bool'", "'func'", "'struct'",
		"'delete'", "'printf'", "'scan'", "'if'", "'else'", "'for'", "'return'",
		"'new'", "'true'", "'false'", "'nil'", "'=='", "'!='", "'>='", "'<='",
		"'||'", "'&&'", "'>'", "'<'", "'='", "'+'", "'-'", "'*'", "'/'", "'('",
		"')'", "'{'", "'}'", "';'", "','", "'.'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "COMMENT", "TYPE", "VAR", "INT", "BOOL", "FUNC", "STRUCT", "DELETE",
		"PRINTF", "SCAN", "IF", "ELSE", "FOR", "RETURN", "NEW", "TRUE", "FALSE",
		"NIL", "EQUALS", "NEQUALS", "GTE", "LTE", "OR", "AND", "GT", "LT", "ASSIGN",
		"PLUS", "MINUS", "ASTERIX", "FSLASH", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "SEMICOLON", "COMMA", "DOT", "NOT", "STRING", "NUM", "ID",
		"WS",
	}
	staticData.ruleNames = []string{
		"COMMENT", "TYPE", "VAR", "INT", "BOOL", "FUNC", "STRUCT", "DELETE",
		"PRINTF", "SCAN", "IF", "ELSE", "FOR", "RETURN", "NEW", "TRUE", "FALSE",
		"NIL", "EQUALS", "NEQUALS", "GTE", "LTE", "OR", "AND", "GT", "LT", "ASSIGN",
		"PLUS", "MINUS", "ASTERIX", "FSLASH", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "SEMICOLON", "COMMA", "DOT", "NOT", "STRING", "NUM", "ID",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 266, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 92, 8, 0, 10, 0, 12, 0,
		95, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 236, 8,
		39, 10, 39, 12, 39, 239, 9, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40,
		246, 8, 40, 10, 40, 12, 40, 249, 9, 40, 3, 40, 251, 8, 40, 1, 41, 1, 41,
		5, 41, 255, 8, 41, 10, 41, 12, 41, 258, 9, 41, 1, 42, 4, 42, 261, 8, 42,
		11, 42, 12, 42, 262, 1, 42, 1, 42, 1, 237, 0, 43, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 1, 0, 6, 2, 0, 10, 10, 13, 13, 1, 0, 49, 57, 1, 0,
		48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 3, 0, 9,
		10, 13, 13, 32, 32, 271, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 1, 87, 1, 0, 0, 0, 3, 98, 1, 0, 0,
		0, 5, 103, 1, 0, 0, 0, 7, 107, 1, 0, 0, 0, 9, 111, 1, 0, 0, 0, 11, 116,
		1, 0, 0, 0, 13, 121, 1, 0, 0, 0, 15, 128, 1, 0, 0, 0, 17, 135, 1, 0, 0,
		0, 19, 142, 1, 0, 0, 0, 21, 147, 1, 0, 0, 0, 23, 150, 1, 0, 0, 0, 25, 155,
		1, 0, 0, 0, 27, 159, 1, 0, 0, 0, 29, 166, 1, 0, 0, 0, 31, 170, 1, 0, 0,
		0, 33, 175, 1, 0, 0, 0, 35, 181, 1, 0, 0, 0, 37, 185, 1, 0, 0, 0, 39, 188,
		1, 0, 0, 0, 41, 191, 1, 0, 0, 0, 43, 194, 1, 0, 0, 0, 45, 197, 1, 0, 0,
		0, 47, 200, 1, 0, 0, 0, 49, 203, 1, 0, 0, 0, 51, 205, 1, 0, 0, 0, 53, 207,
		1, 0, 0, 0, 55, 209, 1, 0, 0, 0, 57, 211, 1, 0, 0, 0, 59, 213, 1, 0, 0,
		0, 61, 215, 1, 0, 0, 0, 63, 217, 1, 0, 0, 0, 65, 219, 1, 0, 0, 0, 67, 221,
		1, 0, 0, 0, 69, 223, 1, 0, 0, 0, 71, 225, 1, 0, 0, 0, 73, 227, 1, 0, 0,
		0, 75, 229, 1, 0, 0, 0, 77, 231, 1, 0, 0, 0, 79, 233, 1, 0, 0, 0, 81, 250,
		1, 0, 0, 0, 83, 252, 1, 0, 0, 0, 85, 260, 1, 0, 0, 0, 87, 88, 5, 47, 0,
		0, 88, 89, 5, 47, 0, 0, 89, 93, 1, 0, 0, 0, 90, 92, 8, 0, 0, 0, 91, 90,
		1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0,
		94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 6, 0, 0, 0, 97, 2, 1, 0,
		0, 0, 98, 99, 5, 116, 0, 0, 99, 100, 5, 121, 0, 0, 100, 101, 5, 112, 0,
		0, 101, 102, 5, 101, 0, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5, 118, 0, 0,
		104, 105, 5, 97, 0, 0, 105, 106, 5, 114, 0, 0, 106, 6, 1, 0, 0, 0, 107,
		108, 5, 105, 0, 0, 108, 109, 5, 110, 0, 0, 109, 110, 5, 116, 0, 0, 110,
		8, 1, 0, 0, 0, 111, 112, 5, 98, 0, 0, 112, 113, 5, 111, 0, 0, 113, 114,
		5, 111, 0, 0, 114, 115, 5, 108, 0, 0, 115, 10, 1, 0, 0, 0, 116, 117, 5,
		102, 0, 0, 117, 118, 5, 117, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5,
		99, 0, 0, 120, 12, 1, 0, 0, 0, 121, 122, 5, 115, 0, 0, 122, 123, 5, 116,
		0, 0, 123, 124, 5, 114, 0, 0, 124, 125, 5, 117, 0, 0, 125, 126, 5, 99,
		0, 0, 126, 127, 5, 116, 0, 0, 127, 14, 1, 0, 0, 0, 128, 129, 5, 100, 0,
		0, 129, 130, 5, 101, 0, 0, 130, 131, 5, 108, 0, 0, 131, 132, 5, 101, 0,
		0, 132, 133, 5, 116, 0, 0, 133, 134, 5, 101, 0, 0, 134, 16, 1, 0, 0, 0,
		135, 136, 5, 112, 0, 0, 136, 137, 5, 114, 0, 0, 137, 138, 5, 105, 0, 0,
		138, 139, 5, 110, 0, 0, 139, 140, 5, 116, 0, 0, 140, 141, 5, 102, 0, 0,
		141, 18, 1, 0, 0, 0, 142, 143, 5, 115, 0, 0, 143, 144, 5, 99, 0, 0, 144,
		145, 5, 97, 0, 0, 145, 146, 5, 110, 0, 0, 146, 20, 1, 0, 0, 0, 147, 148,
		5, 105, 0, 0, 148, 149, 5, 102, 0, 0, 149, 22, 1, 0, 0, 0, 150, 151, 5,
		101, 0, 0, 151, 152, 5, 108, 0, 0, 152, 153, 5, 115, 0, 0, 153, 154, 5,
		101, 0, 0, 154, 24, 1, 0, 0, 0, 155, 156, 5, 102, 0, 0, 156, 157, 5, 111,
		0, 0, 157, 158, 5, 114, 0, 0, 158, 26, 1, 0, 0, 0, 159, 160, 5, 114, 0,
		0, 160, 161, 5, 101, 0, 0, 161, 162, 5, 116, 0, 0, 162, 163, 5, 117, 0,
		0, 163, 164, 5, 114, 0, 0, 164, 165, 5, 110, 0, 0, 165, 28, 1, 0, 0, 0,
		166, 167, 5, 110, 0, 0, 167, 168, 5, 101, 0, 0, 168, 169, 5, 119, 0, 0,
		169, 30, 1, 0, 0, 0, 170, 171, 5, 116, 0, 0, 171, 172, 5, 114, 0, 0, 172,
		173, 5, 117, 0, 0, 173, 174, 5, 101, 0, 0, 174, 32, 1, 0, 0, 0, 175, 176,
		5, 102, 0, 0, 176, 177, 5, 97, 0, 0, 177, 178, 5, 108, 0, 0, 178, 179,
		5, 115, 0, 0, 179, 180, 5, 101, 0, 0, 180, 34, 1, 0, 0, 0, 181, 182, 5,
		110, 0, 0, 182, 183, 5, 105, 0, 0, 183, 184, 5, 108, 0, 0, 184, 36, 1,
		0, 0, 0, 185, 186, 5, 61, 0, 0, 186, 187, 5, 61, 0, 0, 187, 38, 1, 0, 0,
		0, 188, 189, 5, 33, 0, 0, 189, 190, 5, 61, 0, 0, 190, 40, 1, 0, 0, 0, 191,
		192, 5, 62, 0, 0, 192, 193, 5, 61, 0, 0, 193, 42, 1, 0, 0, 0, 194, 195,
		5, 60, 0, 0, 195, 196, 5, 61, 0, 0, 196, 44, 1, 0, 0, 0, 197, 198, 5, 124,
		0, 0, 198, 199, 5, 124, 0, 0, 199, 46, 1, 0, 0, 0, 200, 201, 5, 38, 0,
		0, 201, 202, 5, 38, 0, 0, 202, 48, 1, 0, 0, 0, 203, 204, 5, 62, 0, 0, 204,
		50, 1, 0, 0, 0, 205, 206, 5, 60, 0, 0, 206, 52, 1, 0, 0, 0, 207, 208, 5,
		61, 0, 0, 208, 54, 1, 0, 0, 0, 209, 210, 5, 43, 0, 0, 210, 56, 1, 0, 0,
		0, 211, 212, 5, 45, 0, 0, 212, 58, 1, 0, 0, 0, 213, 214, 5, 42, 0, 0, 214,
		60, 1, 0, 0, 0, 215, 216, 5, 47, 0, 0, 216, 62, 1, 0, 0, 0, 217, 218, 5,
		40, 0, 0, 218, 64, 1, 0, 0, 0, 219, 220, 5, 41, 0, 0, 220, 66, 1, 0, 0,
		0, 221, 222, 5, 123, 0, 0, 222, 68, 1, 0, 0, 0, 223, 224, 5, 125, 0, 0,
		224, 70, 1, 0, 0, 0, 225, 226, 5, 59, 0, 0, 226, 72, 1, 0, 0, 0, 227, 228,
		5, 44, 0, 0, 228, 74, 1, 0, 0, 0, 229, 230, 5, 46, 0, 0, 230, 76, 1, 0,
		0, 0, 231, 232, 5, 33, 0, 0, 232, 78, 1, 0, 0, 0, 233, 237, 5, 34, 0, 0,
		234, 236, 9, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		238, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 241, 5, 34, 0, 0, 241, 80, 1, 0, 0, 0, 242, 251, 5, 48,
		0, 0, 243, 247, 7, 1, 0, 0, 244, 246, 7, 2, 0, 0, 245, 244, 1, 0, 0, 0,
		246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248,
		251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 242, 1, 0, 0, 0, 250, 243,
		1, 0, 0, 0, 251, 82, 1, 0, 0, 0, 252, 256, 7, 3, 0, 0, 253, 255, 7, 4,
		0, 0, 254, 253, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0,
		256, 257, 1, 0, 0, 0, 257, 84, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 261,
		7, 5, 0, 0, 260, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 260, 1, 0,
		0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 6, 42, 0, 0,
		265, 86, 1, 0, 0, 0, 7, 0, 93, 237, 247, 250, 256, 262, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoliteLexerInit initializes any static state used to implement GoliteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoliteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.once.Do(golitelexerLexerInit)
}

// NewGoliteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoliteLexer(input antlr.CharStream) *GoliteLexer {
	GoliteLexerInit()
	l := new(GoliteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &golitelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "GoliteLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoliteLexer tokens.
const (
	GoliteLexerCOMMENT   = 1
	GoliteLexerTYPE      = 2
	GoliteLexerVAR       = 3
	GoliteLexerINT       = 4
	GoliteLexerBOOL      = 5
	GoliteLexerFUNC      = 6
	GoliteLexerSTRUCT    = 7
	GoliteLexerDELETE    = 8
	GoliteLexerPRINTF    = 9
	GoliteLexerSCAN      = 10
	GoliteLexerIF        = 11
	GoliteLexerELSE      = 12
	GoliteLexerFOR       = 13
	GoliteLexerRETURN    = 14
	GoliteLexerNEW       = 15
	GoliteLexerTRUE      = 16
	GoliteLexerFALSE     = 17
	GoliteLexerNIL       = 18
	GoliteLexerEQUALS    = 19
	GoliteLexerNEQUALS   = 20
	GoliteLexerGTE       = 21
	GoliteLexerLTE       = 22
	GoliteLexerOR        = 23
	GoliteLexerAND       = 24
	GoliteLexerGT        = 25
	GoliteLexerLT        = 26
	GoliteLexerASSIGN    = 27
	GoliteLexerPLUS      = 28
	GoliteLexerMINUS     = 29
	GoliteLexerASTERIX   = 30
	GoliteLexerFSLASH    = 31
	GoliteLexerLPAREN    = 32
	GoliteLexerRPAREN    = 33
	GoliteLexerLBRACE    = 34
	GoliteLexerRBRACE    = 35
	GoliteLexerSEMICOLON = 36
	GoliteLexerCOMMA     = 37
	GoliteLexerDOT       = 38
	GoliteLexerNOT       = 39
	GoliteLexerSTRING    = 40
	GoliteLexerNUM       = 41
	GoliteLexerID        = 42
	GoliteLexerWS        = 43
)
