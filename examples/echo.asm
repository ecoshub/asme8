__SERIAL_START__=0xffec
ADDR_PUT_CHAR=__SERIAL_START__
ADDR_READY_CHAR=__SERIAL_START__+1
ADDR_GET_CHAR=__SERIAL_START__+2

start:
    mov c, 0                    ; char index for screen
char_wait:
    mov a, [ADDR_READY_CHAR]    ; char ready addr
    cmp a, 1                    ; char ready mean 1
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [ADDR_GET_CHAR]      ; read from char addr
    jmp char_out                ; jump to char out routine

char_out:
    mov [ADDR_PUT_CHAR], a      ; write content of a in to screen buffer with offset c
    inc c                       ; increment cursor index by 1
    jmp char_wait               ; jump to char out
