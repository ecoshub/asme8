NUMBER=0x9000

; ip is buffer pointer
; a is length
; -> return
; high byte a
; low byte b

start:
    mov a, 4
    mov b, a
    xor c, c                ; buffer index
    mov ip, TEST_BUFFER
loop:
    cmp b, 0
    jz done
    mov a, [ip+c]
    call __C_TO_I__
    cmp a, CHAR_NOT_VALID
    jz input_not_valid
    mov [NUMBER+c], a
    inc c
    dec b
    jmp loop
done:
    mov b, [NUMBER]
    shl b
    shl b
    shl b
    shl b
    mov c, [NUMBER+1]
    or b, c
    mov [NUMBER], b

    mov b, [NUMBER+2]
    shl b
    shl b
    shl b
    shl b
    mov c, [NUMBER+3]
    or b, c
    mov [NUMBER+1], b
    mov a, [NUMBER]
    mov b, [NUMBER+1]
    ret

input_not_valid:
    ret


CHAR_NOT_VALID=0x10

__C_TO_I__:
    ; check if its in rage '0' and '9'
    cmp a, 0x30             ; '0' = 0x30
    js not_valid
    cmp a, 0x40             ; '9' = 0x39 + 1
    js process_number

    ; check if its in rage 'A' and 'Z'
    cmp a, 0x41             ; 'A' = 0x41
    js not_valid
    cmp a, 0x47             ; 'F' = 0x46 + 1
    js process_upper

    ; check if its in rage 'a' and 'f'
    cmp a, 0x61             ; 'a' = 0x61
    js not_valid
    cmp a, 0x67             ; 'f' = 0x66 + 1
    js process_lower
    jmp not_valid

process_lower:
    sub a, 0x57
    ret

process_upper:
    sub a, 0x37
    ret

process_number:
    sub a, 0x30
    ret

not_valid:
    mov a, CHAR_NOT_VALID
    ret


TEST_BUFFER:
.byte 'f', 'f', '9', 'c'
