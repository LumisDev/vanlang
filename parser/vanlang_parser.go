// Generated from ./grammar/VanLang.g4 by ANTLR 4.7.

package parser // VanLang
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []int32{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 234,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14, 2,
	35, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 56, 10,
	4, 12, 4, 14, 4, 59, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 67,
	10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 84, 10, 6, 12, 6, 14, 6, 87, 11, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 96, 10, 7, 12, 7, 14, 7, 99,
	11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 108, 10, 7, 12,
	7, 14, 7, 111, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 118, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 127, 10, 9, 12, 9, 14, 9,
	130, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	141, 10, 9, 12, 9, 14, 9, 144, 11, 9, 3, 9, 3, 9, 3, 9, 5, 9, 149, 10,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10,
	160, 10, 10, 12, 10, 14, 10, 163, 11, 10, 3, 10, 3, 10, 5, 10, 167, 10,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 176, 10, 11,
	12, 11, 14, 11, 179, 11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 194, 10, 11, 3,
	12, 3, 12, 3, 12, 7, 12, 199, 10, 12, 12, 12, 14, 12, 202, 11, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 211, 10, 13, 12, 13, 14,
	13, 214, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 221, 10, 14,
	12, 14, 14, 14, 224, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	5, 15, 232, 10, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 2, 5, 3, 2, 43, 47, 4, 2, 24, 25, 35, 36, 3, 2, 37, 42,
	2, 251, 2, 33, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 62,
	3, 2, 2, 2, 10, 73, 3, 2, 2, 2, 12, 117, 3, 2, 2, 2, 14, 119, 3, 2, 2,
	2, 16, 148, 3, 2, 2, 2, 18, 150, 3, 2, 2, 2, 20, 193, 3, 2, 2, 2, 22, 195,
	3, 2, 2, 2, 24, 205, 3, 2, 2, 2, 26, 215, 3, 2, 2, 2, 28, 231, 3, 2, 2,
	2, 30, 32, 5, 4, 3, 2, 31, 30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 3, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2,
	36, 43, 5, 6, 4, 2, 37, 43, 5, 8, 5, 2, 38, 43, 5, 14, 8, 2, 39, 43, 5,
	16, 9, 2, 40, 43, 5, 18, 10, 2, 41, 43, 5, 20, 11, 2, 42, 36, 3, 2, 2,
	2, 42, 37, 3, 2, 2, 2, 42, 38, 3, 2, 2, 2, 42, 39, 3, 2, 2, 2, 42, 40,
	3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 5, 3, 2, 2, 2, 44, 45, 7, 3, 2, 2,
	45, 46, 7, 49, 2, 2, 46, 47, 7, 54, 2, 2, 47, 48, 5, 12, 7, 2, 48, 57,
	7, 55, 2, 2, 49, 50, 7, 56, 2, 2, 50, 51, 7, 49, 2, 2, 51, 52, 7, 54, 2,
	2, 52, 53, 5, 12, 7, 2, 53, 54, 7, 55, 2, 2, 54, 56, 3, 2, 2, 2, 55, 49,
	3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2,
	58, 60, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 61, 7, 57, 2, 2, 61, 7, 3,
	2, 2, 2, 62, 63, 7, 4, 2, 2, 63, 64, 7, 49, 2, 2, 64, 66, 7, 52, 2, 2,
	65, 67, 5, 10, 6, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3,
	2, 2, 2, 68, 69, 7, 53, 2, 2, 69, 70, 7, 58, 2, 2, 70, 71, 5, 12, 7, 2,
	71, 72, 7, 57, 2, 2, 72, 9, 3, 2, 2, 2, 73, 74, 7, 49, 2, 2, 74, 75, 7,
	54, 2, 2, 75, 76, 5, 12, 7, 2, 76, 85, 7, 55, 2, 2, 77, 78, 7, 56, 2, 2,
	78, 79, 7, 49, 2, 2, 79, 80, 7, 54, 2, 2, 80, 81, 5, 12, 7, 2, 81, 82,
	7, 55, 2, 2, 82, 84, 3, 2, 2, 2, 83, 77, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2,
	85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 85, 3,
	2, 2, 2, 88, 118, 7, 27, 2, 2, 89, 118, 7, 28, 2, 2, 90, 91, 7, 30, 2,
	2, 91, 92, 7, 52, 2, 2, 92, 97, 5, 12, 7, 2, 93, 94, 7, 56, 2, 2, 94, 96,
	5, 12, 7, 2, 95, 93, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2,
	97, 98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101,
	7, 53, 2, 2, 101, 118, 3, 2, 2, 2, 102, 103, 7, 29, 2, 2, 103, 104, 7,
	52, 2, 2, 104, 109, 5, 12, 7, 2, 105, 106, 7, 56, 2, 2, 106, 108, 5, 12,
	7, 2, 107, 105, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112,
	113, 7, 53, 2, 2, 113, 118, 3, 2, 2, 2, 114, 118, 7, 26, 2, 2, 115, 118,
	7, 31, 2, 2, 116, 118, 7, 32, 2, 2, 117, 88, 3, 2, 2, 2, 117, 89, 3, 2,
	2, 2, 117, 90, 3, 2, 2, 2, 117, 102, 3, 2, 2, 2, 117, 114, 3, 2, 2, 2,
	117, 115, 3, 2, 2, 2, 117, 116, 3, 2, 2, 2, 118, 13, 3, 2, 2, 2, 119, 120,
	7, 49, 2, 2, 120, 121, 7, 37, 2, 2, 121, 122, 5, 22, 12, 2, 122, 123, 7,
	57, 2, 2, 123, 15, 3, 2, 2, 2, 124, 128, 7, 14, 2, 2, 125, 127, 5, 4, 3,
	2, 126, 125, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132,
	7, 15, 2, 2, 132, 133, 5, 22, 12, 2, 133, 134, 7, 57, 2, 2, 134, 149, 3,
	2, 2, 2, 135, 136, 7, 16, 2, 2, 136, 137, 5, 22, 12, 2, 137, 138, 7, 17,
	2, 2, 138, 142, 7, 5, 2, 2, 139, 141, 5, 4, 3, 2, 140, 139, 3, 2, 2, 2,
	141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143,
	145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 6, 2, 2, 146, 147,
	7, 57, 2, 2, 147, 149, 3, 2, 2, 2, 148, 124, 3, 2, 2, 2, 148, 135, 3, 2,
	2, 2, 149, 17, 3, 2, 2, 2, 150, 151, 7, 7, 2, 2, 151, 152, 5, 22, 12, 2,
	152, 153, 7, 8, 2, 2, 153, 161, 5, 4, 3, 2, 154, 155, 7, 9, 2, 2, 155,
	156, 5, 22, 12, 2, 156, 157, 7, 8, 2, 2, 157, 158, 5, 4, 3, 2, 158, 160,
	3, 2, 2, 2, 159, 154, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2,
	2, 2, 161, 162, 3, 2, 2, 2, 162, 166, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2,
	164, 165, 7, 10, 2, 2, 165, 167, 5, 4, 3, 2, 166, 164, 3, 2, 2, 2, 166,
	167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 7, 57, 2, 2, 169, 19,
	3, 2, 2, 2, 170, 171, 7, 12, 2, 2, 171, 172, 7, 52, 2, 2, 172, 177, 7,
	51, 2, 2, 173, 174, 7, 56, 2, 2, 174, 176, 5, 22, 12, 2, 175, 173, 3, 2,
	2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2,
	178, 180, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 180, 181, 7, 53, 2, 2, 181,
	194, 7, 57, 2, 2, 182, 183, 7, 49, 2, 2, 183, 184, 7, 37, 2, 2, 184, 185,
	7, 13, 2, 2, 185, 186, 7, 52, 2, 2, 186, 187, 7, 51, 2, 2, 187, 188, 7,
	53, 2, 2, 188, 194, 7, 57, 2, 2, 189, 190, 7, 11, 2, 2, 190, 191, 5, 22,
	12, 2, 191, 192, 7, 57, 2, 2, 192, 194, 3, 2, 2, 2, 193, 170, 3, 2, 2,
	2, 193, 182, 3, 2, 2, 2, 193, 189, 3, 2, 2, 2, 194, 21, 3, 2, 2, 2, 195,
	200, 5, 24, 13, 2, 196, 197, 9, 2, 2, 2, 197, 199, 5, 24, 13, 2, 198, 196,
	3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2,
	2, 2, 201, 203, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 204, 7, 57, 2, 2,
	204, 23, 3, 2, 2, 2, 205, 212, 5, 26, 14, 2, 206, 207, 9, 3, 2, 2, 207,
	211, 5, 26, 14, 2, 208, 209, 9, 4, 2, 2, 209, 211, 5, 26, 14, 2, 210, 206,
	3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 210, 3, 2,
	2, 2, 212, 213, 3, 2, 2, 2, 213, 25, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2,
	215, 222, 5, 28, 15, 2, 216, 217, 9, 3, 2, 2, 217, 221, 5, 28, 15, 2, 218,
	219, 9, 4, 2, 2, 219, 221, 5, 28, 15, 2, 220, 216, 3, 2, 2, 2, 220, 218,
	3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2,
	2, 2, 223, 27, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 232, 7, 49, 2, 2,
	226, 232, 7, 50, 2, 2, 227, 228, 7, 52, 2, 2, 228, 229, 5, 22, 12, 2, 229,
	230, 7, 53, 2, 2, 230, 232, 3, 2, 2, 2, 231, 225, 3, 2, 2, 2, 231, 226,
	3, 2, 2, 2, 231, 227, 3, 2, 2, 2, 232, 29, 3, 2, 2, 2, 23, 33, 42, 57,
	66, 85, 97, 109, 117, 128, 142, 148, 161, 166, 177, 193, 200, 210, 212,
	220, 222, 231,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.Deserialize(parserATN)

var literalNames = []string{
	"", "'var'", "'func'", "'begin'", "'end'", "'if'", "'then'", "'elseif'",
	"'else'", "'return'", "'println'", "'readln'", "'repeat'", "'until'", "'while'",
	"'do'", "'not'", "'and'", "'or'", "'xor'", "'imp'", "'eqv'", "'div'", "'mod'",
	"'bool'", "'int'", "'string'", "'tuple'", "'array'", "'float'", "'unsigned'",
	"'+'", "'-'", "'*'", "'/'", "'='", "'<>'", "'<'", "'>'", "'<='", "'>='",
	"", "", "", "", "", "", "", "", "", "'('", "')'", "'['", "']'", "','",
	"';'", "'->'",
}
var symbolicNames = []string{
	"", "VAR", "FUNC", "BEGIN", "END", "IF", "THEN", "ELSEIF", "ELSE", "RETURN",
	"PRINTLN", "READLN", "REPEAT", "UNTIL", "WHILE", "DO", "NOT", "AND", "OR",
	"XOR", "IMP", "EQV", "DIV", "MOD", "BOOL", "INTTYPE", "STRTYPE", "TUPLETYPE",
	"ARRAYTYPE", "FLOAT", "UNSIGNED", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
	"EQUAL", "NOTEQUAL", "LESS", "GREATER", "LESSEQUAL", "GREATEREQUAL", "AND_OP",
	"OR_OP", "XOR_OP", "EQV_OP", "IMP_OP", "NOT_OP", "ID", "INT", "STRING",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "COMMA", "SEMICOLON", "ARROW",
	"WS",
}

var ruleNames = []string{
	"program", "statement", "variableDeclaration", "functionDeclaration", "argumentList",
	"datatype", "logicStatement", "loopStatement", "conditionStatement", "expressionStatement",
	"expression", "term", "factor", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type VanLangParser struct {
	*antlr.BaseParser
}

func NewVanLangParser(input antlr.TokenStream) *VanLangParser {
	this := new(VanLangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "VanLang.g4"

	return this
}

// VanLangParser tokens.
const (
	VanLangParserEOF          = antlr.TokenEOF
	VanLangParserVAR          = 1
	VanLangParserFUNC         = 2
	VanLangParserBEGIN        = 3
	VanLangParserEND          = 4
	VanLangParserIF           = 5
	VanLangParserTHEN         = 6
	VanLangParserELSEIF       = 7
	VanLangParserELSE         = 8
	VanLangParserRETURN       = 9
	VanLangParserPRINTLN      = 10
	VanLangParserREADLN       = 11
	VanLangParserREPEAT       = 12
	VanLangParserUNTIL        = 13
	VanLangParserWHILE        = 14
	VanLangParserDO           = 15
	VanLangParserNOT          = 16
	VanLangParserAND          = 17
	VanLangParserOR           = 18
	VanLangParserXOR          = 19
	VanLangParserIMP          = 20
	VanLangParserEQV          = 21
	VanLangParserDIV          = 22
	VanLangParserMOD          = 23
	VanLangParserBOOL         = 24
	VanLangParserINTTYPE      = 25
	VanLangParserSTRTYPE      = 26
	VanLangParserTUPLETYPE    = 27
	VanLangParserARRAYTYPE    = 28
	VanLangParserFLOAT        = 29
	VanLangParserUNSIGNED     = 30
	VanLangParserPLUS         = 31
	VanLangParserMINUS        = 32
	VanLangParserMULTIPLY     = 33
	VanLangParserDIVIDE       = 34
	VanLangParserEQUAL        = 35
	VanLangParserNOTEQUAL     = 36
	VanLangParserLESS         = 37
	VanLangParserGREATER      = 38
	VanLangParserLESSEQUAL    = 39
	VanLangParserGREATEREQUAL = 40
	VanLangParserAND_OP       = 41
	VanLangParserOR_OP        = 42
	VanLangParserXOR_OP       = 43
	VanLangParserEQV_OP       = 44
	VanLangParserIMP_OP       = 45
	VanLangParserNOT_OP       = 46
	VanLangParserID           = 47
	VanLangParserINT          = 48
	VanLangParserSTRING       = 49
	VanLangParserLPAREN       = 50
	VanLangParserRPAREN       = 51
	VanLangParserLBRACK       = 52
	VanLangParserRBRACK       = 53
	VanLangParserCOMMA        = 54
	VanLangParserSEMICOLON    = 55
	VanLangParserARROW        = 56
	VanLangParserWS           = 57
)

// VanLangParser rules.
const (
	VanLangParserRULE_program             = 0
	VanLangParserRULE_statement           = 1
	VanLangParserRULE_variableDeclaration = 2
	VanLangParserRULE_functionDeclaration = 3
	VanLangParserRULE_argumentList        = 4
	VanLangParserRULE_datatype            = 5
	VanLangParserRULE_logicStatement      = 6
	VanLangParserRULE_loopStatement       = 7
	VanLangParserRULE_conditionStatement  = 8
	VanLangParserRULE_expressionStatement = 9
	VanLangParserRULE_expression          = 10
	VanLangParserRULE_term                = 11
	VanLangParserRULE_factor              = 12
	VanLangParserRULE_atom                = 13
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
	p.RuleIndex = VanLangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VanLangParserRULE_program)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<VanLangParserVAR)|(1<<VanLangParserFUNC)|(1<<VanLangParserIF)|(1<<VanLangParserRETURN)|(1<<VanLangParserPRINTLN)|(1<<VanLangParserREPEAT)|(1<<VanLangParserWHILE))) != 0) || _la == VanLangParserID {
		{
			p.SetState(28)
			p.Statement()
		}

		p.SetState(33)
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
	p.RuleIndex = VanLangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) LogicStatement() ILogicStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicStatementContext)
}

