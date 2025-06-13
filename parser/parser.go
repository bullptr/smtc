package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func ParseFile(filePath string) (*SMTXLexer, *SMTXParser) {
	input, _ := antlr.NewFileStream(filePath)
	lexer := NewSMTXLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewSMTXParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	return lexer, parser
}

func VisualizeParseTree(parser *SMTXParser, tree IScriptContext, filePath string) {
	visualizer := NewGraphvizVisitor(parser.GetRuleNames(), parser.GetTokenNames())
	visualizer.Visit(tree)

	err := os.WriteFile(filePath+".dot", []byte(visualizer.GetDOT()), 0644)
	if err != nil {
		fmt.Printf("Error writing DOT file: %v\n", err)
	}
}
