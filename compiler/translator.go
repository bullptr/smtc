package compiler

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/smtc/ast"
	"github.com/smtc/parser"
)

/**
 * Builders for compiler components
 */

func BuildCommands(p *parser.Parser) []ast.Stmt {
	stmts := []ast.Stmt{}

	// Create new command visitor
	visitor := NewCommandVisitor()
	visitor.Visit(p.Tree.Script())

	// Walk through parse tree
	// antlr.NewParseTreeWalker().Walk(visitor, p.Tree)
	// p.Tree.AddParseListener(visitor)

	// Check for any errors during visit
	if len(visitor.errors) > 0 {
		// Handle errors appropriately
		return nil
	}

	return stmts
}

// CommandVisitor implements the BaseSMTCListener interface
type CommandVisitor struct {
	*antlr.BaseParseTreeVisitor
	statements *[]ast.Stmt
	errors     []error
}

// NewCommandVisitor creates a new CommandVisitor instance
func NewCommandVisitor() *CommandVisitor {
	return &CommandVisitor{
		statements: &[]ast.Stmt{},
		errors:     []error{},
	}
}

// Visit implementation for command nodes
func (v *CommandVisitor) Visit(tree antlr.ParseTree) {
	ctx, ok := tree.(*parser.ScriptContext)
	if !ok {
		v.errors = append(v.errors, fmt.Errorf("expected ScriptContext, got %T", tree))
		return
	}
	if ctx == nil {
		v.errors = append(v.errors, fmt.Errorf("nil ScriptContext"))
		return
	}

	switch ctx.GetChild(1).(type) {
	case *antlr.TerminalNodeImpl:
		cmd := ctx.GetChild(1).GetPayload().(*antlr.TerminalNodeImpl).GetText()
		println("Command:", cmd)
		// switch cmd {
		// case "check-sat":
		// 	*v.statements = append(*v.statements, &ast.CheckSatStmt{})
		// case "declare-const":
		// 	symbol := ctx.Symbol().GetText()
		// 	sort := ctx.Sort().GetText()
		// 	*v.statements = append(*v.statements, &ast.DeclareConstStmt{
		// 		Symbol: symbol,
		// 		Sort:   sort,
		// 	})
		// case "declare-fun":
		// 	symbol := ctx.Symbol().GetText()
		// 	sorts := []string{}
		// 	for _, s := range ctx.AllSort() {
		// 		sorts = append(sorts, s.GetText())
		// 	}
		// 	returnSort := sorts[len(sorts)-1]
		// 	paramSorts := sorts[:len(sorts)-1]
		// 	*v.statements = append(*v.statements, &ast.DeclareFunStmt{
		// 		Symbol:     symbol,
		// 		ParamSorts: paramSorts,
		// 		ReturnSort: returnSort,
		// 	})
		// case "assert":
		// 	term := ctx.Term().GetText()
		// 	*v.statements = append(*v.statements, &ast.AssertStmt{
		// 		Term: term,
		// 	})
		// case "set-logic":
		// 	logic := ctx.Symbol().GetText()
		// 	*v.statements = append(*v.statements, &ast.SetLogicStmt{
		// 		Logic: logic,
		// 	})
		// case "set-option":
		// 	option := ctx.Option().GetText()
		// 	*v.statements = append(*v.statements, &ast.SetOptionStmt{
		// 		Option: option,
		// 	})
		// case "exit":
		// 	*v.statements = append(*v.statements, &ast.ExitStmt{})
		// }
	}
}

// Error handling
func (v *CommandVisitor) VisitErrorNode(node antlr.ErrorNode) {
	v.errors = append(v.errors, fmt.Errorf("error at %v: %v", node.GetSymbol().GetLine(), node.GetText()))
}
