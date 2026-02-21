__SERIAL_START__=0xffed
ADDR_PUT_CHAR=__SERIAL_START__

start:
    mov a, 'e'          ; move char value of 'e' to register a
    call char_out       ; print value of register a
    mov a, 'c'          ; move char value of 'c' to register a
    call char_out       ; print value of register a
    mov a, 'o'          ; move char value of 'o' to register a
    call char_out       ; print value of register a
    mov a, '!'          ; move char value of '!' to register a
    call char_out       ; print value of register a
    mov a, 0x0d         ; move value (carriage return) to register a
    call char_out       ; print value of register a
    brk

char_out:
    mov [ADDR_PUT_CHAR], a      ; write content of a in to screen buffer with offset c
    ret

