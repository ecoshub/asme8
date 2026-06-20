## ASME8 Instructions Overview

This page is a practical reference for the current assembler implementation.
It focuses on forms that are known to work today.

## Registers

### 8-bit registers (general + low/high parts)
- `a`, `b`, `c`, `d`
- `ipl`, `iph` (instruction pointer low/high)
- `bpl`, `bph` (base pointer low/high)

### 16-bit registers
- `ip` (instruction pointer)
- `bp` (base pointer)
- `sp` (stack pointer)

### Status register
- `sr` (status register)
- Used by: `push sr`, `pop sr`

## Addressing Modes

The assembler supports the following operand styles.

| Mode | Syntax | Typical use |
| --- | --- | --- |
| Implied | `hlt` | No operand |
| Implied reg8 | `inc a` | Single 8-bit register |
| Implied reg16 | `inc ip` | Single 16-bit register |
| Implied status reg | `push sr` | Status register only |
| Implied imm16 | `jmp 0x1234` or `jmp target` | Jumps/calls to address or label |
| reg8, imm8 | `mov a, 0x10` | 8-bit immediate |
| reg8, reg8 | `mov a, b` | 8-bit register transfer |
| reg16, imm16 | `mov ip, 0x1234` | 16-bit immediate move |
| reg16, reg16 | `mov ip, bp` | 16-bit register transfer |
| reg16, imm8 | `add sp, 0x10` | 16-bit add/sub with 8-bit immediate |
| mem direct -> reg | `mov a, [0x1234]` or `mov a, [buffer]` | Absolute memory read |
| reg -> mem direct | `mov [0x1234], a` or `mov [buffer], a` | Absolute memory write |
| mem indexed -> reg | `mov a, [0x1234+b]` | Direct base + reg8 index read |
| reg -> mem indexed | `mov [0x1234+b], a` | Direct base + reg8 index write |
| mem indirect -> reg | `mov a, [ip]` | Read through `ip`/`bp` |
| reg -> mem indirect | `mov [bp], a` | Write through `ip`/`bp` |
| mem indirect+offset -> reg | `mov a, [ip+4]` | Pointer + imm8 offset read |
| reg -> mem indirect+offset | `mov [bp+8], a` | Pointer + imm8 offset write |
| mem reg16-indexed -> reg | `mov a, [ip+c]` | `ip`/`bp` + reg8 index read |
| reg -> mem reg16-indexed | `mov [bp+d], a` | `ip`/`bp` + reg8 index write |

## Instruction Reference

### Data movement
- `mov reg8, reg8`
- `mov reg8, imm8`
- `mov reg16, reg16`
- `mov reg16, imm16`
- `mov reg8, [addr]`
- `mov [addr], reg8`
- `mov reg8, [addr+reg8]`
- `mov [addr+reg8], reg8`
- `mov reg8, [reg16]`
- `mov [reg16], reg8`
- `mov reg8, [reg16+imm8]`
- `mov [reg16+imm8], reg8`
- `mov reg8, [reg16+reg8]`
- `mov [reg16+reg8], reg8`

Example:
```asm
mov a, 0x41
mov [0x1000], a
mov b, [0x1000]
mov c, [ip+2]
```

### Arithmetic and compare
- `add reg8, reg8`
- `add reg8, imm8`
- `add reg16, imm8` (`ip`, `bp`, `sp`)
- `adc reg8, reg8`
- `adc reg8, imm8`
- `sub reg8, reg8`
- `sub reg8, imm8`
- `sub reg16, imm8` (`ip`, `bp`, `sp`)
- `sbb reg8, reg8`
- `sbb reg8, imm8`
- `inc reg8`
- `inc reg16`
- `dec reg8`
- `dec reg16`
- `cmp reg8, reg8`
- `cmp reg8, imm8`

Example:
```asm
add a, 1
adc a, b
sub sp, 0x10
cmp a, 0
```

### Bitwise
- `and reg8, reg8`
- `and reg8, imm8`
- `or reg8, reg8`
- `or reg8, imm8`
- `xor reg8, reg8`
- `xor reg8, imm8`
- `shl reg8`
- `shr reg8`
- `rol reg8`
- `ror reg8`

Example:
```asm
and a, 0x0f
or a, 0x80
shl b
ror c
```

### Flow control
- `jmp imm16|label|reg16`
- `js imm16|label`
- `jns imm16|label`
- `jz imm16|label`
- `jnz imm16|label`
- `jc imm16|label`
- `jnc imm16|label`
- `call imm16|label|reg16`
- `ret`
- `rti`

Example:
```asm
cmp a, 0
jz done
call worker
jmp loop
```

### Stack and flags
- `push reg8`
- `push reg16`
- `push sr`
- `pop reg8`
- `pop reg16`
- `pop sr`
- `clc`
- `cli`
- `sti`

Example:
```asm
push a
push ip
push sr
pop sr
pop ip
```

### System
- `nop`
- `hlt`

## Notes and Limitations

- `add/sub reg16, imm8` uses an 8-bit immediate (`ip`, `bp`, `sp`), not a full 16-bit immediate.
- Labels and variables are accepted wherever an immediate is allowed.
