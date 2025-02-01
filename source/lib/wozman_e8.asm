.segment "SEG_WOZ"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    ; extern __PRINT__
    extern __DEL_CHAR__
    extern __IRQ__
    global WOZMAN

CHAR_DEL=0x7f

WOZMAN:
    xor c, c                    ; char index for screen
start:
    call __GET_CHAR__
    cmp a, CHAR_DEL             ; check if char is 'del'
    jz delete_char                  ; jump to key_del label if char 'del'
    call __PUT_CHAR__
    jmp start                 ; jump to char out routine

delete_char:
    call __DEL_CHAR__
    jmp start
