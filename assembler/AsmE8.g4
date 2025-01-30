grammar AsmE8;

prog: instruction EOF;

instruction: line+ LINE_COMMENT? 
	| line+ SUFFIX_WHITESPACE?
	;

line: 
	variable
	| ('\t' | '    ' | '  ' ) inst
	| segment
	| access
	| label ' ' directives
	| directives
    | label
    | '\n'
    ;

label: tag  ':';

inst: 
	inst_reg_reg
	| inst_reg_imm
	| inst_stack_imm
	| inst_index_imm
	| inst_reg_imm_variable
	| inst_ptr_reg
	| inst_reg_ptr
	| inst_reg_ptr_offset
	| inst_indirect_reg_stack
	| inst_indirect_stack_register
	| inst_indirect_reg_index
	| inst_indirect_index_register
	| inst_ptr_offset_reg
	| inst_single_reg
	| inst_implied_stack
	| inst_implied_index
	| inst_single_imm
	| inst_single_tag
	| inst_single
    ;

inst_reg_reg: mnemonic ' ' reg  ', ' reg;

inst_reg_imm: mnemonic ' ' reg  ', ' imm;

inst_stack_imm: mnemonic ' ' stack  ', ' imm;

inst_index_imm: mnemonic ' ' index  ', ' imm;

inst_reg_imm_variable: mnemonic ' ' reg  ', ' tag;

inst_ptr_reg: mnemonic ' ' ptr ', ' reg;

inst_reg_ptr: mnemonic ' ' reg ', ' ptr;

inst_reg_ptr_offset: mnemonic ' ' reg ', ' ptr_offset;

inst_indirect_reg_stack: mnemonic ' ' reg ', ' stack_offset;

inst_indirect_stack_register: mnemonic ' ' stack_offset ', ' reg;

inst_indirect_reg_index: mnemonic ' ' reg ', ' index_offset;

inst_indirect_index_register: mnemonic ' ' index_offset ', ' reg;

inst_ptr_offset_reg: mnemonic ' ' ptr_offset ', ' reg;

inst_single_reg: mnemonic ' ' reg;

inst_implied_stack: mnemonic ' ' stack;

inst_implied_index: mnemonic ' ' index;

inst_single_imm: mnemonic ' ' imm;

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
	| 'xor'
	| 'and'
	| 'or'
	| 'not'
	| 'shl'
	| 'shr'
	| 'rol'
	| 'ror'
	| 'jmp'
	| 'jz'
	| 'jnz'
	| 'jc'
	| 'js'
	| 'jns'
	| 'call'
	| 'ret'
	| 'brk'
	| 'nop'
	| 'push'
	| 'pop'
	| 'clc'
	;

reg: 
	'a' 
	| 'b' 
	| 'c'
	| 'd'
	| 'ipl'
	| 'iph'
	;

stack: 'sp';

index: 'ip';

ptr: 
	'[' imm ']'
	| '[' tag ']'
	| ptr_virtual_offset
	;

ptr_virtual_offset: '[' STR  ('+' | '-' ) INT ']' ;

ptr_offset: 
	'[' imm '+' reg ']'
	| '[' tag '+' reg ']'
	;

stack_offset: '[' stack ']'
	| '[' stack '+' imm ']'
	| '[' stack '+' reg ']'
	;

index_offset: '[' index ']'
	| '[' index '+' imm ']'
	| '[' index '+' reg ']'
	;

variable: tag  '=' imm;

directives: '.byte '  imm_list
	| '.word '  imm_list
	| '.resb '  imm_list
	| '.org '  imm_list
	;

segment: '.segment ' '"' STR '"' ;

access: '    ' 'global ' tag
	| '    ' 'extern ' tag
	;

imm_list: (imm | tag) (', ' (imm | tag ))*
	;

imm: INT 
	| HEX 
	| BINARY 
	| CHAR
	;

tag: STR;

LINE_COMMENT: (WHITE_SPACE* ';' ~[\r\n]*) -> skip;
SUFFIX_WHITESPACE: ([ ]* [\r\n]) -> skip;

WHITE_SPACE: [ ]+;
BINARY: '0b' ([0-1])+;
INT: [0-9]+;
CHAR: '\' \'' 
	| '"\'"' 
	| '\',\'' 
	| '\'' ~[\r\n'] '\'' 
	;
HEX: '0x' ([a-fA-F0-9])+;
STR: [a-zA-Z_][a-zA-Z0-9_]*;