func (s *StatementContext) LoopStatement() ILoopStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *StatementContext) ConditionStatement() IConditionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VanLangParserRULE_statement)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.VariableDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.FunctionDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.LogicStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.LoopStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(38)
			p.ConditionStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(39)
			p.ExpressionStatement()
		}

	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(VanLangParserVAR, 0)
}

func (s *VariableDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserID)
}

func (s *VariableDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserID, i)
}

func (s *VariableDeclarationContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserLBRACK)
}

func (s *VariableDeclarationContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserLBRACK, i)
}

func (s *VariableDeclarationContext) AllDatatype() []IDatatypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatatypeContext)(nil)).Elem())
	var tst = make([]IDatatypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatatypeContext)
		}
	}

	return tst
}

func (s *VariableDeclarationContext) Datatype(i int) IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *VariableDeclarationContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserRBRACK)
}

func (s *VariableDeclarationContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserRBRACK, i)
}

func (s *VariableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *VariableDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserCOMMA)
}

func (s *VariableDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserCOMMA, i)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VanLangParserRULE_variableDeclaration)
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
		p.SetState(42)
		p.Match(VanLangParserVAR)
	}
	{
		p.SetState(43)
		p.Match(VanLangParserID)
	}
	{
		p.SetState(44)
		p.Match(VanLangParserLBRACK)
	}
	{
		p.SetState(45)
		p.Datatype()
	}
	{
		p.SetState(46)
		p.Match(VanLangParserRBRACK)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VanLangParserCOMMA {
		{
			p.SetState(47)
			p.Match(VanLangParserCOMMA)
		}
		{
			p.SetState(48)
			p.Match(VanLangParserID)
		}
		{
			p.SetState(49)
			p.Match(VanLangParserLBRACK)
		}
		{
			p.SetState(50)
			p.Datatype()
		}
		{
			p.SetState(51)
			p.Match(VanLangParserRBRACK)
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(VanLangParserSEMICOLON)
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FUNC() antlr.TerminalNode {
	return s.GetToken(VanLangParserFUNC, 0)
}

func (s *FunctionDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VanLangParserID, 0)
}

func (s *FunctionDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserLPAREN, 0)
}

func (s *FunctionDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserRPAREN, 0)
}

