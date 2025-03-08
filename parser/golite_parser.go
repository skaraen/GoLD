// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoliteParser struct {
	*antlr.BaseParser
}

var goliteparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goliteparserParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.literalNames = []string{
		"", "'type'", "'var'", "'int'", "'bool'", "'func'", "'struct'", "'delete'",
		"'printf'", "'scan'", "'if'", "'else'", "'for'", "'return'", "'new'",
		"'true'", "'false'", "'nil'", "'=='", "'!='", "'>='", "'<='", "'||'",
		"'&&'", "'>'", "'<'", "'='", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
		"'{'", "'}'", "';'", "','", "'.'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "TYPE", "VAR", "INT", "BOOL", "FUNC", "STRUCT", "DELETE", "PRINTF",
		"SCAN", "IF", "ELSE", "FOR", "RETURN", "NEW", "TRUE", "FALSE", "NIL",
		"EQUALS", "NEQUALS", "GTE", "LTE", "OR", "AND", "GT", "LT", "ASSIGN",
		"PLUS", "MINUS", "ASTERIX", "FSLASH", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "SEMICOLON", "COMMA", "DOT", "NOT", "STRING", "NUM", "ID",
		"WS",
	}
	staticData.ruleNames = []string{
		"program", "userTypes", "typeDecl", "fields", "fieldsPrime", "decl",
		"reference", "type", "declarations", "declaration", "ids", "idsPrime",
		"functions", "function", "params", "paramsPrime", "returnType", "statements",
		"statement", "block", "delete", "read", "assignment", "print", "conditional",
		"loop", "return", "invocation", "args", "argsPrime", "lValue", "lValuePrime",
		"expression", "exprPrime", "boolTerm", "boolTermPrime", "equalTermPrime",
		"equalTerm", "relTermPrime", "relTerm", "simpleTermPrime", "simpleTerm",
		"termPrime", "term", "unaryTerm", "selectorTerm", "factor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 375, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 101, 8, 1, 10, 1, 12, 1, 104, 9, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 116,
		8, 3, 10, 3, 12, 3, 119, 9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 133, 8, 7, 1, 8, 5, 8, 136, 8, 8, 10,
		8, 12, 8, 139, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10,
		148, 8, 10, 10, 10, 12, 10, 151, 9, 10, 1, 11, 1, 11, 1, 11, 1, 12, 5,
		12, 157, 8, 12, 10, 12, 12, 12, 160, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		3, 13, 166, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 5, 14, 176, 8, 14, 10, 14, 12, 14, 179, 9, 14, 3, 14, 181, 8, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 5, 17, 191, 8, 17,
		10, 17, 12, 17, 194, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 204, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 227, 8, 23, 10, 23, 12, 23, 230,
		9, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 242, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 3, 26, 252, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 5, 28, 263, 8, 28, 10, 28, 12, 28, 266, 9, 28, 3, 28,
		268, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 277,
		8, 30, 10, 30, 12, 30, 280, 9, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5,
		32, 287, 8, 32, 10, 32, 12, 32, 290, 9, 32, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 5, 34, 297, 8, 34, 10, 34, 12, 34, 300, 9, 34, 1, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 310, 8, 37, 10, 37, 12, 37,
		313, 9, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 320, 8, 39, 10, 39,
		12, 39, 323, 9, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 5, 41, 330, 8, 41,
		10, 41, 12, 41, 333, 9, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 5, 43, 340,
		8, 43, 10, 43, 12, 43, 343, 9, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3,
		44, 350, 8, 44, 1, 45, 1, 45, 5, 45, 354, 8, 45, 10, 45, 12, 45, 357, 9,
		45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 365, 8, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 373, 8, 46, 1, 46, 0, 0, 47,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 4, 1, 0, 18, 19, 2, 0, 20, 21,
		24, 25, 1, 0, 27, 28, 1, 0, 29, 30, 367, 0, 94, 1, 0, 0, 0, 2, 102, 1,
		0, 0, 0, 4, 105, 1, 0, 0, 0, 6, 113, 1, 0, 0, 0, 8, 120, 1, 0, 0, 0, 10,
		123, 1, 0, 0, 0, 12, 126, 1, 0, 0, 0, 14, 132, 1, 0, 0, 0, 16, 137, 1,
		0, 0, 0, 18, 140, 1, 0, 0, 0, 20, 145, 1, 0, 0, 0, 22, 152, 1, 0, 0, 0,
		24, 158, 1, 0, 0, 0, 26, 161, 1, 0, 0, 0, 28, 172, 1, 0, 0, 0, 30, 184,
		1, 0, 0, 0, 32, 187, 1, 0, 0, 0, 34, 192, 1, 0, 0, 0, 36, 203, 1, 0, 0,
		0, 38, 205, 1, 0, 0, 0, 40, 209, 1, 0, 0, 0, 42, 213, 1, 0, 0, 0, 44, 217,
		1, 0, 0, 0, 46, 222, 1, 0, 0, 0, 48, 234, 1, 0, 0, 0, 50, 243, 1, 0, 0,
		0, 52, 249, 1, 0, 0, 0, 54, 255, 1, 0, 0, 0, 56, 259, 1, 0, 0, 0, 58, 271,
		1, 0, 0, 0, 60, 274, 1, 0, 0, 0, 62, 281, 1, 0, 0, 0, 64, 284, 1, 0, 0,
		0, 66, 291, 1, 0, 0, 0, 68, 294, 1, 0, 0, 0, 70, 301, 1, 0, 0, 0, 72, 304,
		1, 0, 0, 0, 74, 307, 1, 0, 0, 0, 76, 314, 1, 0, 0, 0, 78, 317, 1, 0, 0,
		0, 80, 324, 1, 0, 0, 0, 82, 327, 1, 0, 0, 0, 84, 334, 1, 0, 0, 0, 86, 337,
		1, 0, 0, 0, 88, 349, 1, 0, 0, 0, 90, 351, 1, 0, 0, 0, 92, 372, 1, 0, 0,
		0, 94, 95, 3, 2, 1, 0, 95, 96, 3, 16, 8, 0, 96, 97, 3, 24, 12, 0, 97, 98,
		5, 0, 0, 1, 98, 1, 1, 0, 0, 0, 99, 101, 3, 4, 2, 0, 100, 99, 1, 0, 0, 0,
		101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103,
		3, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 1, 0, 0, 106, 107, 5,
		41, 0, 0, 107, 108, 5, 6, 0, 0, 108, 109, 5, 33, 0, 0, 109, 110, 3, 6,
		3, 0, 110, 111, 5, 34, 0, 0, 111, 112, 5, 35, 0, 0, 112, 5, 1, 0, 0, 0,
		113, 117, 3, 8, 4, 0, 114, 116, 3, 8, 4, 0, 115, 114, 1, 0, 0, 0, 116,
		119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 7, 1,
		0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 121, 3, 10, 5, 0, 121, 122, 5, 35,
		0, 0, 122, 9, 1, 0, 0, 0, 123, 124, 5, 41, 0, 0, 124, 125, 3, 14, 7, 0,
		125, 11, 1, 0, 0, 0, 126, 127, 5, 29, 0, 0, 127, 128, 5, 41, 0, 0, 128,
		13, 1, 0, 0, 0, 129, 133, 5, 3, 0, 0, 130, 133, 5, 4, 0, 0, 131, 133, 3,
		12, 6, 0, 132, 129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1, 0, 0,
		0, 133, 15, 1, 0, 0, 0, 134, 136, 3, 18, 9, 0, 135, 134, 1, 0, 0, 0, 136,
		139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 17, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 2, 0, 0, 141, 142, 3, 20, 10,
		0, 142, 143, 3, 14, 7, 0, 143, 144, 5, 35, 0, 0, 144, 19, 1, 0, 0, 0, 145,
		149, 5, 41, 0, 0, 146, 148, 3, 22, 11, 0, 147, 146, 1, 0, 0, 0, 148, 151,
		1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 21, 1, 0,
		0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 36, 0, 0, 153, 154, 5, 41, 0,
		0, 154, 23, 1, 0, 0, 0, 155, 157, 3, 26, 13, 0, 156, 155, 1, 0, 0, 0, 157,
		160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 25, 1,
		0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 162, 5, 5, 0, 0, 162, 163, 5, 41, 0,
		0, 163, 165, 3, 28, 14, 0, 164, 166, 3, 32, 16, 0, 165, 164, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 5, 33, 0, 0, 168,
		169, 3, 16, 8, 0, 169, 170, 3, 34, 17, 0, 170, 171, 5, 34, 0, 0, 171, 27,
		1, 0, 0, 0, 172, 180, 5, 31, 0, 0, 173, 177, 3, 10, 5, 0, 174, 176, 3,
		30, 15, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0,
		0, 0, 177, 178, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0,
		180, 173, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		183, 5, 32, 0, 0, 183, 29, 1, 0, 0, 0, 184, 185, 5, 36, 0, 0, 185, 186,
		3, 10, 5, 0, 186, 31, 1, 0, 0, 0, 187, 188, 3, 14, 7, 0, 188, 33, 1, 0,
		0, 0, 189, 191, 3, 36, 18, 0, 190, 189, 1, 0, 0, 0, 191, 194, 1, 0, 0,
		0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 35, 1, 0, 0, 0, 194,
		192, 1, 0, 0, 0, 195, 204, 3, 44, 22, 0, 196, 204, 3, 46, 23, 0, 197, 204,
		3, 42, 21, 0, 198, 204, 3, 40, 20, 0, 199, 204, 3, 48, 24, 0, 200, 204,
		3, 50, 25, 0, 201, 204, 3, 52, 26, 0, 202, 204, 3, 54, 27, 0, 203, 195,
		1, 0, 0, 0, 203, 196, 1, 0, 0, 0, 203, 197, 1, 0, 0, 0, 203, 198, 1, 0,
		0, 0, 203, 199, 1, 0, 0, 0, 203, 200, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0,
		203, 202, 1, 0, 0, 0, 204, 37, 1, 0, 0, 0, 205, 206, 5, 33, 0, 0, 206,
		207, 3, 34, 17, 0, 207, 208, 5, 34, 0, 0, 208, 39, 1, 0, 0, 0, 209, 210,
		5, 7, 0, 0, 210, 211, 3, 64, 32, 0, 211, 212, 5, 35, 0, 0, 212, 41, 1,
		0, 0, 0, 213, 214, 5, 9, 0, 0, 214, 215, 3, 60, 30, 0, 215, 216, 5, 35,
		0, 0, 216, 43, 1, 0, 0, 0, 217, 218, 3, 60, 30, 0, 218, 219, 5, 26, 0,
		0, 219, 220, 3, 64, 32, 0, 220, 221, 5, 35, 0, 0, 221, 45, 1, 0, 0, 0,
		222, 223, 5, 8, 0, 0, 223, 224, 5, 31, 0, 0, 224, 228, 5, 39, 0, 0, 225,
		227, 3, 58, 29, 0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226,
		1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 231, 1, 0, 0, 0, 230, 228, 1, 0,
		0, 0, 231, 232, 5, 32, 0, 0, 232, 233, 5, 35, 0, 0, 233, 47, 1, 0, 0, 0,
		234, 235, 5, 10, 0, 0, 235, 236, 5, 31, 0, 0, 236, 237, 3, 64, 32, 0, 237,
		238, 5, 32, 0, 0, 238, 241, 3, 38, 19, 0, 239, 240, 5, 11, 0, 0, 240, 242,
		3, 38, 19, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 49, 1, 0,
		0, 0, 243, 244, 5, 12, 0, 0, 244, 245, 5, 31, 0, 0, 245, 246, 3, 64, 32,
		0, 246, 247, 5, 32, 0, 0, 247, 248, 3, 38, 19, 0, 248, 51, 1, 0, 0, 0,
		249, 251, 5, 13, 0, 0, 250, 252, 3, 64, 32, 0, 251, 250, 1, 0, 0, 0, 251,
		252, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 5, 35, 0, 0, 254, 53,
		1, 0, 0, 0, 255, 256, 5, 41, 0, 0, 256, 257, 3, 56, 28, 0, 257, 258, 5,
		35, 0, 0, 258, 55, 1, 0, 0, 0, 259, 267, 5, 31, 0, 0, 260, 264, 3, 64,
		32, 0, 261, 263, 3, 58, 29, 0, 262, 261, 1, 0, 0, 0, 263, 266, 1, 0, 0,
		0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266,
		264, 1, 0, 0, 0, 267, 260, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269,
		1, 0, 0, 0, 269, 270, 5, 32, 0, 0, 270, 57, 1, 0, 0, 0, 271, 272, 5, 36,
		0, 0, 272, 273, 3, 64, 32, 0, 273, 59, 1, 0, 0, 0, 274, 278, 5, 41, 0,
		0, 275, 277, 3, 62, 31, 0, 276, 275, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0,
		278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 61, 1, 0, 0, 0, 280, 278,
		1, 0, 0, 0, 281, 282, 5, 37, 0, 0, 282, 283, 5, 41, 0, 0, 283, 63, 1, 0,
		0, 0, 284, 288, 3, 68, 34, 0, 285, 287, 3, 66, 33, 0, 286, 285, 1, 0, 0,
		0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		65, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 292, 5, 22, 0, 0, 292, 293,
		3, 68, 34, 0, 293, 67, 1, 0, 0, 0, 294, 298, 3, 74, 37, 0, 295, 297, 3,
		70, 35, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0,
		0, 0, 298, 299, 1, 0, 0, 0, 299, 69, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0,
		301, 302, 5, 23, 0, 0, 302, 303, 3, 74, 37, 0, 303, 71, 1, 0, 0, 0, 304,
		305, 7, 0, 0, 0, 305, 306, 3, 78, 39, 0, 306, 73, 1, 0, 0, 0, 307, 311,
		3, 78, 39, 0, 308, 310, 3, 72, 36, 0, 309, 308, 1, 0, 0, 0, 310, 313, 1,
		0, 0, 0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 75, 1, 0, 0,
		0, 313, 311, 1, 0, 0, 0, 314, 315, 7, 1, 0, 0, 315, 316, 3, 82, 41, 0,
		316, 77, 1, 0, 0, 0, 317, 321, 3, 82, 41, 0, 318, 320, 3, 76, 38, 0, 319,
		318, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 322,
		1, 0, 0, 0, 322, 79, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 325, 7, 2,
		0, 0, 325, 326, 3, 86, 43, 0, 326, 81, 1, 0, 0, 0, 327, 331, 3, 86, 43,
		0, 328, 330, 3, 80, 40, 0, 329, 328, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0,
		331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 83, 1, 0, 0, 0, 333, 331,
		1, 0, 0, 0, 334, 335, 7, 3, 0, 0, 335, 336, 3, 88, 44, 0, 336, 85, 1, 0,
		0, 0, 337, 341, 3, 88, 44, 0, 338, 340, 3, 84, 42, 0, 339, 338, 1, 0, 0,
		0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342,
		87, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 345, 5, 38, 0, 0, 345, 350,
		3, 90, 45, 0, 346, 347, 5, 28, 0, 0, 347, 350, 3, 90, 45, 0, 348, 350,
		3, 90, 45, 0, 349, 344, 1, 0, 0, 0, 349, 346, 1, 0, 0, 0, 349, 348, 1,
		0, 0, 0, 350, 89, 1, 0, 0, 0, 351, 355, 3, 92, 46, 0, 352, 354, 3, 62,
		31, 0, 353, 352, 1, 0, 0, 0, 354, 357, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0,
		355, 356, 1, 0, 0, 0, 356, 91, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 358, 359,
		5, 31, 0, 0, 359, 360, 3, 64, 32, 0, 360, 361, 5, 32, 0, 0, 361, 373, 1,
		0, 0, 0, 362, 364, 5, 41, 0, 0, 363, 365, 3, 56, 28, 0, 364, 363, 1, 0,
		0, 0, 364, 365, 1, 0, 0, 0, 365, 373, 1, 0, 0, 0, 366, 373, 5, 40, 0, 0,
		367, 368, 5, 14, 0, 0, 368, 373, 5, 41, 0, 0, 369, 373, 5, 15, 0, 0, 370,
		373, 5, 16, 0, 0, 371, 373, 5, 17, 0, 0, 372, 358, 1, 0, 0, 0, 372, 362,
		1, 0, 0, 0, 372, 366, 1, 0, 0, 0, 372, 367, 1, 0, 0, 0, 372, 369, 1, 0,
		0, 0, 372, 370, 1, 0, 0, 0, 372, 371, 1, 0, 0, 0, 373, 93, 1, 0, 0, 0,
		27, 102, 117, 132, 137, 149, 158, 165, 177, 180, 192, 203, 228, 241, 251,
		264, 267, 278, 288, 298, 311, 321, 331, 341, 349, 355, 364, 372,
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

// GoliteParserInit initializes any static state used to implement GoliteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoliteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.once.Do(goliteparserParserInit)
}

// NewGoliteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoliteParser(input antlr.TokenStream) *GoliteParser {
	GoliteParserInit()
	this := new(GoliteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &goliteparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// GoliteParser tokens.
const (
	GoliteParserEOF       = antlr.TokenEOF
	GoliteParserTYPE      = 1
	GoliteParserVAR       = 2
	GoliteParserINT       = 3
	GoliteParserBOOL      = 4
	GoliteParserFUNC      = 5
	GoliteParserSTRUCT    = 6
	GoliteParserDELETE    = 7
	GoliteParserPRINTF    = 8
	GoliteParserSCAN      = 9
	GoliteParserIF        = 10
	GoliteParserELSE      = 11
	GoliteParserFOR       = 12
	GoliteParserRETURN    = 13
	GoliteParserNEW       = 14
	GoliteParserTRUE      = 15
	GoliteParserFALSE     = 16
	GoliteParserNIL       = 17
	GoliteParserEQUALS    = 18
	GoliteParserNEQUALS   = 19
	GoliteParserGTE       = 20
	GoliteParserLTE       = 21
	GoliteParserOR        = 22
	GoliteParserAND       = 23
	GoliteParserGT        = 24
	GoliteParserLT        = 25
	GoliteParserASSIGN    = 26
	GoliteParserPLUS      = 27
	GoliteParserMINUS     = 28
	GoliteParserASTERIX   = 29
	GoliteParserFSLASH    = 30
	GoliteParserLPAREN    = 31
	GoliteParserRPAREN    = 32
	GoliteParserLBRACE    = 33
	GoliteParserRBRACE    = 34
	GoliteParserSEMICOLON = 35
	GoliteParserCOMMA     = 36
	GoliteParserDOT       = 37
	GoliteParserNOT       = 38
	GoliteParserSTRING    = 39
	GoliteParserNUM       = 40
	GoliteParserID        = 41
	GoliteParserWS        = 42
)

// GoliteParser rules.
const (
	GoliteParserRULE_program         = 0
	GoliteParserRULE_userTypes       = 1
	GoliteParserRULE_typeDecl        = 2
	GoliteParserRULE_fields          = 3
	GoliteParserRULE_fieldsPrime     = 4
	GoliteParserRULE_decl            = 5
	GoliteParserRULE_reference       = 6
	GoliteParserRULE_type            = 7
	GoliteParserRULE_declarations    = 8
	GoliteParserRULE_declaration     = 9
	GoliteParserRULE_ids             = 10
	GoliteParserRULE_idsPrime        = 11
	GoliteParserRULE_functions       = 12
	GoliteParserRULE_function        = 13
	GoliteParserRULE_params          = 14
	GoliteParserRULE_paramsPrime     = 15
	GoliteParserRULE_returnType      = 16
	GoliteParserRULE_statements      = 17
	GoliteParserRULE_statement       = 18
	GoliteParserRULE_block           = 19
	GoliteParserRULE_delete          = 20
	GoliteParserRULE_read            = 21
	GoliteParserRULE_assignment      = 22
	GoliteParserRULE_print           = 23
	GoliteParserRULE_conditional     = 24
	GoliteParserRULE_loop            = 25
	GoliteParserRULE_return          = 26
	GoliteParserRULE_invocation      = 27
	GoliteParserRULE_args            = 28
	GoliteParserRULE_argsPrime       = 29
	GoliteParserRULE_lValue          = 30
	GoliteParserRULE_lValuePrime     = 31
	GoliteParserRULE_expression      = 32
	GoliteParserRULE_exprPrime       = 33
	GoliteParserRULE_boolTerm        = 34
	GoliteParserRULE_boolTermPrime   = 35
	GoliteParserRULE_equalTermPrime  = 36
	GoliteParserRULE_equalTerm       = 37
	GoliteParserRULE_relTermPrime    = 38
	GoliteParserRULE_relTerm         = 39
	GoliteParserRULE_simpleTermPrime = 40
	GoliteParserRULE_simpleTerm      = 41
	GoliteParserRULE_termPrime       = 42
	GoliteParserRULE_term            = 43
	GoliteParserRULE_unaryTerm       = 44
	GoliteParserRULE_selectorTerm    = 45
	GoliteParserRULE_factor          = 46
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) UserTypes() IUserTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserTypesContext)
}

