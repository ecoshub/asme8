; ----------------------------------------------------------
; Routine Name.: ECHO
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Implements an echo functionality that reads characters 
;                from input, displays them on the screen, and handles 
;                backspace (delete) functionality. 
;                - Reads input using `__GET_CHAR__`.
;                - Displays the input character using `__PUT_CHAR__`.
;                - If the input is a delete character (`DEL`), it erases 
;                  the last displayed character by overwriting it with a 
;                  space and moving the cursor back.
; 
; Input........: Standard input via `__GET_CHAR__` (ASCII characters).
; Output.......: Displays characters on screen via `__PUT_CHAR__`.
; Dependencies.: __GET_CHAR__, __PUT_CHAR__, __RAM_START__
; Modified.....: Register A (input/output character)
;                Register B (cursor position/index)
; ----------------------------------------------------------

.segment "SEG_ECHO"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    extern __RAM_START__
    global ECHO

CHAR_DEL=0x7f
CHAR_SPACE=' '

ECHO:
    xor b, b                        ; clear input index
start:
    jsr __GET_CHAR__                     ; read from stdin
    cmp a, CHAR_DEL                 ; check if char is 'del'
    jz key_del                      ; jump to key_del label if char 'del'
    jsr __PUT_CHAR__                    ; print char to screen
    inc b                           ; increment input index
    jmp start                       ; start over

key_del:
    cmp b, 0                        ; check if is there char to delete
    jz start                        ; if not jump to wait
    dec b                           ; dec cursor index by 1
    mov a, CHAR_SPACE               ; store space char in to a
    jsr __PUT_CHAR__                    ; print space char to screen
    jmp start                       ; start over
