; capabilities

; read from address
;   1 (enter)
;   4a (enter)
;   700 (enter)
;   48ef (enter)
; read range
;   1.5
;   4a.ff (enter)
;   700.2000 (enter)
;   70f3.e0a9 (enter)

; write to address
;   1:84 (enter)
;   4a:3f (enter)
;   700:9d (enter)
;   48ef:1e (enter)

; run (last read address)
; r (enter)


.segment "SEG_WOZ"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    extern __RETURN__
    extern CONVERTER_BUFFER
    extern WOZMAN_BUFFER
    extern __STR_CONV_HEX__
    extern __HEX_CONV_STR__
    extern CHAR_NOT_VALID
    global WOZMAN


KEY_CODE_ENTER=0x0d


WOZMAN:
    mov a, '>'
    call __PUT_CHAR__
    mov a, ' '
    call __PUT_CHAR__
    xor d, d                ; buffer index
main_loop:
    call __GET_CHAR__
    cmp a, KEY_CODE_ENTER
    jz execute
    call __PUT_CHAR__
    mov [WOZMAN_BUFFER+d], a
    inc d
    jmp main_loop

execute:
    mov ip, WOZMAN_BUFFER
    mov a, d
    call __STR_CONV_HEX__
    cmp a, CHAR_NOT_VALID
    jz invalid_hex_char
    call __RETURN__
    mov a, ':'
    call __PUT_CHAR__
    mov a, ' '
    call __PUT_CHAR__
    mov ipl, [CONVERTER_BUFFER]
    mov iph, [CONVERTER_BUFFER+1]
    mov a, [ip]
    call __HEX_CONV_STR__
    mov a, [CONVERTER_BUFFER]
    call __PUT_CHAR__
    mov a, [CONVERTER_BUFFER+1]
    call __PUT_CHAR__
    call __RETURN__
    jmp WOZMAN

invalid_hex_char:
    mov a, '?'
    call __PUT_CHAR__
    call __RETURN__
    jmp WOZMAN
