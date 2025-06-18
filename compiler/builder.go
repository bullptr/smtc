package compiler

import (
	"fmt"
	"go/token"

	ts "github.com/tree-sitter/go-tree-sitter"

	"github.com/smtc/ast"
	"github.com/smtc/utils"
)

// @TODO: use: https://pkg.go.dev/golang.org/x/tools/go/ast/astutil

// func BuildFileSetFromParser(fset *token.FileSet, filename string, parser *parser.Parser) {
// 	file := fset.AddFile(filename, -1, parser.Tree.GetTokenStream().Size())

// 	// Process all tokens from ANTLR
// 	stream := parser.Tree.GetTokenStream()
// 	// stream.Fill()
// 	println("stream size:", stream.Size())
// 	for i := range stream.Size() {
// 		token := stream.Get(i)
// 		if token.GetChannel() != antlr.TokenDefaultChannel {
// 			continue
// 		}

// 		// pos := file.Pos(token.GetStart())
// 		// end := file.Pos(token.GetStop() + 1)
// 		_pos := token.GetStart()
// 		_end := token.GetStop() + 1

// 		// Create token positions
// 		file.AddLine(int(_pos))

// 		// println("token:", token.GetText(), "pos:", pos, "end:", end)
// 		// println("token:", token.GetText(), "pos:", _pos, "end:", _end)

// 		// Check for newline in token text instead of specific token type
// 		if text := token.GetText(); text == "\n" || text == "\r\n" {
// 			// if token.GetTokenType() == parser.NEWLINE {
// 			file.AddLine(int(_end))
// 			println("newline token: \\n", "pos:", _pos, "end:", _end)
// 		}

// 	}
// }

// func BuildSourceFile(filename string) *ast.SourceFile {
// 	src := utils.ReadFileBytes(filename)
// 	sf := &ast.SourceFile{
// 		Src: src,
// 		// Parser: parser.NewParser(src),
// 	}
// 	return sf
// }

// func BuildSourceFileList(filenames []*string) []*ast.SourceFile {
// 	var sources []*ast.SourceFile
// 	for _, filename := range filenames {
// 		r, err := os.Open(*filename)
// 		if err != nil {
// 			fmt.Printf("Error opening file %s: %v\n", *filename, err)
// 			continue
// 		}
// 		defer r.Close()

// 		sources = append(sources, BuildSourceFile(*filename))
// 	}
// 	return sources
// }

// func BuildGoSourceFile(filename string) *ast.SourceFile {
// 	fset := token.NewFileSet()
// 	f, err := p.ParseFile(fset, filename, nil, 0)
// 	if err != nil {
// 		panic(fmt.Sprintf("Error parsing file %s: %v", filename, err))
// 	}

// 	sf := BuildSourceFile(filename)
// 	sf.Fset = fset
// 	sf.Ast = f

// 	return sf
// }

func BuildMainFunction(stmts *[]ast.Stmt) *ast.FuncDecl {
	if stmts == nil {
		stmts = &[]ast.Stmt{}
	}

	return &ast.FuncDecl{
		Name: &ast.Ident{Name: "main"},
		Type: &ast.FuncType{Params: &ast.FieldList{}},
		Body: &ast.BlockStmt{
			List: *stmts,
		},
	}
}

func BuildCommands(src *[]byte, root *ts.Node) []ast.Stmt {
	var stmts []ast.Stmt
	cursor := root.Walk()

	defer cursor.Close()

	for _, node := range root.Children(cursor) {
		// println(parser.GetSrcByRange(*src, node.NamedChild(0)))
		if namedChild := node.NamedChild(0); namedChild != nil {
			switch namedChild.Kind() {
			case "symbol":
				// stmts = append(stmts, BuildCallExpr(&node))
				// println(namedChild.Kind())
			case "term":
				// println(namedChild.NamedChild(0).Kind())
				// println(parser.GetSrcByRange(*src, namedChild))
				// println(parser.GetSrcByRange(*src, namedChild.Child(0)))
				// println(parser.GetSrcByRange(*src, namedChild.NamedChild(0)))
				// stmts = append(stmts, BuildCallExpr(&node))
			default:
				fmt.Printf("Unknown node: %s\n", namedChild.Kind())
			}
		} else {
			// @TODO: test: unnamed, probably a call expression
			stmts = append(stmts, BuildCallExpr(&node))
		}
	}

	return stmts
}

func BuildCallExpr(node *ts.Node) ast.Stmt {
	funcExprName := utils.ToPascalCase(node.Child(1).Kind())

	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{
				Name:    funcExprName,
				NamePos: token.Pos(node.StartByte()),
			},
		},
	}
}
