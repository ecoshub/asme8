ADDR_PUT_CHAR=0xffec
ADDR_CHAR_RDY=0xffed
ADDR_CHAR_READ=0xffee

start:
    mov c, 0                    ; char index for screen
char_wait:
    mov a, [ADDR_CHAR_RDY]      ; char ready addr
    cmp a, 1                    ; char ready mean 1
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [ADDR_CHAR_READ]     ; read from char addr
    jmp char_out                ; jump to char out routine

char_out:
    mov [ADDR_PUT_CHAR], a      ; write content of a in to screen buffer with offset c
    inc c                       ; increment cursor index by 1
    jmp char_wait               ; jump to char out
