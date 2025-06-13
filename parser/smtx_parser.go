// Code generated from /home/baeyun/smtc/parser/SMTX.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SMTX
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SMTXParser struct {
	*antlr.BaseParser
}

var SMTXParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func smtxParserInit() {
	staticData := &SMTXParserStaticData
	staticData.LiteralNames = []string{
		"", "'assert'", "'check-sat'", "'check-sat-assuming'", "'declare-const'",
		"'declare-datatype'", "'declare-datatypes'", "'declare-fun'", "'declare-sort'",
		"'declare-sort-parameter'", "'define-const'", "'define-fun'", "'define-fun-rec'",
		"'define-funs-rec'", "'define-sort'", "'echo'", "'exit'", "'get-assertions'",
		"'get-assignment'", "'get-info'", "'get-model'", "'get-option'", "'get-proof'",
		"'get-unsat-assumptions'", "'get-unsat-core'", "'get-value'", "'pop'",
		"'push'", "'reset'", "'reset-assertions'", "'set-info'", "'set-logic'",
		"'set-option'", "'par'", "'let'", "'lambda'", "'forall'", "'exists'",
		"'match'", "'!'", "'as'", "'_'", "':produce-assertions'", "", "", "",
		"", "", "", "", "'('", "')'", "", "", "", "'not'", "'Bool'", "'continued-execution'",
		"'error'", "'false'", "'immediate-exit'", "'incomplete'", "'logic'",
		"'memout'", "'sat'", "'success'", "'theory'", "'true'", "'unknown'",
		"'unsupported'", "'unsat'", "':all-statistics'", "':assertion-stack-levels'",
		"':authors'", "':category'", "':chainable'", "':definition'", "':diagnostic-output-channel'",
		"':error-behavior'", "':extensions'", "':funs'", "':funs-description'",
		"':global-declarations'", "':interactive-mode'", "':language'", "':left-assoc'",
		"':license'", "':name'", "':named'", "':notes'", "':pattern'", "':print-success'",
		"':produce-assignments'", "':produce-models'", "':produce-proofs'",
		"':produce-unsat-assumptions'", "':produce-unsat-cores'", "':random-seed'",
		"':reason-unknown'", "':regular-output-channel'", "':reproducible-resource-limit'",
		"':right-assoc'", "':smt-lib-version'", "':sorts'", "':sorts-description'",
		"':source'", "':status'", "':theories'", "':values'", "':verbosity'",
		"':version'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "WS", "COMMENT", "NUMERAL", "DECIMAL",
		"HEXADECIMAL", "BINARY", "STRING", "PAROPEN", "PARCLOSE", "SIMPLE_SYMBOL",
		"QUOTED_SYMBOL", "KEYWORD", "PS_NOT", "PS_BOOL", "PS_CONTINUEDEXECUTION",
		"PS_ERROR", "PS_FALSE", "PS_IMMEDIATEEXIT", "PS_INCOMPLETE", "PS_LOGIC",
		"PS_MEMOUT", "PS_SAT", "PS_SUCCESS", "PS_THEORY", "PS_TRUE", "PS_UNKNOWN",
		"PS_UNSUPPORTED", "PS_UNSAT", "PK_ALL_STATISTICS", "PK_ASSERTION_STACK_LEVELS",
		"PK_AUTHORS", "PK_CATEGORY", "PK_CHAINABLE", "PK_DEFINITION", "PK_DIAGNOSTIC_OUTPUT_CHANNEL",
		"PK_ERROR_BEHAVIOR", "PK_EXTENSIONS", "PK_FUNS", "PK_FUNS_DESCRIPTION",
		"PK_GLOBAL_DECLARATIONS", "PK_INTERACTIVE_MODE", "PK_LANGUAGE", "PK_LEFT_ASSOC",
		"PK_LICENSE", "PK_NAME", "PK_NAMED", "PK_NOTES", "PK_PATTERN", "PK_PRINT_SUCCESS",
		"PK_PRODUCE_ASSIGNMENTS", "PK_PRODUCE_MODELS", "PK_PRODUCE_PROOFS",
		"PK_PRODUCE_UNSAT_ASSUMPTIONS", "PK_PRODUCE_UNSAT_CORES", "PK_RANDOM_SEED",
		"PK_REASON_UNKNOWN", "PK_REGULAR_OUTPUT_CHANNEL", "PK_REPRODUCIBLE_RESOURCE_LIMIT",
		"PK_RIGHT_ASSOC", "PK_SMT_LIB_VERSION", "PK_SORTS", "PK_SORTS_DESCRIPTION",
		"PK_SOURCE", "PK_STATUS", "PK_THEORIES", "PK_VALUES", "PK_VERBOSITY",
		"PK_VERSION",
	}
	staticData.RuleNames = []string{
		"script", "command", "symbol", "sort", "sort_dec", "selector_dec", "constructor_dec",
		"datatype_dec", "function_dec", "function_def", "term", "spec_constant",
		"qual_identifier", "var_binding", "sorted_var", "pattern", "symbol_",
		"match_case", "identifier", "index", "attribute_value", "attribute",
		"info_flag", "option", "b_value", "s_expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 110, 552, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 71, 8, 1, 10, 1, 12, 1, 74,
		9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 94, 8, 1, 11, 1, 12, 1,
		95, 1, 1, 1, 1, 1, 1, 4, 1, 101, 8, 1, 11, 1, 12, 1, 102, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 113, 8, 1, 10, 1, 12, 1, 116, 9,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 4, 1, 154, 8, 1, 11, 1, 12, 1, 155, 1, 1, 1, 1, 1, 1, 4, 1, 161, 8,
		1, 11, 1, 12, 1, 162, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 173, 8, 1, 10, 1, 12, 1, 176, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 4, 1, 220, 8, 1, 11, 1, 12, 1, 221, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 256, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 4,
		3, 264, 8, 3, 11, 3, 12, 3, 265, 1, 3, 1, 3, 3, 3, 270, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5,
		6, 285, 8, 6, 10, 6, 12, 6, 288, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 4, 7, 294,
		8, 7, 11, 7, 12, 7, 295, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 304,
		8, 7, 11, 7, 12, 7, 305, 1, 7, 1, 7, 1, 7, 4, 7, 311, 8, 7, 11, 7, 12,
		7, 312, 1, 7, 1, 7, 1, 7, 3, 7, 318, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8,
		324, 8, 8, 10, 8, 12, 8, 327, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 5, 9, 336, 8, 9, 10, 9, 12, 9, 339, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 350, 8, 10, 11, 10, 12, 10, 351,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 360, 8, 10, 11, 10, 12,
		10, 361, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10,
		372, 8, 10, 11, 10, 12, 10, 373, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 4, 10, 384, 8, 10, 11, 10, 12, 10, 385, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 396, 8, 10, 11, 10, 12,
		10, 397, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		4, 10, 409, 8, 10, 11, 10, 12, 10, 410, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 4, 10, 420, 8, 10, 11, 10, 12, 10, 421, 1, 10, 1, 10,
		3, 10, 426, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 437, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 453, 8,
		15, 11, 15, 12, 15, 454, 1, 15, 1, 15, 3, 15, 459, 8, 15, 1, 16, 1, 16,
		3, 16, 463, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 4, 18, 475, 8, 18, 11, 18, 12, 18, 476, 1, 18, 1, 18,
		3, 18, 481, 8, 18, 1, 19, 1, 19, 3, 19, 485, 8, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 491, 8, 20, 10, 20, 12, 20, 494, 9, 20, 1, 20, 3, 20, 497,
		8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 502, 8, 21, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 535, 8, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 544, 8, 25, 10, 25, 12,
		25, 547, 9, 25, 1, 25, 3, 25, 550, 8, 25, 1, 25, 0, 0, 26, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 0, 4, 1, 0, 52, 53, 1, 0, 45, 49, 6, 0, 54, 54, 71, 73,
		78, 78, 87, 87, 98, 98, 110, 110, 2, 0, 59, 59, 67, 67, 618, 0, 55, 1,
		0, 0, 0, 2, 255, 1, 0, 0, 0, 4, 257, 1, 0, 0, 0, 6, 269, 1, 0, 0, 0, 8,
		271, 1, 0, 0, 0, 10, 276, 1, 0, 0, 0, 12, 281, 1, 0, 0, 0, 14, 317, 1,
		0, 0, 0, 16, 319, 1, 0, 0, 0, 18, 332, 1, 0, 0, 0, 20, 425, 1, 0, 0, 0,
		22, 427, 1, 0, 0, 0, 24, 436, 1, 0, 0, 0, 26, 438, 1, 0, 0, 0, 28, 443,
		1, 0, 0, 0, 30, 458, 1, 0, 0, 0, 32, 462, 1, 0, 0, 0, 34, 464, 1, 0, 0,
		0, 36, 480, 1, 0, 0, 0, 38, 484, 1, 0, 0, 0, 40, 496, 1, 0, 0, 0, 42, 501,
		1, 0, 0, 0, 44, 503, 1, 0, 0, 0, 46, 534, 1, 0, 0, 0, 48, 536, 1, 0, 0,
		0, 50, 549, 1, 0, 0, 0, 52, 54, 3, 2, 1, 0, 53, 52, 1, 0, 0, 0, 54, 57,
		1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 1, 1, 0, 0, 0,
		57, 55, 1, 0, 0, 0, 58, 59, 5, 50, 0, 0, 59, 60, 5, 1, 0, 0, 60, 61, 3,
		20, 10, 0, 61, 62, 5, 51, 0, 0, 62, 256, 1, 0, 0, 0, 63, 64, 5, 50, 0,
		0, 64, 65, 5, 2, 0, 0, 65, 256, 5, 51, 0, 0, 66, 67, 5, 50, 0, 0, 67, 68,
		5, 3, 0, 0, 68, 72, 5, 50, 0, 0, 69, 71, 3, 20, 10, 0, 70, 69, 1, 0, 0,
		0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75,
		1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 51, 0, 0, 76, 256, 5, 51, 0,
		0, 77, 78, 5, 50, 0, 0, 78, 79, 5, 4, 0, 0, 79, 80, 3, 4, 2, 0, 80, 81,
		3, 6, 3, 0, 81, 82, 5, 51, 0, 0, 82, 256, 1, 0, 0, 0, 83, 84, 5, 50, 0,
		0, 84, 85, 5, 5, 0, 0, 85, 86, 3, 4, 2, 0, 86, 87, 3, 14, 7, 0, 87, 88,
		5, 51, 0, 0, 88, 256, 1, 0, 0, 0, 89, 90, 5, 50, 0, 0, 90, 91, 5, 6, 0,
		0, 91, 93, 5, 50, 0, 0, 92, 94, 3, 8, 4, 0, 93, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 98, 5, 51, 0, 0, 98, 100, 5, 50, 0, 0, 99, 101, 3, 14, 7, 0, 100, 99,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0,
		0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 5, 51, 0, 0, 105, 106, 5, 51, 0,
		0, 106, 256, 1, 0, 0, 0, 107, 108, 5, 50, 0, 0, 108, 109, 5, 7, 0, 0, 109,
		110, 3, 4, 2, 0, 110, 114, 5, 50, 0, 0, 111, 113, 3, 6, 3, 0, 112, 111,
		1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 51, 0, 0,
		118, 119, 3, 6, 3, 0, 119, 120, 5, 51, 0, 0, 120, 256, 1, 0, 0, 0, 121,
		122, 5, 50, 0, 0, 122, 123, 5, 8, 0, 0, 123, 124, 3, 4, 2, 0, 124, 125,
		5, 45, 0, 0, 125, 126, 5, 51, 0, 0, 126, 256, 1, 0, 0, 0, 127, 128, 5,
		50, 0, 0, 128, 129, 5, 9, 0, 0, 129, 130, 3, 4, 2, 0, 130, 131, 5, 51,
		0, 0, 131, 256, 1, 0, 0, 0, 132, 133, 5, 50, 0, 0, 133, 134, 5, 10, 0,
		0, 134, 135, 3, 4, 2, 0, 135, 136, 3, 6, 3, 0, 136, 137, 3, 20, 10, 0,
		137, 138, 5, 51, 0, 0, 138, 256, 1, 0, 0, 0, 139, 140, 5, 50, 0, 0, 140,
		141, 5, 11, 0, 0, 141, 142, 3, 18, 9, 0, 142, 143, 5, 51, 0, 0, 143, 256,
		1, 0, 0, 0, 144, 145, 5, 50, 0, 0, 145, 146, 5, 12, 0, 0, 146, 147, 3,
		18, 9, 0, 147, 148, 5, 51, 0, 0, 148, 256, 1, 0, 0, 0, 149, 150, 5, 50,
		0, 0, 150, 151, 5, 13, 0, 0, 151, 153, 5, 50, 0, 0, 152, 154, 3, 16, 8,
		0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155,
		156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 5, 51, 0, 0, 158, 160,
		5, 50, 0, 0, 159, 161, 3, 20, 10, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1,
		0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0,
		0, 164, 165, 5, 51, 0, 0, 165, 166, 5, 51, 0, 0, 166, 256, 1, 0, 0, 0,
		167, 168, 5, 50, 0, 0, 168, 169, 5, 14, 0, 0, 169, 170, 3, 4, 2, 0, 170,
		174, 5, 50, 0, 0, 171, 173, 3, 4, 2, 0, 172, 171, 1, 0, 0, 0, 173, 176,
		1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 177, 178, 5, 51, 0, 0, 178, 179, 3, 6, 3, 0,
		179, 180, 5, 51, 0, 0, 180, 256, 1, 0, 0, 0, 181, 182, 5, 50, 0, 0, 182,
		183, 5, 15, 0, 0, 183, 184, 5, 49, 0, 0, 184, 256, 5, 51, 0, 0, 185, 186,
		5, 50, 0, 0, 186, 187, 5, 16, 0, 0, 187, 256, 5, 51, 0, 0, 188, 189, 5,
		50, 0, 0, 189, 190, 5, 17, 0, 0, 190, 256, 5, 51, 0, 0, 191, 192, 5, 50,
		0, 0, 192, 193, 5, 18, 0, 0, 193, 256, 5, 51, 0, 0, 194, 195, 5, 50, 0,
		0, 195, 196, 5, 19, 0, 0, 196, 197, 3, 44, 22, 0, 197, 198, 5, 51, 0, 0,
		198, 256, 1, 0, 0, 0, 199, 200, 5, 50, 0, 0, 200, 201, 5, 20, 0, 0, 201,
		256, 5, 51, 0, 0, 202, 203, 5, 50, 0, 0, 203, 204, 5, 21, 0, 0, 204, 205,
		5, 54, 0, 0, 205, 256, 5, 51, 0, 0, 206, 207, 5, 50, 0, 0, 207, 208, 5,
		22, 0, 0, 208, 256, 5, 51, 0, 0, 209, 210, 5, 50, 0, 0, 210, 211, 5, 23,
		0, 0, 211, 256, 5, 51, 0, 0, 212, 213, 5, 50, 0, 0, 213, 214, 5, 24, 0,
		0, 214, 256, 5, 51, 0, 0, 215, 216, 5, 50, 0, 0, 216, 217, 5, 25, 0, 0,
		217, 219, 5, 50, 0, 0, 218, 220, 3, 20, 10, 0, 219, 218, 1, 0, 0, 0, 220,
		221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 224, 5, 51, 0, 0, 224, 225, 5, 51, 0, 0, 225, 256, 1,
		0, 0, 0, 226, 227, 5, 50, 0, 0, 227, 228, 5, 26, 0, 0, 228, 229, 5, 45,
		0, 0, 229, 256, 5, 51, 0, 0, 230, 231, 5, 50, 0, 0, 231, 232, 5, 27, 0,
		0, 232, 233, 5, 45, 0, 0, 233, 256, 5, 51, 0, 0, 234, 235, 5, 50, 0, 0,
		235, 236, 5, 28, 0, 0, 236, 256, 5, 51, 0, 0, 237, 238, 5, 50, 0, 0, 238,
		239, 5, 29, 0, 0, 239, 256, 5, 51, 0, 0, 240, 241, 5, 50, 0, 0, 241, 242,
		5, 30, 0, 0, 242, 243, 3, 42, 21, 0, 243, 244, 5, 51, 0, 0, 244, 256, 1,
		0, 0, 0, 245, 246, 5, 50, 0, 0, 246, 247, 5, 31, 0, 0, 247, 248, 3, 4,
		2, 0, 248, 249, 5, 51, 0, 0, 249, 256, 1, 0, 0, 0, 250, 251, 5, 50, 0,
		0, 251, 252, 5, 32, 0, 0, 252, 253, 3, 46, 23, 0, 253, 254, 5, 51, 0, 0,
		254, 256, 1, 0, 0, 0, 255, 58, 1, 0, 0, 0, 255, 63, 1, 0, 0, 0, 255, 66,
		1, 0, 0, 0, 255, 77, 1, 0, 0, 0, 255, 83, 1, 0, 0, 0, 255, 89, 1, 0, 0,
		0, 255, 107, 1, 0, 0, 0, 255, 121, 1, 0, 0, 0, 255, 127, 1, 0, 0, 0, 255,
		132, 1, 0, 0, 0, 255, 139, 1, 0, 0, 0, 255, 144, 1, 0, 0, 0, 255, 149,
		1, 0, 0, 0, 255, 167, 1, 0, 0, 0, 255, 181, 1, 0, 0, 0, 255, 185, 1, 0,
		0, 0, 255, 188, 1, 0, 0, 0, 255, 191, 1, 0, 0, 0, 255, 194, 1, 0, 0, 0,
		255, 199, 1, 0, 0, 0, 255, 202, 1, 0, 0, 0, 255, 206, 1, 0, 0, 0, 255,
		209, 1, 0, 0, 0, 255, 212, 1, 0, 0, 0, 255, 215, 1, 0, 0, 0, 255, 226,
		1, 0, 0, 0, 255, 230, 1, 0, 0, 0, 255, 234, 1, 0, 0, 0, 255, 237, 1, 0,
		0, 0, 255, 240, 1, 0, 0, 0, 255, 245, 1, 0, 0, 0, 255, 250, 1, 0, 0, 0,
		256, 3, 1, 0, 0, 0, 257, 258, 7, 0, 0, 0, 258, 5, 1, 0, 0, 0, 259, 270,
		3, 36, 18, 0, 260, 261, 5, 50, 0, 0, 261, 263, 3, 36, 18, 0, 262, 264,
		3, 6, 3, 0, 263, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 263, 1, 0,
		0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 5, 51, 0, 0,
		268, 270, 1, 0, 0, 0, 269, 259, 1, 0, 0, 0, 269, 260, 1, 0, 0, 0, 270,
		7, 1, 0, 0, 0, 271, 272, 5, 50, 0, 0, 272, 273, 3, 4, 2, 0, 273, 274, 5,
		45, 0, 0, 274, 275, 5, 51, 0, 0, 275, 9, 1, 0, 0, 0, 276, 277, 5, 50, 0,
		0, 277, 278, 3, 4, 2, 0, 278, 279, 3, 6, 3, 0, 279, 280, 5, 51, 0, 0, 280,
		11, 1, 0, 0, 0, 281, 282, 5, 50, 0, 0, 282, 286, 3, 4, 2, 0, 283, 285,
		3, 10, 5, 0, 284, 283, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0,
		0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0,
		289, 290, 5, 51, 0, 0, 290, 13, 1, 0, 0, 0, 291, 293, 5, 50, 0, 0, 292,
		294, 3, 12, 6, 0, 293, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 293,
		1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298, 5, 51,
		0, 0, 298, 318, 1, 0, 0, 0, 299, 300, 5, 50, 0, 0, 300, 301, 5, 33, 0,
		0, 301, 303, 5, 50, 0, 0, 302, 304, 3, 4, 2, 0, 303, 302, 1, 0, 0, 0, 304,
		305, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307,
		1, 0, 0, 0, 307, 308, 5, 51, 0, 0, 308, 310, 5, 50, 0, 0, 309, 311, 3,
		12, 6, 0, 310, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 310, 1, 0, 0,
		0, 312, 313, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 315, 5, 51, 0, 0, 315,
		316, 5, 51, 0, 0, 316, 318, 1, 0, 0, 0, 317, 291, 1, 0, 0, 0, 317, 299,
		1, 0, 0, 0, 318, 15, 1, 0, 0, 0, 319, 320, 5, 50, 0, 0, 320, 321, 3, 4,
		2, 0, 321, 325, 5, 50, 0, 0, 322, 324, 3, 28, 14, 0, 323, 322, 1, 0, 0,
		0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326,
		328, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 329, 5, 51, 0, 0, 329, 330,
		3, 6, 3, 0, 330, 331, 5, 51, 0, 0, 331, 17, 1, 0, 0, 0, 332, 333, 3, 4,
		2, 0, 333, 337, 5, 50, 0, 0, 334, 336, 3, 28, 14, 0, 335, 334, 1, 0, 0,
		0, 336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338,
		340, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 341, 5, 51, 0, 0, 341, 342,
		3, 6, 3, 0, 342, 343, 3, 20, 10, 0, 343, 19, 1, 0, 0, 0, 344, 426, 3, 22,
		11, 0, 345, 426, 3, 24, 12, 0, 346, 347, 5, 50, 0, 0, 347, 349, 3, 24,
		12, 0, 348, 350, 3, 20, 10, 0, 349, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0,
		0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353,
		354, 5, 51, 0, 0, 354, 426, 1, 0, 0, 0, 355, 356, 5, 50, 0, 0, 356, 357,
		5, 34, 0, 0, 357, 359, 5, 50, 0, 0, 358, 360, 3, 26, 13, 0, 359, 358, 1,
		0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 362, 1, 0, 0,
		0, 362, 363, 1, 0, 0, 0, 363, 364, 5, 51, 0, 0, 364, 365, 3, 20, 10, 0,
		365, 366, 5, 51, 0, 0, 366, 426, 1, 0, 0, 0, 367, 368, 5, 50, 0, 0, 368,
		369, 5, 35, 0, 0, 369, 371, 5, 50, 0, 0, 370, 372, 3, 28, 14, 0, 371, 370,
		1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 5, 51, 0, 0, 376, 377, 3, 20, 10,
		0, 377, 378, 5, 51, 0, 0, 378, 426, 1, 0, 0, 0, 379, 380, 5, 50, 0, 0,
		380, 381, 5, 36, 0, 0, 381, 383, 5, 50, 0, 0, 382, 384, 3, 28, 14, 0, 383,
		382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386,
		1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 5, 51, 0, 0, 388, 389, 3, 20,
		10, 0, 389, 390, 5, 51, 0, 0, 390, 426, 1, 0, 0, 0, 391, 392, 5, 50, 0,
		0, 392, 393, 5, 37, 0, 0, 393, 395, 5, 50, 0, 0, 394, 396, 3, 28, 14, 0,
		395, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397,
		398, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400, 5, 51, 0, 0, 400, 401,
		3, 20, 10, 0, 401, 402, 5, 51, 0, 0, 402, 426, 1, 0, 0, 0, 403, 404, 5,
		50, 0, 0, 404, 405, 5, 38, 0, 0, 405, 406, 3, 20, 10, 0, 406, 408, 5, 50,
		0, 0, 407, 409, 3, 34, 17, 0, 408, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0,
		0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412,
		413, 5, 51, 0, 0, 413, 414, 5, 51, 0, 0, 414, 426, 1, 0, 0, 0, 415, 416,
		5, 50, 0, 0, 416, 417, 5, 39, 0, 0, 417, 419, 3, 20, 10, 0, 418, 420, 3,
		42, 21, 0, 419, 418, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 419, 1, 0,
		0, 0, 421, 422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 5, 51, 0, 0,
		424, 426, 1, 0, 0, 0, 425, 344, 1, 0, 0, 0, 425, 345, 1, 0, 0, 0, 425,
		346, 1, 0, 0, 0, 425, 355, 1, 0, 0, 0, 425, 367, 1, 0, 0, 0, 425, 379,
		1, 0, 0, 0, 425, 391, 1, 0, 0, 0, 425, 403, 1, 0, 0, 0, 425, 415, 1, 0,
		0, 0, 426, 21, 1, 0, 0, 0, 427, 428, 7, 1, 0, 0, 428, 23, 1, 0, 0, 0, 429,
		437, 3, 36, 18, 0, 430, 431, 5, 50, 0, 0, 431, 432, 5, 40, 0, 0, 432, 433,
		3, 36, 18, 0, 433, 434, 3, 6, 3, 0, 434, 435, 5, 51, 0, 0, 435, 437, 1,
		0, 0, 0, 436, 429, 1, 0, 0, 0, 436, 430, 1, 0, 0, 0, 437, 25, 1, 0, 0,
		0, 438, 439, 5, 50, 0, 0, 439, 440, 3, 4, 2, 0, 440, 441, 3, 20, 10, 0,
		441, 442, 5, 51, 0, 0, 442, 27, 1, 0, 0, 0, 443, 444, 5, 50, 0, 0, 444,
		445, 3, 4, 2, 0, 445, 446, 3, 6, 3, 0, 446, 447, 5, 51, 0, 0, 447, 29,
		1, 0, 0, 0, 448, 459, 3, 32, 16, 0, 449, 450, 5, 50, 0, 0, 450, 452, 3,
		4, 2, 0, 451, 453, 3, 32, 16, 0, 452, 451, 1, 0, 0, 0, 453, 454, 1, 0,
		0, 0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0,
		456, 457, 5, 51, 0, 0, 457, 459, 1, 0, 0, 0, 458, 448, 1, 0, 0, 0, 458,
		449, 1, 0, 0, 0, 459, 31, 1, 0, 0, 0, 460, 463, 3, 4, 2, 0, 461, 463, 5,
		41, 0, 0, 462, 460, 1, 0, 0, 0, 462, 461, 1, 0, 0, 0, 463, 33, 1, 0, 0,
		0, 464, 465, 5, 50, 0, 0, 465, 466, 3, 30, 15, 0, 466, 467, 3, 20, 10,
		0, 467, 468, 5, 51, 0, 0, 468, 35, 1, 0, 0, 0, 469, 481, 3, 4, 2, 0, 470,
		471, 5, 50, 0, 0, 471, 472, 5, 41, 0, 0, 472, 474, 3, 4, 2, 0, 473, 475,
		3, 38, 19, 0, 474, 473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 474, 1,
		0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 479, 5, 51, 0,
		0, 479, 481, 1, 0, 0, 0, 480, 469, 1, 0, 0, 0, 480, 470, 1, 0, 0, 0, 481,
		37, 1, 0, 0, 0, 482, 485, 5, 45, 0, 0, 483, 485, 3, 4, 2, 0, 484, 482,
		1, 0, 0, 0, 484, 483, 1, 0, 0, 0, 485, 39, 1, 0, 0, 0, 486, 497, 3, 22,
		11, 0, 487, 497, 3, 4, 2, 0, 488, 492, 5, 50, 0, 0, 489, 491, 3, 50, 25,
		0, 490, 489, 1, 0, 0, 0, 491, 494, 1, 0, 0, 0, 492, 490, 1, 0, 0, 0, 492,
		493, 1, 0, 0, 0, 493, 495, 1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 495, 497,
		5, 51, 0, 0, 496, 486, 1, 0, 0, 0, 496, 487, 1, 0, 0, 0, 496, 488, 1, 0,
		0, 0, 497, 41, 1, 0, 0, 0, 498, 502, 5, 54, 0, 0, 499, 500, 5, 54, 0, 0,
		500, 502, 3, 40, 20, 0, 501, 498, 1, 0, 0, 0, 501, 499, 1, 0, 0, 0, 502,
		43, 1, 0, 0, 0, 503, 504, 7, 2, 0, 0, 504, 45, 1, 0, 0, 0, 505, 506, 5,
		77, 0, 0, 506, 535, 5, 49, 0, 0, 507, 508, 5, 82, 0, 0, 508, 535, 3, 48,
		24, 0, 509, 510, 5, 83, 0, 0, 510, 535, 3, 48, 24, 0, 511, 512, 5, 91,
		0, 0, 512, 535, 3, 48, 24, 0, 513, 514, 5, 42, 0, 0, 514, 535, 3, 48, 24,
		0, 515, 516, 5, 92, 0, 0, 516, 535, 3, 48, 24, 0, 517, 518, 5, 93, 0, 0,
		518, 535, 3, 48, 24, 0, 519, 520, 5, 94, 0, 0, 520, 535, 3, 48, 24, 0,
		521, 522, 5, 95, 0, 0, 522, 535, 3, 48, 24, 0, 523, 524, 5, 96, 0, 0, 524,
		535, 3, 48, 24, 0, 525, 526, 5, 97, 0, 0, 526, 535, 5, 45, 0, 0, 527, 528,
		5, 99, 0, 0, 528, 535, 5, 49, 0, 0, 529, 530, 5, 100, 0, 0, 530, 535, 5,
		45, 0, 0, 531, 532, 5, 109, 0, 0, 532, 535, 5, 45, 0, 0, 533, 535, 3, 42,
		21, 0, 534, 505, 1, 0, 0, 0, 534, 507, 1, 0, 0, 0, 534, 509, 1, 0, 0, 0,
		534, 511, 1, 0, 0, 0, 534, 513, 1, 0, 0, 0, 534, 515, 1, 0, 0, 0, 534,
		517, 1, 0, 0, 0, 534, 519, 1, 0, 0, 0, 534, 521, 1, 0, 0, 0, 534, 523,
		1, 0, 0, 0, 534, 525, 1, 0, 0, 0, 534, 527, 1, 0, 0, 0, 534, 529, 1, 0,
		0, 0, 534, 531, 1, 0, 0, 0, 534, 533, 1, 0, 0, 0, 535, 47, 1, 0, 0, 0,
		536, 537, 7, 3, 0, 0, 537, 49, 1, 0, 0, 0, 538, 550, 3, 22, 11, 0, 539,
		550, 3, 4, 2, 0, 540, 550, 5, 54, 0, 0, 541, 545, 5, 50, 0, 0, 542, 544,
		3, 50, 25, 0, 543, 542, 1, 0, 0, 0, 544, 547, 1, 0, 0, 0, 545, 543, 1,
		0, 0, 0, 545, 546, 1, 0, 0, 0, 546, 548, 1, 0, 0, 0, 547, 545, 1, 0, 0,
		0, 548, 550, 5, 51, 0, 0, 549, 538, 1, 0, 0, 0, 549, 539, 1, 0, 0, 0, 549,
		540, 1, 0, 0, 0, 549, 541, 1, 0, 0, 0, 550, 51, 1, 0, 0, 0, 40, 55, 72,
		95, 102, 114, 155, 162, 174, 221, 255, 265, 269, 286, 295, 305, 312, 317,
		325, 337, 351, 361, 373, 385, 397, 410, 421, 425, 436, 454, 458, 462, 476,
		480, 484, 492, 496, 501, 534, 545, 549,
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

// SMTXParserInit initializes any static state used to implement SMTXParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSMTXParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SMTXParserInit() {
	staticData := &SMTXParserStaticData
	staticData.once.Do(smtxParserInit)
}

// NewSMTXParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSMTXParser(input antlr.TokenStream) *SMTXParser {
	SMTXParserInit()
	this := new(SMTXParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SMTXParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SMTX.g4"

	return this
}

// SMTXParser tokens.
const (
	SMTXParserEOF                            = antlr.TokenEOF
	SMTXParserT__0                           = 1
	SMTXParserT__1                           = 2
	SMTXParserT__2                           = 3
	SMTXParserT__3                           = 4
	SMTXParserT__4                           = 5
	SMTXParserT__5                           = 6
	SMTXParserT__6                           = 7
	SMTXParserT__7                           = 8
	SMTXParserT__8                           = 9
	SMTXParserT__9                           = 10
	SMTXParserT__10                          = 11
	SMTXParserT__11                          = 12
	SMTXParserT__12                          = 13
	SMTXParserT__13                          = 14
	SMTXParserT__14                          = 15
	SMTXParserT__15                          = 16
	SMTXParserT__16                          = 17
	SMTXParserT__17                          = 18
	SMTXParserT__18                          = 19
	SMTXParserT__19                          = 20
	SMTXParserT__20                          = 21
	SMTXParserT__21                          = 22
	SMTXParserT__22                          = 23
	SMTXParserT__23                          = 24
	SMTXParserT__24                          = 25
	SMTXParserT__25                          = 26
	SMTXParserT__26                          = 27
	SMTXParserT__27                          = 28
	SMTXParserT__28                          = 29
	SMTXParserT__29                          = 30
	SMTXParserT__30                          = 31
	SMTXParserT__31                          = 32
	SMTXParserT__32                          = 33
	SMTXParserT__33                          = 34
	SMTXParserT__34                          = 35
	SMTXParserT__35                          = 36
	SMTXParserT__36                          = 37
	SMTXParserT__37                          = 38
	SMTXParserT__38                          = 39
	SMTXParserT__39                          = 40
	SMTXParserT__40                          = 41
	SMTXParserT__41                          = 42
	SMTXParserWS                             = 43
	SMTXParserCOMMENT                        = 44
	SMTXParserNUMERAL                        = 45
	SMTXParserDECIMAL                        = 46
	SMTXParserHEXADECIMAL                    = 47
	SMTXParserBINARY                         = 48
	SMTXParserSTRING                         = 49
	SMTXParserPAROPEN                        = 50
	SMTXParserPARCLOSE                       = 51
	SMTXParserSIMPLE_SYMBOL                  = 52
	SMTXParserQUOTED_SYMBOL                  = 53
	SMTXParserKEYWORD                        = 54
	SMTXParserPS_NOT                         = 55
	SMTXParserPS_BOOL                        = 56
	SMTXParserPS_CONTINUEDEXECUTION          = 57
	SMTXParserPS_ERROR                       = 58
	SMTXParserPS_FALSE                       = 59
	SMTXParserPS_IMMEDIATEEXIT               = 60
	SMTXParserPS_INCOMPLETE                  = 61
	SMTXParserPS_LOGIC                       = 62
	SMTXParserPS_MEMOUT                      = 63
	SMTXParserPS_SAT                         = 64
	SMTXParserPS_SUCCESS                     = 65
	SMTXParserPS_THEORY                      = 66
	SMTXParserPS_TRUE                        = 67
	SMTXParserPS_UNKNOWN                     = 68
	SMTXParserPS_UNSUPPORTED                 = 69
	SMTXParserPS_UNSAT                       = 70
	SMTXParserPK_ALL_STATISTICS              = 71
	SMTXParserPK_ASSERTION_STACK_LEVELS      = 72
	SMTXParserPK_AUTHORS                     = 73
	SMTXParserPK_CATEGORY                    = 74
	SMTXParserPK_CHAINABLE                   = 75
	SMTXParserPK_DEFINITION                  = 76
	SMTXParserPK_DIAGNOSTIC_OUTPUT_CHANNEL   = 77
	SMTXParserPK_ERROR_BEHAVIOR              = 78
	SMTXParserPK_EXTENSIONS                  = 79
	SMTXParserPK_FUNS                        = 80
	SMTXParserPK_FUNS_DESCRIPTION            = 81
	SMTXParserPK_GLOBAL_DECLARATIONS         = 82
	SMTXParserPK_INTERACTIVE_MODE            = 83
	SMTXParserPK_LANGUAGE                    = 84
	SMTXParserPK_LEFT_ASSOC                  = 85
	SMTXParserPK_LICENSE                     = 86
	SMTXParserPK_NAME                        = 87
	SMTXParserPK_NAMED                       = 88
	SMTXParserPK_NOTES                       = 89
	SMTXParserPK_PATTERN                     = 90
	SMTXParserPK_PRINT_SUCCESS               = 91
	SMTXParserPK_PRODUCE_ASSIGNMENTS         = 92
	SMTXParserPK_PRODUCE_MODELS              = 93
	SMTXParserPK_PRODUCE_PROOFS              = 94
	SMTXParserPK_PRODUCE_UNSAT_ASSUMPTIONS   = 95
	SMTXParserPK_PRODUCE_UNSAT_CORES         = 96
	SMTXParserPK_RANDOM_SEED                 = 97
	SMTXParserPK_REASON_UNKNOWN              = 98
	SMTXParserPK_REGULAR_OUTPUT_CHANNEL      = 99
	SMTXParserPK_REPRODUCIBLE_RESOURCE_LIMIT = 100
	SMTXParserPK_RIGHT_ASSOC                 = 101
	SMTXParserPK_SMT_LIB_VERSION             = 102
	SMTXParserPK_SORTS                       = 103
	SMTXParserPK_SORTS_DESCRIPTION           = 104
	SMTXParserPK_SOURCE                      = 105
	SMTXParserPK_STATUS                      = 106
	SMTXParserPK_THEORIES                    = 107
	SMTXParserPK_VALUES                      = 108
	SMTXParserPK_VERBOSITY                   = 109
	SMTXParserPK_VERSION                     = 110
)

// SMTXParser rules.
const (
	SMTXParserRULE_script          = 0
	SMTXParserRULE_command         = 1
	SMTXParserRULE_symbol          = 2
	SMTXParserRULE_sort            = 3
	SMTXParserRULE_sort_dec        = 4
	SMTXParserRULE_selector_dec    = 5
	SMTXParserRULE_constructor_dec = 6
	SMTXParserRULE_datatype_dec    = 7
	SMTXParserRULE_function_dec    = 8
	SMTXParserRULE_function_def    = 9
	SMTXParserRULE_term            = 10
	SMTXParserRULE_spec_constant   = 11
	SMTXParserRULE_qual_identifier = 12
	SMTXParserRULE_var_binding     = 13
	SMTXParserRULE_sorted_var      = 14
	SMTXParserRULE_pattern         = 15
	SMTXParserRULE_symbol_         = 16
	SMTXParserRULE_match_case      = 17
	SMTXParserRULE_identifier      = 18
	SMTXParserRULE_index           = 19
	SMTXParserRULE_attribute_value = 20
	SMTXParserRULE_attribute       = 21
	SMTXParserRULE_info_flag       = 22
	SMTXParserRULE_option          = 23
	SMTXParserRULE_b_value         = 24
	SMTXParserRULE_s_expr          = 25
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommand() []ICommandContext
	Command(i int) ICommandContext

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_script
	return p
}

func InitEmptyScriptContext(p *ScriptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_script
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
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

	return t.(ICommandContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *SMTXParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SMTXParserRULE_script)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SMTXParserPAROPEN {
		{
			p.SetState(52)
			p.Command()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPAROPEN() []antlr.TerminalNode
	PAROPEN(i int) antlr.TerminalNode
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllPARCLOSE() []antlr.TerminalNode
	PARCLOSE(i int) antlr.TerminalNode
	AllSymbol() []ISymbolContext
	Symbol(i int) ISymbolContext
	AllSort() []ISortContext
	Sort(i int) ISortContext
	AllDatatype_dec() []IDatatype_decContext
	Datatype_dec(i int) IDatatype_decContext
	AllSort_dec() []ISort_decContext
	Sort_dec(i int) ISort_decContext
	NUMERAL() antlr.TerminalNode
	Function_def() IFunction_defContext
	AllFunction_dec() []IFunction_decContext
	Function_dec(i int) IFunction_decContext
	STRING() antlr.TerminalNode
	Info_flag() IInfo_flagContext
	KEYWORD() antlr.TerminalNode
	Attribute() IAttributeContext
	Option() IOptionContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) AllPAROPEN() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPAROPEN)
}

func (s *CommandContext) PAROPEN(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, i)
}

