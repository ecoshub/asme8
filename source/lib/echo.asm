.segment "SEG_ECHO"

    extern CHAR_IN
    extern CHAR_OUT

CHAR_DEL=0x7f
CHAR_SPACE=' '

    xor b, b
start:
    jsr CHAR_IN
    cmp a, CHAR_DEL             ; check if char is 'del'
    jz key_del                  ; jump to key_del label if char 'del'
    jsr CHAR_OUT
    inc b
    jmp start

key_del:
    cmp b, 0                    ; check if is there char to delete
    jz start                    ; if not jump to wait
    dec b                       ; dec cursor index by 1
    mov a, CHAR_SPACE           ; store space char in to a
    jsr CHAR_OUT
    jmp start
