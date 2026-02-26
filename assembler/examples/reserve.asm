RESULT: .resb 2

start:
    mov a, 0xbb
    mov [RESULT], a
    mov a, 0xaa
    mov [RESULT+1], a
    hlt

