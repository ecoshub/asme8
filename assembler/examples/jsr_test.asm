start:
    call hello
    mov a, 0x10
    hlt

hello:
    mov a, 0x20
    mov b, 0x20
    ret
    hlt