func (s *FunctionDeclarationContext) ARROW() antlr.TerminalNode {
	return s.GetToken(VanLangParserARROW, 0)
}

func (s *FunctionDeclarationContext) Datatype() IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *FunctionDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *FunctionDeclarationContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VanLangParserRULE_functionDeclaration)
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
		p.SetState(60)
		p.Match(VanLangParserFUNC)
	}
	{
		p.SetState(61)
		p.Match(VanLangParserID)
	}
	{
		p.SetState(62)
		p.Match(VanLangParserLPAREN)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VanLangParserID {
		{
			p.SetState(63)
			p.ArgumentList()
		}

	}
	{
		p.SetState(66)
		p.Match(VanLangParserRPAREN)
	}
	{
		p.SetState(67)
		p.Match(VanLangParserARROW)
	}
	{
		p.SetState(68)
		p.Datatype()
	}
	{
		p.SetState(69)
		p.Match(VanLangParserSEMICOLON)
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserID)
}

func (s *ArgumentListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserID, i)
}

func (s *ArgumentListContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserLBRACK)
}

func (s *ArgumentListContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserLBRACK, i)
}

func (s *ArgumentListContext) AllDatatype() []IDatatypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatatypeContext)(nil)).Elem())
	var tst = make([]IDatatypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatatypeContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) Datatype(i int) IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *ArgumentListContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserRBRACK)
}

