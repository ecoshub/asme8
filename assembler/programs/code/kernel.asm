__START_ADDR__=0xfffc

ADDR_CHAR_OUT=0x7000
ADDR_CHAR_RDY=0x6ffe
ADDR_CHAR_READ=0x6fff

__CHAR_OUT__:
.include "char_out.asm"

__CHAR_IN__:
.include "char_in.asm"

; start running
.include "start.asm"
