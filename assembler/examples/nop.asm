start:
    mov a, 0
    mov b, 1

loop:
    nop
    nop
    add a, b
    cmp a, 5
    jz end
    jmp loop
    nop

end:
    brk
