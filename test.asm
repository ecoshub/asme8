; ----------------------------------------------------------
; Routine Name.: STR_TO_HEX
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Converts an ASCII character in register A to its
;                hexadecimal equivalent. 
;                - If the input is a valid hex character, the 
;                  converted value is stored in register A, and
;                  register B is set to 1 (valid).
;                - If the input is not a valid hex character,
;                  register B is set to 0 (invalid).
; 
; Input........: Register A (ASCII value)
; Output.......: Register A (hexadecimal value, if valid)
;                Register B (validation flag: 1 for valid, 0 for invalid)
; Modified.....: Register A, Register B
; ----------------------------------------------------------

STR_TO_HEX:
    mov a, 'b'
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
    mov b, 1
    brk

process_upper:
    sub a, 0x37
    mov b, 1
    brk

process_number:
    sub a, 0x30
    mov b, 1
    brk

not_valid:
    mov b, 0
    brk
