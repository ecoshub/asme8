grammar AsmE8;

prog: instruction EOF;

instruction: line+ LINE_COMMENT? 
	| line+ SUFFIX_WHITESPACE?
	;

line: 
	variable_reference
	| variable
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
	addr_mode_imm16
	| addr_mode_imm16_tag
	| addr_mode_reg_to_mem_indexed
	| addr_mode_reg_to_mem_direct
	| addr_mode_reg_to_mem_indirect_offset
	| addr_mode_reg_to_mem_reg16_indexed
	| addr_mode_reg_to_mem_indirect
	| addr_mode_implied_reg8
	| addr_mode_reg8_imm8
	| addr_mode_reg8_imm8_tag
	| addr_mode_mem_to_reg_indexed
	| addr_mode_mem_to_reg_direct
	| addr_mode_mem_to_reg_indirect_offset
	| addr_mode_mem_to_reg_reg16_indexed
	| addr_mode_mem_to_reg_indirect
	| addr_mode_reg8_reg8
	| addr_mode_implied
	| addr_mode_implied_reg16
	| addr_mode_reg16_imm
	| addr_mode_reg16_imm_tag
	| addr_mode_reg16_reg16
	| addr_mode_reg16_stack
	| addr_mode_implied_stack
	| addr_mode_stack_imm8
	| addr_mode_stack_imm8_tag
	| addr_mode_stack_reg16
    ;

addr_mode_imm16: mnemonic ' ' imm;
addr_mode_imm16_tag: mnemonic ' ' tag;
addr_mode_reg_to_mem_indexed: mnemonic ' ' direct_offset ', ' reg8;
addr_mode_reg_to_mem_direct: mnemonic ' ' direct ', ' reg8;
addr_mode_reg_to_mem_indirect_offset: mnemonic ' ' indirect_offset ', ' reg8;
addr_mode_reg_to_mem_reg16_indexed: mnemonic ' ' reg16_indexed ', ' reg8;
addr_mode_reg_to_mem_indirect: mnemonic ' ' indirect ', ' reg8;
addr_mode_implied_reg8: mnemonic ' ' reg8;
addr_mode_reg8_imm8: mnemonic ' ' reg8 ', ' imm;
addr_mode_reg8_imm8_tag: mnemonic ' ' reg8 ', ' tag;
addr_mode_mem_to_reg_indexed: mnemonic ' ' reg8 ', ' direct_offset;
addr_mode_mem_to_reg_direct: mnemonic ' ' reg8 ', ' direct;
addr_mode_mem_to_reg_indirect_offset: mnemonic ' ' reg8 ', ' indirect_offset;
addr_mode_mem_to_reg_reg16_indexed: mnemonic ' ' reg8 ', ' reg16_indexed;
addr_mode_mem_to_reg_indirect: mnemonic ' ' reg8 ', ' indirect;
addr_mode_reg8_reg8: mnemonic ' ' reg8 ', ' reg8;
addr_mode_implied: mnemonic;
addr_mode_implied_reg16: mnemonic ' ' reg16;
addr_mode_reg16_imm: mnemonic ' ' reg16 ', ' imm;
addr_mode_reg16_imm_tag: mnemonic ' ' reg16 ', ' tag;
addr_mode_reg16_reg16: mnemonic ' ' reg16 ', ' reg16;
addr_mode_reg16_stack: mnemonic ' ' reg16 ', ' stack ;
addr_mode_implied_stack: mnemonic ' ' stack;
addr_mode_stack_imm8: mnemonic ' ' stack ', ' imm;
addr_mode_stack_imm8_tag: mnemonic ' ' stack ', ' tag;
addr_mode_stack_reg16: mnemonic ' ' stack ', ' reg16;

mnemonic: 
'adc'
| 'add'
| 'and'
| 'brk'
| 'call'
| 'clc'
| 'cmp'
| 'dec'
| 'inc'
| 'jc'
| 'jmp'
| 'jnc'
| 'jns'
| 'jnz'
| 'js'
| 'jz'
| 'mov'
| 'nop'
| 'not'
| 'or'
| 'pop'
| 'push'
| 'ret'
| 'rol'
| 'ror'
| 'rti'
| 'sbb'
| 'shl'
| 'shr'
| 'sub'
| 'xor'
	;

reg8: 
	'a' 
	| 'b' 
	| 'c'
	| 'd'
	| 'e'
	| 'ipl'
	| 'iph'
	| 'bpl'
	| 'bph'
	;

reg16: 
	'ip'
	| 'bp'
	;

stack: 'sp';

direct: 
	'[' imm ']'
	| '[' tag ']'
	| direct_virtual_offset
	;

direct_virtual_offset: '[' STR  ('+' | '-' ) INT ']' ;

direct_offset: 
	'[' imm '+' reg8 ']'
	| '[' tag '+' reg8 ']'
	;

indirect_offset: '[' reg16 '+' imm ']';
reg16_indexed: '[' reg16 '+' reg8 ']';
indirect: '[' reg16 ']';

variable: tag  '=' imm;

variable_reference: tag  '=' reference;

reference: STR
	| STR ('+'|'-') imm
	;

directives: '.asciiz ' ascii
	| '.ascii ' ascii
	| '.byte '  imm_list
	| '.word '  imm_list
	| '.resb '  imm_list
	| '.org '  imm_list
	;

segment: '.segment ' ANY_STR ;

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

ascii: ANY_STR;

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
ANY_STR: '"' ~[\r\n"]* '"';
HEX: '0x' ([a-fA-F0-9])+;
STR: [a-zA-Z_][a-zA-Z0-9_]*;
