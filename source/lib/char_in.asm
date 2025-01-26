.segment "SEG_CHAR_IN"

    extern ADDR_CHAR_READY
    extern ADDR_CHAR_READ
    global CHAR_IN

    xor a, a
CHAR_IN:
char_wait:
    mov a, [ADDR_CHAR_READY]    ; char ready addr
    cmp a, 1                    ; 1 mean char ready
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [ADDR_CHAR_READ]     ; read from char addr
    rts