func (s *CommandContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *CommandContext) AllPARCLOSE() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPARCLOSE)
}

func (s *CommandContext) PARCLOSE(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, i)
}

func (s *CommandContext) AllSymbol() []ISymbolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISymbolContext); ok {
			len++
		}
	}

	tst := make([]ISymbolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISymbolContext); ok {
			tst[i] = t.(ISymbolContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Symbol(i int) ISymbolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
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

	return t.(ISymbolContext)
}

func (s *CommandContext) AllSort() []ISortContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortContext); ok {
			len++
		}
	}

	tst := make([]ISortContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortContext); ok {
			tst[i] = t.(ISortContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Sort(i int) ISortContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
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

	return t.(ISortContext)
}

func (s *CommandContext) AllDatatype_dec() []IDatatype_decContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatatype_decContext); ok {
			len++
		}
	}

	tst := make([]IDatatype_decContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatatype_decContext); ok {
			tst[i] = t.(IDatatype_decContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Datatype_dec(i int) IDatatype_decContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatatype_decContext); ok {
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

	return t.(IDatatype_decContext)
}

func (s *CommandContext) AllSort_dec() []ISort_decContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISort_decContext); ok {
			len++
		}
	}

	tst := make([]ISort_decContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISort_decContext); ok {
			tst[i] = t.(ISort_decContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Sort_dec(i int) ISort_decContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISort_decContext); ok {
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

	return t.(ISort_decContext)
}

func (s *CommandContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserNUMERAL, 0)
}

