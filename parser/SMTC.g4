// This grammar defines the syntax for the SMT-LIB v2.7 (v2.7-r2025-04-09) language excluding the
// heories, Logics, and Command responses rules.
grammar SMTC;

// Parser rules
script: command*;

command:
	PAROPEN 'assert' term PARCLOSE
	| PAROPEN 'check-sat' PARCLOSE
	| PAROPEN 'check-sat-assuming' PAROPEN term* PARCLOSE PARCLOSE
	| PAROPEN 'declare-const' symbol sort PARCLOSE
	| PAROPEN 'declare-datatype' symbol datatype_dec PARCLOSE
	| PAROPEN 'declare-datatypes' PAROPEN sort_dec+ PARCLOSE PAROPEN datatype_dec+ PARCLOSE PARCLOSE
	| PAROPEN 'declare-fun' symbol PAROPEN sort* PARCLOSE sort PARCLOSE
	| PAROPEN 'declare-sort' symbol NUMERAL PARCLOSE
	| PAROPEN 'declare-sort-parameter' symbol PARCLOSE
	| PAROPEN 'define-const' symbol sort term PARCLOSE
	| PAROPEN 'define-fun' function_def PARCLOSE
	| PAROPEN 'define-fun-rec' function_def PARCLOSE
	| PAROPEN 'define-funs-rec' PAROPEN function_dec+ PARCLOSE PAROPEN term+ PARCLOSE PARCLOSE
	| PAROPEN 'define-sort' symbol PAROPEN symbol* PARCLOSE sort PARCLOSE
	| PAROPEN 'echo' STRING PARCLOSE
	| PAROPEN 'exit' PARCLOSE
	| PAROPEN 'get-assertions' PARCLOSE
	| PAROPEN 'get-assignment' PARCLOSE
	| PAROPEN 'get-info' info_flag PARCLOSE
	| PAROPEN 'get-model' PARCLOSE
	| PAROPEN 'get-option' KEYWORD PARCLOSE
	| PAROPEN 'get-proof' PARCLOSE
	| PAROPEN 'get-unsat-assumptions' PARCLOSE
	| PAROPEN 'get-unsat-core' PARCLOSE
	| PAROPEN 'get-value' PAROPEN term+ PARCLOSE PARCLOSE
	| PAROPEN 'pop' NUMERAL PARCLOSE
	| PAROPEN 'push' NUMERAL PARCLOSE
	| PAROPEN 'reset' PARCLOSE
	| PAROPEN 'reset-assertions' PARCLOSE
	| PAROPEN 'set-info' attribute PARCLOSE
	| PAROPEN 'set-logic' symbol PARCLOSE
	| PAROPEN 'set-option' option PARCLOSE;

symbol: SIMPLE_SYMBOL | QUOTED_SYMBOL;

sort: identifier | PAROPEN identifier sort+ PARCLOSE;

sort_dec: PAROPEN symbol NUMERAL PARCLOSE;

selector_dec: PAROPEN symbol sort PARCLOSE;

constructor_dec: PAROPEN symbol selector_dec* PARCLOSE;

datatype_dec:
	PAROPEN constructor_dec+ PARCLOSE
	| PAROPEN 'par' PAROPEN symbol+ PARCLOSE PAROPEN constructor_dec+ PARCLOSE PARCLOSE;

function_dec:
	PAROPEN symbol PAROPEN sorted_var* PARCLOSE sort PARCLOSE;

function_def: symbol PAROPEN sorted_var* PARCLOSE sort term;

term:
	spec_constant
	| qual_identifier
	| PAROPEN qual_identifier term+ PARCLOSE
	| PAROPEN 'let' PAROPEN var_binding+ PARCLOSE term PARCLOSE
	| PAROPEN 'lambda' PAROPEN sorted_var+ PARCLOSE term PARCLOSE
	| PAROPEN 'forall' PAROPEN sorted_var+ PARCLOSE term PARCLOSE
	| PAROPEN 'exists' PAROPEN sorted_var+ PARCLOSE term PARCLOSE
	| PAROPEN 'match' term PAROPEN match_case+ PARCLOSE PARCLOSE
	| PAROPEN '!' term attribute+ PARCLOSE;

spec_constant:
	NUMERAL
	| DECIMAL
	| HEXADECIMAL
	| BINARY
	| STRING;

