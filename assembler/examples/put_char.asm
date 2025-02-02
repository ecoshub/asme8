PUT_CHAR_ADDR=0xf86d

start:
    mov a, 'A'
    jmp char_out
loop:
    inc a
char_out:
    mov [PUT_CHAR_ADDR], a
    cmp a, 'Z'
    jz done
    jmp loop

done:
    jmp start
