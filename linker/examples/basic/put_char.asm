.segment "put_char"

    extern __SERIAL_START__
    global put_char

ADDR_PUT_CHAR=__SERIAL_START__

put_char:
    mov [ADDR_PUT_CHAR+b], a
    ret
