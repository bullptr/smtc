package main

import (
	"os"

	arg "github.com/alexflint/go-arg"

	"github.com/smtx/config"
	p "github.com/smtx/parser"
)

var args config.CmdArgs

func main() {
	argParser := arg.MustParse(&args)

	switch {
	case args.Check != nil:
		if len(args.Check.Files) == 0 {
			argParser.WriteUsage(os.Stderr)
			os.Exit(0)
		}

		for _, file := range args.Check.Files {
			lexer, parser := p.ParseFile(file)
			tree := parser.Script()

			if args.Graph {
				p.VisualizeParseTree(parser, tree, file)
			}

			lexer.Reset()
		}

		// compiler.NewCompilerFromArgs(args)
		// c.CheckTypes()

	default:
		// argParser.WriteUsage(os.Stderr)
	}
}
