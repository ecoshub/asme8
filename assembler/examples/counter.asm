start:
    mov a, 10
loop:
    dec a
    jz end
    jmp loop
end:
    brk
