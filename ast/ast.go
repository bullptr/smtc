package ast

import (
	"go/ast"

	"github.com/smtx/parser"
)

type SourceFile struct {
	// Fset     *token.FileSet // Compiler.Fset ref
	Filename string
	Ast      *ast.File
	Parser   *parser.Parser
}

type (
	File      = ast.File
	Ident     = ast.Ident
	Decl      = ast.Decl
	FuncDecl  = ast.FuncDecl
	FuncType  = ast.FuncType
	FieldList = ast.FieldList
	BlockStmt = ast.BlockStmt
	Stmt      = ast.Stmt
	ExprStmt  = ast.ExprStmt
	CallExpr  = ast.CallExpr
)