func (s *ProgramContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoliteParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GoliteParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoliteParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.UserTypes()
	}
	{
		p.SetState(95)
		p.Declarations()
	}
	{
		p.SetState(96)
		p.Functions()
	}
	{
		p.SetState(97)
		p.Match(GoliteParserEOF)
	}

	return localctx
}

// IUserTypesContext is an interface to support dynamic dispatch.
type IUserTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUserTypesContext differentiates from other interfaces.
	IsUserTypesContext()
}

type UserTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserTypesContext() *UserTypesContext {
	var p = new(UserTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_userTypes
	return p
}

func (*UserTypesContext) IsUserTypesContext() {}

func NewUserTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserTypesContext {
	var p = new(UserTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_userTypes

	return p
}

func (s *UserTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *UserTypesContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *UserTypesContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *UserTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUserTypes(s)
	}
}

func (s *UserTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUserTypes(s)
	}
}

func (p *GoliteParser) UserTypes() (localctx IUserTypesContext) {
	this := p
	_ = this

	localctx = NewUserTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoliteParserRULE_userTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserTYPE {
		{
			p.SetState(99)
			p.TypeDecl()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTYPE, 0)
}

func (s *TypeDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypeDeclContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRUCT, 0)
}

func (s *TypeDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *TypeDeclContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *TypeDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *TypeDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *GoliteParser) TypeDecl() (localctx ITypeDeclContext) {
	this := p
	_ = this

	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoliteParserRULE_typeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(GoliteParserTYPE)
	}
	{
		p.SetState(106)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(107)
		p.Match(GoliteParserSTRUCT)
	}
	{
		p.SetState(108)
		p.Match(GoliteParserLBRACE)
	}
	{
		p.SetState(109)
		p.Fields()
	}
	{
		p.SetState(110)
		p.Match(GoliteParserRBRACE)
	}
	{
		p.SetState(111)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllFieldsPrime() []IFieldsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IFieldsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldsPrimeContext); ok {
			tst[i] = t.(IFieldsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) FieldsPrime(i int) IFieldsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsPrimeContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *GoliteParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoliteParserRULE_fields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.FieldsPrime()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserID {
		{
			p.SetState(114)
			p.FieldsPrime()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldsPrimeContext is an interface to support dynamic dispatch.
type IFieldsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDeclr returns the declr rule contexts.
	GetDeclr() IDeclContext

	// SetDeclr sets the declr rule contexts.
	SetDeclr(IDeclContext)

	// IsFieldsPrimeContext differentiates from other interfaces.
	IsFieldsPrimeContext()
}

type FieldsPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	declr  IDeclContext
}

func NewEmptyFieldsPrimeContext() *FieldsPrimeContext {
	var p = new(FieldsPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_fieldsPrime
	return p
}

func (*FieldsPrimeContext) IsFieldsPrimeContext() {}

func NewFieldsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsPrimeContext {
	var p = new(FieldsPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fieldsPrime

	return p
}

func (s *FieldsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsPrimeContext) GetDeclr() IDeclContext { return s.declr }

func (s *FieldsPrimeContext) SetDeclr(v IDeclContext) { s.declr = v }

func (s *FieldsPrimeContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *FieldsPrimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFieldsPrime(s)
	}
}

func (s *FieldsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFieldsPrime(s)
	}
}

func (p *GoliteParser) FieldsPrime() (localctx IFieldsPrimeContext) {
	this := p
	_ = this

	localctx = NewFieldsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoliteParserRULE_fieldsPrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)

		var _x = p.Decl()

		localctx.(*FieldsPrimeContext).declr = _x
	}
	{
		p.SetState(121)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *GoliteParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoliteParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(124)
		p.Type_()
	}

	return localctx
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_reference
	return p
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) GetName() antlr.Token { return s.name }

