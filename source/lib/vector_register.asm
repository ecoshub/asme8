; ----------------------------------------------------------
; Segment Name.: Register Vector
; Author.......: eco
; Date.........: 23.02.2026
; Description..: Provides vector handler register functionality
; Modified....: ip, bp
; ----------------------------------------------------------
.segment "SEG_VEC_REG"

    extern INT_VEC_0
    extern INT_VEC_1
    extern INT_VEC_2
    extern INT_VEC_3
    global __REGISTER_VEC_0__
    global __REGISTER_VEC_1__
    global __REGISTER_VEC_2__
    global __REGISTER_VEC_3__

__REGISTER_VEC_0__:
    mov bp, INT_VEC_0
    jmp register_vec_core

__REGISTER_VEC_1__:
    mov bp, INT_VEC_1
    jmp register_vec_core

__REGISTER_VEC_2__:
    mov bp, INT_VEC_2
    jmp register_vec_core

__REGISTER_VEC_3__:
    mov bp, INT_VEC_3
    jmp register_vec_core

register_vec_core:
    mov [bp], ipl
    mov [bp+1], iph
    ret
