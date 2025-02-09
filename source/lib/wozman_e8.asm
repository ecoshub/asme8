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
KEY_CODE_DOT='.'

MODE_READ=0x00
MODE_READ_RANGE=0x01
MODE_WRITE=0x02
MODE_WRITE_RANGE=0x03
MODE_RUN=0x04

MODE=WOZMAN_BUFFER
SEPARATOR_INDEX=WOZMAN_BUFFER+1
RANGE_START=WOZMAN_BUFFER+2
RANGE_END=WOZMAN_BUFFER+4
CHAR_BUFFER=WOZMAN_BUFFER+6

WOZMAN:
    mov a, MODE_READ                    ; set mode as read
    mov [MODE], a

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
    cmp a, KEY_CODE_DOT
    jz set_mode_read_range
    mov [CHAR_BUFFER+d], a              ; add char in to buffer
    inc d                               ; increment index
    jmp main_loop                       ; start read next char

set_mode_read_range:
    mov a, MODE_READ_RANGE
    mov [MODE], a
    mov [SEPARATOR_INDEX], d
    jmp main_loop

execute:
    mov a, [MODE]                       ; get mode in to a
    cmp a, MODE_READ                    ; compare mode with MODE_READ
    jz execute_read                     ; if it is jump to execute_read
    cmp a, MODE_READ_RANGE
    jz execute_read_range
    jmp main_loop                       ; start read next char

execute_read_range:
    push d
    mov a, [SEPARATOR_INDEX]
    mov ip, CHAR_BUFFER
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 
    cmp a, CHAR_NOT_VALID               ; if it returns with CHAR_NOT_VALID
    jz operation_failed                 ; jump to operation failed

    mov a, [CONVERTER_BUFFER]
    mov [RANGE_START], a
    mov a, [CONVERTER_BUFFER+1]
    mov [RANGE_START+1], a

    mov ip, CHAR_BUFFER
    mov a, [SEPARATOR_INDEX]
    add ipl, a
    adc iph, 0
    pop d
    sub d, a
    mov a, d
    inc a
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 

    mov a, [CONVERTER_BUFFER]
    mov [RANGE_END], a
    mov a, [CONVERTER_BUFFER+1]
    mov [RANGE_END+1], a

    call _print_range

    jmp WOZMAN

_print_range:
    call __RETURN__
    mov ipl, [RANGE_START]
    mov iph, [RANGE_START+1]
    mov b, 0
print_range_loop:
    mov a, [RANGE_END+1]
    cmp a, iph
    js print_range_break
    mov a, [RANGE_END]
    cmp a, ipl
    js print_range_break
    push b
    call _print_byte
    pop b
    mov a, ' '
    call __PUT_CHAR__
    add ipl, 1
    adc iph, 0
    cmp b, 7
    jz print_new_line
    inc b
    jmp print_range_loop

print_range_break:
    mov a, 0
    mov [RANGE_START], a
    mov [RANGE_START+1], a
    mov [RANGE_END], a
    mov [RANGE_END+1], a
    call __RETURN__
    ret

print_new_line:
    call __RETURN__
    mov b, 0
    jmp print_range_loop

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
    call _print_byte
    call __RETURN__                     ; print return char
    jmp WOZMAN                          ; start over

_print_byte:
    mov a, [ip]                         ; read from address that returned from str_to_hex subroutine
    call __HEX_CONV_STR__               ; convert number to chars
    mov a, [CONVERTER_BUFFER]           ; get converted first digit in to a
    call __PUT_CHAR__                   ; print first char
    mov a, [CONVERTER_BUFFER+1]         ; get converted second digit in to a
    call __PUT_CHAR__                   ; print second char
    ret

operation_failed:                       ; print '?\n' and start over
    mov a, '?'
    call __PUT_CHAR__
    call __RETURN__
    jmp WOZMAN