func (s *ReferenceContext) SetName(v antlr.Token) { s.name = v }

func (s *ReferenceContext) ASTERIX() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERIX, 0)
}

func (s *ReferenceContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *GoliteParser) Reference() (localctx IReferenceContext) {
	this := p
	_ = this

	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoliteParserRULE_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(GoliteParserASTERIX)
	}
	{
		p.SetState(127)

		var _m = p.Match(GoliteParserID)

		localctx.(*ReferenceContext).name = _m
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInteger returns the integer token.
	GetInteger() antlr.Token

	// GetBoolean returns the boolean token.
	GetBoolean() antlr.Token

	// SetInteger sets the integer token.
	SetInteger(antlr.Token)

	// SetBoolean sets the boolean token.
	SetBoolean(antlr.Token)

	// GetPointer returns the pointer rule contexts.
	GetPointer() IReferenceContext

	// SetPointer sets the pointer rule contexts.
	SetPointer(IReferenceContext)

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	integer antlr.Token
	boolean antlr.Token
	pointer IReferenceContext
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) GetInteger() antlr.Token { return s.integer }

func (s *TypeContext) GetBoolean() antlr.Token { return s.boolean }

func (s *TypeContext) SetInteger(v antlr.Token) { s.integer = v }

func (s *TypeContext) SetBoolean(v antlr.Token) { s.boolean = v }

func (s *TypeContext) GetPointer() IReferenceContext { return s.pointer }

func (s *TypeContext) SetPointer(v IReferenceContext) { s.pointer = v }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL, 0)
}

func (s *TypeContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GoliteParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoliteParserRULE_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)

			var _m = p.Match(GoliteParserINT)

			localctx.(*TypeContext).integer = _m
		}

	case GoliteParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)

			var _m = p.Match(GoliteParserBOOL)

			localctx.(*TypeContext).boolean = _m
		}

	case GoliteParserASTERIX:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)

			var _x = p.Reference()

			localctx.(*TypeContext).pointer = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
	return p
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationsContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclarations(s)
	}
}

func (s *DeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclarations(s)
	}
}

func (p *GoliteParser) Declarations() (localctx IDeclarationsContext) {
	this := p
	_ = this

	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoliteParserRULE_declarations)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(134)
			p.Declaration()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GoliteParserVAR, 0)
}

func (s *DeclarationContext) Ids() IIdsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsContext)
}

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *GoliteParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoliteParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(GoliteParserVAR)
	}
	{
		p.SetState(141)
		p.Ids()
	}
	{
		p.SetState(142)
		p.Type_()
	}
	{
		p.SetState(143)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IIdsContext is an interface to support dynamic dispatch.
type IIdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdsContext differentiates from other interfaces.
	IsIdsContext()
}

type IdsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdsContext() *IdsContext {
	var p = new(IdsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
	return p
}

func (*IdsContext) IsIdsContext() {}

func NewIdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsContext {
	var p = new(IdsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_ids

	return p
}

func (s *IdsContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *IdsContext) AllIdsPrime() []IIdsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IIdsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdsPrimeContext); ok {
			tst[i] = t.(IIdsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *IdsContext) IdsPrime(i int) IIdsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsPrimeContext)
}

func (s *IdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIds(s)
	}
}

func (s *IdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIds(s)
	}
}

func (p *GoliteParser) Ids() (localctx IIdsContext) {
	this := p
	_ = this

	localctx = NewIdsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoliteParserRULE_ids)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(GoliteParserID)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(146)
			p.IdsPrime()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIdsPrimeContext is an interface to support dynamic dispatch.
type IIdsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsIdsPrimeContext differentiates from other interfaces.
	IsIdsPrimeContext()
}

type IdsPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyIdsPrimeContext() *IdsPrimeContext {
	var p = new(IdsPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_idsPrime
	return p
}

func (*IdsPrimeContext) IsIdsPrimeContext() {}

func NewIdsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsPrimeContext {
	var p = new(IdsPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_idsPrime

	return p
}

func (s *IdsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsPrimeContext) GetId() antlr.Token { return s.id }

func (s *IdsPrimeContext) SetId(v antlr.Token) { s.id = v }

func (s *IdsPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *IdsPrimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *IdsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIdsPrime(s)
	}
}

func (s *IdsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIdsPrime(s)
	}
}

