
section .data                       ; defining constants
    msg db "hello world", 10        ; define byte 'h','e'.....'l','d',10
    msgLen equ $-msg                ; dummy comment
    name db 'e', 'c', 'o', 0

section .bss                        ; reserving bytes
    buffer resb 128                 ; reserve 128 byte for buffer

section .text
    global _start

_start:
    ; Code to use data from .data and .bss
    ; Example: Writing message to stdout (Linux syscall)
    mov edx, 13                      ; Length of the message
    mov ecx, message                 ; Address of the message
    mov ebx, 1                       ; File descriptor (stdout)
    mov eax, 4                       ; Syscall number for sys_write
    int 0x80                         ; Call kernel

    ; Exit program
    mov eax, 1                       ; Syscall number for sys_exit
    xor ebx, ebx                     ; Exit code 0
    int 0x80                         ; Call kernel
