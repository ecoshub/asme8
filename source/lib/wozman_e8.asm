; ----------------------------------------------------------
; Segment Name.: SEG_WOZ
; Author.......: eco
; Date.........: 10.02.2025
; Description..: Implements the WOZ monitor, a simple command 
;                parser for reading, writing, and executing 
;                memory operations using hex input.
; ----------------------------------------------------------
; Subroutines:
; ----------------------------------------------------------
;    Name........: WOZMAN
;    Description.: Entry point for the WOZ monitor, handling 
;                  user input and executing commands.
;    Input.......: User input (via __GET_CHAR__), stored in 
;                  WOZMAN_BUFFER.
;    Output......: Executes memory read, write, or run commands.
;    Modified....: Registers A, B, C, D, IP
; ----------------------------------------------------------

.segment "SEG_WOZ"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    extern __RETURN__
    extern CONVERTER_BUFFER
    extern WOZMAN_BUFFER
    extern __STR_CONV_HEX__
    extern __HEX_CONV_STR__
    extern CHAR_NOT_VALID
    global WOZMAN


KEY_CODE_ENTER=0x0d
KEY_CODE_DOT='.'
KEY_CODE_COLUMN=':'
KEY_CODE_R='r'

MODE_READ=0x00
MODE_READ_RANGE=0x01
MODE_WRITE=0x02
MODE_RUN=0x03

MODE=WOZMAN_BUFFER
SEPARATOR_INDEX=WOZMAN_BUFFER+1
RANGE_START=WOZMAN_BUFFER+2
RANGE_END=WOZMAN_BUFFER+4
CHAR_BUFFER=WOZMAN_BUFFER+6

WOZMAN:
    xor a, a                            ; clear all range values
    mov [RANGE_START], a
    mov [RANGE_START+1], a
    mov [RANGE_END], a
    mov [RANGE_END+1], a

    mov a, MODE_READ                    ; set mode to MODE_READ
    mov [MODE], a

    mov a, '>'                          ; print terminal start chars '> '
    call __PUT_CHAR__
    mov a, ' '
    call __PUT_CHAR__

    xor d, d                            ; clear buffer index
main_loop:
    call __GET_CHAR__                   ; read char from stdin
    cmp a, KEY_CODE_ENTER               ; cmp it with key enter
    jz execute                          ; if it is enter jump execute
    call __PUT_CHAR__                   ; print the char (echo)
    cmp a, KEY_CODE_DOT                 ; check if char is '.'
    jz set_mode_read_range              ; set the mode MODE_READ_RANGE
    cmp a, KEY_CODE_COLUMN              ; check if char is ':'
    jz set_mode_write                   ; set the mode MODE_WRITE
    cmp a, KEY_CODE_R                   ; check if char is 'r'
    jz set_mode_run                     ; set the mode MODE_RUN
    mov [CHAR_BUFFER+d], a              ; add char in to buffer
    inc d                               ; increment index
    jmp main_loop                       ; start read next char

set_mode_write:
    mov a, MODE_WRITE                   ; set mode to MODE_WRITE
    mov [MODE], a
    mov [SEPARATOR_INDEX], d            ; save separator index
    jmp main_loop                       ; return to main_loop

set_mode_run:
    mov a, MODE_RUN                     ; set mode to MODE_RUN
    mov [MODE], a
    jmp main_loop                       ; return to main_loop

set_mode_read_range:
    mov a, MODE_READ_RANGE              ; set mode to MODE_READ_RANGE
    mov [MODE], a
    mov [SEPARATOR_INDEX], d            ; save separator index
    jmp main_loop                       ; return to main_loop

execute:
    mov a, [MODE]                       ; get mode in to a
    cmp a, MODE_READ                    ; compare mode with MODE_READ
    jz execute_read                     ; if it is jump to execute_read
    cmp a, MODE_READ_RANGE
    jz execute_read_range
    cmp a, MODE_RUN
    jz execute_run
    cmp a, MODE_WRITE
    jz execute_write
    jmp operation_failed

execute_write:
    push d                              ; save the buffer index to stack
    mov a, [SEPARATOR_INDEX]            ; read separator index
    mov ip, CHAR_BUFFER                 ; pass char buffer
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 
    cmp a, CHAR_NOT_VALID               ; if it returns with CHAR_NOT_VALID
    jz operation_failed                 ; jump to operation failed

    mov a, [CONVERTER_BUFFER]           ; read low byte from converter buffer
    mov [RANGE_START], a                ; store it in to range_start low byte
    mov a, [CONVERTER_BUFFER+1]         ; read high byte from converter buffer
    mov [RANGE_START+1], a              ; store it in to range_start high byte

    mov ip, CHAR_BUFFER                 ; pass char buffer
    mov a, [SEPARATOR_INDEX]            ; read separator index
    add ipl, a                          ; add index to ip, ip+=a
    adc iph, 0

; check if remaining buffer is bigger than 3
    pop d                               ; restore d (buffer size)
    sub d, a                            ; move index as a
    xor c, c                            ; clear c, its gonna be write array index

