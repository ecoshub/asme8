.segment "IO"

ADDR_CHAR_OUT=0x7000
ADDR_CHAR_RDY=0x6ffe
ADDR_CHAR_READ=0x6fff

CHAR_DEL=0x7f
CHAR_SPACE=' '

start:
    mov c, 0                    ; char index for screen
char_wait:
    mov a, [ADDR_CHAR_RDY]      ; char ready addr
    cmp a, 1                    ; char ready mean 1
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [ADDR_CHAR_READ]     ; read from char addr
    cmp a, CHAR_DEL             ; check if char is 'del'
    jz key_del                  ; jump to key_del label if char 'del'
    jmp char_out                ; jump to char out routine

key_del:
    cmp c, 0                    ; check if is there char to delete
    jz char_wait                ; if not jump to wait
    dec c                       ; dec cursor index by 1
    mov a, CHAR_SPACE           ; store space char in to a
    mov [ADDR_CHAR_OUT+c], a    ; write screen buffer to a with offset c
    jmp char_wait               ; jump to char wait

char_out:
    mov [ADDR_CHAR_OUT+c], a    ; write content of a in to screen buffer with offset c
    inc c                       ; increment cursor index by 1
    jmp char_wait               ; jump to char out
