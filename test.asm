    sti
start:
    mov a, 0x10
    mov b, 0x11
    brk

.org 0x1000
    mov b, 0x31
    mov c, 0x32
    rti
