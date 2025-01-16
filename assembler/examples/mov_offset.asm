ADDR=0x90ff
TEST_ADDR=0x9100

start:
    mov b, 5
loop:
    mov a, [ADDR+b]
    mov c, [TEST_ADDR+b]
    dec b
    jz done
    jmp loop

done:
    brk
