
    global __SERIAL_START__

ADDR_PUT_CHAR=__SERIAL_START__

    jmp start

start:
    mov a, 'a'
    call char_out
    brk

char_out:
    mov [ADDR_PUT_CHAR], a
    ret

