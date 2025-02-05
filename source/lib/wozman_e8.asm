.segment "SEG_WOZ"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    extern __DEL_CHAR__
    extern __STR_CONV_HEX__
    extern __GO_TO_ROW__
    extern __PUT_CURSOR_HOME__
    global WOZMAN

CHAR_DEL=0x7f

WOZMAN:
    mov a, '>'
    mov b, 1
main_loop:
    cmp b, 10
    jz done
    mov a, 'H'
    call __PUT_CHAR__
    mov a, 'e'
    call __PUT_CHAR__
    mov a, 'l'
    call __PUT_CHAR__
    mov a, 'l'
    call __PUT_CHAR__
    mov a, 'o'
    call __PUT_CHAR__
    mov a, b
    call __GO_TO_ROW__
    inc b
    jmp main_loop

    ; mov ip, TEST_BUFFER
    ; mov a, 4
    ; call __STR_CONV_HEX__
    ; brk
done:
    brk

    call __GET_CHAR__
    cmp a, CHAR_DEL         ; check if char is 'del'
    jz delete_char          ; jump to key_del label if char 'del'
    call __PUT_CHAR__
    jmp WOZMAN               ; jump to char out routine

delete_char:
    call __DEL_CHAR__
    jmp WOZMAN

TEST_BUFFER:
.byte 'e', '7', 'a', '8'
