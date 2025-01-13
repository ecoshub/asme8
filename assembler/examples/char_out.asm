start:
    mov a, 65
    jmp char_out
loop:
    add a, 1
char_out:
    cmp a, 91
    jz done
    mov [0x7000], a
    jmp loop

done:
    brk
