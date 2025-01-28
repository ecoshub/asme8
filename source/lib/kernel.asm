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

    global ADDR_PUT_CHAR
    global ADDR_GET_CHAR
    global ADDR_READY_CHAR
    extern WOZMAN

ADDR_PUT_CHAR=0xf7ee
ADDR_GET_CHAR=0xffef
ADDR_READY_CHAR=0xffee

    jmp WOZMAN       ; start with wozman
