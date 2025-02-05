.segment "put_char"

    extern __VIDEO_START__
    global put_char

ADDR_PUT_CHAR=__VIDEO_START__

put_char:
    mov [ADDR_PUT_CHAR+b], a
    ret
