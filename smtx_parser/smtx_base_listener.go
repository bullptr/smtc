// Code generated from /home/baeyun/smtx/smtx_parser/SMTX.g4 by ANTLR 4.13.1. DO NOT EDIT.

package smtx_parser // SMTX
import "github.com/antlr4-go/antlr/v4"

// BaseSMTXListener is a complete listener for a parse tree produced by SMTXParser.
type BaseSMTXListener struct{}

var _ SMTXListener = &BaseSMTXListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSMTXListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSMTXListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSMTXListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSMTXListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseSMTXListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseSMTXListener) ExitScript(ctx *ScriptContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseSMTXListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseSMTXListener) ExitCommand(ctx *CommandContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseSMTXListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseSMTXListener) ExitSymbol(ctx *SymbolContext) {}

// EnterSort is called when production sort is entered.
func (s *BaseSMTXListener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BaseSMTXListener) ExitSort(ctx *SortContext) {}

// EnterSort_dec is called when production sort_dec is entered.
func (s *BaseSMTXListener) EnterSort_dec(ctx *Sort_decContext) {}

// ExitSort_dec is called when production sort_dec is exited.
func (s *BaseSMTXListener) ExitSort_dec(ctx *Sort_decContext) {}

// EnterSelector_dec is called when production selector_dec is entered.
func (s *BaseSMTXListener) EnterSelector_dec(ctx *Selector_decContext) {}

// ExitSelector_dec is called when production selector_dec is exited.
func (s *BaseSMTXListener) ExitSelector_dec(ctx *Selector_decContext) {}

// EnterConstructor_dec is called when production constructor_dec is entered.
func (s *BaseSMTXListener) EnterConstructor_dec(ctx *Constructor_decContext) {}

// ExitConstructor_dec is called when production constructor_dec is exited.
func (s *BaseSMTXListener) ExitConstructor_dec(ctx *Constructor_decContext) {}

// EnterDatatype_dec is called when production datatype_dec is entered.
func (s *BaseSMTXListener) EnterDatatype_dec(ctx *Datatype_decContext) {}

// ExitDatatype_dec is called when production datatype_dec is exited.
func (s *BaseSMTXListener) ExitDatatype_dec(ctx *Datatype_decContext) {}

// EnterFunction_dec is called when production function_dec is entered.
func (s *BaseSMTXListener) EnterFunction_dec(ctx *Function_decContext) {}

// ExitFunction_dec is called when production function_dec is exited.
func (s *BaseSMTXListener) ExitFunction_dec(ctx *Function_decContext) {}

// EnterFunction_def is called when production function_def is entered.
func (s *BaseSMTXListener) EnterFunction_def(ctx *Function_defContext) {}

// ExitFunction_def is called when production function_def is exited.
func (s *BaseSMTXListener) ExitFunction_def(ctx *Function_defContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSMTXListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSMTXListener) ExitTerm(ctx *TermContext) {}

// EnterSpec_constant is called when production spec_constant is entered.
func (s *BaseSMTXListener) EnterSpec_constant(ctx *Spec_constantContext) {}

// ExitSpec_constant is called when production spec_constant is exited.
func (s *BaseSMTXListener) ExitSpec_constant(ctx *Spec_constantContext) {}

// EnterQual_identifier is called when production qual_identifier is entered.
func (s *BaseSMTXListener) EnterQual_identifier(ctx *Qual_identifierContext) {}

// ExitQual_identifier is called when production qual_identifier is exited.
func (s *BaseSMTXListener) ExitQual_identifier(ctx *Qual_identifierContext) {}

// EnterVar_binding is called when production var_binding is entered.
func (s *BaseSMTXListener) EnterVar_binding(ctx *Var_bindingContext) {}

// ExitVar_binding is called when production var_binding is exited.
func (s *BaseSMTXListener) ExitVar_binding(ctx *Var_bindingContext) {}

// EnterSorted_var is called when production sorted_var is entered.
func (s *BaseSMTXListener) EnterSorted_var(ctx *Sorted_varContext) {}

// ExitSorted_var is called when production sorted_var is exited.
func (s *BaseSMTXListener) ExitSorted_var(ctx *Sorted_varContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSMTXListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSMTXListener) ExitPattern(ctx *PatternContext) {}

// EnterSymbol_ is called when production symbol_ is entered.
func (s *BaseSMTXListener) EnterSymbol_(ctx *Symbol_Context) {}

// ExitSymbol_ is called when production symbol_ is exited.
func (s *BaseSMTXListener) ExitSymbol_(ctx *Symbol_Context) {}

// EnterMatch_case is called when production match_case is entered.
func (s *BaseSMTXListener) EnterMatch_case(ctx *Match_caseContext) {}

// ExitMatch_case is called when production match_case is exited.
func (s *BaseSMTXListener) ExitMatch_case(ctx *Match_caseContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSMTXListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSMTXListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseSMTXListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseSMTXListener) ExitIndex(ctx *IndexContext) {}

// EnterAttribute_value is called when production attribute_value is entered.
func (s *BaseSMTXListener) EnterAttribute_value(ctx *Attribute_valueContext) {}

// ExitAttribute_value is called when production attribute_value is exited.
func (s *BaseSMTXListener) ExitAttribute_value(ctx *Attribute_valueContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseSMTXListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseSMTXListener) ExitAttribute(ctx *AttributeContext) {}

// EnterInfo_flag is called when production info_flag is entered.
func (s *BaseSMTXListener) EnterInfo_flag(ctx *Info_flagContext) {}

// ExitInfo_flag is called when production info_flag is exited.
func (s *BaseSMTXListener) ExitInfo_flag(ctx *Info_flagContext) {}

// EnterOption is called when production option is entered.
func (s *BaseSMTXListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseSMTXListener) ExitOption(ctx *OptionContext) {}

// EnterB_value is called when production b_value is entered.
func (s *BaseSMTXListener) EnterB_value(ctx *B_valueContext) {}

// ExitB_value is called when production b_value is exited.
func (s *BaseSMTXListener) ExitB_value(ctx *B_valueContext) {}

// EnterS_expr is called when production s_expr is entered.
func (s *BaseSMTXListener) EnterS_expr(ctx *S_exprContext) {}

// ExitS_expr is called when production s_expr is exited.
func (s *BaseSMTXListener) ExitS_expr(ctx *S_exprContext) {}
