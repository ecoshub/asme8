; ----------------------------------------------------------
; Segment Name.: SEG_CONV
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Provides subroutines for character and 
;                integer conversion, including hex-string 
;                to integer and vice versa.
; ----------------------------------------------------------
; Subroutines:
; ----------------------------------------------------------
;    Name........: __C_TO_I__
;    Description.: Converts a single hex character ('0'-'9', 'A'-'F', 'a'-'f')
;                  into its integer value.
;    Input.......: Register A (Character to convert)
;    Output......: Register A (Converted integer value)
;                  Returns CHAR_NOT_VALID (0x10) if invalid
;    Modified....: Register A
; ----------------------------------------------------------
;    Name........: __I_TO_C__
;    Description.: Converts a 4-bit integer (0-15) to its hex character.
;    Input.......: Register A (4-bit integer)
;    Output......: Register B (Converted character)
;    Modified....: Register B
; ----------------------------------------------------------
;    Name........: __STR_CONV_HEX__
;    Description.: Converts a hexadecimal string to an integer.
;    Input.......: IP (Pointer to hex string), A (String length)
;    Output......: CONVERTER_BUFFER (16-bit integer result)
;                  Returns CHAR_NOT_VALID (0x10) if invalid
;    Modified....: Registers A, B, C, D
; ----------------------------------------------------------
;    Name........: __HEX_CONV_STR__
;    Description.: Converts a 16-bit integer into a two-character hex string.
;    Input.......: Register A (16-bit integer)
;    Output......: CONVERTER_BUFFER (Two-character hex string)
;    Modified....: Register A, B
; ----------------------------------------------------------

.segment "SEG_CONV"

    extern CONVERTER_BUFFER
    global CHAR_NOT_VALID
    global __C_TO_I__
    global __I_TO_C__
    global __STR_CONV_HEX__
    global __HEX_CONV_STR__

CHAR_NOT_VALID=0x10


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
    call __I_TO_C__                 ; param in to a so we can execute int_to_char immediately
    mov [CONVERTER_BUFFER+1], b     ; store result into second part of buffer because its right most digit
    shr a                           ; shift a right 4 times
    shr a
    shr a
    shr a
    call __I_TO_C__                 ; a's second octet is ready. execute int_to_char
    mov [CONVERTER_BUFFER], b       ; store result in to first part because its left most digit
    ret

__I_TO_C__:
    mov b, a
    and b, 0x0f                     ; get first octet
    cmp b, 0xa                      ; < 10
    js process_number_char          ; process as number
    add b, 0x57                     ; convert to char ( A, B, C, D)
    ret

process_number_char:
    add b, 0x30                     ; convert to char
    ret


__C_TO_I__:
    ; check if its in rage '0' and '9'
    cmp a, 0x30                     ; '0' = 0x30
    js not_valid
    cmp a, 0x40                     ; '9' = 0x39 + 1
    js process_char_number

    ; check if its in rage 'A' and 'Z'
    cmp a, 0x41                     ; 'A' = 0x41
    js not_valid
    cmp a, 0x47                     ; 'F' = 0x46 + 1
    js process_char_upper

    ; check if its in rage 'a' and 'f'
    cmp a, 0x61                     ; 'a' = 0x61
    js not_valid
    cmp a, 0x67                     ; 'f' = 0x66 + 1
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
    mov a, CHAR_NOT_VALID
    ret