func (s *CommandContext) Function_def() IFunction_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_defContext)
}

func (s *CommandContext) AllFunction_dec() []IFunction_decContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_decContext); ok {
			len++
		}
	}

	tst := make([]IFunction_decContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_decContext); ok {
			tst[i] = t.(IFunction_decContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Function_dec(i int) IFunction_decContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_decContext); ok {
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

	return t.(IFunction_decContext)
}

func (s *CommandContext) STRING() antlr.TerminalNode {
	return s.GetToken(SMTXParserSTRING, 0)
}

func (s *CommandContext) Info_flag() IInfo_flagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfo_flagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfo_flagContext)
}

func (s *CommandContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SMTXParserKEYWORD, 0)
}

func (s *CommandContext) Attribute() IAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *CommandContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *SMTXParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SMTXParserRULE_command)
	var _la int

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(SMTXParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Term()
		}
		{
			p.SetState(61)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(SMTXParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(SMTXParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15727414323707904) != 0 {
			{
				p.SetState(69)
				p.Term()
			}

			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(75)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(77)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(SMTXParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Symbol()
		}
		{
			p.SetState(80)
			p.Sort()
		}
		{
			p.SetState(81)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(SMTXParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Symbol()
		}
		{
			p.SetState(86)
			p.Datatype_dec()
		}
		{
			p.SetState(87)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(89)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(SMTXParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(92)
				p.Sort_dec()
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(99)
				p.Datatype_dec()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(104)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(107)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(SMTXParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Symbol()
		}
		{
			p.SetState(110)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14636698788954112) != 0 {
			{
				p.SetState(111)
				p.Sort()
			}

			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(117)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Sort()
		}
		{
			p.SetState(119)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(121)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(SMTXParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Symbol()
		}
		{
			p.SetState(124)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(127)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(SMTXParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Symbol()
		}
		{
			p.SetState(130)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(132)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(SMTXParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Symbol()
		}
		{
			p.SetState(135)
			p.Sort()
		}
		{
			p.SetState(136)
			p.Term()
		}
		{
			p.SetState(137)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(139)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(SMTXParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Function_def()
		}
		{
			p.SetState(142)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(144)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(SMTXParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Function_def()
		}
		{
			p.SetState(147)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(149)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(SMTXParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(152)
				p.Function_dec()
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(157)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15727414323707904) != 0) {
			{
				p.SetState(159)
				p.Term()
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(164)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(167)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(SMTXParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Symbol()
		}
		{
			p.SetState(170)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SMTXParserSIMPLE_SYMBOL || _la == SMTXParserQUOTED_SYMBOL {
			{
				p.SetState(171)
				p.Symbol()
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(177)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Sort()
		}
		{
			p.SetState(179)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(181)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(SMTXParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(SMTXParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(185)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(SMTXParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(188)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(SMTXParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(191)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(SMTXParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(194)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(SMTXParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Info_flag()
		}
		{
			p.SetState(197)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(199)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(SMTXParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(202)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(SMTXParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(SMTXParserKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(206)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(SMTXParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(209)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(SMTXParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(212)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(SMTXParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(215)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(SMTXParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15727414323707904) != 0) {
			{
				p.SetState(218)
				p.Term()
			}

			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(223)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(226)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(SMTXParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(230)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(SMTXParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(234)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Match(SMTXParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(237)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(SMTXParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(240)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(SMTXParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Attribute()
		}
		{
			p.SetState(243)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(245)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(SMTXParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Symbol()
		}
		{
			p.SetState(248)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 32:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(250)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.Match(SMTXParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Option()
		}
		{
			p.SetState(253)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIMPLE_SYMBOL() antlr.TerminalNode
	QUOTED_SYMBOL() antlr.TerminalNode

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_symbol
	return p
}

func InitEmptySymbolContext(p *SymbolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_symbol
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) SIMPLE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(SMTXParserSIMPLE_SYMBOL, 0)
}

func (s *SymbolContext) QUOTED_SYMBOL() antlr.TerminalNode {
	return s.GetToken(SMTXParserQUOTED_SYMBOL, 0)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *SMTXParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SMTXParserRULE_symbol)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SMTXParserSIMPLE_SYMBOL || _la == SMTXParserQUOTED_SYMBOL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISortContext is an interface to support dynamic dispatch.
type ISortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	PAROPEN() antlr.TerminalNode
	PARCLOSE() antlr.TerminalNode
	AllSort() []ISortContext
	Sort(i int) ISortContext

	// IsSortContext differentiates from other interfaces.
	IsSortContext()
}

type SortContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortContext() *SortContext {
	var p = new(SortContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_sort
	return p
}

func InitEmptySortContext(p *SortContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_sort
}

func (*SortContext) IsSortContext() {}

func NewSortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortContext {
	var p = new(SortContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_sort

	return p
}

func (s *SortContext) GetParser() antlr.Parser { return s.parser }

func (s *SortContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SortContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *SortContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *SortContext) AllSort() []ISortContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortContext); ok {
			len++
		}
	}

	tst := make([]ISortContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortContext); ok {
			tst[i] = t.(ISortContext)
			i++
		}
	}

	return tst
}

func (s *SortContext) Sort(i int) ISortContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
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

	return t.(ISortContext)
}

func (s *SortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSort(s)
	}
}

func (s *SortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSort(s)
	}
}

func (p *SMTXParser) Sort() (localctx ISortContext) {
	localctx = NewSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SMTXParserRULE_sort)
	var _la int

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Identifier()
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14636698788954112) != 0) {
			{
				p.SetState(262)
				p.Sort()
			}

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(267)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISort_decContext is an interface to support dynamic dispatch.
type ISort_decContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAROPEN() antlr.TerminalNode
	Symbol() ISymbolContext
	NUMERAL() antlr.TerminalNode
	PARCLOSE() antlr.TerminalNode

	// IsSort_decContext differentiates from other interfaces.
	IsSort_decContext()
}

type Sort_decContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySort_decContext() *Sort_decContext {
	var p = new(Sort_decContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_sort_dec
	return p
}

func InitEmptySort_decContext(p *Sort_decContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_sort_dec
}

func (*Sort_decContext) IsSort_decContext() {}

func NewSort_decContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sort_decContext {
	var p = new(Sort_decContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_sort_dec

	return p
}

func (s *Sort_decContext) GetParser() antlr.Parser { return s.parser }

func (s *Sort_decContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Sort_decContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Sort_decContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserNUMERAL, 0)
}

func (s *Sort_decContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Sort_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sort_decContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sort_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSort_dec(s)
	}
}

func (s *Sort_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSort_dec(s)
	}
}

func (p *SMTXParser) Sort_dec() (localctx ISort_decContext) {
	localctx = NewSort_decContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SMTXParserRULE_sort_dec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Symbol()
	}
	{
		p.SetState(273)
		p.Match(SMTXParserNUMERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelector_decContext is an interface to support dynamic dispatch.
type ISelector_decContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAROPEN() antlr.TerminalNode
	Symbol() ISymbolContext
	Sort() ISortContext
	PARCLOSE() antlr.TerminalNode

	// IsSelector_decContext differentiates from other interfaces.
	IsSelector_decContext()
}

type Selector_decContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelector_decContext() *Selector_decContext {
	var p = new(Selector_decContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_selector_dec
	return p
}

func InitEmptySelector_decContext(p *Selector_decContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_selector_dec
}

func (*Selector_decContext) IsSelector_decContext() {}

func NewSelector_decContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Selector_decContext {
	var p = new(Selector_decContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_selector_dec

	return p
}

func (s *Selector_decContext) GetParser() antlr.Parser { return s.parser }

func (s *Selector_decContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Selector_decContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Selector_decContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *Selector_decContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Selector_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Selector_decContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Selector_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSelector_dec(s)
	}
}

func (s *Selector_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSelector_dec(s)
	}
}

func (p *SMTXParser) Selector_dec() (localctx ISelector_decContext) {
	localctx = NewSelector_decContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SMTXParserRULE_selector_dec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Symbol()
	}
	{
		p.SetState(278)
		p.Sort()
	}
	{
		p.SetState(279)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructor_decContext is an interface to support dynamic dispatch.
type IConstructor_decContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAROPEN() antlr.TerminalNode
	Symbol() ISymbolContext
	PARCLOSE() antlr.TerminalNode
	AllSelector_dec() []ISelector_decContext
	Selector_dec(i int) ISelector_decContext

	// IsConstructor_decContext differentiates from other interfaces.
	IsConstructor_decContext()
}

type Constructor_decContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructor_decContext() *Constructor_decContext {
	var p = new(Constructor_decContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_constructor_dec
	return p
}

func InitEmptyConstructor_decContext(p *Constructor_decContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_constructor_dec
}

func (*Constructor_decContext) IsConstructor_decContext() {}

func NewConstructor_decContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constructor_decContext {
	var p = new(Constructor_decContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_constructor_dec

	return p
}

func (s *Constructor_decContext) GetParser() antlr.Parser { return s.parser }

func (s *Constructor_decContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Constructor_decContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Constructor_decContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Constructor_decContext) AllSelector_dec() []ISelector_decContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelector_decContext); ok {
			len++
		}
	}

	tst := make([]ISelector_decContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelector_decContext); ok {
			tst[i] = t.(ISelector_decContext)
			i++
		}
	}

	return tst
}

func (s *Constructor_decContext) Selector_dec(i int) ISelector_decContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelector_decContext); ok {
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

	return t.(ISelector_decContext)
}

func (s *Constructor_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constructor_decContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constructor_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterConstructor_dec(s)
	}
}

func (s *Constructor_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitConstructor_dec(s)
	}
}

func (p *SMTXParser) Constructor_dec() (localctx IConstructor_decContext) {
	localctx = NewConstructor_decContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SMTXParserRULE_constructor_dec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Symbol()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SMTXParserPAROPEN {
		{
			p.SetState(283)
			p.Selector_dec()
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(289)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatatype_decContext is an interface to support dynamic dispatch.
type IDatatype_decContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPAROPEN() []antlr.TerminalNode
	PAROPEN(i int) antlr.TerminalNode
	AllPARCLOSE() []antlr.TerminalNode
	PARCLOSE(i int) antlr.TerminalNode
	AllConstructor_dec() []IConstructor_decContext
	Constructor_dec(i int) IConstructor_decContext
	AllSymbol() []ISymbolContext
	Symbol(i int) ISymbolContext

	// IsDatatype_decContext differentiates from other interfaces.
	IsDatatype_decContext()
}

type Datatype_decContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatype_decContext() *Datatype_decContext {
	var p = new(Datatype_decContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_datatype_dec
	return p
}

func InitEmptyDatatype_decContext(p *Datatype_decContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_datatype_dec
}

func (*Datatype_decContext) IsDatatype_decContext() {}

func NewDatatype_decContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datatype_decContext {
	var p = new(Datatype_decContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_datatype_dec

	return p
}

func (s *Datatype_decContext) GetParser() antlr.Parser { return s.parser }

func (s *Datatype_decContext) AllPAROPEN() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPAROPEN)
}

func (s *Datatype_decContext) PAROPEN(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, i)
}

func (s *Datatype_decContext) AllPARCLOSE() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPARCLOSE)
}

func (s *Datatype_decContext) PARCLOSE(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, i)
}

func (s *Datatype_decContext) AllConstructor_dec() []IConstructor_decContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstructor_decContext); ok {
			len++
		}
	}

	tst := make([]IConstructor_decContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstructor_decContext); ok {
			tst[i] = t.(IConstructor_decContext)
			i++
		}
	}

	return tst
}

func (s *Datatype_decContext) Constructor_dec(i int) IConstructor_decContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructor_decContext); ok {
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

	return t.(IConstructor_decContext)
}

func (s *Datatype_decContext) AllSymbol() []ISymbolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISymbolContext); ok {
			len++
		}
	}

	tst := make([]ISymbolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISymbolContext); ok {
			tst[i] = t.(ISymbolContext)
			i++
		}
	}

	return tst
}

func (s *Datatype_decContext) Symbol(i int) ISymbolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
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

	return t.(ISymbolContext)
}

func (s *Datatype_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datatype_decContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datatype_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterDatatype_dec(s)
	}
}

func (s *Datatype_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitDatatype_dec(s)
	}
}

func (p *SMTXParser) Datatype_dec() (localctx IDatatype_decContext) {
	localctx = NewDatatype_decContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SMTXParserRULE_datatype_dec)
	var _la int

	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(292)
				p.Constructor_dec()
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(297)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(SMTXParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserSIMPLE_SYMBOL || _la == SMTXParserQUOTED_SYMBOL {
			{
				p.SetState(302)
				p.Symbol()
			}

			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(307)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(309)
				p.Constructor_dec()
			}

			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(314)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_decContext is an interface to support dynamic dispatch.
type IFunction_decContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPAROPEN() []antlr.TerminalNode
	PAROPEN(i int) antlr.TerminalNode
	Symbol() ISymbolContext
	AllPARCLOSE() []antlr.TerminalNode
	PARCLOSE(i int) antlr.TerminalNode
	Sort() ISortContext
	AllSorted_var() []ISorted_varContext
	Sorted_var(i int) ISorted_varContext

	// IsFunction_decContext differentiates from other interfaces.
	IsFunction_decContext()
}

type Function_decContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_decContext() *Function_decContext {
	var p = new(Function_decContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_function_dec
	return p
}

func InitEmptyFunction_decContext(p *Function_decContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_function_dec
}

func (*Function_decContext) IsFunction_decContext() {}

func NewFunction_decContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_decContext {
	var p = new(Function_decContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_function_dec

	return p
}

func (s *Function_decContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_decContext) AllPAROPEN() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPAROPEN)
}

func (s *Function_decContext) PAROPEN(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, i)
}

func (s *Function_decContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Function_decContext) AllPARCLOSE() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPARCLOSE)
}

func (s *Function_decContext) PARCLOSE(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, i)
}

func (s *Function_decContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *Function_decContext) AllSorted_var() []ISorted_varContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISorted_varContext); ok {
			len++
		}
	}

	tst := make([]ISorted_varContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISorted_varContext); ok {
			tst[i] = t.(ISorted_varContext)
			i++
		}
	}

	return tst
}

