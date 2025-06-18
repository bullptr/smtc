// Code generated from /home/baeyun/smtc/parser/SMTC.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SMTC
import "github.com/antlr4-go/antlr/v4"

// BaseSMTCListener is a complete listener for a parse tree produced by SMTCParser.
type BaseSMTCListener struct{}

var _ SMTCListener = &BaseSMTCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSMTCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSMTCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSMTCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSMTCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseSMTCListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseSMTCListener) ExitScript(ctx *ScriptContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseSMTCListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseSMTCListener) ExitCommand(ctx *CommandContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseSMTCListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseSMTCListener) ExitSymbol(ctx *SymbolContext) {}

// EnterSort is called when production sort is entered.
func (s *BaseSMTCListener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BaseSMTCListener) ExitSort(ctx *SortContext) {}

// EnterSort_dec is called when production sort_dec is entered.
func (s *BaseSMTCListener) EnterSort_dec(ctx *Sort_decContext) {}

// ExitSort_dec is called when production sort_dec is exited.
func (s *BaseSMTCListener) ExitSort_dec(ctx *Sort_decContext) {}

// EnterSelector_dec is called when production selector_dec is entered.
func (s *BaseSMTCListener) EnterSelector_dec(ctx *Selector_decContext) {}

// ExitSelector_dec is called when production selector_dec is exited.
func (s *BaseSMTCListener) ExitSelector_dec(ctx *Selector_decContext) {}

// EnterConstructor_dec is called when production constructor_dec is entered.
func (s *BaseSMTCListener) EnterConstructor_dec(ctx *Constructor_decContext) {}

// ExitConstructor_dec is called when production constructor_dec is exited.
func (s *BaseSMTCListener) ExitConstructor_dec(ctx *Constructor_decContext) {}

// EnterDatatype_dec is called when production datatype_dec is entered.
func (s *BaseSMTCListener) EnterDatatype_dec(ctx *Datatype_decContext) {}

// ExitDatatype_dec is called when production datatype_dec is exited.
func (s *BaseSMTCListener) ExitDatatype_dec(ctx *Datatype_decContext) {}

// EnterFunction_dec is called when production function_dec is entered.
func (s *BaseSMTCListener) EnterFunction_dec(ctx *Function_decContext) {}

// ExitFunction_dec is called when production function_dec is exited.
func (s *BaseSMTCListener) ExitFunction_dec(ctx *Function_decContext) {}

// EnterFunction_def is called when production function_def is entered.
func (s *BaseSMTCListener) EnterFunction_def(ctx *Function_defContext) {}

// ExitFunction_def is called when production function_def is exited.
func (s *BaseSMTCListener) ExitFunction_def(ctx *Function_defContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSMTCListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSMTCListener) ExitTerm(ctx *TermContext) {}

// EnterSpec_constant is called when production spec_constant is entered.
func (s *BaseSMTCListener) EnterSpec_constant(ctx *Spec_constantContext) {}

// ExitSpec_constant is called when production spec_constant is exited.
func (s *BaseSMTCListener) ExitSpec_constant(ctx *Spec_constantContext) {}

// EnterQual_identifier is called when production qual_identifier is entered.
func (s *BaseSMTCListener) EnterQual_identifier(ctx *Qual_identifierContext) {}

// ExitQual_identifier is called when production qual_identifier is exited.
func (s *BaseSMTCListener) ExitQual_identifier(ctx *Qual_identifierContext) {}

// EnterVar_binding is called when production var_binding is entered.
func (s *BaseSMTCListener) EnterVar_binding(ctx *Var_bindingContext) {}

// ExitVar_binding is called when production var_binding is exited.
func (s *BaseSMTCListener) ExitVar_binding(ctx *Var_bindingContext) {}

// EnterSorted_var is called when production sorted_var is entered.
func (s *BaseSMTCListener) EnterSorted_var(ctx *Sorted_varContext) {}

// ExitSorted_var is called when production sorted_var is exited.
func (s *BaseSMTCListener) ExitSorted_var(ctx *Sorted_varContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSMTCListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSMTCListener) ExitPattern(ctx *PatternContext) {}

// EnterSymbol_ is called when production symbol_ is entered.
func (s *BaseSMTCListener) EnterSymbol_(ctx *Symbol_Context) {}

// ExitSymbol_ is called when production symbol_ is exited.
func (s *BaseSMTCListener) ExitSymbol_(ctx *Symbol_Context) {}

// EnterMatch_case is called when production match_case is entered.
func (s *BaseSMTCListener) EnterMatch_case(ctx *Match_caseContext) {}

// ExitMatch_case is called when production match_case is exited.
func (s *BaseSMTCListener) ExitMatch_case(ctx *Match_caseContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSMTCListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSMTCListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseSMTCListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseSMTCListener) ExitIndex(ctx *IndexContext) {}

// EnterAttribute_value is called when production attribute_value is entered.
func (s *BaseSMTCListener) EnterAttribute_value(ctx *Attribute_valueContext) {}

// ExitAttribute_value is called when production attribute_value is exited.
func (s *BaseSMTCListener) ExitAttribute_value(ctx *Attribute_valueContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseSMTCListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseSMTCListener) ExitAttribute(ctx *AttributeContext) {}

// EnterInfo_flag is called when production info_flag is entered.
func (s *BaseSMTCListener) EnterInfo_flag(ctx *Info_flagContext) {}

// ExitInfo_flag is called when production info_flag is exited.
func (s *BaseSMTCListener) ExitInfo_flag(ctx *Info_flagContext) {}

// EnterOption is called when production option is entered.
func (s *BaseSMTCListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseSMTCListener) ExitOption(ctx *OptionContext) {}

// EnterB_value is called when production b_value is entered.
func (s *BaseSMTCListener) EnterB_value(ctx *B_valueContext) {}

// ExitB_value is called when production b_value is exited.
func (s *BaseSMTCListener) ExitB_value(ctx *B_valueContext) {}

// EnterS_expr is called when production s_expr is entered.
func (s *BaseSMTCListener) EnterS_expr(ctx *S_exprContext) {}

// ExitS_expr is called when production s_expr is exited.
func (s *BaseSMTCListener) ExitS_expr(ctx *S_exprContext) {}
