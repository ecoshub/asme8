grammar AsmE8;

prog: instruction EOF;

instruction: line+ LINE_COMMENT? ;

line: 
	variable
	| ('\t' | '    ' | '  ' ) inst
    | label
    | directives
    | '\n'
    ;

label: tag  ':';

inst: 
	inst_reg_reg
	| inst_reg_inm
	| inst_reg_inm_variable
	| inst_single_reg
	| inst_single_inm
	| inst_single_tag 
	| inst_single
	| inst_ptr_reg
    | inst_reg_ptr
	| inst_reg_ptr_offset
	| inst_ptr_offset_reg
    ;

inst_reg_reg: mnemonic ' ' reg  ', ' reg;

inst_reg_inm: mnemonic ' ' reg  ', ' inm;

inst_reg_inm_variable: mnemonic ' ' reg  ', ' tag;

inst_ptr_reg: mnemonic ' ' ptr ', ' reg;

inst_reg_ptr: mnemonic ' ' reg ', ' ptr;

inst_reg_ptr_offset: mnemonic ' ' reg ', ' ptr_offset;

inst_ptr_offset_reg: mnemonic ' ' ptr_offset ', ' reg;

inst_single_reg: mnemonic ' ' reg;

inst_single_inm: mnemonic ' ' inm;

inst_single_tag: mnemonic ' ' tag;

inst_single: mnemonic;

mnemonic: 
	'mov'
	| 'add'
	| 'adc'
	| 'sub'
	| 'sbb'
	| 'cmp'
	| 'inc'
	| 'dec'
	| 'jmp'
	| 'jz'
	| 'jnz'
	| 'jsr'
	| 'rts'
	| 'brk'
	| 'nop'
	| 'push'
	| 'pop'
	;

reg: 
	'a' 
	| 'b' 
	| 'c'
	| 'd'
	;

ptr: 
	'[' WHITE_SPACE? inm WHITE_SPACE? ']'
	| '[' WHITE_SPACE? tag WHITE_SPACE? ']'
	;

ptr_offset: 
	'[' WHITE_SPACE? inm WHITE_SPACE? '+' WHITE_SPACE? reg WHITE_SPACE? ']'
	| '[' WHITE_SPACE? tag WHITE_SPACE? '+' WHITE_SPACE? reg WHITE_SPACE? ']'
	;

variable: tag  WHITE_SPACE? '=' WHITE_SPACE? inm;

directives: '.org ' inm_list
	| '.byte '  inm_list
	| '.word '  inm_list
	;

inm_list: inm ((', '| ',') (inm | tag ))*
	;

inm: INT 
	| HEX 
	| BINARY 
	| '\'' CHAR '\'' 
	| '\' \''
	;

tag: STR;

LINE_COMMENT: WHITE_SPACE* ';' ~[\r\n]* -> skip;

INT: [0-9]+;
CHAR: ~[\r\n];
STR: [a-zA-Z_][a-zA-Z0-9_]*;
HEX: '0x' ([a-fA-F0-9])+;
BINARY: '0b' ([0-1])+;
WHITE_SPACE: [ ]+;