func (s *Function_decContext) Sorted_var(i int) ISorted_varContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISorted_varContext); ok {
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

	return t.(ISorted_varContext)
}

func (s *Function_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_decContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterFunction_dec(s)
	}
}

func (s *Function_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitFunction_dec(s)
	}
}

func (p *SMTXParser) Function_dec() (localctx IFunction_decContext) {
	localctx = NewFunction_decContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SMTXParserRULE_function_dec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Symbol()
	}
	{
		p.SetState(321)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SMTXParserPAROPEN {
		{
			p.SetState(322)
			p.Sorted_var()
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(328)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Sort()
	}
	{
		p.SetState(330)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_defContext is an interface to support dynamic dispatch.
type IFunction_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Symbol() ISymbolContext
	PAROPEN() antlr.TerminalNode
	PARCLOSE() antlr.TerminalNode
	Sort() ISortContext
	Term() ITermContext
	AllSorted_var() []ISorted_varContext
	Sorted_var(i int) ISorted_varContext

	// IsFunction_defContext differentiates from other interfaces.
	IsFunction_defContext()
}

type Function_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_defContext() *Function_defContext {
	var p = new(Function_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_function_def
	return p
}

func InitEmptyFunction_defContext(p *Function_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_function_def
}

func (*Function_defContext) IsFunction_defContext() {}

func NewFunction_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_defContext {
	var p = new(Function_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_function_def

	return p
}

func (s *Function_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_defContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Function_defContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Function_defContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Function_defContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *Function_defContext) Term() ITermContext {
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

func (s *Function_defContext) AllSorted_var() []ISorted_varContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISorted_varContext); ok {
			len++
		}
	}

	tst := make([]ISorted_varContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISorted_varContext); ok {
			tst[i] = t.(ISorted_varContext)
			i++
		}
	}

	return tst
}

func (s *Function_defContext) Sorted_var(i int) ISorted_varContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISorted_varContext); ok {
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

	return t.(ISorted_varContext)
}

