; ----------------------------------------------------------
; Routine Name.: __PUT_CHAR__
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Outputs the character stored in register A 
;                to the screen buffer by writing it to the 
;                `ADDR_PUT_CHAR` address.
; 
; Input........: Register A (ASCII value of the character)
; Output.......: Writes the character to `ADDR_PUT_CHAR`
; Dependencies.: ADDR_PUT_CHAR
; Modified.....: Register A
; ----------------------------------------------------------

.segment "SEG_PUT_CHAR"

    extern ADDR_PUT_CHAR
    global __PUT_CHAR__
    global __RETURN__

__PUT_CHAR__:
    mov [ADDR_PUT_CHAR], a
    ret

__RETURN__:
    mov a, 0x0d
    mov [ADDR_PUT_CHAR], a
    ret
