start:
    mov a, 5
    mov b, 6    ; 41 01 06
    cmp a, b    ; 3a 10
    jz l1       ; 15 0e 00
    jmp l2      ; 10 1a 00

l1:
    mov a, 1    ; 41 00 01
    jmp end     ; 10 1a 00

l2:
    mov a, 2    ; 41 00 02
    jmp end     ; 10 1a 00

end:
    brk         ; 00