func (p *GoliteParser) IdsPrime() (localctx IIdsPrimeContext) {
	this := p
	_ = this

	localctx = NewIdsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoliteParserRULE_idsPrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(153)

		var _m = p.Match(GoliteParserID)

		localctx.(*IdsPrimeContext).id = _m
	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionsContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *GoliteParser) Functions() (localctx IFunctionsContext) {
	this := p
	_ = this

	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoliteParserRULE_functions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFUNC {
		{
			p.SetState(155)
			p.Function()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoliteParserFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FunctionContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *FunctionContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *FunctionContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *FunctionContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *GoliteParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoliteParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(GoliteParserFUNC)
	}
	{
		p.SetState(162)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(163)
		p.Params()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536870936) != 0 {
		{
			p.SetState(164)
			p.ReturnType()
		}

	}
	{
		p.SetState(167)
		p.Match(GoliteParserLBRACE)
	}
	{
		p.SetState(168)
		p.Declarations()
	}
	{
		p.SetState(169)
		p.Statements()
	}
	{
		p.SetState(170)
		p.Match(GoliteParserRBRACE)
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ParamsContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParamsContext) AllParamsPrime() []IParamsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IParamsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamsPrimeContext); ok {
			tst[i] = t.(IParamsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) ParamsPrime(i int) IParamsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsPrimeContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *GoliteParser) Params() (localctx IParamsContext) {
	this := p
	_ = this

	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoliteParserRULE_params)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(GoliteParserLPAREN)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserID {
		{
			p.SetState(173)
			p.Decl()
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(174)
				p.ParamsPrime()
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(182)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IParamsPrimeContext is an interface to support dynamic dispatch.
type IParamsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDeclr returns the declr rule contexts.
	GetDeclr() IDeclContext

	// SetDeclr sets the declr rule contexts.
	SetDeclr(IDeclContext)

	// IsParamsPrimeContext differentiates from other interfaces.
	IsParamsPrimeContext()
}

type ParamsPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	declr  IDeclContext
}

func NewEmptyParamsPrimeContext() *ParamsPrimeContext {
	var p = new(ParamsPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_paramsPrime
	return p
}

func (*ParamsPrimeContext) IsParamsPrimeContext() {}

func NewParamsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsPrimeContext {
	var p = new(ParamsPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_paramsPrime

	return p
}

func (s *ParamsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsPrimeContext) GetDeclr() IDeclContext { return s.declr }

func (s *ParamsPrimeContext) SetDeclr(v IDeclContext) { s.declr = v }

func (s *ParamsPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ParamsPrimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParamsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParamsPrime(s)
	}
}

func (s *ParamsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParamsPrime(s)
	}
}

func (p *GoliteParser) ParamsPrime() (localctx IParamsPrimeContext) {
	this := p
	_ = this

	localctx = NewParamsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoliteParserRULE_paramsPrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(185)

		var _x = p.Decl()

		localctx.(*ParamsPrimeContext).declr = _x
	}

	return localctx
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRt returns the rt rule contexts.
	GetRt() ITypeContext

	// SetRt sets the rt rule contexts.
	SetRt(ITypeContext)

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	rt     ITypeContext
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
	return p
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) GetRt() ITypeContext { return s.rt }

func (s *ReturnTypeContext) SetRt(v ITypeContext) { s.rt = v }

func (s *ReturnTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *GoliteParser) ReturnType() (localctx IReturnTypeContext) {
	this := p
	_ = this

	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoliteParserRULE_returnType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)

		var _x = p.Type_()

		localctx.(*ReturnTypeContext).rt = _x
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStmts returns the stmts rule contexts.
	GetStmts() IStatementContext

	// SetStmts sets the stmts rule contexts.
	SetStmts(IStatementContext)

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	stmts  IStatementContext
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) GetStmts() IStatementContext { return s.stmts }

func (s *StatementsContext) SetStmts(v IStatementContext) { s.stmts = v }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *GoliteParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoliteParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023269760) != 0 {
		{
			p.SetState(189)

			var _x = p.Statement()

			localctx.(*StatementsContext).stmts = _x
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *StatementContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoliteParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoliteParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.Print_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Read()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)
			p.Delete_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(199)
			p.Conditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(200)
			p.Loop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(201)
			p.Return_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(202)
			p.Invocation()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GoliteParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoliteParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(GoliteParserLBRACE)
	}
	{
		p.SetState(206)
		p.Statements()
	}
	{
		p.SetState(207)
		p.Match(GoliteParserRBRACE)
	}

	return localctx
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
	return p
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoliteParserDELETE, 0)
}

func (s *DeleteContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *GoliteParser) Delete_() (localctx IDeleteContext) {
	this := p
	_ = this

	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoliteParserRULE_delete)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(GoliteParserDELETE)
	}
	{
		p.SetState(210)
		p.Expression()
	}
	{
		p.SetState(211)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_read
	return p
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) SCAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserSCAN, 0)
}

func (s *ReadContext) LValue() ILValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValueContext)
}

func (s *ReadContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *GoliteParser) Read() (localctx IReadContext) {
	this := p
	_ = this

	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoliteParserRULE_read)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(GoliteParserSCAN)
	}
	{
		p.SetState(214)
		p.LValue()
	}
	{
		p.SetState(215)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) LValue() ILValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValueContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoliteParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GoliteParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoliteParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.LValue()
	}
	{
		p.SetState(218)
		p.Match(GoliteParserASSIGN)
	}
	{
		p.SetState(219)
		p.Expression()
	}
	{
		p.SetState(220)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_print
	return p
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GoliteParserPRINTF, 0)
}

func (s *PrintContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *PrintContext) STRING() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRING, 0)
}

func (s *PrintContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *PrintContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *PrintContext) AllArgsPrime() []IArgsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IArgsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgsPrimeContext); ok {
			tst[i] = t.(IArgsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) ArgsPrime(i int) IArgsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsPrimeContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *GoliteParser) Print_() (localctx IPrintContext) {
	this := p
	_ = this

	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoliteParserRULE_print)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(GoliteParserPRINTF)
	}
	{
		p.SetState(223)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(224)
		p.Match(GoliteParserSTRING)
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(225)
			p.ArgsPrime()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(231)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(232)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(GoliteParserIF, 0)
}

func (s *ConditionalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ConditionalContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ConditionalContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserELSE, 0)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *GoliteParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoliteParserRULE_conditional)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(GoliteParserIF)
	}
	{
		p.SetState(235)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(236)
		p.Expression()
	}
	{
		p.SetState(237)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(238)
		p.Block()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserELSE {
		{
			p.SetState(239)
			p.Match(GoliteParserELSE)
		}
		{
			p.SetState(240)
			p.Block()
		}

	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoliteParserFOR, 0)
}

func (s *LoopContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *GoliteParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoliteParserRULE_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(GoliteParserFOR)
	}
	{
		p.SetState(244)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(245)
		p.Expression()
	}
	{
		p.SetState(246)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(247)
		p.Block()
	}

	return localctx
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_return
	return p
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRETURN, 0)
}

func (s *ReturnContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReturnContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *GoliteParser) Return_() (localctx IReturnContext) {
	this := p
	_ = this

	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoliteParserRULE_return)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(GoliteParserRETURN)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3575828955136) != 0 {
		{
			p.SetState(250)
			p.Expression()
		}

	}
	{
		p.SetState(253)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *InvocationContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *InvocationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (p *GoliteParser) Invocation() (localctx IInvocationContext) {
	this := p
	_ = this

	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoliteParserRULE_invocation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(256)
		p.Args()
	}
	{
		p.SetState(257)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ArgsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ArgsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgsContext) AllArgsPrime() []IArgsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IArgsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgsPrimeContext); ok {
			tst[i] = t.(IArgsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) ArgsPrime(i int) IArgsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsPrimeContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *GoliteParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoliteParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(GoliteParserLPAREN)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3575828955136) != 0 {
		{
			p.SetState(260)
			p.Expression()
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(261)
				p.ArgsPrime()
			}

			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(269)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IArgsPrimeContext is an interface to support dynamic dispatch.
type IArgsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsArgsPrimeContext differentiates from other interfaces.
	IsArgsPrimeContext()
}

type ArgsPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyArgsPrimeContext() *ArgsPrimeContext {
	var p = new(ArgsPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_argsPrime
	return p
}

func (*ArgsPrimeContext) IsArgsPrimeContext() {}

func NewArgsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsPrimeContext {
	var p = new(ArgsPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_argsPrime

	return p
}

func (s *ArgsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsPrimeContext) GetExpr() IExpressionContext { return s.expr }

func (s *ArgsPrimeContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *ArgsPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ArgsPrimeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArgsPrime(s)
	}
}

func (s *ArgsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArgsPrime(s)
	}
}

