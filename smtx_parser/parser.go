package smtx_parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func ParseFile(filePath string) *SMTXParser {
	input, _ := antlr.NewFileStream(filePath)
	lexer := NewSMTXLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewSMTXParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := parser.Script()

	visualizer := NewGraphvizVisitor(parser.GetRuleNames(), parser.GetTokenNames())
	visualizer.Visit(tree)

	// Generate DOT output
	// fmt.Println(visualizer.GetDOT())
	// write visualizer.GetDOT() to a file named filePath + ".dot"
	err := os.WriteFile(filePath+".dot", []byte(visualizer.GetDOT()), 0644)
	if err != nil {
		fmt.Printf("Error writing DOT file: %v\n", err)
		return nil
	}

	lexer.Reset()

	return parser
}