qual_identifier:
	identifier
	| PAROPEN 'as' identifier sort PARCLOSE;

var_binding: PAROPEN symbol term PARCLOSE;

sorted_var: PAROPEN symbol sort PARCLOSE;

pattern: symbol_ | PAROPEN symbol symbol_+ PARCLOSE;

symbol_: symbol | '_';

match_case: PAROPEN pattern term PARCLOSE;

identifier: symbol | PAROPEN '_' symbol index+ PARCLOSE;

index: NUMERAL | symbol;

attribute_value:
	spec_constant
	| symbol
	| PAROPEN s_expr* PARCLOSE;

attribute: KEYWORD | KEYWORD attribute_value;

info_flag:
	':all-statistics'
	| ':assertion-stack-levels'
	| ':authors'
	| ':error-behavior'
	| ':name'
	| ':reason-unknown'
	| ':version'
	| KEYWORD;

option:
	':diagnostic-output-channel' STRING
	| ':global-declarations' b_value
	| ':interactive-mode' b_value
	| ':print-success' b_value
	| ':produce-assertions' b_value
	| ':produce-assignments' b_value
	| ':produce-models' b_value
	| ':produce-proofs' b_value
	| ':produce-unsat-assumptions' b_value
	| ':produce-unsat-cores' b_value
	| ':random-seed' NUMERAL
	| ':regular-output-channel' STRING
	| ':reproducible-resource-limit' NUMERAL
	| ':verbosity' NUMERAL
	| attribute;

b_value: 'true' | 'false';

s_expr:
	spec_constant
	| symbol
	| KEYWORD
	| PAROPEN s_expr* PARCLOSE;

// Lexer rules
WS: [ \t\r\n]+ -> skip;
COMMENT: ';' ~[\r\n]* -> skip;

NUMERAL: '0' | [1-9][0-9]*;
DECIMAL: NUMERAL '.' '0'* NUMERAL;
HEXADECIMAL: '#x' [0-9a-fA-F]+;
BINARY: '#b' [01]+;
STRING: '"' (CHAR_NODQUOTE | CHAR_WHITESPACE)+ '"';

PAROPEN: '(';
PARCLOSE: ')';

SIMPLE_SYMBOL:
	PS_NOT
	| PS_BOOL
	| PS_CONTINUEDEXECUTION
	| PS_ERROR
	| PS_FALSE
	| PS_IMMEDIATEEXIT
	| PS_INCOMPLETE
	| PS_LOGIC
	| PS_MEMOUT
	| PS_SAT
	| PS_SUCCESS
	| PS_THEORY
	| PS_TRUE
	| PS_UNKNOWN
	| PS_UNSUPPORTED
	| PS_UNSAT
	| SYM ([0-9] | SYM)*;

QUOTED_SYMBOL: '|' (CHAR_NOBACKSLASH | CHAR_WHITESPACE)+ '|';

KEYWORD:
	PK_ALL_STATISTICS
	| PK_ASSERTION_STACK_LEVELS
	| PK_AUTHORS
	| PK_CATEGORY
	| PK_CHAINABLE
	| PK_DEFINITION
	| PK_DIAGNOSTIC_OUTPUT_CHANNEL
	| PK_ERROR_BEHAVIOR
	| PK_EXTENSIONS
	| PK_FUNS
	| PK_FUNS_DESCRIPTION
	| PK_GLOBAL_DECLARATIONS
	| PK_INTERACTIVE_MODE
	| PK_LANGUAGE
	| PK_LEFT_ASSOC
	| PK_LICENSE
	| PK_NAME
	| PK_NAMED
	| PK_NOTES
	| PK_PATTERN
	| PK_PRINT_SUCCESS
	| PK_PRODUCE_ASSIGNMENTS
	| PK_PRODUCE_MODELS
	| PK_PRODUCE_PROOFS
	| PK_PRODUCE_UNSAT_ASSUMPTIONS
	| PK_PRODUCE_UNSAT_CORES
	| PK_RANDOM_SEED
	| PK_REASON_UNKNOWN
	| PK_REGULAR_OUTPUT_CHANNEL
	| PK_REPRODUCIBLE_RESOURCE_LIMIT
	| PK_RIGHT_ASSOC
	| PK_SMT_LIB_VERSION
	| PK_SORTS
	| PK_SORTS_DESCRIPTION
	| PK_SOURCE
	| PK_STATUS
	| PK_THEORIES
	| PK_VALUES
	| PK_VERBOSITY
	| PK_VERSION
	| ':' SIMPLE_SYMBOL;

