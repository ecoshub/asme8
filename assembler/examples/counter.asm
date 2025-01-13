start:
    mov a, 0
    mov b, 1

loop:
    add a, b
    cmp a, 5
    jz end
    jmp loop

end:
    brk
