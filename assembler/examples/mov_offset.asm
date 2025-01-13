start:
    mov b, 0
loop:
    mov a, [0x90ff+b]
    add b, 1
    cmp b, 5
    jz done
    jmp loop

done:
    brk
