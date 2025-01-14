VAL = 0x31          ; test value
ADDR = 0x80ff       ; test addr

start:
    mov a, VAL
    jmp ADDR
