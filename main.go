package main

import (
	"os"

	arg "github.com/alexflint/go-arg"

	// "github.com/smtx/compiler"
	"github.com/smtx/config"
	"github.com/smtx/parser"
	smtx "github.com/smtx/smtx_parser"
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
			smtx.ParseFile(file)
			// tree := parser.Start_()
			// println(tree.ToStringTree(nil, parser))
		}

		// compiler.NewCompilerFromArgs(args)
		// c.CheckTypes()

	default:
		// argParser.WriteUsage(os.Stderr)
		parser.TestBindings()
	}
}
