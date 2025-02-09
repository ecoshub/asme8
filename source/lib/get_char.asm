; ----------------------------------------------------------
; Segment Name.: SEG_GET_CHAR
; Includes.....: '__GET_CHAR__'
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Implements the functionality to check if a character
;                is available for reading and then retrieves it.
; ----------------------------------------------------------
; Subroutines:
; ----------------------------------------------------------
;    Name........: __GET_CHAR__
;    Description.: "stdin" for e8 computer. Checks if a character is ready to be read from the input buffer.
;                  If ready, it reads the character and returns. If not ready, it
;                  waits until a character becomes available.
;    Input.......: None (reads from memory addresses `ADDR_READY_CHAR` and `ADDR_GET_CHAR`)
;    Output......: Register A (contains the character read from `ADDR_GET_CHAR`)
;    Modified....: Register A
; ----------------------------------------------------------

.segment "SEG_GET_CHAR"

    extern ADDR_READY_CHAR
    extern ADDR_GET_CHAR
    global __GET_CHAR__

__GET_CHAR__:
    mov a, [ADDR_READY_CHAR]    ; char ready addr
    cmp a, 1                    ; 1 mean char ready
    jz char_read                ; jump to read label if ready
    jmp __GET_CHAR__            ; jump to char wait

char_read:
    mov a, [ADDR_GET_CHAR]     ; read from char addr
    ret
