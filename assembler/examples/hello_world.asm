ADDR_PUT_CHAR=0xf86d

start:
    mov b, 0
loop:
    mov a, [message+b]
    cmp a, 0
    jz done
    call print_char
    inc b
    jmp loop

print_char:
    mov [ADDR_PUT_CHAR+b], a
    ret

done:
    brk

message:
.asciiz "Hello, World!"
