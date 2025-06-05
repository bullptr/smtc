package smtx_parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const (
	TOKEN_COLOR = "darkgoldenrod1"
	RULE_COLOR  = "lightblue"
)

type GraphvizVisitor struct {
	*antlr.BaseParseTreeVisitor
	nodes      []string
	edges      []string
	ruleNames  []string
	tokenNames []string
	counter    int
}

func NewGraphvizVisitor(ruleNames, tokenNames []string) *GraphvizVisitor {
	return &GraphvizVisitor{
		nodes:      make([]string, 0),
		edges:      make([]string, 0),
		ruleNames:  ruleNames,
		tokenNames: tokenNames,
		counter:    0,
	}
}

func (v *GraphvizVisitor) Visit(tree antlr.ParseTree) interface{} {
	nodeId := v.counter
	v.counter++

	switch node := tree.(type) {
	case antlr.TerminalNode:
		// tokenName := v.tokenNames[node.GetSymbol().GetTokenIndex()]
		// tokenType := node.GetSymbol().GetTokenType()
		label := node.GetText()
		v.nodes = append(v.nodes, fmt.Sprintf("    node%d [fillcolor=%s, label=\"%s\"]",
			nodeId, TOKEN_COLOR, label))
	case antlr.RuleNode:
		label := v.ruleNames[node.GetRuleContext().GetRuleIndex()]
		v.nodes = append(v.nodes, fmt.Sprintf("    node%d [fillcolor=%s, label=\"%s\"]",
			nodeId, RULE_COLOR, label))
	}

	for i := range tree.GetChildCount() {
		child := tree.GetChild(i)
		childId := v.counter
		if parseTreeChild, ok := child.(antlr.ParseTree); ok {
			v.Visit(parseTreeChild)
			v.edges = append(v.edges, fmt.Sprintf("    node%d -> node%d",
				nodeId, childId))
		}
	}

	return nil
}

// func (v *GraphvizVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
// 	nodeId := v.counter
// 	v.counter++

// 	// Create node label with token type and text
// 	label := fmt.Sprintf("%d\\n'%s'", node.GetSymbol().GetTokenType(),
// 		strings.ReplaceAll(node.GetText(), "\"", "\\\""))

// 	// Add node to graph
// 	v.nodes = append(v.nodes, fmt.Sprintf("    node%d [label=\"%s\", style=filled, fillcolor=%s]",
// 		nodeId, label, TOKEN_COLOR))

// 	return nil
// }

// func (v *GraphvizVisitor) EnterEveryRuleNode(node antlr.RuleNode) {
// 	nodeId := v.counter
// 	v.counter++

// 	// Create node label with rule name
// 	label := fmt.Sprintf("%d", node.GetRuleContext().GetRuleIndex())

// 	// Add node to graph
// 	v.nodes = append(v.nodes, fmt.Sprintf("    node%d [label=\"%s\", style=filled, fillcolor=%s]",
// 		nodeId, label, RULE_COLOR))
// }

func (v *GraphvizVisitor) GetDOT() string {
	var sb strings.Builder
	sb.WriteString("digraph ParseTree {\n")
	sb.WriteString("    node [shape=box, style=\"rounded, filled\", fontname=\"Arial\"];\n")
	sb.WriteString("    edge [arrowsize=0.5];\n\n")

	// Add all nodes
	sb.WriteString("    // Nodes\n")
	for _, node := range v.nodes {
		sb.WriteString(node + "\n")
	}
	sb.WriteString("\n")

	// Add all edges
	sb.WriteString("    // Edges\n")
	for _, edge := range v.edges {
		sb.WriteString(edge + "\n")
	}

	sb.WriteString("}")
	return sb.String()
}
