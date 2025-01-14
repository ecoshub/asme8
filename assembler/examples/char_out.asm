CHAR_OUT_ADDR = 0x7000

start:
    mov a, 'A'
    jmp char_out
loop:
    inc a
char_out:
    mov [CHAR_OUT_ADDR], a
    cmp a, 'Z'
    jz done
    jmp loop

done:
    brk
