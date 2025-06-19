package parser

import (
	"fmt"
	"go/token"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
	Src   *string
	Lexer *SMTCLexer
	Tree  *SMTCParser
}

func (p *Parser) ResetLexer() {
	p.Lexer.Reset()
}

func NewParser(fset *token.FileSet, filename string) *Parser {
	input, _ := antlr.NewFileStream(filename)
	src := input.String()
	lexer := NewSMTCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewSMTCParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// stream.Fill() // necessary to for fset building
	BuildFileSet(fset, filename, []byte(src))

	return &Parser{
		Src:   &src,
		Lexer: lexer,
		Tree:  parser,
	}
}

func BuildFileSet(fset *token.FileSet, filename string, content []byte) {
	file := fset.AddFile(filename, -1, len(content))

	file.SetLinesForContent(content)
}

func VisualizeParserGraph(parser *Parser, filename string) {
	println("Visualizing parser graph for:", filename)
	tree := parser.Tree.Script()
	visualizer := NewGraphvizVisitor(
		parser.Tree.GetRuleNames(),
		parser.Tree.GetTokenNames(),
	)
	visualizer.Visit(tree)

	err := os.WriteFile(filename+".dot", []byte(visualizer.GetDOT()), 0644)
	if err != nil {
		fmt.Printf("Error writing DOT file: %v\n", err)
	}

	parser.ResetLexer()
}

// type FileStream struct {
// 	InputStream
// 	filename string
// }

// // NewInputStream creates a new input stream from the given string
// func NewInputStream(content []byte) *antlr.InputStream {
// 	return antlr.NewInputStream(string(content))
// }