func (s *Function_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterFunction_def(s)
	}
}

func (s *Function_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitFunction_def(s)
	}
}

func (p *SMTXParser) Function_def() (localctx IFunction_defContext) {
	localctx = NewFunction_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SMTXParserRULE_function_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Symbol()
	}
	{
		p.SetState(333)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SMTXParserPAROPEN {
		{
			p.SetState(334)
			p.Sorted_var()
		}

		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(340)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Sort()
	}
	{
		p.SetState(342)
		p.Term()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Spec_constant() ISpec_constantContext
	Qual_identifier() IQual_identifierContext
	AllPAROPEN() []antlr.TerminalNode
	PAROPEN(i int) antlr.TerminalNode
	AllPARCLOSE() []antlr.TerminalNode
	PARCLOSE(i int) antlr.TerminalNode
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllVar_binding() []IVar_bindingContext
	Var_binding(i int) IVar_bindingContext
	AllSorted_var() []ISorted_varContext
	Sorted_var(i int) ISorted_varContext
	AllMatch_case() []IMatch_caseContext
	Match_case(i int) IMatch_caseContext
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Spec_constant() ISpec_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpec_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpec_constantContext)
}

func (s *TermContext) Qual_identifier() IQual_identifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQual_identifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQual_identifierContext)
}

