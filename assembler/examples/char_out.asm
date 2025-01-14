CHAR_OUT_ADDR = 0x7000

start:
    mov a, 65
    jmp char_out
loop:
    add a, 1
char_out:
    cmp a, 91
    jz done
    mov [CHAR_OUT_ADDR], a
    jmp loop

done:
    brk
