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
; Modified.....: None
; ----------------------------------------------------------

.segment "SEG_PUT_CHAR"

    extern ADDR_CURSOR_INDEX_L          ; address of cursor index low byte
    extern ADDR_CURSOR_INDEX_H          ; address of cursor index high byte
    extern ADDR_ROW_COUNTER             ; reserved
    extern ADDR_ROW_COUNT               ; reserved
    extern __VIDEO_START__
    global __PUT_CHAR__
    global __PUT_CURSOR_HOME__
    global __PRINT__

SCREEN_BUFFER_START=__VIDEO_START__

__PUT_CHAR__:
    mov ip, SCREEN_BUFFER_START
    mov b, [ADDR_CURSOR_INDEX_L]
    add ipl, b 
    mov b, [ADDR_CURSOR_INDEX_H]
    adc iph, b 
    mov [ip], a
; increment cursor index
    mov ipl, [ADDR_CURSOR_INDEX_L]
    mov iph, [ADDR_CURSOR_INDEX_H]
    add ip, 1
    mov [ADDR_CURSOR_INDEX_L], ipl
    mov [ADDR_CURSOR_INDEX_H], iph
    ret

__PUT_CURSOR_HOME__:
    xor a, a
    mov [ADDR_CURSOR_INDEX_L], a
    mov [ADDR_CURSOR_INDEX_H], a
    ret

__PRINT__:
    xor c, c
loop_print:
    mov a, [ip+c]
    cmp a, 0
    jz done_print
    push ip
    call __PUT_CHAR__
    pop ip
    inc c
    jmp loop_print
done_print:
    ret