func (s *ArgumentListContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserRBRACK, i)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VanLangParserRULE_argumentList)
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
		p.SetState(71)
		p.Match(VanLangParserID)
	}
	{
		p.SetState(72)
		p.Match(VanLangParserLBRACK)
	}
	{
		p.SetState(73)
		p.Datatype()
	}
	{
		p.SetState(74)
		p.Match(VanLangParserRBRACK)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VanLangParserCOMMA {
		{
			p.SetState(75)
			p.Match(VanLangParserCOMMA)
		}
		{
			p.SetState(76)
			p.Match(VanLangParserID)
		}
		{
			p.SetState(77)
			p.Match(VanLangParserLBRACK)
		}
		{
			p.SetState(78)
			p.Datatype()
		}
		{
			p.SetState(79)
			p.Match(VanLangParserRBRACK)
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDatatypeContext is an interface to support dynamic dispatch.
type IDatatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeContext() *DatatypeContext {
	var p = new(DatatypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_datatype
	return p
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) INTTYPE() antlr.TerminalNode {
	return s.GetToken(VanLangParserINTTYPE, 0)
}

func (s *DatatypeContext) STRTYPE() antlr.TerminalNode {
	return s.GetToken(VanLangParserSTRTYPE, 0)
}

func (s *DatatypeContext) ARRAYTYPE() antlr.TerminalNode {
	return s.GetToken(VanLangParserARRAYTYPE, 0)
}

func (s *DatatypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserLPAREN, 0)
}

func (s *DatatypeContext) AllDatatype() []IDatatypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatatypeContext)(nil)).Elem())
	var tst = make([]IDatatypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatatypeContext)
		}
	}

	return tst
}

