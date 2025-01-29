; ----------------------------------------------------------
; Routine Name.: STR_TO_HEX
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Converts an ASCII character in register A to its
;                hexadecimal equivalent. 
;                - If the input is a valid hex character, the 
;                  converted value is stored in register A.
;                - If the input is not a valid hex character,
;                  register A is set to 0 (invalid).
; 
; Input........: Register A (ASCII value)
; Output.......: Register A (hexadecimal value, if not valid = 0)
; Modified.....: Register A
; ----------------------------------------------------------


.segment "SEG_STR_TO_HEX"

    global STR_TO_HEX

STR_TO_HEX:
start:
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
    mov a, 0
    ret
