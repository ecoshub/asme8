start:
    mov a, 0x10
    jnz here
    mov a, 0x20
    brk
here:
    mov a, 0x30
