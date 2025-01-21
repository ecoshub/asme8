start:
    jsr hello
    mov a, 0x10
    brk

hello:
    mov a, 0x20
    mov b, 0x20
    rts
    brk
