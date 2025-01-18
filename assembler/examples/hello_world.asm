ADDR_CHAR_OUT=0x7000

start:
    mov b, 0
loop:
    mov a, [message+b]
    cmp a, 0
    jz done
    jsr print_char
    inc b
    jmp loop

print_char:
    mov [ADDR_CHAR_OUT+b], a
    rts

done:
    brk

message:
.byte 'H', 'e', 'l', 'l', 'o', ',', ' ', 'W', 'o', 'r', 'l', 'd', '!', 0
