package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
	Lexer *SMTXLexer
	Tree  *SMTXParser
}

func (p *Parser) ResetLexer() {
	p.Lexer.Reset()
}

func NewParser(filePath string) *Parser {
	input, _ := antlr.NewFileStream(filePath)
	lexer := NewSMTXLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewSMTXParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	return &Parser{
		Lexer: lexer,
		Tree:  parser,
	}
}

func VisualizeParserGraph(parser *Parser, filePath string) {
	tree := parser.Tree.Script()
	visualizer := NewGraphvizVisitor(
		parser.Tree.GetRuleNames(),
		parser.Tree.GetTokenNames(),
	)
	visualizer.Visit(tree)

	err := os.WriteFile(filePath+".dot", []byte(visualizer.GetDOT()), 0644)
	if err != nil {
		fmt.Printf("Error writing DOT file: %v\n", err)
	}
}
