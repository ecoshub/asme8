start:
	xor c, 0x10
	and c, b
	and c, 0x10
	or c, b
	or c, 0x10
	not b
	shl b
	shr b
	rol b
	ror b
