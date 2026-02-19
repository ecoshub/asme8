
    sti
start:
    mov a, 0x10
    mov a, 0x01
    jmp start

.org 0x1000
    mov b, 0x31
    rti
