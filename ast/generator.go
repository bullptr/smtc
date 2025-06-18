package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"

	"github.com/smtc/utils"
)

func Generate(fset *token.FileSet, sf *SourceFile) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, sf.Ast)

	return buf.String()
}

func GenerateFile(fset *token.FileSet, sf *SourceFile) {
	sourceCode := Generate(fset, sf)
	err := utils.WriteFileBytes(sf.Filename, []byte(sourceCode))
	if err != nil {
		panic(err)
	}
}

func PrintAst(fset *token.FileSet, Ast *File) {
	utils.PrintSectionHeader("AST")
	if Ast == nil {
		panic("AST is nil")
	}

	ast.Print(fset, Ast)
}

func PrintSourceFile(fset *token.FileSet, sf *SourceFile) {
	utils.PrintSectionHeader("Source File")
	fmt.Println(Generate(fset, sf))
}
