CHAR_OUT_ADDR = 0x7000
CHAR_RDY_ADDR = 0x7100
CHAR_READ_ADDR = 0x7101

start:
    mov c, 0                    ; char index for screen
char_wait:
    mov a, [CHAR_RDY_ADDR]      ; char ready addr
    cmp a, 1                    ; char ready mean 1
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [CHAR_READ_ADDR]     ; read from char addr
    cmp a, 0x7f                 ; check if char is 'del'
    jz key_del                  ; jump to key_del label if char 'del'
    jmp char_out                ; jump to char out routine

key_del:
    cmp c, 1                    ; check if is there char to delete
    jz char_wait                ; if not jump to wait
    sub c, 1                    ; dec cursor index by 1
    mov a, 32                   ; store space char in to a
    mov [CHAR_OUT_ADDR+c], a    ; write screen buffer to a with offset c
    jmp char_wait               ; jump to char wait

char_out:
    mov [CHAR_OUT_ADDR+c], a    ; write content of a in to screen buffer with offset c
    add c, 1                    ; increment cursor index by 1
    jmp char_wait               ; jump to char out
