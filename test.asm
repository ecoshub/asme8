ADDR_LOW=0x8000

start:
    mov a, 0x10
    mov [ADDR_LOW], a
    mov a, 0x20
    mov [ADDR_LOW+1], a
    mov a, 0x30
    mov [ADDR_LOW+2], a
    mov a, 0x40
    mov [ADDR_LOW+3], a
    mov a, 0x50
    mov [ADDR_LOW+4], a

    mov ip, ADDR_LOW
    xor b, b
loop:
    cmp b, 5
    jz done
    mov a, [ip+b] 
    inc b
    jmp loop

done:
    brk
