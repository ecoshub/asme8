__SERIAL_START__=0xffed
ADDR_PUT_CHAR=__SERIAL_START__

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
    mov [ADDR_PUT_CHAR], a
    ret

done:
    brk

message:
.asciiz "Hello, World!"
