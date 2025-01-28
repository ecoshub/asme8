.segment "SEG_WOZ"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    extern __RAM_START__
    extern STR_TO_HEX
    global WOZMAN

CHAR_DEL=0x7f
CHAR_SPACE=' '

WOZMAN:
    xor b, b                        ; b in index
start:
    jsr __GET_CHAR__                     ; read from stdin
    cmp a, CHAR_DEL                 ; check if char is 'del'
    jz key_del                      ; jump to key_del label if char 'del'
    jsr __PUT_CHAR__                    ; print char to screen
    mov [__RAM_START__+b], a        ; write char to input buffer 
    cmp b, 0x10                     ; if its hit the 16 bytes
    jz clear_index                  ; clear index and start over
    inc b                           ; increment input index
    ; jsr print_as_number
    jmp start                       ; start over

; NOTE: testing
; print_as_number:
;     jsr STR_TO_HEX
;     cmp a, 0
;     jz exit_without_print
;     mov [0x2020], a
; exit_without_print:
;     rts

clear_index:
    xor b, b                        ; clear index
    jmp start                       ; start over

key_del:
    cmp b, 0                        ; check if is there char to delete
    jz rotate_index                 ; if not jump to wait
continue_delete:
    dec b                           ; dec cursor index by 1
    mov a, CHAR_SPACE               ; store space char in to a
    jsr __PUT_CHAR__                    ; print space char to screen
    mov [__RAM_START__+b], a        ; write char to input buffer 
    jmp start                       ; start over

rotate_index:
    mov b, 0x11
    jmp continue_delete
