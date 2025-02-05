.segment "upper"

    global upper

upper:
    cmp a, 'a'
    js done
    cmp a, 'z'
    js to_upper
    jz to_upper
    jmp done

to_upper:
    sub a, 0x20
    ret

done:
    ret
