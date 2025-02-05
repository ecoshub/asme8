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
; Modified.....: Register A, Register IP
; ----------------------------------------------------------

.segment "SEG_PUT_CHAR"

    extern SCREEN_ATTR_BUFFER
    extern SCREEN_WIDTH
    extern ADDR_CURSOR_INDEX_L          ; address of cursor index low byte
    extern ADDR_CURSOR_INDEX_H          ; address of cursor index high byte
    extern ADDR_ROW_COUNTER             ; reserved
    extern ADDR_ROW_COUNT               ; reserved
    extern __VIDEO_START__
    global __PUT_CHAR__
    global __DEL_CHAR__
    global __PUT_CURSOR_HOME__
    global __GO_TO_ROW__

SCREEN_BUFFER_START=__VIDEO_START__
ROW_COUNT=SCREEN_ATTR_BUFFER
ROW_COUNTER=SCREEN_ATTR_BUFFER

__PUT_CHAR__:
    mov ip, SCREEN_BUFFER_START
    add ipl, [ADDR_CURSOR_INDEX_L]
    adc iph, [ADDR_CURSOR_INDEX_H]
    mov [ip], a
; increment cursor index
    add [ADDR_CURSOR_INDEX_L], 1
    adc [ADDR_CURSOR_INDEX_H], 0
    ret

__DEL_CHAR__:
    mov ipl, [ADDR_CURSOR_INDEX_L]
    cmp ipl, 0
    jz ipl_zero
keep:
    sub ip, 1
    mov [ADDR_CURSOR_INDEX_L], ipl
    mov [ADDR_CURSOR_INDEX_H], iph
    mov a, ' '
    call __PUT_CHAR__
    sub ip, 1
    mov [ADDR_CURSOR_INDEX_L], ipl
    mov [ADDR_CURSOR_INDEX_H], iph
del_finish:
    ret

ipl_zero:
    mov iph, [ADDR_CURSOR_INDEX_H]
    cmp iph, 0
    jz del_finish
    jmp keep

__PUT_CURSOR_HOME__:
    mov [ADDR_CURSOR_INDEX_L], 0
    mov [ADDR_CURSOR_INDEX_H], 0
    ret

__GO_TO_ROW__:
    mov [ADDR_CURSOR_INDEX_L], 0
    mov [ADDR_CURSOR_INDEX_H], 0
    mov ipl, [ADDR_CURSOR_INDEX_L]
    mov iph, [ADDR_CURSOR_INDEX_H]
go_row_loop:
    cmp a, 0
    jz go_row_done
    add ipl, SCREEN_WIDTH
    adc iph, 0
    dec a
    jmp go_row_loop
go_row_done:
    mov [ADDR_CURSOR_INDEX_L], ipl
    mov [ADDR_CURSOR_INDEX_H], iph
    ret

__RETURN__:
    ret
