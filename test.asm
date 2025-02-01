ADDR=0x4000
start:
    mov a, 0x10
    mov [ADDR], a
    mov a, 0x11
    mov [ADDR+1], a
    mov a, 0x12
    mov [ADDR+2], a
    mov a, 0x13
    mov [ADDR+3], a
    mov ip, ADDR
    mov a, [ip+b]
    ; add ip, 1
    ; mov b, [ip]
    ; add ip, 1
    ; mov c, [ip]
    ; add ip, 1
    ; mov d, [ip]