func (p *GoliteParser) ArgsPrime() (localctx IArgsPrimeContext) {
	this := p
	_ = this

	localctx = NewArgsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoliteParserRULE_argsPrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(272)

		var _x = p.Expression()

		localctx.(*ArgsPrimeContext).expr = _x
	}

	return localctx
}

// ILValueContext is an interface to support dynamic dispatch.
type ILValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLValueContext differentiates from other interfaces.
	IsLValueContext()
}

type LValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLValueContext() *LValueContext {
	var p = new(LValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_lValue
	return p
}

func (*LValueContext) IsLValueContext() {}

func NewLValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LValueContext {
	var p = new(LValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lValue

	return p
}

func (s *LValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LValueContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *LValueContext) AllLValuePrime() []ILValuePrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILValuePrimeContext); ok {
			len++
		}
	}

	tst := make([]ILValuePrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILValuePrimeContext); ok {
			tst[i] = t.(ILValuePrimeContext)
			i++
		}
	}

	return tst
}

func (s *LValueContext) LValuePrime(i int) ILValuePrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValuePrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValuePrimeContext)
}

func (s *LValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLValue(s)
	}
}

func (s *LValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLValue(s)
	}
}

func (p *GoliteParser) LValue() (localctx ILValueContext) {
	this := p
	_ = this

	localctx = NewLValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoliteParserRULE_lValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(GoliteParserID)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(275)
			p.LValuePrime()
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILValuePrimeContext is an interface to support dynamic dispatch.
type ILValuePrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPid returns the pid token.
	GetPid() antlr.Token

	// SetPid sets the pid token.
	SetPid(antlr.Token)

	// IsLValuePrimeContext differentiates from other interfaces.
	IsLValuePrimeContext()
}

type LValuePrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	pid    antlr.Token
}

func NewEmptyLValuePrimeContext() *LValuePrimeContext {
	var p = new(LValuePrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_lValuePrime
	return p
}

func (*LValuePrimeContext) IsLValuePrimeContext() {}

func NewLValuePrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LValuePrimeContext {
	var p = new(LValuePrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lValuePrime

	return p
}

func (s *LValuePrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *LValuePrimeContext) GetPid() antlr.Token { return s.pid }

func (s *LValuePrimeContext) SetPid(v antlr.Token) { s.pid = v }

func (s *LValuePrimeContext) DOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, 0)
}

func (s *LValuePrimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *LValuePrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LValuePrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LValuePrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLValuePrime(s)
	}
}

func (s *LValuePrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLValuePrime(s)
	}
}

func (p *GoliteParser) LValuePrime() (localctx ILValuePrimeContext) {
	this := p
	_ = this

	localctx = NewLValuePrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoliteParserRULE_lValuePrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(GoliteParserDOT)
	}
	{
		p.SetState(282)

		var _m = p.Match(GoliteParserID)

		localctx.(*LValuePrimeContext).pid = _m
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) BoolTerm() IBoolTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTermContext)
}

func (s *ExpressionContext) AllExprPrime() []IExprPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprPrimeContext); ok {
			len++
		}
	}

	tst := make([]IExprPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprPrimeContext); ok {
			tst[i] = t.(IExprPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) ExprPrime(i int) IExprPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPrimeContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GoliteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoliteParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.BoolTerm()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserOR {
		{
			p.SetState(285)
			p.ExprPrime()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprPrimeContext is an interface to support dynamic dispatch.
type IExprPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBt returns the bt rule contexts.
	GetBt() IBoolTermContext

	// SetBt sets the bt rule contexts.
	SetBt(IBoolTermContext)

	// IsExprPrimeContext differentiates from other interfaces.
	IsExprPrimeContext()
}

type ExprPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	bt     IBoolTermContext
}

func NewEmptyExprPrimeContext() *ExprPrimeContext {
	var p = new(ExprPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_exprPrime
	return p
}

func (*ExprPrimeContext) IsExprPrimeContext() {}

func NewExprPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprPrimeContext {
	var p = new(ExprPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_exprPrime

	return p
}

func (s *ExprPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprPrimeContext) GetBt() IBoolTermContext { return s.bt }

func (s *ExprPrimeContext) SetBt(v IBoolTermContext) { s.bt = v }

func (s *ExprPrimeContext) OR() antlr.TerminalNode {
	return s.GetToken(GoliteParserOR, 0)
}

func (s *ExprPrimeContext) BoolTerm() IBoolTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTermContext)
}

func (s *ExprPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExprPrime(s)
	}
}

func (s *ExprPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExprPrime(s)
	}
}

func (p *GoliteParser) ExprPrime() (localctx IExprPrimeContext) {
	this := p
	_ = this

	localctx = NewExprPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoliteParserRULE_exprPrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(GoliteParserOR)
	}
	{
		p.SetState(292)

		var _x = p.BoolTerm()

		localctx.(*ExprPrimeContext).bt = _x
	}

	return localctx
}

// IBoolTermContext is an interface to support dynamic dispatch.
type IBoolTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolTermContext differentiates from other interfaces.
	IsBoolTermContext()
}

type BoolTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTermContext() *BoolTermContext {
	var p = new(BoolTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_boolTerm
	return p
}

func (*BoolTermContext) IsBoolTermContext() {}

func NewBoolTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTermContext {
	var p = new(BoolTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolTerm

	return p
}

func (s *BoolTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTermContext) EqualTerm() IEqualTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermContext)
}

func (s *BoolTermContext) AllBoolTermPrime() []IBoolTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]IBoolTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolTermPrimeContext); ok {
			tst[i] = t.(IBoolTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *BoolTermContext) BoolTermPrime(i int) IBoolTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTermPrimeContext)
}

func (s *BoolTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolTerm(s)
	}
}

func (s *BoolTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolTerm(s)
	}
}

func (p *GoliteParser) BoolTerm() (localctx IBoolTermContext) {
	this := p
	_ = this

	localctx = NewBoolTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoliteParserRULE_boolTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.EqualTerm()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserAND {
		{
			p.SetState(295)
			p.BoolTermPrime()
		}

		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolTermPrimeContext is an interface to support dynamic dispatch.
type IBoolTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEt returns the et rule contexts.
	GetEt() IEqualTermContext

	// SetEt sets the et rule contexts.
	SetEt(IEqualTermContext)

	// IsBoolTermPrimeContext differentiates from other interfaces.
	IsBoolTermPrimeContext()
}

type BoolTermPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	et     IEqualTermContext
}

func NewEmptyBoolTermPrimeContext() *BoolTermPrimeContext {
	var p = new(BoolTermPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_boolTermPrime
	return p
}

func (*BoolTermPrimeContext) IsBoolTermPrimeContext() {}

func NewBoolTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTermPrimeContext {
	var p = new(BoolTermPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolTermPrime

	return p
}

func (s *BoolTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTermPrimeContext) GetEt() IEqualTermContext { return s.et }

func (s *BoolTermPrimeContext) SetEt(v IEqualTermContext) { s.et = v }

func (s *BoolTermPrimeContext) AND() antlr.TerminalNode {
	return s.GetToken(GoliteParserAND, 0)
}

func (s *BoolTermPrimeContext) EqualTerm() IEqualTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermContext)
}

func (s *BoolTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolTermPrime(s)
	}
}

func (s *BoolTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolTermPrime(s)
	}
}

