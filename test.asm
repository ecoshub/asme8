start:
	mov ipl, 0x20
	mov iph, 0x80
	mov b, 0x10
	mov c, 0x5
	mov [ip+c], b
	mov a, [ip+c]
