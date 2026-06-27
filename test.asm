ADDR=0x1234
READ=ADDR+1

start:
    mov c, 0xff
    mov b, 0x8
    mov ip, ADDR
    mov [ip+b], c
    mov ip, READ
    mov [ip+b], c
    hlt
