start:
    mov a, 0
    mov c, 10
loop:
    add a, 1
    cmp a, c
    jz end
    jmp loop

end:
    brk
