start:
    mov ip, 0x2200      ; set stack start to ip
    mov sp, ip          ; copy its value to sp
    mov a, 0x10         ; store 0x10 in reg a
    push a              ; push a to stack
    mov b, [0x2200]     ; move 0x2200 value to b to see if its equal to a
    
