# ASME8 Assembler Keywords and Directives

This guide explains the language features used around instructions: labels, variables, directives, symbols, and literal formats.

## Labels

Labels define positions in code or data.

Usage:
```asm
name:
```

Example:
```asm
start:
    jmp done

done:
    hlt
```

## Variables and References

### Constant variable

Define a compile-time value:

```asm
NAME=VALUE
```

Example:
```asm
BUFFER_SIZE=16
ADDR_PUT_CHAR=0xffec
```

### Reference variable

Define a variable from another symbol with optional offset:

```asm
NEW=OLD
NEW=OLD+4
NEW=OLD-2
```

Both `+` and `-` offsets are supported.

Example:
```asm
BASE=0x1000
NEXT=BASE+0x20
PREV=NEXT-0x10
```

## Directives

### `.org`

Set current output offset:

```asm
.org 0x6000
```

### `.resb`

Reserve bytes (filled with `0x00` by current assembler behavior):

```asm
buffer:
    .resb 16
```

### `.byte`

Emit one byte per value:

```asm
.byte 0x41, 0x42, 'C'
```

### `.word`

Emit 16-bit values (little-endian):

```asm
.word 0x1234, 0xabcd
```

### `.ascii`

Emit string bytes without terminator:

```asm
.ascii "HELLO"
```

### `.asciiz`

Emit string bytes plus trailing `0x00`:

```asm
.asciiz "HELLO"
```

### `.segment`

Set segment name metadata:

```asm
.segment "SEG_CODE"
```

### Placement

Directives can appear before, between, or after instructions. Output is generated in source order.

Example:
```asm
start:
    mov a, 0x10
    .byte 0xaa
    mov b, 0x20
```

## Access Modifiers

These are used for linking across files.

### `global`

Export a symbol:

```asm
    global main
```

### `extern`

Declare an imported symbol:

```asm
    extern __PUT_CHAR__
```

Note: current grammar expects indentation before `global` / `extern` (commonly 4 spaces).

## Literal Formats

- Hex: `0x10`, `0xabcd`
- Decimal: `10`, `255`
- Binary: `0b1010`, `0b11110000`
- Character: `'A'`, `','`

## Practical Notes

- Labels and variables can be used where immediates are accepted by the instruction form.
- For 8-bit operands, values are encoded as 8-bit (low byte).
- For 16-bit operands/directives, values are encoded little-endian.