func (p *GoliteParser) BoolTermPrime() (localctx IBoolTermPrimeContext) {
	this := p
	_ = this

	localctx = NewBoolTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoliteParserRULE_boolTermPrime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(GoliteParserAND)
	}
	{
		p.SetState(302)

		var _x = p.EqualTerm()

		localctx.(*BoolTermPrimeContext).et = _x
	}

	return localctx
}

// IEqualTermPrimeContext is an interface to support dynamic dispatch.
type IEqualTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRt returns the rt rule contexts.
	GetRt() IRelTermContext

	// SetRt sets the rt rule contexts.
	SetRt(IRelTermContext)

	// IsEqualTermPrimeContext differentiates from other interfaces.
	IsEqualTermPrimeContext()
}

type EqualTermPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	rt     IRelTermContext
}

func NewEmptyEqualTermPrimeContext() *EqualTermPrimeContext {
	var p = new(EqualTermPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTermPrime
	return p
}

func (*EqualTermPrimeContext) IsEqualTermPrimeContext() {}

func NewEqualTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualTermPrimeContext {
	var p = new(EqualTermPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalTermPrime

	return p
}

func (s *EqualTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualTermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *EqualTermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualTermPrimeContext) GetRt() IRelTermContext { return s.rt }

func (s *EqualTermPrimeContext) SetRt(v IRelTermContext) { s.rt = v }

func (s *EqualTermPrimeContext) RelTerm() IRelTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelTermContext)
}

func (s *EqualTermPrimeContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserEQUALS, 0)
}

func (s *EqualTermPrimeContext) NEQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEQUALS, 0)
}

func (s *EqualTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualTermPrime(s)
	}
}

func (s *EqualTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualTermPrime(s)
	}
}

func (p *GoliteParser) EqualTermPrime() (localctx IEqualTermPrimeContext) {
	this := p
	_ = this

	localctx = NewEqualTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoliteParserRULE_equalTermPrime)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*EqualTermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserEQUALS || _la == GoliteParserNEQUALS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*EqualTermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(305)

		var _x = p.RelTerm()

		localctx.(*EqualTermPrimeContext).rt = _x
	}

	return localctx
}

// IEqualTermContext is an interface to support dynamic dispatch.
type IEqualTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualTermContext differentiates from other interfaces.
	IsEqualTermContext()
}

type EqualTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualTermContext() *EqualTermContext {
	var p = new(EqualTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTerm
	return p
}

func (*EqualTermContext) IsEqualTermContext() {}

func NewEqualTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualTermContext {
	var p = new(EqualTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalTerm

	return p
}

func (s *EqualTermContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualTermContext) RelTerm() IRelTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelTermContext)
}

func (s *EqualTermContext) AllEqualTermPrime() []IEqualTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]IEqualTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualTermPrimeContext); ok {
			tst[i] = t.(IEqualTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *EqualTermContext) EqualTermPrime(i int) IEqualTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermPrimeContext)
}

func (s *EqualTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualTerm(s)
	}
}

func (s *EqualTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualTerm(s)
	}
}

func (p *GoliteParser) EqualTerm() (localctx IEqualTermContext) {
	this := p
	_ = this

	localctx = NewEqualTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoliteParserRULE_equalTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.RelTerm()
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserEQUALS || _la == GoliteParserNEQUALS {
		{
			p.SetState(308)
			p.EqualTermPrime()
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelTermPrimeContext is an interface to support dynamic dispatch.
type IRelTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetSt returns the st rule contexts.
	GetSt() ISimpleTermContext

	// SetSt sets the st rule contexts.
	SetSt(ISimpleTermContext)

	// IsRelTermPrimeContext differentiates from other interfaces.
	IsRelTermPrimeContext()
}

type RelTermPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	st     ISimpleTermContext
}

func NewEmptyRelTermPrimeContext() *RelTermPrimeContext {
	var p = new(RelTermPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_relTermPrime
	return p
}

func (*RelTermPrimeContext) IsRelTermPrimeContext() {}

func NewRelTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelTermPrimeContext {
	var p = new(RelTermPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relTermPrime

	return p
}

func (s *RelTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RelTermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *RelTermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelTermPrimeContext) GetSt() ISimpleTermContext { return s.st }

func (s *RelTermPrimeContext) SetSt(v ISimpleTermContext) { s.st = v }

func (s *RelTermPrimeContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *RelTermPrimeContext) GT() antlr.TerminalNode {
	return s.GetToken(GoliteParserGT, 0)
}

func (s *RelTermPrimeContext) LT() antlr.TerminalNode {
	return s.GetToken(GoliteParserLT, 0)
}

func (s *RelTermPrimeContext) GTE() antlr.TerminalNode {
	return s.GetToken(GoliteParserGTE, 0)
}

func (s *RelTermPrimeContext) LTE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLTE, 0)
}

func (s *RelTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelTermPrime(s)
	}
}

func (s *RelTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelTermPrime(s)
	}
}

func (p *GoliteParser) RelTermPrime() (localctx IRelTermPrimeContext) {
	this := p
	_ = this

	localctx = NewRelTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoliteParserRULE_relTermPrime)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*RelTermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&53477376) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*RelTermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(315)

		var _x = p.SimpleTerm()

		localctx.(*RelTermPrimeContext).st = _x
	}

	return localctx
}

// IRelTermContext is an interface to support dynamic dispatch.
type IRelTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelTermContext differentiates from other interfaces.
	IsRelTermContext()
}

type RelTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelTermContext() *RelTermContext {
	var p = new(RelTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_relTerm
	return p
}

func (*RelTermContext) IsRelTermContext() {}

func NewRelTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelTermContext {
	var p = new(RelTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relTerm

	return p
}

func (s *RelTermContext) GetParser() antlr.Parser { return s.parser }

func (s *RelTermContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *RelTermContext) AllRelTermPrime() []IRelTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]IRelTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelTermPrimeContext); ok {
			tst[i] = t.(IRelTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *RelTermContext) RelTermPrime(i int) IRelTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelTermPrimeContext)
}

func (s *RelTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelTerm(s)
	}
}

func (s *RelTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelTerm(s)
	}
}

func (p *GoliteParser) RelTerm() (localctx IRelTermContext) {
	this := p
	_ = this

	localctx = NewRelTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoliteParserRULE_relTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.SimpleTerm()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&53477376) != 0 {
		{
			p.SetState(318)
			p.RelTermPrime()
		}

		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISimpleTermPrimeContext is an interface to support dynamic dispatch.
type ISimpleTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITermContext

	// SetT sets the t rule contexts.
	SetT(ITermContext)

	// IsSimpleTermPrimeContext differentiates from other interfaces.
	IsSimpleTermPrimeContext()
}

type SimpleTermPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	t      ITermContext
}

func NewEmptySimpleTermPrimeContext() *SimpleTermPrimeContext {
	var p = new(SimpleTermPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTermPrime
	return p
}

func (*SimpleTermPrimeContext) IsSimpleTermPrimeContext() {}

func NewSimpleTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermPrimeContext {
	var p = new(SimpleTermPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleTermPrime

	return p
}

func (s *SimpleTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *SimpleTermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *SimpleTermPrimeContext) GetT() ITermContext { return s.t }

func (s *SimpleTermPrimeContext) SetT(v ITermContext) { s.t = v }

func (s *SimpleTermPrimeContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleTermPrimeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserPLUS, 0)
}

func (s *SimpleTermPrimeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *SimpleTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleTermPrime(s)
	}
}

func (s *SimpleTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleTermPrime(s)
	}
}

