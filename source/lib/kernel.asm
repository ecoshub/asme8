; ----------------------------------------------------------
; Segment Name.: SEG_KERNEL
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Defines essential memory addresses and initializes 
; Input........: None
; Output.......: Program control transferred to `WOZMAN`.
; Dependencies.: WOZMAN
; Modified.....: None
; ----------------------------------------------------------

.segment "SEG_KERNEL"

    global ADDR_GET_CHAR                ; char read address from keyboard interface
    global ADDR_READY_CHAR              ; char ready address from keyboard interface
    global ADDR_CURSOR_INDEX_L          ; address of cursor index low byte
    global ADDR_CURSOR_INDEX_H          ; address of cursor index high byte
    global ADDR_ROW_COUNTER             ; reserved
    global ADDR_ROW_COUNT               ; reserved
    extern __RAM_START__
    extern __KEYBOARD_START__
    extern WOZMAN

ADDR_CURSOR_INDEX_L=__RAM_START__+0x100
ADDR_CURSOR_INDEX_H=__RAM_START__+0x101
ADDR_ROW_COUNTER=__RAM_START__+0x102
ADDR_ROW_COUNT=__RAM_START__+0x103

ADDR_READY_CHAR=__KEYBOARD_START__
ADDR_GET_CHAR=__KEYBOARD_START__+1

    jmp WOZMAN       ; start with wozman
