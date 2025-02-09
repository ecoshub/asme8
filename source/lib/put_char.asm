; ----------------------------------------------------------
; Segment Name.: Character Output
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Provides functionality to output a character to a specified
;                address. Includes a special subroutine to return a carriage return (0x0D).
; ----------------------------------------------------------
; Subroutines:
; ----------------------------------------------------------
;    Name........: __PUT_CHAR__
;    Description.: "stdout" for e8 computer. Writes the character in register A to the memory address `ADDR_PUT_CHAR`.
;    Input.......: Register A (Character to write)
;    Output......: None (Character is written to memory at `ADDR_PUT_CHAR`)
;    Modified....: None
; ----------------------------------------------------------
;    Name........: __RETURN__
;    Description.: Outputs a carriage return (0x0D) to the memory address `ADDR_PUT_CHAR`.
;    Input.......: None
;    Output......: 0x0D (Carriage return written to `ADDR_PUT_CHAR`)
;    Modified....: None
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
