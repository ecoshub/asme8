__VEC_0_ADDR__=0x1000
__SERIAL_START__=0xffed
ADDR_PUT_CHAR=__SERIAL_START__
ADDR_READY_CHAR=__SERIAL_START__+1
ADDR_GET_CHAR=__SERIAL_START__+2

    sti
start:
    xor c, c                    ; clear screen index (char out)
loop:
    nop
    nop
    nop
    jmp loop

.org __VEC_0_ADDR__
vec_0_start:
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
