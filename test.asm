__REGION_RAM__=0x8000

.org 0x0100
start:
	mov a, 0xff
    mov [INPUT_NUMBER], a
	mov a, 0xee
    mov [INPUT_NUMBER+1], a
	mov a, 0xdd
    mov [INPUT_BUFFER], a
	mov a, 0xcc
    mov [INPUT_BUFFER+1], a
	mov a, 0xbb
    mov [INPUT_BUFFER+2], a
	mov a, 0xaa
    mov [INPUT_BUFFER+3], a
end:
    brk

.org 0xff00
init:
    jmp start

.org __REGION_RAM__
INPUT_NUMBER: .resb 2
INPUT_BUFFER: .resb 4
