start:
    mov a, 5
    mov b, 5
    add a, b
    jmp fun
return:
    hlt

fun:
    mov c, 4
    jmp return
