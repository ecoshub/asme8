VAL=31


start:
    jmp 0x6000


.byte 0x41, 0x00, 0x0a
.byte 0x41, 0x02, 0x0a
ADDR: .byte 0x00

message_1:
.byte 0x4
TEST: .word 0x0faa, 0x0011, 0x1100
    mov a, 0x30
message_2:
.byte 0x5, 0x5
.byte 'E', 0x5, VAL, 'A'
.byte 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'
code:
    mov a, 10

data:
CHAR_ADDR: .byte 0x99
.word 0x0faa
    mov b, 10

bss:
.word 0x11aa
.byte 0x66
    mov d, 10