func (s *DatatypeContext) Datatype(i int) IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *DatatypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserRPAREN, 0)
}

func (s *DatatypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserCOMMA)
}

func (s *DatatypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserCOMMA, i)
}

func (s *DatatypeContext) TUPLETYPE() antlr.TerminalNode {
	return s.GetToken(VanLangParserTUPLETYPE, 0)
}

func (s *DatatypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(VanLangParserBOOL, 0)
}

func (s *DatatypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(VanLangParserFLOAT, 0)
}

func (s *DatatypeContext) UNSIGNED() antlr.TerminalNode {
	return s.GetToken(VanLangParserUNSIGNED, 0)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitDatatype(s)
	}
}

func (s *DatatypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitDatatype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Datatype() (localctx IDatatypeContext) {
	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VanLangParserRULE_datatype)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VanLangParserINTTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(VanLangParserINTTYPE)
		}

	case VanLangParserSTRTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(VanLangParserSTRTYPE)
		}

	case VanLangParserARRAYTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(VanLangParserARRAYTYPE)
		}
		{
			p.SetState(89)
			p.Match(VanLangParserLPAREN)
		}
		{
			p.SetState(90)
			p.Datatype()
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == VanLangParserCOMMA {
			{
				p.SetState(91)
				p.Match(VanLangParserCOMMA)
			}
			{
				p.SetState(92)
				p.Datatype()
			}

			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(98)
			p.Match(VanLangParserRPAREN)
		}

	case VanLangParserTUPLETYPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)
			p.Match(VanLangParserTUPLETYPE)
		}
		{
			p.SetState(101)
			p.Match(VanLangParserLPAREN)
		}
		{
			p.SetState(102)
			p.Datatype()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == VanLangParserCOMMA {
			{
				p.SetState(103)
				p.Match(VanLangParserCOMMA)
			}
			{
				p.SetState(104)
				p.Datatype()
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(110)
			p.Match(VanLangParserRPAREN)
		}

	case VanLangParserBOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Match(VanLangParserBOOL)
		}

	case VanLangParserFLOAT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.Match(VanLangParserFLOAT)
		}

	case VanLangParserUNSIGNED:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(114)
			p.Match(VanLangParserUNSIGNED)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILogicStatementContext is an interface to support dynamic dispatch.
type ILogicStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicStatementContext differentiates from other interfaces.
	IsLogicStatementContext()
}

type LogicStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicStatementContext() *LogicStatementContext {
	var p = new(LogicStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_logicStatement
	return p
}

func (*LogicStatementContext) IsLogicStatementContext() {}

func NewLogicStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicStatementContext {
	var p = new(LogicStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_logicStatement

	return p
}

func (s *LogicStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(VanLangParserID, 0)
}

func (s *LogicStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(VanLangParserEQUAL, 0)
}

func (s *LogicStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *LogicStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterLogicStatement(s)
	}
}

func (s *LogicStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitLogicStatement(s)
	}
}

