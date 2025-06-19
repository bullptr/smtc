package compiler

import (
	"fmt"
	"go/token"
	"log"
	"os"

	gast "go/ast"

	"github.com/smtc/ast"
	"github.com/smtc/config"
	"github.com/smtc/parser"
	"github.com/smtc/types"
	"github.com/smtc/utils"
)

type Config struct {
	Include []string
	Exclude []string
	Logics  []string
	Check   bool // check for type errors
	Verbose bool
}

type DiagnosticType int

const (
	Error DiagnosticType = iota
	SyntaxError
	CompilerError
	InternalError
	TypeError
	Warning
	Info
)

type Diagnostic struct {
	Pos     token.Position
	Message string
	Type    DiagnosticType
}

type Compiler struct {
	Config      *Config
	Fset        *token.FileSet
	Files       []*ast.SourceFile
	Diagnostics []*Diagnostic
}

func NewCompiler() *Compiler {
	return &Compiler{
		Fset: token.NewFileSet(),
	}
}

func NewCompilerFromArgs(args config.CmdArgs) *Compiler {
	c := &Compiler{
		Config: &Config{
			Include: []string{},
			Exclude: []string{},
			Logics:  []string{},
			Check:   true,
			Verbose: args.Verbose,
		},
		Fset: token.NewFileSet(),
	}
	// filenames := utils.GetFilesToCheck(args.Check.Files)

	for _, filename := range args.Check.Files {
		c.CompileSourceFile(filename)

	}

	if args.Graph {
		for _, sf := range c.Files {
			parser.VisualizeParserGraph(sf.Parser, sf.Filename)
		}
	}

	// c.Fset.Iterate(func(f *token.File) bool {
	// 	fmt.Printf("%v\n", f)
	// 	return false // continue iterating
	// })

	if args.Verbose {
		fmt.Printf("Files: %d\n", len(c.Files))
	}

	return c
}

func (c *Compiler) CompileSourceFile(filename string) {
	pkg := ast.Ident{Name: "main"}
	p := parser.NewParser(c.Fset, filename)
	sf := &ast.SourceFile{
		Filename: filename,
		Parser:   p,
		Ast: &ast.File{
			Name:      &pkg,
			Scope:     gast.NewScope(nil),
			FileStart: token.Pos(0),
			FileEnd:   token.Pos(len(*p.Src)),
			GoVersion: "1.20",
			// @TODO: build and import prelude library
			// Imports:   c.LoadPrelude(),
		},
	}

	cmds := BuildCommands(sf.Parser)
	println("Commands:", len(cmds))
	mainFunc := BuildMainFunction(&cmds)
	// sf.Ast.Decls = append(sf.Ast.Decls, topLevelDelcs...)
	sf.Ast.Decls = append(sf.Ast.Decls, mainFunc)

	// ast.PrintSourceFile(c.Fset, sf)
	// ast.PrintAst(c.Fset, sf.Ast)
	sf.Parser.ResetLexer()
	c.Files = append(c.Files, sf)
	// c.CheckTypes()
}

// @TODO: check for types for all
func (c *Compiler) CheckTypes() {
	utils.PrintSectionHeader("Check Types")
	for _, sf := range c.Files {
		_, err := types.CheckSourceFile(c.Fset, sf)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
	}

}

// @TODO: load prelude library from logics and theories
func (c *Compiler) LoadPrelude() []*gast.ImportSpec {
	return []*gast.ImportSpec{
		// gast.NewImportSpec(nil, nil, nil, &gast.Ident{Name: "prelude"}),
		// gast.NewImportSpec(nil, nil, nil, &gast.Ident{Name: "smtc/prelude"}),
	}
}
