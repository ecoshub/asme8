; ----------------------------------------------------------
; Routine Name.: __GET_CHAR__
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Waits for a character input to become ready, 
;                then reads the character and stores it in 
;                register A. 
;                - Continuously checks the `ADDR_READY_CHAR` 
;                  address until a character is available (value = 1).
;                - Once ready, reads the character from 
;                  `ADDR_CHAR_READ` and stores it in register A.
; 
; Input........: None
; Output.......: Register A (ASCII value of the character)
; Dependencies.: ADDR_READY_CHAR, ADDR_CHAR_READ
; Modified.....: Register A
; ----------------------------------------------------------

.segment "SEG_GET_CHAR"

    extern ADDR_READY_CHAR
    extern ADDR_GET_CHAR
    global __GET_CHAR__

__GET_CHAR__:
char_wait:
    mov a, [ADDR_READY_CHAR]    ; char ready addr
    cmp a, 1                    ; 1 mean char ready
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [ADDR_GET_CHAR]     ; read from char addr
    rts
