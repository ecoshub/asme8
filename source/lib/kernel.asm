; ----------------------------------------------------------
; Segment Name.: SEG_KERNEL
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Defines essential memory addresses and initializes 
; Dependencies.: WOZMON
; ----------------------------------------------------------

.segment "SEG_KERNEL"

    global ADDR_PUT_CHAR                ; char put address for serial terminals screen
    global ADDR_GET_CHAR                ; char read address from keyboard interface
    global ADDR_READY_CHAR              ; char ready address from keyboard interface
    global CONVERTER_BUFFER             ; converter utils buffer
    extern __RAM_START__
    extern __SERIAL_START__
    extern WOZMON

ADDR_PUT_CHAR=__SERIAL_START__
ADDR_READY_CHAR=__SERIAL_START__+1
ADDR_GET_CHAR=__SERIAL_START__+2

CONVERTER_BUFFER=__RAM_START__+0x100    ; converter utils buffer (16 bytes)
WOZMON_BUFFER=__RAM_START__+0x110       ; wozmon ram area (256 bytes)

    jmp WOZMON                          ; start with wozmon
