; put value of register A to the screen buffer
; .segment "__CHAR_OUT__"
start:
    mov [ADDR_CHAR_OUT], a
    jmp done
done:
    rts
