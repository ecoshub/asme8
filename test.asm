ADDR_LOW=0x8000

start:
    mov a, 0b00110100
    mov [ADDR_LOW], a
    mov b, 0b11101111
    mov [ADDR_LOW+1], b

    mov c, [ADDR_LOW]
    mov d, [ADDR_LOW+1]
    shl c
    rol d
    shl c
    rol d
    shl c
    rol d
    shl c
    rol d
    ; shl c
    ; rol d
    ; shl c
    ; rol d
    ; shl c
    ; rol d
    ; mov [ADDR_LOW+2], c
    ; mov [ADDR_LOW+3], d

    brk