func (s *TermContext) AllPAROPEN() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPAROPEN)
}

func (s *TermContext) PAROPEN(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, i)
}

func (s *TermContext) AllPARCLOSE() []antlr.TerminalNode {
	return s.GetTokens(SMTXParserPARCLOSE)
}

func (s *TermContext) PARCLOSE(i int) antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, i)
}

func (s *TermContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *TermContext) AllVar_binding() []IVar_bindingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_bindingContext); ok {
			len++
		}
	}

	tst := make([]IVar_bindingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_bindingContext); ok {
			tst[i] = t.(IVar_bindingContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Var_binding(i int) IVar_bindingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_bindingContext); ok {
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

	return t.(IVar_bindingContext)
}

func (s *TermContext) AllSorted_var() []ISorted_varContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISorted_varContext); ok {
			len++
		}
	}

	tst := make([]ISorted_varContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISorted_varContext); ok {
			tst[i] = t.(ISorted_varContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Sorted_var(i int) ISorted_varContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISorted_varContext); ok {
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

	return t.(ISorted_varContext)
}

func (s *TermContext) AllMatch_case() []IMatch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatch_caseContext); ok {
			len++
		}
	}

	tst := make([]IMatch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatch_caseContext); ok {
			tst[i] = t.(IMatch_caseContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Match_case(i int) IMatch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatch_caseContext); ok {
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

	return t.(IMatch_caseContext)
}

func (s *TermContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
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

	return t.(IAttributeContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *SMTXParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SMTXParserRULE_term)
	var _la int

	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.Spec_constant()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Qual_identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(346)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Qual_identifier()
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15727414323707904) != 0) {
			{
				p.SetState(348)
				p.Term()
			}

			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(353)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(355)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(SMTXParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(358)
				p.Var_binding()
			}

			p.SetState(361)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(363)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Term()
		}
		{
			p.SetState(365)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(367)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Match(SMTXParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(370)
				p.Sorted_var()
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(375)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.Term()
		}
		{
			p.SetState(377)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(379)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Match(SMTXParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(382)
				p.Sorted_var()
			}

			p.SetState(385)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(387)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Term()
		}
		{
			p.SetState(389)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(391)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(SMTXParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(394)
				p.Sorted_var()
			}

			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(399)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Term()
		}
		{
			p.SetState(401)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(403)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(SMTXParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Term()
		}
		{
			p.SetState(406)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserPAROPEN {
			{
				p.SetState(407)
				p.Match_case()
			}

			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(412)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(415)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(SMTXParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Term()
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SMTXParserKEYWORD {
			{
				p.SetState(418)
				p.Attribute()
			}

			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(423)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpec_constantContext is an interface to support dynamic dispatch.
type ISpec_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERAL() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	HEXADECIMAL() antlr.TerminalNode
	BINARY() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsSpec_constantContext differentiates from other interfaces.
	IsSpec_constantContext()
}

type Spec_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpec_constantContext() *Spec_constantContext {
	var p = new(Spec_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_spec_constant
	return p
}

func InitEmptySpec_constantContext(p *Spec_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_spec_constant
}

func (*Spec_constantContext) IsSpec_constantContext() {}

func NewSpec_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spec_constantContext {
	var p = new(Spec_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_spec_constant

	return p
}

func (s *Spec_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Spec_constantContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserNUMERAL, 0)
}

func (s *Spec_constantContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserDECIMAL, 0)
}

func (s *Spec_constantContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserHEXADECIMAL, 0)
}

func (s *Spec_constantContext) BINARY() antlr.TerminalNode {
	return s.GetToken(SMTXParserBINARY, 0)
}

func (s *Spec_constantContext) STRING() antlr.TerminalNode {
	return s.GetToken(SMTXParserSTRING, 0)
}

func (s *Spec_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spec_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spec_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSpec_constant(s)
	}
}

func (s *Spec_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSpec_constant(s)
	}
}

func (p *SMTXParser) Spec_constant() (localctx ISpec_constantContext) {
	localctx = NewSpec_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SMTXParserRULE_spec_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1090715534753792) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQual_identifierContext is an interface to support dynamic dispatch.
type IQual_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	PAROPEN() antlr.TerminalNode
	Sort() ISortContext
	PARCLOSE() antlr.TerminalNode

	// IsQual_identifierContext differentiates from other interfaces.
	IsQual_identifierContext()
}

type Qual_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQual_identifierContext() *Qual_identifierContext {
	var p = new(Qual_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_qual_identifier
	return p
}

func InitEmptyQual_identifierContext(p *Qual_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_qual_identifier
}

func (*Qual_identifierContext) IsQual_identifierContext() {}

func NewQual_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Qual_identifierContext {
	var p = new(Qual_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_qual_identifier

	return p
}

func (s *Qual_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Qual_identifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Qual_identifierContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Qual_identifierContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *Qual_identifierContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Qual_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qual_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Qual_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterQual_identifier(s)
	}
}

func (s *Qual_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitQual_identifier(s)
	}
}

func (p *SMTXParser) Qual_identifier() (localctx IQual_identifierContext) {
	localctx = NewQual_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SMTXParserRULE_qual_identifier)
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(SMTXParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Identifier()
		}
		{
			p.SetState(433)
			p.Sort()
		}
		{
			p.SetState(434)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_bindingContext is an interface to support dynamic dispatch.
type IVar_bindingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAROPEN() antlr.TerminalNode
	Symbol() ISymbolContext
	Term() ITermContext
	PARCLOSE() antlr.TerminalNode

	// IsVar_bindingContext differentiates from other interfaces.
	IsVar_bindingContext()
}

type Var_bindingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_bindingContext() *Var_bindingContext {
	var p = new(Var_bindingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_var_binding
	return p
}

func InitEmptyVar_bindingContext(p *Var_bindingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_var_binding
}

func (*Var_bindingContext) IsVar_bindingContext() {}

func NewVar_bindingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_bindingContext {
	var p = new(Var_bindingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_var_binding

	return p
}

func (s *Var_bindingContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_bindingContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Var_bindingContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Var_bindingContext) Term() ITermContext {
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

func (s *Var_bindingContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Var_bindingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_bindingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_bindingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterVar_binding(s)
	}
}

func (s *Var_bindingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitVar_binding(s)
	}
}

func (p *SMTXParser) Var_binding() (localctx IVar_bindingContext) {
	localctx = NewVar_bindingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SMTXParserRULE_var_binding)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Symbol()
	}
	{
		p.SetState(440)
		p.Term()
	}
	{
		p.SetState(441)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISorted_varContext is an interface to support dynamic dispatch.
type ISorted_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAROPEN() antlr.TerminalNode
	Symbol() ISymbolContext
	Sort() ISortContext
	PARCLOSE() antlr.TerminalNode

	// IsSorted_varContext differentiates from other interfaces.
	IsSorted_varContext()
}

type Sorted_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySorted_varContext() *Sorted_varContext {
	var p = new(Sorted_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_sorted_var
	return p
}

func InitEmptySorted_varContext(p *Sorted_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_sorted_var
}

func (*Sorted_varContext) IsSorted_varContext() {}

func NewSorted_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sorted_varContext {
	var p = new(Sorted_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_sorted_var

	return p
}

func (s *Sorted_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Sorted_varContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Sorted_varContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Sorted_varContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *Sorted_varContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Sorted_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sorted_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sorted_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSorted_var(s)
	}
}

func (s *Sorted_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSorted_var(s)
	}
}

func (p *SMTXParser) Sorted_var() (localctx ISorted_varContext) {
	localctx = NewSorted_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SMTXParserRULE_sorted_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.Symbol()
	}
	{
		p.SetState(445)
		p.Sort()
	}
	{
		p.SetState(446)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSymbol_() []ISymbol_Context
	Symbol_(i int) ISymbol_Context
	PAROPEN() antlr.TerminalNode
	Symbol() ISymbolContext
	PARCLOSE() antlr.TerminalNode

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_pattern
	return p
}

func InitEmptyPatternContext(p *PatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_pattern
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) AllSymbol_() []ISymbol_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISymbol_Context); ok {
			len++
		}
	}

	tst := make([]ISymbol_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISymbol_Context); ok {
			tst[i] = t.(ISymbol_Context)
			i++
		}
	}

	return tst
}

