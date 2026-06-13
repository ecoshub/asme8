start:
    mov b, 0x8
    mov c, 0x24
    mov ip, 0x2200
    mov [ip+b], c
    mov d, [ip+b]
    hlt
