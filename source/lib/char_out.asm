.segment "SEG_CHAR_OUT"

    extern ADDR_CHAR_OUT
    global CHAR_OUT

CHAR_OUT:
    mov [ADDR_CHAR_OUT+b], a
    rts