func (p *GoliteParser) SimpleTermPrime() (localctx ISimpleTermPrimeContext) {
	this := p
	_ = this

	localctx = NewSimpleTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoliteParserRULE_simpleTermPrime)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*SimpleTermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserPLUS || _la == GoliteParserMINUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*SimpleTermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(325)

		var _x = p.Term()

		localctx.(*SimpleTermPrimeContext).t = _x
	}

	return localctx
}

// ISimpleTermContext is an interface to support dynamic dispatch.
type ISimpleTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleTermContext differentiates from other interfaces.
	IsSimpleTermContext()
}

type SimpleTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTermContext() *SimpleTermContext {
	var p = new(SimpleTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTerm
	return p
}

func (*SimpleTermContext) IsSimpleTermContext() {}

func NewSimpleTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermContext {
	var p = new(SimpleTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleTerm

	return p
}

func (s *SimpleTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleTermContext) AllSimpleTermPrime() []ISimpleTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]ISimpleTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleTermPrimeContext); ok {
			tst[i] = t.(ISimpleTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *SimpleTermContext) SimpleTermPrime(i int) ISimpleTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermPrimeContext)
}

func (s *SimpleTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleTerm(s)
	}
}

func (s *SimpleTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleTerm(s)
	}
}

func (p *GoliteParser) SimpleTerm() (localctx ISimpleTermContext) {
	this := p
	_ = this

	localctx = NewSimpleTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GoliteParserRULE_simpleTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Term()
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPLUS || _la == GoliteParserMINUS {
		{
			p.SetState(328)
			p.SimpleTermPrime()
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermPrimeContext is an interface to support dynamic dispatch.
type ITermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetUt returns the ut rule contexts.
	GetUt() IUnaryTermContext

	// SetUt sets the ut rule contexts.
	SetUt(IUnaryTermContext)

	// IsTermPrimeContext differentiates from other interfaces.
	IsTermPrimeContext()
}

type TermPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	ut     IUnaryTermContext
}

func NewEmptyTermPrimeContext() *TermPrimeContext {
	var p = new(TermPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_termPrime
	return p
}

func (*TermPrimeContext) IsTermPrimeContext() {}

func NewTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermPrimeContext {
	var p = new(TermPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_termPrime

	return p
}

func (s *TermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *TermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *TermPrimeContext) GetUt() IUnaryTermContext { return s.ut }

func (s *TermPrimeContext) SetUt(v IUnaryTermContext) { s.ut = v }

func (s *TermPrimeContext) UnaryTerm() IUnaryTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryTermContext)
}

func (s *TermPrimeContext) ASTERIX() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERIX, 0)
}

func (s *TermPrimeContext) FSLASH() antlr.TerminalNode {
	return s.GetToken(GoliteParserFSLASH, 0)
}

func (s *TermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTermPrime(s)
	}
}

func (s *TermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTermPrime(s)
	}
}

func (p *GoliteParser) TermPrime() (localctx ITermPrimeContext) {
	this := p
	_ = this

	localctx = NewTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GoliteParserRULE_termPrime)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserASTERIX || _la == GoliteParserFSLASH) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(335)

		var _x = p.UnaryTerm()

		localctx.(*TermPrimeContext).ut = _x
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) UnaryTerm() IUnaryTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryTermContext)
}

func (s *TermContext) AllTermPrime() []ITermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermPrimeContext); ok {
			len++
		}
	}

	tst := make([]ITermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermPrimeContext); ok {
			tst[i] = t.(ITermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) TermPrime(i int) ITermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermPrimeContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GoliteParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GoliteParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.UnaryTerm()
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserASTERIX || _la == GoliteParserFSLASH {
		{
			p.SetState(338)
			p.TermPrime()
		}

		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnaryTermContext is an interface to support dynamic dispatch.
type IUnaryTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryTermContext differentiates from other interfaces.
	IsUnaryTermContext()
}

type UnaryTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryTermContext() *UnaryTermContext {
	var p = new(UnaryTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryTerm
	return p
}

func (*UnaryTermContext) IsUnaryTermContext() {}

func NewUnaryTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryTermContext {
	var p = new(UnaryTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_unaryTerm

	return p
}

func (s *UnaryTermContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryTermContext) NOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserNOT, 0)
}

func (s *UnaryTermContext) SelectorTerm() ISelectorTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTermContext)
}

func (s *UnaryTermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *UnaryTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUnaryTerm(s)
	}
}

func (s *UnaryTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUnaryTerm(s)
	}
}

func (p *GoliteParser) UnaryTerm() (localctx IUnaryTermContext) {
	this := p
	_ = this

	localctx = NewUnaryTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GoliteParserRULE_unaryTerm)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(349)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.Match(GoliteParserNOT)
		}
		{
			p.SetState(345)
			p.SelectorTerm()
		}

	case GoliteParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Match(GoliteParserMINUS)
		}
		{
			p.SetState(347)
			p.SelectorTerm()
		}

	case GoliteParserNEW, GoliteParserTRUE, GoliteParserFALSE, GoliteParserNIL, GoliteParserLPAREN, GoliteParserNUM, GoliteParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)
			p.SelectorTerm()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectorTermContext is an interface to support dynamic dispatch.
type ISelectorTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorTermContext differentiates from other interfaces.
	IsSelectorTermContext()
}

type SelectorTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTermContext() *SelectorTermContext {
	var p = new(SelectorTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTerm
	return p
}

func (*SelectorTermContext) IsSelectorTermContext() {}

func NewSelectorTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTermContext {
	var p = new(SelectorTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorTerm

	return p
}

func (s *SelectorTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SelectorTermContext) AllLValuePrime() []ILValuePrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILValuePrimeContext); ok {
			len++
		}
	}

	tst := make([]ILValuePrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILValuePrimeContext); ok {
			tst[i] = t.(ILValuePrimeContext)
			i++
		}
	}

	return tst
}

func (s *SelectorTermContext) LValuePrime(i int) ILValuePrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValuePrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValuePrimeContext)
}

func (s *SelectorTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorTerm(s)
	}
}

func (s *SelectorTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorTerm(s)
	}
}

func (p *GoliteParser) SelectorTerm() (localctx ISelectorTermContext) {
	this := p
	_ = this

	localctx = NewSelectorTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GoliteParserRULE_selectorTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Factor()
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(352)
			p.LValuePrime()
		}

		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FactorContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FactorContext) NUM() antlr.TerminalNode {
	return s.GetToken(GoliteParserNUM, 0)
}

func (s *FactorContext) NEW() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEW, 0)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserFALSE, 0)
}

func (s *FactorContext) NIL() antlr.TerminalNode {
	return s.GetToken(GoliteParserNIL, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GoliteParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GoliteParserRULE_factor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(372)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.Match(GoliteParserLPAREN)
		}
		{
			p.SetState(359)
			p.Expression()
		}
		{
			p.SetState(360)
			p.Match(GoliteParserRPAREN)
		}

	case GoliteParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Match(GoliteParserID)
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoliteParserLPAREN {
			{
				p.SetState(363)
				p.Args()
			}

		}

	case GoliteParserNUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(366)
			p.Match(GoliteParserNUM)
		}

	case GoliteParserNEW:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(367)
			p.Match(GoliteParserNEW)
		}
		{
			p.SetState(368)
			p.Match(GoliteParserID)
		}

	case GoliteParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(369)
			p.Match(GoliteParserTRUE)
		}

	case GoliteParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(370)
			p.Match(GoliteParserFALSE)
		}

	case GoliteParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(371)
			p.Match(GoliteParserNIL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