func (s *LogicStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitLogicStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) LogicStatement() (localctx ILogicStatementContext) {
	localctx = NewLogicStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VanLangParserRULE_logicStatement)

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
		p.SetState(117)
		p.Match(VanLangParserID)
	}
	{
		p.SetState(118)
		p.Match(VanLangParserEQUAL)
	}
	{
		p.SetState(119)
		p.Expression()
	}
	{
		p.SetState(120)
		p.Match(VanLangParserSEMICOLON)
	}

	return localctx
}

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_loopStatement
	return p
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(VanLangParserREPEAT, 0)
}

func (s *LoopStatementContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(VanLangParserUNTIL, 0)
}

func (s *LoopStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *LoopStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *LoopStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LoopStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(VanLangParserWHILE, 0)
}

func (s *LoopStatementContext) DO() antlr.TerminalNode {
	return s.GetToken(VanLangParserDO, 0)
}

func (s *LoopStatementContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(VanLangParserBEGIN, 0)
}

func (s *LoopStatementContext) END() antlr.TerminalNode {
	return s.GetToken(VanLangParserEND, 0)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) LoopStatement() (localctx ILoopStatementContext) {
	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VanLangParserRULE_loopStatement)
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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VanLangParserREPEAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(VanLangParserREPEAT)
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<VanLangParserVAR)|(1<<VanLangParserFUNC)|(1<<VanLangParserIF)|(1<<VanLangParserRETURN)|(1<<VanLangParserPRINTLN)|(1<<VanLangParserREPEAT)|(1<<VanLangParserWHILE))) != 0) || _la == VanLangParserID {
			{
				p.SetState(123)
				p.Statement()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(VanLangParserUNTIL)
		}
		{
			p.SetState(130)
			p.Expression()
		}
		{
			p.SetState(131)
			p.Match(VanLangParserSEMICOLON)
		}

	case VanLangParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(VanLangParserWHILE)
		}
		{
			p.SetState(134)
			p.Expression()
		}
		{
			p.SetState(135)
			p.Match(VanLangParserDO)
		}
		{
			p.SetState(136)
			p.Match(VanLangParserBEGIN)
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<VanLangParserVAR)|(1<<VanLangParserFUNC)|(1<<VanLangParserIF)|(1<<VanLangParserRETURN)|(1<<VanLangParserPRINTLN)|(1<<VanLangParserREPEAT)|(1<<VanLangParserWHILE))) != 0) || _la == VanLangParserID {
			{
				p.SetState(137)
				p.Statement()
			}

			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(143)
			p.Match(VanLangParserEND)
		}
		{
			p.SetState(144)
			p.Match(VanLangParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionStatementContext is an interface to support dynamic dispatch.
type IConditionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionStatementContext differentiates from other interfaces.
	IsConditionStatementContext()
}

type ConditionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionStatementContext() *ConditionStatementContext {
	var p = new(ConditionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_conditionStatement
	return p
}

func (*ConditionStatementContext) IsConditionStatementContext() {}

func NewConditionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionStatementContext {
	var p = new(ConditionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_conditionStatement

	return p
}

func (s *ConditionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(VanLangParserIF, 0)
}

func (s *ConditionStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionStatementContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserTHEN)
}

func (s *ConditionStatementContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserTHEN, i)
}

func (s *ConditionStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ConditionStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ConditionStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *ConditionStatementContext) AllELSEIF() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserELSEIF)
}

func (s *ConditionStatementContext) ELSEIF(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserELSEIF, i)
}

func (s *ConditionStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VanLangParserELSE, 0)
}

func (s *ConditionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterConditionStatement(s)
	}
}

func (s *ConditionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitConditionStatement(s)
	}
}

func (s *ConditionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitConditionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) ConditionStatement() (localctx IConditionStatementContext) {
	localctx = NewConditionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VanLangParserRULE_conditionStatement)
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
		p.SetState(148)
		p.Match(VanLangParserIF)
	}
	{
		p.SetState(149)
		p.Expression()
	}
	{
		p.SetState(150)
		p.Match(VanLangParserTHEN)
	}
	{
		p.SetState(151)
		p.Statement()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VanLangParserELSEIF {
		{
			p.SetState(152)
			p.Match(VanLangParserELSEIF)
		}
		{
			p.SetState(153)
			p.Expression()
		}
		{
			p.SetState(154)
			p.Match(VanLangParserTHEN)
		}
		{
			p.SetState(155)
			p.Statement()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VanLangParserELSE {
		{
			p.SetState(162)
			p.Match(VanLangParserELSE)
		}
		{
			p.SetState(163)
			p.Statement()
		}

	}
	{
		p.SetState(166)
		p.Match(VanLangParserSEMICOLON)
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(VanLangParserPRINTLN, 0)
}

func (s *ExpressionStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserLPAREN, 0)
}

func (s *ExpressionStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(VanLangParserSTRING, 0)
}

func (s *ExpressionStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserRPAREN, 0)
}

func (s *ExpressionStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *ExpressionStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserCOMMA)
}