func (s *PatternContext) Symbol_(i int) ISymbol_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbol_Context); ok {
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

	return t.(ISymbol_Context)
}

func (s *PatternContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *PatternContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *PatternContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *SMTXParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SMTXParserRULE_pattern)
	var _la int

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserT__40, SMTXParserSIMPLE_SYMBOL, SMTXParserQUOTED_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			p.Symbol_()
		}

	case SMTXParserPAROPEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.Symbol()
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13512997905367040) != 0) {
			{
				p.SetState(451)
				p.Symbol_()
			}

			p.SetState(454)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(456)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISymbol_Context is an interface to support dynamic dispatch.
type ISymbol_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Symbol() ISymbolContext

	// IsSymbol_Context differentiates from other interfaces.
	IsSymbol_Context()
}

type Symbol_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbol_Context() *Symbol_Context {
	var p = new(Symbol_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_symbol_
	return p
}

func InitEmptySymbol_Context(p *Symbol_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_symbol_
}

func (*Symbol_Context) IsSymbol_Context() {}

func NewSymbol_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Symbol_Context {
	var p = new(Symbol_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_symbol_

	return p
}

func (s *Symbol_Context) GetParser() antlr.Parser { return s.parser }

func (s *Symbol_Context) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Symbol_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Symbol_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Symbol_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterSymbol_(s)
	}
}

func (s *Symbol_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitSymbol_(s)
	}
}

func (p *SMTXParser) Symbol_() (localctx ISymbol_Context) {
	localctx = NewSymbol_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SMTXParserRULE_symbol_)
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserSIMPLE_SYMBOL, SMTXParserQUOTED_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Symbol()
		}

	case SMTXParserT__40:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.Match(SMTXParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatch_caseContext is an interface to support dynamic dispatch.
type IMatch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAROPEN() antlr.TerminalNode
	Pattern() IPatternContext
	Term() ITermContext
	PARCLOSE() antlr.TerminalNode

	// IsMatch_caseContext differentiates from other interfaces.
	IsMatch_caseContext()
}

type Match_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatch_caseContext() *Match_caseContext {
	var p = new(Match_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_match_case
	return p
}

func InitEmptyMatch_caseContext(p *Match_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_match_case
}

func (*Match_caseContext) IsMatch_caseContext() {}

func NewMatch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Match_caseContext {
	var p = new(Match_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_match_case

	return p
}

func (s *Match_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Match_caseContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Match_caseContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *Match_caseContext) Term() ITermContext {
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

func (s *Match_caseContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Match_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Match_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Match_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterMatch_case(s)
	}
}

func (s *Match_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitMatch_case(s)
	}
}

func (p *SMTXParser) Match_case() (localctx IMatch_caseContext) {
	localctx = NewMatch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SMTXParserRULE_match_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Match(SMTXParserPAROPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Pattern()
	}
	{
		p.SetState(466)
		p.Term()
	}
	{
		p.SetState(467)
		p.Match(SMTXParserPARCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Symbol() ISymbolContext
	PAROPEN() antlr.TerminalNode
	PARCLOSE() antlr.TerminalNode
	AllIndex() []IIndexContext
	Index(i int) IIndexContext

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *IdentifierContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *IdentifierContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *IdentifierContext) AllIndex() []IIndexContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexContext); ok {
			len++
		}
	}

	tst := make([]IIndexContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexContext); ok {
			tst[i] = t.(IIndexContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierContext) Index(i int) IIndexContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
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

	return t.(IIndexContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *SMTXParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SMTXParserRULE_identifier)
	var _la int

	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserSIMPLE_SYMBOL, SMTXParserQUOTED_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(469)
			p.Symbol()
		}

	case SMTXParserPAROPEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(470)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(SMTXParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.Symbol()
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13545983254200320) != 0) {
			{
				p.SetState(473)
				p.Index()
			}

			p.SetState(476)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(478)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERAL() antlr.TerminalNode
	Symbol() ISymbolContext

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserNUMERAL, 0)
}

func (s *IndexContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *SMTXParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SMTXParserRULE_index)
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserNUMERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(482)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserSIMPLE_SYMBOL, SMTXParserQUOTED_SYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(483)
			p.Symbol()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttribute_valueContext is an interface to support dynamic dispatch.
type IAttribute_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Spec_constant() ISpec_constantContext
	Symbol() ISymbolContext
	PAROPEN() antlr.TerminalNode
	PARCLOSE() antlr.TerminalNode
	AllS_expr() []IS_exprContext
	S_expr(i int) IS_exprContext

	// IsAttribute_valueContext differentiates from other interfaces.
	IsAttribute_valueContext()
}

type Attribute_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribute_valueContext() *Attribute_valueContext {
	var p = new(Attribute_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_attribute_value
	return p
}

func InitEmptyAttribute_valueContext(p *Attribute_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_attribute_value
}

func (*Attribute_valueContext) IsAttribute_valueContext() {}

func NewAttribute_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attribute_valueContext {
	var p = new(Attribute_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_attribute_value

	return p
}

func (s *Attribute_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Attribute_valueContext) Spec_constant() ISpec_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpec_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpec_constantContext)
}

func (s *Attribute_valueContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Attribute_valueContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *Attribute_valueContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *Attribute_valueContext) AllS_expr() []IS_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IS_exprContext); ok {
			len++
		}
	}

	tst := make([]IS_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IS_exprContext); ok {
			tst[i] = t.(IS_exprContext)
			i++
		}
	}

	return tst
}

func (s *Attribute_valueContext) S_expr(i int) IS_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IS_exprContext); ok {
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

	return t.(IS_exprContext)
}

func (s *Attribute_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attribute_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attribute_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterAttribute_value(s)
	}
}

func (s *Attribute_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitAttribute_value(s)
	}
}

func (p *SMTXParser) Attribute_value() (localctx IAttribute_valueContext) {
	localctx = NewAttribute_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SMTXParserRULE_attribute_value)
	var _la int

	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserNUMERAL, SMTXParserDECIMAL, SMTXParserHEXADECIMAL, SMTXParserBINARY, SMTXParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.Spec_constant()
		}

	case SMTXParserSIMPLE_SYMBOL, SMTXParserQUOTED_SYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.Symbol()
		}

	case SMTXParserPAROPEN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33741812833189888) != 0 {
			{
				p.SetState(489)
				p.S_expr()
			}

			p.SetState(494)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(495)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEYWORD() antlr.TerminalNode
	Attribute_value() IAttribute_valueContext

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SMTXParserKEYWORD, 0)
}

func (s *AttributeContext) Attribute_value() IAttribute_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribute_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribute_valueContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *SMTXParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SMTXParserRULE_attribute)
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(498)
			p.Match(SMTXParserKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(499)
			p.Match(SMTXParserKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.Attribute_value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInfo_flagContext is an interface to support dynamic dispatch.
type IInfo_flagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PK_ALL_STATISTICS() antlr.TerminalNode
	PK_ASSERTION_STACK_LEVELS() antlr.TerminalNode
	PK_AUTHORS() antlr.TerminalNode
	PK_ERROR_BEHAVIOR() antlr.TerminalNode
	PK_NAME() antlr.TerminalNode
	PK_REASON_UNKNOWN() antlr.TerminalNode
	PK_VERSION() antlr.TerminalNode
	KEYWORD() antlr.TerminalNode

	// IsInfo_flagContext differentiates from other interfaces.
	IsInfo_flagContext()
}

type Info_flagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfo_flagContext() *Info_flagContext {
	var p = new(Info_flagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_info_flag
	return p
}

func InitEmptyInfo_flagContext(p *Info_flagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_info_flag
}

func (*Info_flagContext) IsInfo_flagContext() {}

func NewInfo_flagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Info_flagContext {
	var p = new(Info_flagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_info_flag

	return p
}

func (s *Info_flagContext) GetParser() antlr.Parser { return s.parser }

func (s *Info_flagContext) PK_ALL_STATISTICS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_ALL_STATISTICS, 0)
}

func (s *Info_flagContext) PK_ASSERTION_STACK_LEVELS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_ASSERTION_STACK_LEVELS, 0)
}

