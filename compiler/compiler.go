package compiler

import (
	"fmt"
	"go/token"
	"log"
	"os"

	gast "go/ast"

	"github.com/smtx/ast"
	"github.com/smtx/config"
	"github.com/smtx/parser"
	"github.com/smtx/types"
	"github.com/smtx/utils"
)

/**
 * @TODOs
 * - add new import logic
 */

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
	// 	return true // continue iterating
	// })

	if args.Verbose {
		fmt.Printf("Files: %d\n", len(c.Files))
	}

	return c
}

func (c *Compiler) CompileSourceFile(filename string) {
	pkg := ast.Ident{
		Name: "main",
	}
	sf := &ast.SourceFile{
		Filename: filename,
		Parser:   parser.NewParser(c.Fset, filename),
		Ast: &ast.File{
			Name:  &pkg,
			Scope: gast.NewScope(nil),
		},
	}

	// cmds := BuildCommands(&sf.Src, sf.Parser.RootNode())
	cmds := []gast.Stmt{}
	mainFunc := BuildMainFunction(&cmds)
	// sf.Ast.Decls = append(sf.Ast.Decls, topLevelDelcs...)
	sf.Ast.Decls = append(sf.Ast.Decls, mainFunc)

	// ast.PrintAst(sf.Ast)
	// ast.PrintSourceFile(sf)
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
func (c *Compiler) LoadPrelude() {}
