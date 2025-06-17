package ast

import (
	"go/ast"
	"go/token"

	"github.com/smtx/parser"
)

type SourceFile struct {
	Src []byte
	// reference to Compiler.Fset
	Fset   *token.FileSet
	Ast    *ast.File
	Parser *parser.Parser
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
