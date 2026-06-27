; ----------------------------------------------------------
; Segment Name.: SEG_KERNEL
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Defines essential memory addresses and initializes 
; Dependencies.: WOZMON
; ----------------------------------------------------------

; 32 bytes allocated for this segment
.segment "SEG_INIT"

    global ADDR_PUT_CHAR                ; char put address for serial terminals screen
    global ADDR_GET_CHAR                ; char read address from keyboard interface
    global ADDR_READY_CHAR              ; char ready address from keyboard interface
    global CONVERTER_BUFFER             ; converter utils buffer
    global WOZMON_BUFFER                ; wozmon screen buffer
    global INT_VEC_0_LOW                ; interrupt vector 0 low address
    global INT_VEC_0_HIGH               ; interrupt vector 0 high address
    extern __RAM_START__
    extern __SERIAL_START__
    extern WOZMON

; serial area (3 byte)
ADDR_PUT_CHAR=__SERIAL_START__
ADDR_READY_CHAR=__SERIAL_START__+1
ADDR_GET_CHAR=__SERIAL_START__+2

; stack area (256 byte)
STACK_START=__RAM_START__+0xff

; vector proxy area (64 byte)
INT_VEC_0_LOW=0xfffe
INT_VEC_0_HIGH=0xffff

; shared converter buffer area (16 byte)
CONVERTER_BUFFER=__RAM_START__+0x140

; wozmon char buffer area (256 byte)
WOZMON_BUFFER=__RAM_START__+0x150       ; wozmon ram area (256 bytes)

; set stack
    mov ip, STACK_START
    mov sp, ip

; clear index pointer
    xor ipl, ipl
    xor iph, iph

    jmp WOZMON                          ; start with wozmon
