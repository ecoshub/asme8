INT_VEC_0=0x0020

__SERIAL_START__=0xffed
ADDR_PUT_CHAR=__SERIAL_START__
ADDR_READY_CHAR=__SERIAL_START__+1
ADDR_GET_CHAR=__SERIAL_START__+2

; register interrupt
init:
    
    mov ip, vec_0_handler
    mov bp, INT_VEC_0
    mov [bp], ipl
    mov [bp+1], iph

; enable interrupts
    sti
; start the program
start:
    xor c, c                    ; clear screen index (char out)
loop:
    nop
    nop
    nop
    jmp loop

vec_0_handler:
; check keyboard interrupt
    mov a, [ADDR_READY_CHAR]    ; char ready addr
    cmp a, 1                    ; char ready mean 1
    jnz done                    ; if its not ready finish execution
    call char_read
    call char_out
done:
    rti

char_read:
    mov a, [ADDR_GET_CHAR]      ; read from char addr
    ret                         ; jump to char out routine

char_out:
    mov [ADDR_PUT_CHAR], a      ; write content of a in to screen buffer with offset c
    inc c                       ; increment cursor index by 1
    ret
