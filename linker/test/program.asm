.segment "SEG_PRG"

    global __START__
    extern ADDR_PUT_CHAR

__START__:
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
    hlt

message:
.asciiz "Hello, World!"