func (s *ExpressionStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserCOMMA, i)
}

func (s *ExpressionStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(VanLangParserID, 0)
}

func (s *ExpressionStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(VanLangParserEQUAL, 0)
}

func (s *ExpressionStatementContext) READLN() antlr.TerminalNode {
	return s.GetToken(VanLangParserREADLN, 0)
}

func (s *ExpressionStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(VanLangParserRETURN, 0)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VanLangParserRULE_expressionStatement)
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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VanLangParserPRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(VanLangParserPRINTLN)
		}
		{
			p.SetState(169)
			p.Match(VanLangParserLPAREN)
		}
		{
			p.SetState(170)
			p.Match(VanLangParserSTRING)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == VanLangParserCOMMA {
			{
				p.SetState(171)
				p.Match(VanLangParserCOMMA)
			}
			{
				p.SetState(172)
				p.Expression()
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(178)
			p.Match(VanLangParserRPAREN)
		}
		{
			p.SetState(179)
			p.Match(VanLangParserSEMICOLON)
		}

	case VanLangParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Match(VanLangParserID)
		}
		{
			p.SetState(181)
			p.Match(VanLangParserEQUAL)
		}
		{
			p.SetState(182)
			p.Match(VanLangParserREADLN)
		}
		{
			p.SetState(183)
			p.Match(VanLangParserLPAREN)
		}
		{
			p.SetState(184)
			p.Match(VanLangParserSTRING)
		}
		{
			p.SetState(185)
			p.Match(VanLangParserRPAREN)
		}
		{
			p.SetState(186)
			p.Match(VanLangParserSEMICOLON)
		}

	case VanLangParserRETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(VanLangParserRETURN)
		}
		{
			p.SetState(188)
			p.Expression()
		}
		{
			p.SetState(189)
			p.Match(VanLangParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = VanLangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VanLangParserSEMICOLON, 0)
}

func (s *ExpressionContext) AllAND_OP() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserAND_OP)
}

func (s *ExpressionContext) AND_OP(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserAND_OP, i)
}

func (s *ExpressionContext) AllOR_OP() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserOR_OP)
}

func (s *ExpressionContext) OR_OP(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserOR_OP, i)
}

func (s *ExpressionContext) AllXOR_OP() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserXOR_OP)
}

func (s *ExpressionContext) XOR_OP(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserXOR_OP, i)
}

func (s *ExpressionContext) AllIMP_OP() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserIMP_OP)
}

func (s *ExpressionContext) IMP_OP(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserIMP_OP, i)
}

func (s *ExpressionContext) AllEQV_OP() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserEQV_OP)
}

func (s *ExpressionContext) EQV_OP(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserEQV_OP, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VanLangParserRULE_expression)
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
		p.SetState(193)
		p.Term()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(VanLangParserAND_OP-41))|(1<<(VanLangParserOR_OP-41))|(1<<(VanLangParserXOR_OP-41))|(1<<(VanLangParserEQV_OP-41))|(1<<(VanLangParserIMP_OP-41)))) != 0 {
		p.SetState(194)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(VanLangParserAND_OP-41))|(1<<(VanLangParserOR_OP-41))|(1<<(VanLangParserXOR_OP-41))|(1<<(VanLangParserEQV_OP-41))|(1<<(VanLangParserIMP_OP-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(195)
			p.Term()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(VanLangParserSEMICOLON)
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
	p.RuleIndex = VanLangParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMULTIPLY() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserMULTIPLY)
}

func (s *TermContext) MULTIPLY(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserMULTIPLY, i)
}

func (s *TermContext) AllDIVIDE() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserDIVIDE)
}

func (s *TermContext) DIVIDE(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserDIVIDE, i)
}

func (s *TermContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserDIV)
}

func (s *TermContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserDIV, i)
}

func (s *TermContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserMOD)
}

func (s *TermContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserMOD, i)
}

func (s *TermContext) AllEQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserEQUAL)
}

func (s *TermContext) EQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserEQUAL, i)
}

func (s *TermContext) AllNOTEQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserNOTEQUAL)
}

func (s *TermContext) NOTEQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserNOTEQUAL, i)
}

func (s *TermContext) AllLESS() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserLESS)
}

func (s *TermContext) LESS(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserLESS, i)
}

func (s *TermContext) AllGREATER() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserGREATER)
}

func (s *TermContext) GREATER(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserGREATER, i)
}

func (s *TermContext) AllLESSEQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserLESSEQUAL)
}

func (s *TermContext) LESSEQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserLESSEQUAL, i)
}

func (s *TermContext) AllGREATEREQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserGREATEREQUAL)
}

