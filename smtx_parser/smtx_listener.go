// Code generated from /home/baeyun/smtx/smtx_parser/SMTX.g4 by ANTLR 4.13.1. DO NOT EDIT.

package smtx_parser // SMTX
import "github.com/antlr4-go/antlr/v4"

// SMTXListener is a complete listener for a parse tree produced by SMTXParser.
type SMTXListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterSort is called when entering the sort production.
	EnterSort(c *SortContext)

	// EnterSort_dec is called when entering the sort_dec production.
	EnterSort_dec(c *Sort_decContext)

	// EnterSelector_dec is called when entering the selector_dec production.
	EnterSelector_dec(c *Selector_decContext)

	// EnterConstructor_dec is called when entering the constructor_dec production.
	EnterConstructor_dec(c *Constructor_decContext)

	// EnterDatatype_dec is called when entering the datatype_dec production.
	EnterDatatype_dec(c *Datatype_decContext)

	// EnterFunction_dec is called when entering the function_dec production.
	EnterFunction_dec(c *Function_decContext)

	// EnterFunction_def is called when entering the function_def production.
	EnterFunction_def(c *Function_defContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterSpec_constant is called when entering the spec_constant production.
	EnterSpec_constant(c *Spec_constantContext)

	// EnterQual_identifier is called when entering the qual_identifier production.
	EnterQual_identifier(c *Qual_identifierContext)

	// EnterVar_binding is called when entering the var_binding production.
	EnterVar_binding(c *Var_bindingContext)

	// EnterSorted_var is called when entering the sorted_var production.
	EnterSorted_var(c *Sorted_varContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterSymbol_ is called when entering the symbol_ production.
	EnterSymbol_(c *Symbol_Context)

	// EnterMatch_case is called when entering the match_case production.
	EnterMatch_case(c *Match_caseContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterAttribute_value is called when entering the attribute_value production.
	EnterAttribute_value(c *Attribute_valueContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterInfo_flag is called when entering the info_flag production.
	EnterInfo_flag(c *Info_flagContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterB_value is called when entering the b_value production.
	EnterB_value(c *B_valueContext)

	// EnterS_expr is called when entering the s_expr production.
	EnterS_expr(c *S_exprContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitSort is called when exiting the sort production.
	ExitSort(c *SortContext)

	// ExitSort_dec is called when exiting the sort_dec production.
	ExitSort_dec(c *Sort_decContext)

	// ExitSelector_dec is called when exiting the selector_dec production.
	ExitSelector_dec(c *Selector_decContext)

	// ExitConstructor_dec is called when exiting the constructor_dec production.
	ExitConstructor_dec(c *Constructor_decContext)

	// ExitDatatype_dec is called when exiting the datatype_dec production.
	ExitDatatype_dec(c *Datatype_decContext)

	// ExitFunction_dec is called when exiting the function_dec production.
	ExitFunction_dec(c *Function_decContext)

	// ExitFunction_def is called when exiting the function_def production.
	ExitFunction_def(c *Function_defContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitSpec_constant is called when exiting the spec_constant production.
	ExitSpec_constant(c *Spec_constantContext)

	// ExitQual_identifier is called when exiting the qual_identifier production.
	ExitQual_identifier(c *Qual_identifierContext)

	// ExitVar_binding is called when exiting the var_binding production.
	ExitVar_binding(c *Var_bindingContext)

	// ExitSorted_var is called when exiting the sorted_var production.
	ExitSorted_var(c *Sorted_varContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitSymbol_ is called when exiting the symbol_ production.
	ExitSymbol_(c *Symbol_Context)

	// ExitMatch_case is called when exiting the match_case production.
	ExitMatch_case(c *Match_caseContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitAttribute_value is called when exiting the attribute_value production.
	ExitAttribute_value(c *Attribute_valueContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitInfo_flag is called when exiting the info_flag production.
	ExitInfo_flag(c *Info_flagContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitB_value is called when exiting the b_value production.
	ExitB_value(c *B_valueContext)

	// ExitS_expr is called when exiting the s_expr production.
	ExitS_expr(c *S_exprContext)
}