func (s *Info_flagContext) PK_AUTHORS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_AUTHORS, 0)
}

func (s *Info_flagContext) PK_ERROR_BEHAVIOR() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_ERROR_BEHAVIOR, 0)
}

func (s *Info_flagContext) PK_NAME() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_NAME, 0)
}

func (s *Info_flagContext) PK_REASON_UNKNOWN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_REASON_UNKNOWN, 0)
}

func (s *Info_flagContext) PK_VERSION() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_VERSION, 0)
}

func (s *Info_flagContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SMTXParserKEYWORD, 0)
}

func (s *Info_flagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Info_flagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Info_flagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterInfo_flag(s)
	}
}

func (s *Info_flagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitInfo_flag(s)
	}
}

func (p *SMTXParser) Info_flag() (localctx IInfo_flagContext) {
	localctx = NewInfo_flagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SMTXParserRULE_info_flag)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-54)) & ^0x3f) == 0 && ((int64(1)<<(_la-54))&72075194831601665) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PK_DIAGNOSTIC_OUTPUT_CHANNEL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	PK_GLOBAL_DECLARATIONS() antlr.TerminalNode
	B_value() IB_valueContext
	PK_INTERACTIVE_MODE() antlr.TerminalNode
	PK_PRINT_SUCCESS() antlr.TerminalNode
	PK_PRODUCE_ASSIGNMENTS() antlr.TerminalNode
	PK_PRODUCE_MODELS() antlr.TerminalNode
	PK_PRODUCE_PROOFS() antlr.TerminalNode
	PK_PRODUCE_UNSAT_ASSUMPTIONS() antlr.TerminalNode
	PK_PRODUCE_UNSAT_CORES() antlr.TerminalNode
	PK_RANDOM_SEED() antlr.TerminalNode
	NUMERAL() antlr.TerminalNode
	PK_REGULAR_OUTPUT_CHANNEL() antlr.TerminalNode
	PK_REPRODUCIBLE_RESOURCE_LIMIT() antlr.TerminalNode
	PK_VERBOSITY() antlr.TerminalNode
	Attribute() IAttributeContext

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) PK_DIAGNOSTIC_OUTPUT_CHANNEL() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_DIAGNOSTIC_OUTPUT_CHANNEL, 0)
}

func (s *OptionContext) STRING() antlr.TerminalNode {
	return s.GetToken(SMTXParserSTRING, 0)
}

func (s *OptionContext) PK_GLOBAL_DECLARATIONS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_GLOBAL_DECLARATIONS, 0)
}

func (s *OptionContext) B_value() IB_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IB_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IB_valueContext)
}

func (s *OptionContext) PK_INTERACTIVE_MODE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_INTERACTIVE_MODE, 0)
}

func (s *OptionContext) PK_PRINT_SUCCESS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_PRINT_SUCCESS, 0)
}

func (s *OptionContext) PK_PRODUCE_ASSIGNMENTS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_PRODUCE_ASSIGNMENTS, 0)
}

func (s *OptionContext) PK_PRODUCE_MODELS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_PRODUCE_MODELS, 0)
}

func (s *OptionContext) PK_PRODUCE_PROOFS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_PRODUCE_PROOFS, 0)
}

func (s *OptionContext) PK_PRODUCE_UNSAT_ASSUMPTIONS() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_PRODUCE_UNSAT_ASSUMPTIONS, 0)
}

func (s *OptionContext) PK_PRODUCE_UNSAT_CORES() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_PRODUCE_UNSAT_CORES, 0)
}

func (s *OptionContext) PK_RANDOM_SEED() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_RANDOM_SEED, 0)
}

func (s *OptionContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(SMTXParserNUMERAL, 0)
}

func (s *OptionContext) PK_REGULAR_OUTPUT_CHANNEL() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_REGULAR_OUTPUT_CHANNEL, 0)
}

func (s *OptionContext) PK_REPRODUCIBLE_RESOURCE_LIMIT() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_REPRODUCIBLE_RESOURCE_LIMIT, 0)
}

func (s *OptionContext) PK_VERBOSITY() antlr.TerminalNode {
	return s.GetToken(SMTXParserPK_VERBOSITY, 0)
}

func (s *OptionContext) Attribute() IAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *SMTXParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SMTXParserRULE_option)
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserPK_DIAGNOSTIC_OUTPUT_CHANNEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(505)
			p.Match(SMTXParserPK_DIAGNOSTIC_OUTPUT_CHANNEL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)
			p.Match(SMTXParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserPK_GLOBAL_DECLARATIONS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(507)
			p.Match(SMTXParserPK_GLOBAL_DECLARATIONS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)
			p.B_value()
		}

	case SMTXParserPK_INTERACTIVE_MODE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(509)
			p.Match(SMTXParserPK_INTERACTIVE_MODE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(510)
			p.B_value()
		}

	case SMTXParserPK_PRINT_SUCCESS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(511)
			p.Match(SMTXParserPK_PRINT_SUCCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.B_value()
		}

	case SMTXParserT__41:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(513)
			p.Match(SMTXParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.B_value()
		}

	case SMTXParserPK_PRODUCE_ASSIGNMENTS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(515)
			p.Match(SMTXParserPK_PRODUCE_ASSIGNMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.B_value()
		}

	case SMTXParserPK_PRODUCE_MODELS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(517)
			p.Match(SMTXParserPK_PRODUCE_MODELS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.B_value()
		}

	case SMTXParserPK_PRODUCE_PROOFS:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(519)
			p.Match(SMTXParserPK_PRODUCE_PROOFS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.B_value()
		}

	case SMTXParserPK_PRODUCE_UNSAT_ASSUMPTIONS:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(521)
			p.Match(SMTXParserPK_PRODUCE_UNSAT_ASSUMPTIONS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)
			p.B_value()
		}

	case SMTXParserPK_PRODUCE_UNSAT_CORES:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(523)
			p.Match(SMTXParserPK_PRODUCE_UNSAT_CORES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.B_value()
		}

	case SMTXParserPK_RANDOM_SEED:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(525)
			p.Match(SMTXParserPK_RANDOM_SEED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserPK_REGULAR_OUTPUT_CHANNEL:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(527)
			p.Match(SMTXParserPK_REGULAR_OUTPUT_CHANNEL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Match(SMTXParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserPK_REPRODUCIBLE_RESOURCE_LIMIT:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(529)
			p.Match(SMTXParserPK_REPRODUCIBLE_RESOURCE_LIMIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserPK_VERBOSITY:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(531)
			p.Match(SMTXParserPK_VERBOSITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.Match(SMTXParserNUMERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserKEYWORD:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(533)
			p.Attribute()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IB_valueContext is an interface to support dynamic dispatch.
type IB_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PS_TRUE() antlr.TerminalNode
	PS_FALSE() antlr.TerminalNode

	// IsB_valueContext differentiates from other interfaces.
	IsB_valueContext()
}

type B_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyB_valueContext() *B_valueContext {
	var p = new(B_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_b_value
	return p
}

func InitEmptyB_valueContext(p *B_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_b_value
}

func (*B_valueContext) IsB_valueContext() {}

func NewB_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *B_valueContext {
	var p = new(B_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_b_value

	return p
}

func (s *B_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *B_valueContext) PS_TRUE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPS_TRUE, 0)
}

func (s *B_valueContext) PS_FALSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPS_FALSE, 0)
}

func (s *B_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *B_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *B_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterB_value(s)
	}
}

func (s *B_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitB_value(s)
	}
}

func (p *SMTXParser) B_value() (localctx IB_valueContext) {
	localctx = NewB_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SMTXParserRULE_b_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(536)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SMTXParserPS_FALSE || _la == SMTXParserPS_TRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IS_exprContext is an interface to support dynamic dispatch.
type IS_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Spec_constant() ISpec_constantContext
	Symbol() ISymbolContext
	KEYWORD() antlr.TerminalNode
	PAROPEN() antlr.TerminalNode
	PARCLOSE() antlr.TerminalNode
	AllS_expr() []IS_exprContext
	S_expr(i int) IS_exprContext

	// IsS_exprContext differentiates from other interfaces.
	IsS_exprContext()
}

type S_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyS_exprContext() *S_exprContext {
	var p = new(S_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_s_expr
	return p
}

func InitEmptyS_exprContext(p *S_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SMTXParserRULE_s_expr
}

func (*S_exprContext) IsS_exprContext() {}

func NewS_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *S_exprContext {
	var p = new(S_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SMTXParserRULE_s_expr

	return p
}

func (s *S_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *S_exprContext) Spec_constant() ISpec_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpec_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpec_constantContext)
}

func (s *S_exprContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *S_exprContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SMTXParserKEYWORD, 0)
}

func (s *S_exprContext) PAROPEN() antlr.TerminalNode {
	return s.GetToken(SMTXParserPAROPEN, 0)
}

func (s *S_exprContext) PARCLOSE() antlr.TerminalNode {
	return s.GetToken(SMTXParserPARCLOSE, 0)
}

func (s *S_exprContext) AllS_expr() []IS_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IS_exprContext); ok {
			len++
		}
	}

	tst := make([]IS_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IS_exprContext); ok {
			tst[i] = t.(IS_exprContext)
			i++
		}
	}

	return tst
}

func (s *S_exprContext) S_expr(i int) IS_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IS_exprContext); ok {
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

	return t.(IS_exprContext)
}

func (s *S_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *S_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.EnterS_expr(s)
	}
}

func (s *S_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SMTXListener); ok {
		listenerT.ExitS_expr(s)
	}
}

func (p *SMTXParser) S_expr() (localctx IS_exprContext) {
	localctx = NewS_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SMTXParserRULE_s_expr)
	var _la int

	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SMTXParserNUMERAL, SMTXParserDECIMAL, SMTXParserHEXADECIMAL, SMTXParserBINARY, SMTXParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.Spec_constant()
		}

	case SMTXParserSIMPLE_SYMBOL, SMTXParserQUOTED_SYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(539)
			p.Symbol()
		}

	case SMTXParserKEYWORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(540)
			p.Match(SMTXParserKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SMTXParserPAROPEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(541)
			p.Match(SMTXParserPAROPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33741812833189888) != 0 {
			{
				p.SetState(542)
				p.S_expr()
			}

			p.SetState(547)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(548)
			p.Match(SMTXParserPARCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
