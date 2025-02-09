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

MODE_READ=0x00
MODE_READ_RANGE=0x01
MODE_WRITE=0x02
MODE_WRITE_RANGE=0x03
MODE_RUN=0x04

MODE=WOZMAN_BUFFER
CHAR_BUFFER=WOZMAN_BUFFER+4

WOZMAN:
    mov a, '>'                          ; print terminal start chars '> '
    call __PUT_CHAR__
    mov a, ' '
    call __PUT_CHAR__

    xor d, d                            ; clear buffer index
main_loop:
    call __GET_CHAR__                   ; read char from stdin
    cmp a, KEY_CODE_ENTER               ; cmp it with key enter
    jz execute                          ; if it is enter jump execute
    call __PUT_CHAR__                   ; print the char (echo)
    mov [CHAR_BUFFER+d], a              ; add char in to buffer
    inc d                               ; increment index
    jmp main_loop                       ; start read next char

execute:
    mov a, [MODE]                       ; get mode in to a
    cmp a, MODE_READ                    ; compare mode with MODE_READ
    jz execute_read                     ; if it is jump to execute_read
    jmp main_loop                       ; start read next char

execute_read:
    mov ip, CHAR_BUFFER                 ; move buffer in to ip
    mov a, d                            ; move buffer size in to a
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 
    cmp a, CHAR_NOT_VALID               ; if it returns with CHAR_NOT_VALID
    jz operation_failed                 ; jump to operation failed
    call __RETURN__                     ; print return char
    mov a, ':'                          
    call __PUT_CHAR__                   ; print column
    mov a, ' '
    call __PUT_CHAR__                   ; print space
    mov ipl, [CONVERTER_BUFFER]         ; get converted number low part from converter buffer
    mov iph, [CONVERTER_BUFFER+1]       ; get converted number high part from converter buffer
    mov a, [ip]                         ; read from address that returned from str_to_hex subroutine
    call __HEX_CONV_STR__               ; convert number to chars
    mov a, [CONVERTER_BUFFER]           ; get converted first digit in to a
    call __PUT_CHAR__                   ; print first char
    mov a, [CONVERTER_BUFFER+1]         ; get converted second digit in to a
    call __PUT_CHAR__                   ; print second char
    call __RETURN__                     ; print return char
    jmp WOZMAN                          ; start over

operation_failed:                       ; print '?\n' and start over
    mov a, '?'
    call __PUT_CHAR__
    call __RETURN__
    jmp WOZMAN
