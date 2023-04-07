// Code generated from Mini.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

/* package declaration here */

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MiniLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var minilexerLexerStaticData struct {
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

func minilexerLexerInit() {
	staticData := &minilexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "ID", "INTEGER", "WS",
		"COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 257, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 5, 38,
		225, 8, 38, 10, 38, 12, 38, 228, 9, 38, 1, 39, 1, 39, 1, 39, 5, 39, 233,
		8, 39, 10, 39, 12, 39, 236, 9, 39, 3, 39, 238, 8, 39, 1, 40, 4, 40, 241,
		8, 40, 11, 40, 12, 40, 242, 1, 40, 1, 40, 1, 41, 1, 41, 5, 41, 249, 8,
		41, 10, 41, 12, 41, 252, 9, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 250, 0,
		42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 41, 83, 42, 1, 0, 5, 2, 0, 65, 90, 97, 122, 3,
		0, 48, 57, 65, 90, 97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 9, 10, 12,
		13, 32, 32, 261, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83,
		1, 0, 0, 0, 1, 85, 1, 0, 0, 0, 3, 92, 1, 0, 0, 0, 5, 94, 1, 0, 0, 0, 7,
		96, 1, 0, 0, 0, 9, 98, 1, 0, 0, 0, 11, 102, 1, 0, 0, 0, 13, 107, 1, 0,
		0, 0, 15, 109, 1, 0, 0, 0, 17, 113, 1, 0, 0, 0, 19, 115, 1, 0, 0, 0, 21,
		117, 1, 0, 0, 0, 23, 122, 1, 0, 0, 0, 25, 124, 1, 0, 0, 0, 27, 129, 1,
		0, 0, 0, 29, 135, 1, 0, 0, 0, 31, 140, 1, 0, 0, 0, 33, 143, 1, 0, 0, 0,
		35, 148, 1, 0, 0, 0, 37, 154, 1, 0, 0, 0, 39, 161, 1, 0, 0, 0, 41, 168,
		1, 0, 0, 0, 43, 170, 1, 0, 0, 0, 45, 172, 1, 0, 0, 0, 47, 174, 1, 0, 0,
		0, 49, 176, 1, 0, 0, 0, 51, 178, 1, 0, 0, 0, 53, 180, 1, 0, 0, 0, 55, 182,
		1, 0, 0, 0, 57, 184, 1, 0, 0, 0, 59, 187, 1, 0, 0, 0, 61, 190, 1, 0, 0,
		0, 63, 193, 1, 0, 0, 0, 65, 196, 1, 0, 0, 0, 67, 199, 1, 0, 0, 0, 69, 202,
		1, 0, 0, 0, 71, 207, 1, 0, 0, 0, 73, 213, 1, 0, 0, 0, 75, 217, 1, 0, 0,
		0, 77, 222, 1, 0, 0, 0, 79, 237, 1, 0, 0, 0, 81, 240, 1, 0, 0, 0, 83, 246,
		1, 0, 0, 0, 85, 86, 5, 115, 0, 0, 86, 87, 5, 116, 0, 0, 87, 88, 5, 114,
		0, 0, 88, 89, 5, 117, 0, 0, 89, 90, 5, 99, 0, 0, 90, 91, 5, 116, 0, 0,
		91, 2, 1, 0, 0, 0, 92, 93, 5, 123, 0, 0, 93, 4, 1, 0, 0, 0, 94, 95, 5,
		125, 0, 0, 95, 6, 1, 0, 0, 0, 96, 97, 5, 59, 0, 0, 97, 8, 1, 0, 0, 0, 98,
		99, 5, 105, 0, 0, 99, 100, 5, 110, 0, 0, 100, 101, 5, 116, 0, 0, 101, 10,
		1, 0, 0, 0, 102, 103, 5, 98, 0, 0, 103, 104, 5, 111, 0, 0, 104, 105, 5,
		111, 0, 0, 105, 106, 5, 108, 0, 0, 106, 12, 1, 0, 0, 0, 107, 108, 5, 44,
		0, 0, 108, 14, 1, 0, 0, 0, 109, 110, 5, 102, 0, 0, 110, 111, 5, 117, 0,
		0, 111, 112, 5, 110, 0, 0, 112, 16, 1, 0, 0, 0, 113, 114, 5, 40, 0, 0,
		114, 18, 1, 0, 0, 0, 115, 116, 5, 41, 0, 0, 116, 20, 1, 0, 0, 0, 117, 118,
		5, 118, 0, 0, 118, 119, 5, 111, 0, 0, 119, 120, 5, 105, 0, 0, 120, 121,
		5, 100, 0, 0, 121, 22, 1, 0, 0, 0, 122, 123, 5, 61, 0, 0, 123, 24, 1, 0,
		0, 0, 124, 125, 5, 114, 0, 0, 125, 126, 5, 101, 0, 0, 126, 127, 5, 97,
		0, 0, 127, 128, 5, 100, 0, 0, 128, 26, 1, 0, 0, 0, 129, 130, 5, 112, 0,
		0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 105, 0, 0, 132, 133, 5, 110, 0,
		0, 133, 134, 5, 116, 0, 0, 134, 28, 1, 0, 0, 0, 135, 136, 5, 101, 0, 0,
		136, 137, 5, 110, 0, 0, 137, 138, 5, 100, 0, 0, 138, 139, 5, 108, 0, 0,
		139, 30, 1, 0, 0, 0, 140, 141, 5, 105, 0, 0, 141, 142, 5, 102, 0, 0, 142,
		32, 1, 0, 0, 0, 143, 144, 5, 101, 0, 0, 144, 145, 5, 108, 0, 0, 145, 146,
		5, 115, 0, 0, 146, 147, 5, 101, 0, 0, 147, 34, 1, 0, 0, 0, 148, 149, 5,
		119, 0, 0, 149, 150, 5, 104, 0, 0, 150, 151, 5, 105, 0, 0, 151, 152, 5,
		108, 0, 0, 152, 153, 5, 101, 0, 0, 153, 36, 1, 0, 0, 0, 154, 155, 5, 100,
		0, 0, 155, 156, 5, 101, 0, 0, 156, 157, 5, 108, 0, 0, 157, 158, 5, 101,
		0, 0, 158, 159, 5, 116, 0, 0, 159, 160, 5, 101, 0, 0, 160, 38, 1, 0, 0,
		0, 161, 162, 5, 114, 0, 0, 162, 163, 5, 101, 0, 0, 163, 164, 5, 116, 0,
		0, 164, 165, 5, 117, 0, 0, 165, 166, 5, 114, 0, 0, 166, 167, 5, 110, 0,
		0, 167, 40, 1, 0, 0, 0, 168, 169, 5, 46, 0, 0, 169, 42, 1, 0, 0, 0, 170,
		171, 5, 45, 0, 0, 171, 44, 1, 0, 0, 0, 172, 173, 5, 33, 0, 0, 173, 46,
		1, 0, 0, 0, 174, 175, 5, 42, 0, 0, 175, 48, 1, 0, 0, 0, 176, 177, 5, 47,
		0, 0, 177, 50, 1, 0, 0, 0, 178, 179, 5, 43, 0, 0, 179, 52, 1, 0, 0, 0,
		180, 181, 5, 60, 0, 0, 181, 54, 1, 0, 0, 0, 182, 183, 5, 62, 0, 0, 183,
		56, 1, 0, 0, 0, 184, 185, 5, 60, 0, 0, 185, 186, 5, 61, 0, 0, 186, 58,
		1, 0, 0, 0, 187, 188, 5, 62, 0, 0, 188, 189, 5, 61, 0, 0, 189, 60, 1, 0,
		0, 0, 190, 191, 5, 61, 0, 0, 191, 192, 5, 61, 0, 0, 192, 62, 1, 0, 0, 0,
		193, 194, 5, 33, 0, 0, 194, 195, 5, 61, 0, 0, 195, 64, 1, 0, 0, 0, 196,
		197, 5, 38, 0, 0, 197, 198, 5, 38, 0, 0, 198, 66, 1, 0, 0, 0, 199, 200,
		5, 124, 0, 0, 200, 201, 5, 124, 0, 0, 201, 68, 1, 0, 0, 0, 202, 203, 5,
		116, 0, 0, 203, 204, 5, 114, 0, 0, 204, 205, 5, 117, 0, 0, 205, 206, 5,
		101, 0, 0, 206, 70, 1, 0, 0, 0, 207, 208, 5, 102, 0, 0, 208, 209, 5, 97,
		0, 0, 209, 210, 5, 108, 0, 0, 210, 211, 5, 115, 0, 0, 211, 212, 5, 101,
		0, 0, 212, 72, 1, 0, 0, 0, 213, 214, 5, 110, 0, 0, 214, 215, 5, 101, 0,
		0, 215, 216, 5, 119, 0, 0, 216, 74, 1, 0, 0, 0, 217, 218, 5, 110, 0, 0,
		218, 219, 5, 117, 0, 0, 219, 220, 5, 108, 0, 0, 220, 221, 5, 108, 0, 0,
		221, 76, 1, 0, 0, 0, 222, 226, 7, 0, 0, 0, 223, 225, 7, 1, 0, 0, 224, 223,
		1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0,
		0, 0, 227, 78, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 238, 5, 48, 0, 0,
		230, 234, 7, 2, 0, 0, 231, 233, 7, 3, 0, 0, 232, 231, 1, 0, 0, 0, 233,
		236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 238,
		1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 229, 1, 0, 0, 0, 237, 230, 1, 0,
		0, 0, 238, 80, 1, 0, 0, 0, 239, 241, 7, 4, 0, 0, 240, 239, 1, 0, 0, 0,
		241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243,
		244, 1, 0, 0, 0, 244, 245, 6, 40, 0, 0, 245, 82, 1, 0, 0, 0, 246, 250,
		5, 35, 0, 0, 247, 249, 9, 0, 0, 0, 248, 247, 1, 0, 0, 0, 249, 252, 1, 0,
		0, 0, 250, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0,
		252, 250, 1, 0, 0, 0, 253, 254, 5, 10, 0, 0, 254, 255, 1, 0, 0, 0, 255,
		256, 6, 41, 0, 0, 256, 84, 1, 0, 0, 0, 6, 0, 226, 234, 237, 242, 250, 1,
		6, 0, 0,
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

// MiniLexerInit initializes any static state used to implement MiniLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMiniLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniLexerInit() {
	staticData := &minilexerLexerStaticData
	staticData.once.Do(minilexerLexerInit)
}

// NewMiniLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMiniLexer(input antlr.CharStream) *MiniLexer {
	MiniLexerInit()
	l := new(MiniLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &minilexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Mini.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MiniLexer tokens.
const (
	MiniLexerT__0    = 1
	MiniLexerT__1    = 2
	MiniLexerT__2    = 3
	MiniLexerT__3    = 4
	MiniLexerT__4    = 5
	MiniLexerT__5    = 6
	MiniLexerT__6    = 7
	MiniLexerT__7    = 8
	MiniLexerT__8    = 9
	MiniLexerT__9    = 10
	MiniLexerT__10   = 11
	MiniLexerT__11   = 12
	MiniLexerT__12   = 13
	MiniLexerT__13   = 14
	MiniLexerT__14   = 15
	MiniLexerT__15   = 16
	MiniLexerT__16   = 17
	MiniLexerT__17   = 18
	MiniLexerT__18   = 19
	MiniLexerT__19   = 20
	MiniLexerT__20   = 21
	MiniLexerT__21   = 22
	MiniLexerT__22   = 23
	MiniLexerT__23   = 24
	MiniLexerT__24   = 25
	MiniLexerT__25   = 26
	MiniLexerT__26   = 27
	MiniLexerT__27   = 28
	MiniLexerT__28   = 29
	MiniLexerT__29   = 30
	MiniLexerT__30   = 31
	MiniLexerT__31   = 32
	MiniLexerT__32   = 33
	MiniLexerT__33   = 34
	MiniLexerT__34   = 35
	MiniLexerT__35   = 36
	MiniLexerT__36   = 37
	MiniLexerT__37   = 38
	MiniLexerID      = 39
	MiniLexerINTEGER = 40
	MiniLexerWS      = 41
	MiniLexerCOMMENT = 42
)
