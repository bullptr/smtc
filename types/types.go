package types

import (
	"go/importer"
	"go/token"
	"go/types"

	"github.com/smtc/ast"
	"github.com/smtc/config"
)

func CheckSourceFile(fset *token.FileSet, sf *ast.SourceFile) (*types.Package, error) {
	// A Config controls various options of the type checker.
	// The defaults work fine except for one setting:
	// we must specify how to deal with imports.
	conf := types.Config{Importer: importer.Default()}

	// Type-check the package containing only file sf.Ast.
	// Check returns a *types.Package.
	return conf.Check(config.PKGNAME, fset, []*ast.File{sf.Ast}, nil)
}