func (s *TermContext) GREATEREQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserGREATEREQUAL, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VanLangParserRULE_term)
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
		p.SetState(203)
		p.Factor()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(VanLangParserDIV-22))|(1<<(VanLangParserMOD-22))|(1<<(VanLangParserMULTIPLY-22))|(1<<(VanLangParserDIVIDE-22))|(1<<(VanLangParserEQUAL-22))|(1<<(VanLangParserNOTEQUAL-22))|(1<<(VanLangParserLESS-22))|(1<<(VanLangParserGREATER-22))|(1<<(VanLangParserLESSEQUAL-22))|(1<<(VanLangParserGREATEREQUAL-22)))) != 0 {
		p.SetState(208)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case VanLangParserDIV, VanLangParserMOD, VanLangParserMULTIPLY, VanLangParserDIVIDE:
			p.SetState(204)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(VanLangParserDIV-22))|(1<<(VanLangParserMOD-22))|(1<<(VanLangParserMULTIPLY-22))|(1<<(VanLangParserDIVIDE-22)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(205)
				p.Factor()
			}

		case VanLangParserEQUAL, VanLangParserNOTEQUAL, VanLangParserLESS, VanLangParserGREATER, VanLangParserLESSEQUAL, VanLangParserGREATEREQUAL:
			p.SetState(206)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(VanLangParserEQUAL-35))|(1<<(VanLangParserNOTEQUAL-35))|(1<<(VanLangParserLESS-35))|(1<<(VanLangParserGREATER-35))|(1<<(VanLangParserLESSEQUAL-35))|(1<<(VanLangParserGREATEREQUAL-35)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(207)
				p.Factor()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(212)
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
	p.RuleIndex = VanLangParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *FactorContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *FactorContext) AllMULTIPLY() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserMULTIPLY)
}

func (s *FactorContext) MULTIPLY(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserMULTIPLY, i)
}

func (s *FactorContext) AllDIVIDE() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserDIVIDE)
}

func (s *FactorContext) DIVIDE(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserDIVIDE, i)
}

func (s *FactorContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserDIV)
}

func (s *FactorContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserDIV, i)
}

func (s *FactorContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserMOD)
}

func (s *FactorContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserMOD, i)
}

func (s *FactorContext) AllEQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserEQUAL)
}

func (s *FactorContext) EQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserEQUAL, i)
}

func (s *FactorContext) AllNOTEQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserNOTEQUAL)
}

func (s *FactorContext) NOTEQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserNOTEQUAL, i)
}

func (s *FactorContext) AllLESS() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserLESS)
}

func (s *FactorContext) LESS(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserLESS, i)
}

func (s *FactorContext) AllGREATER() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserGREATER)
}

func (s *FactorContext) GREATER(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserGREATER, i)
}

func (s *FactorContext) AllLESSEQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserLESSEQUAL)
}

func (s *FactorContext) LESSEQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserLESSEQUAL, i)
}

func (s *FactorContext) AllGREATEREQUAL() []antlr.TerminalNode {
	return s.GetTokens(VanLangParserGREATEREQUAL)
}

func (s *FactorContext) GREATEREQUAL(i int) antlr.TerminalNode {
	return s.GetToken(VanLangParserGREATEREQUAL, i)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VanLangParserRULE_factor)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Atom()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(218)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case VanLangParserDIV, VanLangParserMOD, VanLangParserMULTIPLY, VanLangParserDIVIDE:
				p.SetState(214)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(VanLangParserDIV-22))|(1<<(VanLangParserMOD-22))|(1<<(VanLangParserMULTIPLY-22))|(1<<(VanLangParserDIVIDE-22)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(215)
					p.Atom()
				}

			case VanLangParserEQUAL, VanLangParserNOTEQUAL, VanLangParserLESS, VanLangParserGREATER, VanLangParserLESSEQUAL, VanLangParserGREATEREQUAL:
				p.SetState(216)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(VanLangParserEQUAL-35))|(1<<(VanLangParserNOTEQUAL-35))|(1<<(VanLangParserLESS-35))|(1<<(VanLangParserGREATER-35))|(1<<(VanLangParserLESSEQUAL-35))|(1<<(VanLangParserGREATEREQUAL-35)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(217)
					p.Atom()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VanLangParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VanLangParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(VanLangParserID, 0)
}

func (s *AtomContext) INT() antlr.TerminalNode {
	return s.GetToken(VanLangParserINT, 0)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserLPAREN, 0)
}

func (s *AtomContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VanLangParserRPAREN, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VanLangListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VanLangVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VanLangParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VanLangParserRULE_atom)

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

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VanLangParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(VanLangParserID)
		}

	case VanLangParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(VanLangParserINT)
		}

	case VanLangParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Match(VanLangParserLPAREN)
		}
		{
			p.SetState(226)
			p.Expression()
		}
		{
			p.SetState(227)
			p.Match(VanLangParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
