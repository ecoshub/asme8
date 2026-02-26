.segment "main"

    extern upper
    extern put_char

start:
    mov ip, message
    xor b, b
loop:
    mov a, [ip+b]
    cmp a, 0
    jz done
    call upper
    call put_char
    inc b
    jmp loop

done:
    hlt

message:
.asciiz "hello"
