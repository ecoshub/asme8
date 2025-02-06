
start:
    mov ip, message_2
    mov a, [ip]
    brk

message_1:
.asciiz "hello"

message_2:
.ascii "hello world"
