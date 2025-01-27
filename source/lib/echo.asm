.segment "SEG_ECHO"

    extern CHAR_IN
    extern CHAR_OUT
    extern __MEM_RAM_START__

CHAR_DEL=0x7f
CHAR_SPACE=' '

; test section
; linker global plus optional offset test
    mov b, 0x20                     ; test value
    mov [__MEM_RAM_START__+1], b    ; write mem start + 1 

    xor b, b                        ; clear input index
start:
    jsr CHAR_IN                     ; read from stdin
    cmp a, CHAR_DEL                 ; check if char is 'del'
    jz key_del                      ; jump to key_del label if char 'del'
    jsr CHAR_OUT                    ; print char to screen
    inc b                           ; increment input index
    jmp start                       ; start over

key_del:
    cmp b, 0                        ; check if is there char to delete
    jz start                        ; if not jump to wait
    dec b                           ; dec cursor index by 1
    mov a, CHAR_SPACE               ; store space char in to a
    jsr CHAR_OUT                    ; print space char to screen
    jmp start                       ; start over
