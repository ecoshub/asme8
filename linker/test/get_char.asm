.segment "SEG_GET_CHAR"

    extern ADDR_READY_CHAR
    extern ADDR_GET_CHAR
    global __GET_CHAR__

__GET_CHAR__:
    mov a, [ADDR_READY_CHAR]    ; char ready addr
    cmp a, 1                    ; 1 mean char ready
    jz char_read                ; jump to read label if ready
    jmp __GET_CHAR__            ; jump to char wait

char_read:
    mov a, [ADDR_GET_CHAR]     ; read from char addr
    ret
