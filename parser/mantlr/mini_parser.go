// Code generated from Mini.g4 by ANTLR 4.12.0. DO NOT EDIT.

package mantlr // Mini
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

/* package declaration here */

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MiniParser struct {
	*antlr.BaseParser
}

var miniParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func miniParserInit() {
	staticData := &miniParserStaticData
	staticData.literalNames = []string{
		"", "'struct'", "'{'", "'}'", "';'", "'int'", "'bool'", "','", "'fun'",
		"'('", "')'", "'void'", "'='", "'read'", "'print'", "'endl'", "'if'",
		"'else'", "'while'", "'delete'", "'return'", "'.'", "'-'", "'!'", "'*'",
		"'/'", "'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "'true'", "'false'", "'new'", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "ID", "INTEGER", "WS", "COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "types", "typeDeclaration", "nestedDecl", "decl", "type",
		"declarations", "declaration", "functions", "function", "parameters",
		"returnType", "statement", "block", "statementList", "lvalue", "expression",
		"arguments",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 251, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 43,
		8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 1, 3, 1, 49, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 61, 8, 3, 11, 3, 12, 3, 62,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 72, 8, 5, 1, 6, 5, 6, 75,
		8, 6, 10, 6, 12, 6, 78, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 84, 8, 7, 10,
		7, 12, 7, 87, 9, 7, 1, 7, 1, 7, 1, 8, 5, 8, 92, 8, 8, 10, 8, 12, 8, 95,
		9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 5, 10, 110, 8, 10, 10, 10, 12, 10, 113, 9, 10, 3, 10, 115,
		8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 121, 8, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 128, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 148, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 162, 8, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 171, 8, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 5, 14, 178, 8, 14, 10, 14, 12, 14, 181, 9,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 189, 8, 15, 10, 15,
		12, 15, 192, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 3, 16, 213, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 236, 8, 16, 10, 16, 12, 16, 239, 9,
		16, 1, 17, 1, 17, 1, 17, 5, 17, 244, 8, 17, 10, 17, 12, 17, 247, 9, 17,
		3, 17, 249, 8, 17, 1, 17, 0, 2, 30, 32, 18, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 5, 1, 0, 22, 23, 1, 0, 24, 25,
		2, 0, 22, 22, 26, 26, 1, 0, 27, 30, 1, 0, 31, 32, 273, 0, 36, 1, 0, 0,
		0, 2, 48, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6, 60, 1, 0, 0, 0, 8, 64, 1, 0,
		0, 0, 10, 71, 1, 0, 0, 0, 12, 76, 1, 0, 0, 0, 14, 79, 1, 0, 0, 0, 16, 93,
		1, 0, 0, 0, 18, 96, 1, 0, 0, 0, 20, 105, 1, 0, 0, 0, 22, 120, 1, 0, 0,
		0, 24, 170, 1, 0, 0, 0, 26, 172, 1, 0, 0, 0, 28, 179, 1, 0, 0, 0, 30, 182,
		1, 0, 0, 0, 32, 212, 1, 0, 0, 0, 34, 248, 1, 0, 0, 0, 36, 37, 3, 2, 1,
		0, 37, 38, 3, 12, 6, 0, 38, 39, 3, 16, 8, 0, 39, 40, 5, 0, 0, 1, 40, 1,
		1, 0, 0, 0, 41, 43, 3, 4, 2, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0,
		44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 49, 1, 0, 0, 0, 46, 44, 1,
		0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 44, 1, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49,
		3, 1, 0, 0, 0, 50, 51, 5, 1, 0, 0, 51, 52, 5, 39, 0, 0, 52, 53, 5, 2, 0,
		0, 53, 54, 3, 6, 3, 0, 54, 55, 5, 3, 0, 0, 55, 56, 5, 4, 0, 0, 56, 5, 1,
		0, 0, 0, 57, 58, 3, 8, 4, 0, 58, 59, 5, 4, 0, 0, 59, 61, 1, 0, 0, 0, 60,
		57, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0,
		0, 63, 7, 1, 0, 0, 0, 64, 65, 3, 10, 5, 0, 65, 66, 5, 39, 0, 0, 66, 9,
		1, 0, 0, 0, 67, 72, 5, 5, 0, 0, 68, 72, 5, 6, 0, 0, 69, 70, 5, 1, 0, 0,
		70, 72, 5, 39, 0, 0, 71, 67, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 72, 11, 1, 0, 0, 0, 73, 75, 3, 14, 7, 0, 74, 73, 1, 0, 0, 0, 75,
		78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 13, 1, 0, 0,
		0, 78, 76, 1, 0, 0, 0, 79, 80, 3, 10, 5, 0, 80, 85, 5, 39, 0, 0, 81, 82,
		5, 7, 0, 0, 82, 84, 5, 39, 0, 0, 83, 81, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0,
		85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1,
		0, 0, 0, 88, 89, 5, 4, 0, 0, 89, 15, 1, 0, 0, 0, 90, 92, 3, 18, 9, 0, 91,
		90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0,
		0, 94, 17, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 5, 8, 0, 0, 97, 98,
		5, 39, 0, 0, 98, 99, 3, 20, 10, 0, 99, 100, 3, 22, 11, 0, 100, 101, 5,
		2, 0, 0, 101, 102, 3, 12, 6, 0, 102, 103, 3, 28, 14, 0, 103, 104, 5, 3,
		0, 0, 104, 19, 1, 0, 0, 0, 105, 114, 5, 9, 0, 0, 106, 111, 3, 8, 4, 0,
		107, 108, 5, 7, 0, 0, 108, 110, 3, 8, 4, 0, 109, 107, 1, 0, 0, 0, 110,
		113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 115,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 106, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 5, 10, 0, 0, 117, 21, 1, 0, 0, 0,
		118, 121, 3, 10, 5, 0, 119, 121, 5, 11, 0, 0, 120, 118, 1, 0, 0, 0, 120,
		119, 1, 0, 0, 0, 121, 23, 1, 0, 0, 0, 122, 171, 3, 26, 13, 0, 123, 124,
		3, 30, 15, 0, 124, 127, 5, 12, 0, 0, 125, 128, 3, 32, 16, 0, 126, 128,
		5, 13, 0, 0, 127, 125, 1, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0,
		0, 0, 129, 130, 5, 4, 0, 0, 130, 171, 1, 0, 0, 0, 131, 132, 5, 14, 0, 0,
		132, 133, 3, 32, 16, 0, 133, 134, 5, 4, 0, 0, 134, 171, 1, 0, 0, 0, 135,
		136, 5, 14, 0, 0, 136, 137, 3, 32, 16, 0, 137, 138, 5, 15, 0, 0, 138, 139,
		5, 4, 0, 0, 139, 171, 1, 0, 0, 0, 140, 141, 5, 16, 0, 0, 141, 142, 5, 9,
		0, 0, 142, 143, 3, 32, 16, 0, 143, 144, 5, 10, 0, 0, 144, 147, 3, 26, 13,
		0, 145, 146, 5, 17, 0, 0, 146, 148, 3, 26, 13, 0, 147, 145, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 171, 1, 0, 0, 0, 149, 150, 5, 18, 0, 0, 150,
		151, 5, 9, 0, 0, 151, 152, 3, 32, 16, 0, 152, 153, 5, 10, 0, 0, 153, 154,
		3, 26, 13, 0, 154, 171, 1, 0, 0, 0, 155, 156, 5, 19, 0, 0, 156, 157, 3,
		32, 16, 0, 157, 158, 5, 4, 0, 0, 158, 171, 1, 0, 0, 0, 159, 161, 5, 20,
		0, 0, 160, 162, 3, 32, 16, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 171, 5, 4, 0, 0, 164, 165, 5, 39, 0, 0, 165,
		166, 5, 9, 0, 0, 166, 167, 3, 34, 17, 0, 167, 168, 5, 10, 0, 0, 168, 169,
		5, 4, 0, 0, 169, 171, 1, 0, 0, 0, 170, 122, 1, 0, 0, 0, 170, 123, 1, 0,
		0, 0, 170, 131, 1, 0, 0, 0, 170, 135, 1, 0, 0, 0, 170, 140, 1, 0, 0, 0,
		170, 149, 1, 0, 0, 0, 170, 155, 1, 0, 0, 0, 170, 159, 1, 0, 0, 0, 170,
		164, 1, 0, 0, 0, 171, 25, 1, 0, 0, 0, 172, 173, 5, 2, 0, 0, 173, 174, 3,
		28, 14, 0, 174, 175, 5, 3, 0, 0, 175, 27, 1, 0, 0, 0, 176, 178, 3, 24,
		12, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0,
		179, 180, 1, 0, 0, 0, 180, 29, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 183,
		6, 15, -1, 0, 183, 184, 5, 39, 0, 0, 184, 190, 1, 0, 0, 0, 185, 186, 10,
		1, 0, 0, 186, 187, 5, 21, 0, 0, 187, 189, 5, 39, 0, 0, 188, 185, 1, 0,
		0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0,
		191, 31, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 6, 16, -1, 0, 194,
		195, 5, 39, 0, 0, 195, 196, 5, 9, 0, 0, 196, 197, 3, 34, 17, 0, 197, 198,
		5, 10, 0, 0, 198, 213, 1, 0, 0, 0, 199, 200, 7, 0, 0, 0, 200, 213, 3, 32,
		16, 14, 201, 213, 5, 39, 0, 0, 202, 213, 5, 40, 0, 0, 203, 213, 5, 35,
		0, 0, 204, 213, 5, 36, 0, 0, 205, 206, 5, 37, 0, 0, 206, 213, 5, 39, 0,
		0, 207, 213, 5, 38, 0, 0, 208, 209, 5, 9, 0, 0, 209, 210, 3, 32, 16, 0,
		210, 211, 5, 10, 0, 0, 211, 213, 1, 0, 0, 0, 212, 193, 1, 0, 0, 0, 212,
		199, 1, 0, 0, 0, 212, 201, 1, 0, 0, 0, 212, 202, 1, 0, 0, 0, 212, 203,
		1, 0, 0, 0, 212, 204, 1, 0, 0, 0, 212, 205, 1, 0, 0, 0, 212, 207, 1, 0,
		0, 0, 212, 208, 1, 0, 0, 0, 213, 237, 1, 0, 0, 0, 214, 215, 10, 13, 0,
		0, 215, 216, 7, 1, 0, 0, 216, 236, 3, 32, 16, 14, 217, 218, 10, 12, 0,
		0, 218, 219, 7, 2, 0, 0, 219, 236, 3, 32, 16, 13, 220, 221, 10, 11, 0,
		0, 221, 222, 7, 3, 0, 0, 222, 236, 3, 32, 16, 12, 223, 224, 10, 10, 0,
		0, 224, 225, 7, 4, 0, 0, 225, 236, 3, 32, 16, 11, 226, 227, 10, 9, 0, 0,
		227, 228, 5, 33, 0, 0, 228, 236, 3, 32, 16, 10, 229, 230, 10, 8, 0, 0,
		230, 231, 5, 34, 0, 0, 231, 236, 3, 32, 16, 9, 232, 233, 10, 15, 0, 0,
		233, 234, 5, 21, 0, 0, 234, 236, 5, 39, 0, 0, 235, 214, 1, 0, 0, 0, 235,
		217, 1, 0, 0, 0, 235, 220, 1, 0, 0, 0, 235, 223, 1, 0, 0, 0, 235, 226,
		1, 0, 0, 0, 235, 229, 1, 0, 0, 0, 235, 232, 1, 0, 0, 0, 236, 239, 1, 0,
		0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 33, 1, 0, 0, 0,
		239, 237, 1, 0, 0, 0, 240, 245, 3, 32, 16, 0, 241, 242, 5, 7, 0, 0, 242,
		244, 3, 32, 16, 0, 243, 241, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0,
		0, 0, 248, 240, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 35, 1, 0, 0, 0,
		21, 44, 48, 62, 71, 76, 85, 93, 111, 114, 120, 127, 147, 161, 170, 179,
		190, 212, 235, 237, 245, 248,
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

// MiniParserInit initializes any static state used to implement MiniParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMiniParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniParserInit() {
	staticData := &miniParserStaticData
	staticData.once.Do(miniParserInit)
}

// NewMiniParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMiniParser(input antlr.TokenStream) *MiniParser {
	MiniParserInit()
	this := new(MiniParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &miniParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mini.g4"

	return this
}

// MiniParser tokens.
const (
	MiniParserEOF     = antlr.TokenEOF
	MiniParserT__0    = 1
	MiniParserT__1    = 2
	MiniParserT__2    = 3
	MiniParserT__3    = 4
	MiniParserT__4    = 5
	MiniParserT__5    = 6
	MiniParserT__6    = 7
	MiniParserT__7    = 8
	MiniParserT__8    = 9
	MiniParserT__9    = 10
	MiniParserT__10   = 11
	MiniParserT__11   = 12
	MiniParserT__12   = 13
	MiniParserT__13   = 14
	MiniParserT__14   = 15
	MiniParserT__15   = 16
	MiniParserT__16   = 17
	MiniParserT__17   = 18
	MiniParserT__18   = 19
	MiniParserT__19   = 20
	MiniParserT__20   = 21
	MiniParserT__21   = 22
	MiniParserT__22   = 23
	MiniParserT__23   = 24
	MiniParserT__24   = 25
	MiniParserT__25   = 26
	MiniParserT__26   = 27
	MiniParserT__27   = 28
	MiniParserT__28   = 29
	MiniParserT__29   = 30
	MiniParserT__30   = 31
	MiniParserT__31   = 32
	MiniParserT__32   = 33
	MiniParserT__33   = 34
	MiniParserT__34   = 35
	MiniParserT__35   = 36
	MiniParserT__36   = 37
	MiniParserT__37   = 38
	MiniParserID      = 39
	MiniParserINTEGER = 40
	MiniParserWS      = 41
	MiniParserCOMMENT = 42
)

// MiniParser rules.
const (
	MiniParserRULE_program         = 0
	MiniParserRULE_types           = 1
	MiniParserRULE_typeDeclaration = 2
	MiniParserRULE_nestedDecl      = 3
	MiniParserRULE_decl            = 4
	MiniParserRULE_type            = 5
	MiniParserRULE_declarations    = 6
	MiniParserRULE_declaration     = 7
	MiniParserRULE_functions       = 8
	MiniParserRULE_function        = 9
	MiniParserRULE_parameters      = 10
	MiniParserRULE_returnType      = 11
	MiniParserRULE_statement       = 12
	MiniParserRULE_block           = 13
	MiniParserRULE_statementList   = 14
	MiniParserRULE_lvalue          = 15
	MiniParserRULE_expression      = 16
	MiniParserRULE_arguments       = 17
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Types() ITypesContext
	Declarations() IDeclarationsContext
	Functions() IFunctionsContext
	EOF() antlr.TerminalNode

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
	p.RuleIndex = MiniParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
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
	return s.GetToken(MiniParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MiniParserRULE_program)

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
		p.SetState(36)
		p.Types()
	}
	{
		p.SetState(37)
		p.Declarations()
	}
	{
		p.SetState(38)
		p.Functions()
	}
	{
		p.SetState(39)
		p.Match(MiniParserEOF)
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeDeclaration() []ITypeDeclarationContext
	TypeDeclaration(i int) ITypeDeclarationContext

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypesContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
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

	return t.(ITypeDeclarationContext)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Types() (localctx ITypesContext) {
	this := p
	_ = this

	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MiniParserRULE_types)

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

	var _alt int

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(41)
					p.TypeDeclaration()
				}

			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	NestedDecl() INestedDeclContext

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

func (s *TypeDeclarationContext) NestedDecl() INestedDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedDeclContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	this := p
	_ = this

	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MiniParserRULE_typeDeclaration)

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
		p.SetState(50)
		p.Match(MiniParserT__0)
	}
	{
		p.SetState(51)
		p.Match(MiniParserID)
	}
	{
		p.SetState(52)
		p.Match(MiniParserT__1)
	}
	{
		p.SetState(53)
		p.NestedDecl()
	}
	{
		p.SetState(54)
		p.Match(MiniParserT__2)
	}
	{
		p.SetState(55)
		p.Match(MiniParserT__3)
	}

	return localctx
}

