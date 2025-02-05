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

    global SCREEN_WIDTH
    global SCREEN_HEIGHT
    global ADDR_GET_CHAR                ; char read address from keyboard interface
    global ADDR_READY_CHAR              ; char ready address from keyboard interface
    global ADDR_CURSOR_INDEX_L          ; address of cursor index low byte
    global ADDR_CURSOR_INDEX_H          ; address of cursor index high byte
    global ADDR_ROW_COUNTER
    global ADDR_ROW_COUNT
    global CONVERTER_BUFFER
    extern __RAM_START__
    extern __KEYBOARD_START__
    extern WOZMAN


SCREEN_WIDTH=80
SCREEN_HEIGHT=24

ADDR_CURSOR_INDEX_L=__RAM_START__+0x100     ; size 1 byte
ADDR_CURSOR_INDEX_H=__RAM_START__+0x101     ; size 1 byte
ADDR_ROW_COUNTER=__RAM_START__+0x102        ; size 1 byte
ADDR_ROW_COUNT=__RAM_START__+0x103          ; size 1 byte
CONVERTER_BUFFER=__RAM_START__+0x104        ; size 8 byte
WOZMAN_BUFFER=__RAM_START__+0x10c           ; size 16 byte
SCREEN_ATTR_BUFFER=__RAM_START__+0x11c      ; size 4 byte

ADDR_READY_CHAR=__KEYBOARD_START__
ADDR_GET_CHAR=__KEYBOARD_START__+1

    jmp WOZMAN       ; start with wozman
