start:
	mov a, 0xff
	mov b, 0x02
    add a, b
    jc here
    mov c, 0x20
    brk
here:
    mov c, 0x30