// INestedDeclContext is an interface to support dynamic dispatch.
type INestedDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDecl() []IDeclContext
	Decl(i int) IDeclContext

	// IsNestedDeclContext differentiates from other interfaces.
	IsNestedDeclContext()
}

type NestedDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedDeclContext() *NestedDeclContext {
	var p = new(NestedDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_nestedDecl
	return p
}

func (*NestedDeclContext) IsNestedDeclContext() {}

func NewNestedDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedDeclContext {
	var p = new(NestedDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_nestedDecl

	return p
}

func (s *NestedDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedDeclContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *NestedDeclContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
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

	return t.(IDeclContext)
}

func (s *NestedDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) NestedDecl() (localctx INestedDeclContext) {
	this := p
	_ = this

	localctx = NewNestedDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MiniParserRULE_nestedDecl)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&98) != 0) {
		{
			p.SetState(57)
			p.Decl()
		}
		{
			p.SetState(58)
			p.Match(MiniParserT__3)
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

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
	p.RuleIndex = MiniParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

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

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MiniParserRULE_decl)

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
		p.SetState(64)
		p.Type_()
	}
	{
		p.SetState(65)
		p.Match(MiniParserID)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyFrom(ctx *TypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolTypeContext struct {
	*TypeContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

type StructTypeContext struct {
	*TypeContext
}

func NewStructTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeContext {
	var p = new(StructTypeContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

type IntTypeContext struct {
	*TypeContext
}

func NewIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTypeContext {
	var p = new(IntTypeContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (p *MiniParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MiniParserRULE_type)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MiniParserT__4:
		localctx = NewIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(MiniParserT__4)
		}

	case MiniParserT__5:
		localctx = NewBoolTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Match(MiniParserT__5)
		}

	case MiniParserT__0:
		localctx = NewStructTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Match(MiniParserT__0)
		}
		{
			p.SetState(70)
			p.Match(MiniParserID)
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

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

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
	p.RuleIndex = MiniParserRULE_declarations
	return p
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_declarations

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

func (p *MiniParser) Declarations() (localctx IDeclarationsContext) {
	this := p
	_ = this

	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MiniParserRULE_declarations)
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
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&98) != 0 {
		{
			p.SetState(73)
			p.Declaration()
		}

		p.SetState(78)
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

	// Getter signatures
	Type_() ITypeContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

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

func (s *DeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(MiniParserID)
}

func (s *DeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(MiniParserID, i)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MiniParserRULE_declaration)
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
		p.SetState(79)
		p.Type_()
	}
	{
		p.SetState(80)
		p.Match(MiniParserID)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MiniParserT__6 {
		{
			p.SetState(81)
			p.Match(MiniParserT__6)
		}
		{
			p.SetState(82)
			p.Match(MiniParserID)
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(MiniParserT__3)
	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunction() []IFunctionContext
	Function(i int) IFunctionContext

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
	p.RuleIndex = MiniParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_functions

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

func (p *MiniParser) Functions() (localctx IFunctionsContext) {
	this := p
	_ = this

	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MiniParserRULE_functions)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MiniParserT__7 {
		{
			p.SetState(90)
			p.Function()
		}

		p.SetState(95)
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

	// Getter signatures
	ID() antlr.TerminalNode
	Parameters() IParametersContext
	ReturnType() IReturnTypeContext
	Declarations() IDeclarationsContext
	StatementList() IStatementListContext

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
	p.RuleIndex = MiniParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
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

func (s *FunctionContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MiniParserRULE_function)

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
		p.SetState(96)
		p.Match(MiniParserT__7)
	}
	{
		p.SetState(97)
		p.Match(MiniParserID)
	}
	{
		p.SetState(98)
		p.Parameters()
	}
	{
		p.SetState(99)
		p.ReturnType()
	}
	{
		p.SetState(100)
		p.Match(MiniParserT__1)
	}
	{
		p.SetState(101)
		p.Declarations()
	}
	{
		p.SetState(102)
		p.StatementList()
	}
	{
		p.SetState(103)
		p.Match(MiniParserT__2)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDecl() []IDeclContext
	Decl(i int) IDeclContext

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
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

	return t.(IDeclContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Parameters() (localctx IParametersContext) {
	this := p
	_ = this

	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MiniParserRULE_parameters)
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
		p.SetState(105)
		p.Match(MiniParserT__8)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&98) != 0 {
		{
			p.SetState(106)
			p.Decl()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MiniParserT__6 {
			{
				p.SetState(107)
				p.Match(MiniParserT__6)
			}
			{
				p.SetState(108)
				p.Decl()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(116)
		p.Match(MiniParserT__9)
	}

	return localctx
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_returnType
	return p
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) CopyFrom(ctx *ReturnTypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnTypeVoidContext struct {
	*ReturnTypeContext
}

func NewReturnTypeVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnTypeVoidContext {
	var p = new(ReturnTypeVoidContext)

	p.ReturnTypeContext = NewEmptyReturnTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReturnTypeContext))

	return p
}

func (s *ReturnTypeVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

type ReturnTypeRealContext struct {
	*ReturnTypeContext
}

func NewReturnTypeRealContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnTypeRealContext {
	var p = new(ReturnTypeRealContext)

	p.ReturnTypeContext = NewEmptyReturnTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReturnTypeContext))

	return p
}

func (s *ReturnTypeRealContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeRealContext) Type_() ITypeContext {
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

func (p *MiniParser) ReturnType() (localctx IReturnTypeContext) {
	this := p
	_ = this

	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MiniParserRULE_returnType)

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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MiniParserT__0, MiniParserT__4, MiniParserT__5:
		localctx = NewReturnTypeRealContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Type_()
		}

	case MiniParserT__10:
		localctx = NewReturnTypeVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Match(MiniParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = MiniParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentContext struct {
	*StatementContext
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
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

type DeleteContext struct {
	*StatementContext
}

func NewDeleteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteContext {
	var p = new(DeleteContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
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

type PrintContext struct {
	*StatementContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) Expression() IExpressionContext {
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

type ReturnContext struct {
	*StatementContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
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

type InvocationContext struct {
	*StatementContext
}

func NewInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationContext {
	var p = new(InvocationContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

func (s *InvocationContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

type PrintLnContext struct {
	*StatementContext
}

func NewPrintLnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintLnContext {
	var p = new(PrintLnContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PrintLnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintLnContext) Expression() IExpressionContext {
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

type ConditionalContext struct {
	*StatementContext
	thenBlock IBlockContext
	elseBlock IBlockContext
}

func NewConditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalContext {
	var p = new(ConditionalContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ConditionalContext) GetThenBlock() IBlockContext { return s.thenBlock }

func (s *ConditionalContext) GetElseBlock() IBlockContext { return s.elseBlock }

func (s *ConditionalContext) SetThenBlock(v IBlockContext) { s.thenBlock = v }

func (s *ConditionalContext) SetElseBlock(v IBlockContext) { s.elseBlock = v }

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
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

type NestedBlockContext struct {
	*StatementContext
}

func NewNestedBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedBlockContext {
	var p = new(NestedBlockContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NestedBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedBlockContext) Block() IBlockContext {
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

type WhileContext struct {
	*StatementContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) Expression() IExpressionContext {
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

func (s *WhileContext) Block() IBlockContext {
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

func (p *MiniParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MiniParserRULE_statement)
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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNestedBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Block()
		}

	case 2:
		localctx = NewAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.lvalue(0)
		}
		{
			p.SetState(124)
			p.Match(MiniParserT__11)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MiniParserT__8, MiniParserT__21, MiniParserT__22, MiniParserT__34, MiniParserT__35, MiniParserT__36, MiniParserT__37, MiniParserID, MiniParserINTEGER:
			{
				p.SetState(125)
				p.expression(0)
			}

		case MiniParserT__12:
			{
				p.SetState(126)
				p.Match(MiniParserT__12)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(129)
			p.Match(MiniParserT__3)
		}

	case 3:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(MiniParserT__13)
		}
		{
			p.SetState(132)
			p.expression(0)
		}
		{
			p.SetState(133)
			p.Match(MiniParserT__3)
		}

	case 4:
		localctx = NewPrintLnContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.Match(MiniParserT__13)
		}
		{
			p.SetState(136)
			p.expression(0)
		}
		{
			p.SetState(137)
			p.Match(MiniParserT__14)
		}
		{
			p.SetState(138)
			p.Match(MiniParserT__3)
		}

	case 5:
		localctx = NewConditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(140)
			p.Match(MiniParserT__15)
		}
		{
			p.SetState(141)
			p.Match(MiniParserT__8)
		}
		{
			p.SetState(142)
			p.expression(0)
		}
		{
			p.SetState(143)
			p.Match(MiniParserT__9)
		}
		{
			p.SetState(144)

			var _x = p.Block()

			localctx.(*ConditionalContext).thenBlock = _x
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MiniParserT__16 {
			{
				p.SetState(145)
				p.Match(MiniParserT__16)
			}
			{
				p.SetState(146)

				var _x = p.Block()

				localctx.(*ConditionalContext).elseBlock = _x
			}

		}

	case 6:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(149)
			p.Match(MiniParserT__17)
		}
		{
			p.SetState(150)
			p.Match(MiniParserT__8)
		}
		{
			p.SetState(151)
			p.expression(0)
		}
		{
			p.SetState(152)
			p.Match(MiniParserT__9)
		}
		{
			p.SetState(153)
			p.Block()
		}

	case 7:
		localctx = NewDeleteContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Match(MiniParserT__18)
		}
		{
			p.SetState(156)
			p.expression(0)
		}
		{
			p.SetState(157)
			p.Match(MiniParserT__3)
		}

	case 8:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(159)
			p.Match(MiniParserT__19)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2164676100608) != 0 {
			{
				p.SetState(160)
				p.expression(0)
			}

		}
		{
			p.SetState(163)
			p.Match(MiniParserT__3)
		}

	case 9:
		localctx = NewInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(164)
			p.Match(MiniParserID)
		}
		{
			p.SetState(165)
			p.Match(MiniParserT__8)
		}
		{
			p.SetState(166)
			p.Arguments()
		}
		{
			p.SetState(167)
			p.Match(MiniParserT__9)
		}
		{
			p.SetState(168)
			p.Match(MiniParserT__3)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementList() IStatementListContext

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
	p.RuleIndex = MiniParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MiniParserRULE_block)

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
		p.Match(MiniParserT__1)
	}
	{
		p.SetState(173)
		p.StatementList()
	}
	{
		p.SetState(174)
		p.Match(MiniParserT__2)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
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

func (s *StatementListContext) Statement(i int) IStatementContext {
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

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) StatementList() (localctx IStatementListContext) {
	this := p
	_ = this

	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MiniParserRULE_statementList)
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549757730820) != 0 {
		{
			p.SetState(176)
			p.Statement()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) CopyFrom(ctx *LvalueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LvalueIdContext struct {
	*LvalueContext
}

func NewLvalueIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LvalueIdContext {
	var p = new(LvalueIdContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LvalueIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueIdContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

type LvalueDotContext struct {
	*LvalueContext
}

func NewLvalueDotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LvalueDotContext {
	var p = new(LvalueDotContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LvalueDotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueDotContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *LvalueDotContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

func (p *MiniParser) Lvalue() (localctx ILvalueContext) {
	return p.lvalue(0)
}

func (p *MiniParser) lvalue(_p int) (localctx ILvalueContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILvalueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, MiniParserRULE_lvalue, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewLvalueIdContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(183)
		p.Match(MiniParserID)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLvalueDotContext(p, NewLvalueContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_lvalue)
			p.SetState(185)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(186)
				p.Match(MiniParserT__20)
			}
			{
				p.SetState(187)
				p.Match(MiniParserID)
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.RuleIndex = MiniParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerExprContext struct {
	*ExpressionContext
}

func NewIntegerExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerExprContext {
	var p = new(IntegerExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntegerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerExprContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MiniParserINTEGER, 0)
}

type TrueExprContext struct {
	*ExpressionContext
}

func NewTrueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueExprContext {
	var p = new(TrueExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TrueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

type IdentifierExprContext struct {
	*ExpressionContext
}

func NewIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExprContext {
	var p = new(IdentifierExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExprContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

type BinaryExprContext struct {
	*ExpressionContext
	lft IExpressionContext
	op  antlr.Token
	rht IExpressionContext
}

func NewBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExprContext {
	var p = new(BinaryExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExprContext) GetLft() IExpressionContext { return s.lft }

func (s *BinaryExprContext) GetRht() IExpressionContext { return s.rht }

func (s *BinaryExprContext) SetLft(v IExpressionContext) { s.lft = v }

func (s *BinaryExprContext) SetRht(v IExpressionContext) { s.rht = v }

func (s *BinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type NewExprContext struct {
	*ExpressionContext
}

func NewNewExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewExprContext {
	var p = new(NewExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NewExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewExprContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

type NestedExprContext struct {
	*ExpressionContext
}

func NewNestedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExprContext {
	var p = new(NestedExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NestedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExprContext) Expression() IExpressionContext {
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

type DotExprContext struct {
	*ExpressionContext
}

func NewDotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotExprContext {
	var p = new(DotExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExprContext) Expression() IExpressionContext {
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

func (s *DotExprContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

type UnaryExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
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

type InvocationExprContext struct {
	*ExpressionContext
}

func NewInvocationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationExprContext {
	var p = new(InvocationExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InvocationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationExprContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniParserID, 0)
}

func (s *InvocationExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

type FalseExprContext struct {
	*ExpressionContext
}

func NewFalseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseExprContext {
	var p = new(FalseExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FalseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

type NullExprContext struct {
	*ExpressionContext
}

func NewNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprContext {
	var p = new(NullExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (p *MiniParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MiniParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, MiniParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInvocationExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(194)
			p.Match(MiniParserID)
		}
		{
			p.SetState(195)
			p.Match(MiniParserT__8)
		}
		{
			p.SetState(196)
			p.Arguments()
		}
		{
			p.SetState(197)
			p.Match(MiniParserT__9)
		}

	case 2:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(199)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MiniParserT__21 || _la == MiniParserT__22) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(200)
			p.expression(14)
		}

	case 3:
		localctx = NewIdentifierExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(201)
			p.Match(MiniParserID)
		}

	case 4:
		localctx = NewIntegerExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(202)
			p.Match(MiniParserINTEGER)
		}

	case 5:
		localctx = NewTrueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(203)
			p.Match(MiniParserT__34)
		}

	case 6:
		localctx = NewFalseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)
			p.Match(MiniParserT__35)
		}

	case 7:
		localctx = NewNewExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(205)
			p.Match(MiniParserT__36)
		}
		{
			p.SetState(206)
			p.Match(MiniParserID)
		}

	case 8:
		localctx = NewNullExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(207)
			p.Match(MiniParserT__37)
		}

	case 9:
		localctx = NewNestedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Match(MiniParserT__8)
		}
		{
			p.SetState(209)
			p.expression(0)
		}
		{
			p.SetState(210)
			p.Match(MiniParserT__9)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).lft = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(215)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MiniParserT__23 || _la == MiniParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(216)

					var _x = p.expression(14)

					localctx.(*BinaryExprContext).rht = _x
				}

			case 2:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).lft = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(218)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MiniParserT__21 || _la == MiniParserT__25) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(219)

					var _x = p.expression(13)

					localctx.(*BinaryExprContext).rht = _x
				}

			case 3:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).lft = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(221)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2013265920) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(222)

					var _x = p.expression(12)

					localctx.(*BinaryExprContext).rht = _x
				}

			case 4:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).lft = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(224)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MiniParserT__30 || _la == MiniParserT__31) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(225)

					var _x = p.expression(11)

					localctx.(*BinaryExprContext).rht = _x
				}

			case 5:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).lft = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(227)

					var _m = p.Match(MiniParserT__32)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(228)

					var _x = p.expression(10)

					localctx.(*BinaryExprContext).rht = _x
				}

			case 6:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).lft = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(230)

					var _m = p.Match(MiniParserT__33)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(231)

					var _x = p.expression(9)

					localctx.(*BinaryExprContext).rht = _x
				}

			case 7:
				localctx = NewDotExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniParserRULE_expression)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}

				{
					p.SetState(233)
					p.Match(MiniParserT__20)
				}
				{
					p.SetState(234)
					p.Match(MiniParserID)
				}

			}

		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MiniParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MiniParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MiniParserRULE_arguments)
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
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2164676100608) != 0 {
		{
			p.SetState(240)
			p.expression(0)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MiniParserT__6 {
			{
				p.SetState(241)
				p.Match(MiniParserT__6)
			}
			{
				p.SetState(242)
				p.expression(0)
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

func (p *MiniParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *LvalueContext = nil
		if localctx != nil {
			t = localctx.(*LvalueContext)
		}
		return p.Lvalue_Sempred(t, predIndex)

	case 16:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MiniParser) Lvalue_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MiniParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
