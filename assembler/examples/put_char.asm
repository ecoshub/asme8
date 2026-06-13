__SERIAL_START__=0xffed
ADDR_PUT_CHAR=__SERIAL_START__

start:
    mov a, 'A'
    jmp char_out
loop:
    inc a
char_out:
    mov [ADDR_PUT_CHAR], a
    cmp a, 'Z'
    jz done
    jmp loop

done:
    jmp start
