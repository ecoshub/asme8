start:
    mov a, 0xff
    mov [0x3010], a
    mov ipl, 0x10
    mov iph, 0x30
    mov b, [ip]
    jmp start