// Predefined Symbols
PS_NOT: 'not';
PS_BOOL: 'Bool';
PS_CONTINUEDEXECUTION: 'continued-execution';
PS_ERROR: 'error';
PS_FALSE: 'false';
PS_IMMEDIATEEXIT: 'immediate-exit';
PS_INCOMPLETE: 'incomplete';
PS_LOGIC: 'logic';
PS_MEMOUT: 'memout';
PS_SAT: 'sat';
PS_SUCCESS: 'success';
PS_THEORY: 'theory';
PS_TRUE: 'true';
PS_UNKNOWN: 'unknown';
PS_UNSUPPORTED: 'unsupported';
PS_UNSAT: 'unsat';

// Predefined Keywords @TODO: is this even necessary?
PK_ALL_STATISTICS: ':all-statistics';
PK_ASSERTION_STACK_LEVELS: ':assertion-stack-levels';
PK_AUTHORS: ':authors';
PK_CATEGORY: ':category';
PK_CHAINABLE: ':chainable';
PK_DEFINITION: ':definition';
PK_DIAGNOSTIC_OUTPUT_CHANNEL: ':diagnostic-output-channel';
PK_ERROR_BEHAVIOR: ':error-behavior';
PK_EXTENSIONS: ':extensions';
PK_FUNS: ':funs';
PK_FUNS_DESCRIPTION: ':funs-description';
PK_GLOBAL_DECLARATIONS: ':global-declarations';
PK_INTERACTIVE_MODE: ':interactive-mode';
PK_LANGUAGE: ':language';
PK_LEFT_ASSOC: ':left-assoc';
PK_LICENSE: ':license';
PK_NAME: ':name';
PK_NAMED: ':named';
PK_NOTES: ':notes';
PK_PATTERN: ':pattern';
PK_PRINT_SUCCESS: ':print-success';
PK_PRODUCE_ASSIGNMENTS: ':produce-assignments';
PK_PRODUCE_MODELS: ':produce-models';
PK_PRODUCE_PROOFS: ':produce-proofs';
PK_PRODUCE_UNSAT_ASSUMPTIONS: ':produce-unsat-assumptions';
PK_PRODUCE_UNSAT_CORES: ':produce-unsat-cores';
PK_RANDOM_SEED: ':random-seed';
PK_REASON_UNKNOWN: ':reason-unknown';
PK_REGULAR_OUTPUT_CHANNEL: ':regular-output-channel';
PK_REPRODUCIBLE_RESOURCE_LIMIT: ':reproducible-resource-limit';
PK_RIGHT_ASSOC: ':right-assoc';
PK_SMT_LIB_VERSION: ':smt-lib-version';
PK_SORTS: ':sorts';
PK_SORTS_DESCRIPTION: ':sorts-description';
PK_SOURCE: ':source';
PK_STATUS: ':status';
PK_THEORIES: ':theories';
PK_VALUES: ':values';
PK_VERBOSITY: ':verbosity';
PK_VERSION: ':version';

fragment SYM:
	'a' ..'z'
	| 'A' .. 'Z'
	| '+'
	| '='
	| '/'
	| '*'
	| '%'
	| '?'
	| '!'
	| '$'
	| '-'
	| '_'
	| '~'
	| '&'
	| '^'
	| '<'
	| '>'
	| '@'
	| '.';

fragment CHAR:
	'\u0020' .. '\u007E'
	| '\u0080' .. '\uffff'
	| ESCAPEDSPACE;

fragment CHAR_NODQUOTE:
	'\u0020' .. '\u0021'
	| '\u0023' .. '\u007E'
	| '\u0080' .. '\uffff'
	| ESCAPEDSPACE;

fragment CHAR_NOBACKSLASH:
	'\u0020' .. '\u005B'
	| '\u005D' .. '\u007B'
	| '\u007D' .. '\u007E'
	| '\u0080' .. '\uffff'
	| ESCAPEDSPACE;

fragment ESCAPEDSPACE: '""';

fragment CHAR_WHITESPACE:
	'\u0009'
	| '\u000A'
	| '\u000D'
	| '\u0020';