start:
    mov a, 65
    jmp print
loop:
    add a, 1
print:
    cmp a, 91
    jz done
    mov [0x7000], a
    jmp loop

done:
    brk
