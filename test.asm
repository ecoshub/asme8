start:
    mov a, 0x5
loop:
    jz end
    dec a
    jmp loop
end:
    brk
