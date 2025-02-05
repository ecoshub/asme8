; ----------------------------------------------------------
; Segment Name.: SEG_CONV
; Includes.....: '__C_TO_I__', '__STR_CONV_HEX__'
; Author.......: eco
; Date.........: 28.01.2025
; ----------------------------------------------------------

.segment "SEG_CONV"

    extern CONVERTER_BUFFER
    global __C_TO_I__
    global __STR_CONV_HEX__

CHAR_NOT_VALID=0x10

; parameters:
;       ip  -> buffer pointer
;       a   -> buffer length
; return:
;       a   -> high byte of value
;       b   -> low byte of value
__STR_CONV_HEX__:
    mov b, a                ; get length of buffer in b
    xor c, c                ; buffer index
hex_conv_loop:
    cmp b, 0                ; final condition
    jz conv_done            ; conversion done
    mov a, [ip+c]
    call __C_TO_I__
    cmp a, CHAR_NOT_VALID
    jz buffer_not_valid
    mov [CONVERTER_BUFFER+c], a
    inc c
    dec b
    jmp hex_conv_loop
conv_done:
    ; [0] = [0] << 4 | [1]
    mov b, [CONVERTER_BUFFER]
    shl b
    shl b
    shl b
    shl b
    mov c, [CONVERTER_BUFFER+1]
    or b, c
    mov [CONVERTER_BUFFER], b

    ; [1] = [2] << 4 | [3]
    mov b, [CONVERTER_BUFFER+2]
    shl b
    shl b
    shl b
    shl b
    mov c, [CONVERTER_BUFFER+3]
    or b, c
    mov [CONVERTER_BUFFER+1], b

    mov a, [CONVERTER_BUFFER]
    mov b, [CONVERTER_BUFFER+1]
buffer_not_valid:
    ret

__C_TO_I__:
    ; check if its in rage '0' and '9'
    cmp a, 0x30             ; '0' = 0x30
    js not_valid
    cmp a, 0x40             ; '9' = 0x39 + 1
    js process_number

    ; check if its in rage 'A' and 'Z'
    cmp a, 0x41             ; 'A' = 0x41
    js not_valid
    cmp a, 0x47             ; 'F' = 0x46 + 1
    js process_upper

    ; check if its in rage 'a' and 'f'
    cmp a, 0x61             ; 'a' = 0x61
    js not_valid
    cmp a, 0x67             ; 'f' = 0x66 + 1
    js process_lower
    jmp not_valid

process_lower:
    sub a, 0x57
    ret

process_upper:
    sub a, 0x37
    ret

process_number:
    sub a, 0x30
    ret

not_valid:
    mov a, CHAR_NOT_VALID
    ret
