ADDR_CHAR_OUT=0x7000
ADDR_CHAR_RDY=0x6ffe
ADDR_CHAR_READ=0x6fff

CHAR_DEL=0x7f
CHAR_SPACE=0x20
CHAR_ENTER=0x0d

INPUT_NUMBER: .resb 2
INPUT_BUFFER: .resb 16

start:
    xor c, c                    ; c is char index for buffers (screen and input)
                                ; lets clear char index

char_wait:
    mov a, [ADDR_CHAR_RDY]      ; char ready addr
    cmp a, 1                    ; char ready mean 1
    jz char_read                ; jump to read label if ready
    jmp char_wait               ; jump to char wait

char_read:
    mov a, [ADDR_CHAR_READ]     ; read from char addr
    cmp a, CHAR_DEL             ; check if char is 'del'
    jz key_del                  ; run key_del subroutine if char 'del'
    cmp a, CHAR_ENTER           ; check if char is 'enter'
    jz key_enter                ; run key_enter subroutine if char 'enter'
    jsr char_out                ; print char to screen
    jsr write_to_buffer         ; store char to input buffer
    inc c                       ; increment cursor index by 1
    jmp char_wait               ; jump to char wait

key_del:
    cmp c, 0                    ; check if is there char to delete
    jz char_wait                ; if not jump to wait
    dec c                       ; dec cursor index by 1
    mov a, CHAR_SPACE           ; store space char in to a
    mov [ADDR_CHAR_OUT+c], a    ; write screen buffer to a with offset c
    mov [INPUT_BUFFER+c], a     ; write input buffer a zero
    jmp char_wait               ; jump to char wait

key_enter:
    mov a, c
key_enter_loop:
    cmp a, c
    jnz key_enter_quit
    mov b, [INPUT_BUFFER+a]
    dec a 
    sub b, 0x30
    add [INPUT_NUMBER], b
    ; jnc  key_enter_loop
    ; jsr mul:
    jmp key_enter_loop

key_enter_quit:
    xor a, a
    add a, 0x30
    mov a, [INPUT_NUMBER]
    jsr char_out
    jmp char_wait

char_out:
    mov [ADDR_CHAR_OUT+c], a    ; write content of a in to screen buffer with offset c
    rts

write_to_buffer:
    mov [INPUT_BUFFER+c], a    ; write content of a in to screen buffer with offset c
    rts
