; ----------------------------------------------------------
; Segment Name.: SEG_CONV
; Includes.....: '__C_TO_I__', '__STR_CONV_HEX__'
; Author.......: eco
; Date.........: 28.01.2025
; ----------------------------------------------------------

.segment "SEG_CONV"

    extern CONVERTER_BUFFER
    global CHAR_NOT_VALID
    global __C_TO_I__
    global __I_TO_C__
    global __STR_CONV_HEX__
    global __HEX_CONV_STR__

CHAR_NOT_VALID=0x10


; parameters:
;       ip  -> buffer pointer
;       a   -> buffer length
; return:
;       a   -> high byte of value
;       b   -> low byte of value
__STR_CONV_HEX__:
    xor c, c                        ; buffer index
    mov [CONVERTER_BUFFER], c       ; clear buffer
    mov [CONVERTER_BUFFER+1], c     ; clear buffer

    mov b, a                        ; get length of buffer in b
hex_conv_loop:
    cmp b, 0                        ; final condition
    jz conv_done                    ; conversion done
    mov a, [ip+c]                   ; get char from buffer plus offset
    call __C_TO_I__                 ; run conversion subroutine
    cmp a, CHAR_NOT_VALID           ; check if char is valid
    jz buffer_not_valid             ; return if not
    push c                          ; save c
    mov c, a                        ; get a to c
    mov c, [CONVERTER_BUFFER]       ; get lower part of result
    mov d, [CONVERTER_BUFFER+1]     ; get upper part of result
    shl c                           ; shift lower and upper  (16 bit) 4 times 
    rol d
    shl c
    rol d
    shl c
    rol d
    shl c
    rol d
    or c, a                         ; add the first octet of a to c
    mov [CONVERTER_BUFFER], c       ; store lower part of result
    mov [CONVERTER_BUFFER+1], d     ; store upper part of result
    pop c                           ; restore c
    inc c                           ; increment index
    dec b                           ; decrement process length
    jmp hex_conv_loop               ; start over for next char

conv_done:
    mov a, 0                        ; 0 means everything OK
    ret

buffer_not_valid:
    mov a, CHAR_NOT_VALID           ; return the error code 0x10 (char not valid)
    ret

__HEX_CONV_STR__:
    call __I_TO_C__
    mov [CONVERTER_BUFFER+1], b
    shr a
    shr a
    shr a
    shr a
    call __I_TO_C__
    mov [CONVERTER_BUFFER], b
    ret

__I_TO_C__:
    mov b, a
    and b, 0x0f                 ; get first octet
    cmp b, 0xa                  ; < 10
    js process_number_char      ; process as number
    add b, 0x57                 ; convert to char ( A, B, C, D)
    ret

process_number_char:
    add b, 0x30                 ; convert to char
    ret


__C_TO_I__:
    ; check if its in rage '0' and '9'
    cmp a, 0x30             ; '0' = 0x30
    js not_valid
    cmp a, 0x40             ; '9' = 0x39 + 1
    js process_char_number

    ; check if its in rage 'A' and 'Z'
    cmp a, 0x41             ; 'A' = 0x41
    js not_valid
    cmp a, 0x47             ; 'F' = 0x46 + 1
    js process_char_upper

    ; check if its in rage 'a' and 'f'
    cmp a, 0x61             ; 'a' = 0x61
    js not_valid
    cmp a, 0x67             ; 'f' = 0x66 + 1
    js process_char_lower
    jmp not_valid

process_char_lower:
    sub a, 0x57
    ret

process_char_upper:
    sub a, 0x37
    ret

process_char_number:
    sub a, 0x30
    ret

not_valid:
    mov d, a
    mov a, CHAR_NOT_VALID
    ret