write_loop:
    cmp d, 3                            ; check if there is at least 3 char left
    js break_execute_write              ; if its not means we are end of the buffer
    add ipl, 1                          ; skip the space char
    adc iph, 0
    mov a, 2                            ; read always 2 byte for a hex value
    push d                              ; save d
    push c                              ; save c
    call __STR_CONV_HEX__               ; execute conversion subroutine
    pop c                               ; restore c
    pop d                               ; restore d
    push ip                             ; save ip
    mov ipl, [RANGE_START]              ; set ip=range_start(16 bit)
    mov iph, [RANGE_START+1]
    mov a, [CONVERTER_BUFFER]           ; read converted value from converter buffer
    mov [ip+c], a                       ; set a to memory address plus offset
    pop ip                              ; restore ip
    add ipl, 2                          ; increment char buffer start as 2 byte
    adc iph, 0
    sub d, 3                            ; sub 3 from total buffer size
    inc c                               ; increment write array index
    jmp write_loop

break_execute_write:
    call __RETURN__                     ; print '\n' and start over
    jmp WOZMAN

execute_run:
    call __RETURN__                     ; print '\n' and jump to given address
    mov a, [SEPARATOR_INDEX]            ; read separator index
    mov ip, CHAR_BUFFER                 ; pass char buffer
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 
    cmp a, CHAR_NOT_VALID               ; if it returns with CHAR_NOT_VALID
    jz operation_failed                 ; jump to operation failed
    mov ipl, [CONVERTER_BUFFER]           ; read low byte from converter buffer
    mov iph, [CONVERTER_BUFFER+1]         ; read high byte from converter buffer
    jmp ip

execute_read_range:
    push d                              ; save the buffer index to stack
    mov a, [SEPARATOR_INDEX]            ; read separator index
    mov ip, CHAR_BUFFER                 ; pass char buffer
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 
    cmp a, CHAR_NOT_VALID               ; if it returns with CHAR_NOT_VALID
    jz operation_failed                 ; jump to operation failed

    mov a, [CONVERTER_BUFFER]           ; read low byte from converter buffer
    mov [RANGE_START], a                ; store it in to range_start low byte
    mov a, [CONVERTER_BUFFER+1]         ; read high byte from converter buffer
    mov [RANGE_START+1], a              ; store it in to range_start high byte

    mov ip, CHAR_BUFFER                 ; pass char buffer
    mov a, [SEPARATOR_INDEX]            ; read separator index
    add ipl, a                          ; add index to ip, ip+=a
    adc iph, 0
    pop d                               ; restore buffer size
    sub d, a                            ; sub a from total buffer size to get remaining index
    mov a, d
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 

    mov a, [CONVERTER_BUFFER]           ; read low byte from converter buffer
    mov [RANGE_END], a                  ; store it in to range_end low byte
    mov a, [CONVERTER_BUFFER+1]         ; read high byte from converter buffer
    mov [RANGE_END+1], a                ; store it in to range_end high byte

    call _print_range                   ; print from range_start to range_end

    jmp WOZMAN                          ; start over

_print_range:
    call __RETURN__                     ; print a carriage return
    mov ipl, [RANGE_START]              ; ip=range_start(16)
    mov iph, [RANGE_START+1]
    mov b, 0                            ; print count
print_range_loop:
    mov a, [RANGE_END+1]                ; read high byte of range_end
    cmp a, iph                          ; compare it to current value
    js print_range_break                ; if current value is bigger than range_end break
    mov a, [RANGE_END]                  ; read low byte of range_end
    cmp a, ipl                          ; compare it to current value
    js print_range_break
    push b                              ; save b in to stack
    call _print_byte                    ; call print byte routine
    pop b                               ; restore b
    mov a, ' '                          ; print space
    call __PUT_CHAR__
    add ipl, 1                          ; inc ip, ip++
    adc iph, 0
    cmp b, 7                            ; if print count is 7 print a carriage return
    jz print_new_line
    inc b                               ; inc print counter
    jmp print_range_loop                ; read next

print_range_break:
    call __RETURN__                     ; print a carriage return 
    ret

print_new_line:
    call __RETURN__                     ; print a carriage return 
    mov b, 0                            ; clear print counter
    jmp print_range_loop                ; read next

execute_read:
    mov ip, CHAR_BUFFER                 ; move buffer in to ip
    mov a, d                            ; move buffer size in to a
    call __STR_CONV_HEX__               ; execute str_to_hex subroutine 
    cmp a, CHAR_NOT_VALID               ; if it returns with CHAR_NOT_VALID
    jz operation_failed                 ; jump to operation failed
    call __RETURN__                     ; print return char
    mov a, ':'                          
    call __PUT_CHAR__                   ; print column
    mov a, ' '
    call __PUT_CHAR__                   ; print space
    mov ipl, [CONVERTER_BUFFER]         ; get converted number low part from converter buffer
    mov iph, [CONVERTER_BUFFER+1]       ; get converted number high part from converter buffer
    call _print_byte                    ; print byte
    call __RETURN__                     ; print return char
    jmp WOZMAN                          ; start over

_print_byte:
    mov a, [ip]                         ; value that ip point
    call __HEX_CONV_STR__               ; convert number to chars
    mov a, [CONVERTER_BUFFER]           ; get converted first digit in to a
    call __PUT_CHAR__                   ; print first char
    mov a, [CONVERTER_BUFFER+1]         ; get converted second digit in to a
    call __PUT_CHAR__                   ; print second char
    ret

operation_failed:                       ; print '?\n' and start over
    mov a, '?'
    call __PUT_CHAR__
    call __RETURN__
    jmp WOZMAN
