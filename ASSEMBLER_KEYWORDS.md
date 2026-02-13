# AsmE8 Assembler Keywords & Directives

## Labels

`<label>:`

Labels mark execution points and can be referenced by jump and call instructions. They serve as symbolic addresses that the assembler resolves to memory offsets.

```asm
start:
    mov a, 5
    jmp done

done:
    brk
```

---

## Variables

### Simple Variable

`<name>=<value>`

Compile-time constant assigned a literal value (hex, decimal, binary, or character).

```asm
BUFFER_SIZE=16
ADDR_PUT_CHAR=0xffec
FLAG=0x01

start:
    mov a, BUFFER_SIZE
    mov [ADDR_PUT_CHAR], a
```

---

### Variable with Offset

`<name>=<reference>+/-<offset>`

Variables can reference other variables or labels with addition, allowing relative address calculations.

```asm
BASE=0x1000
OFFSET=0x10

; Declare derived addresses
BUFFER=BASE+OFFSET          ; 0x1010
NEXT_BUFFER=BUFFER+0x20     ; 0x1030

start:
    mov a, [BUFFER]
    mov [NEXT_BUFFER], a
```

References can point to labels:

```asm
data_start:
    .byte 0x01, 0x02, 0x03

data_middle=data_start+1
data_end=data_start+2

start:
    mov a, [data_middle]    ; reads 0x02
```

**Note:** Only addition (`+`) is supported. Subtraction is not implemented.

---

## Directives

### `.resb <count>`

**Byte Reserve.** Allocates `<count>` uninitialized bytes in memory.

```asm
buffer: .resb 2

start:
    mov a, 0xbb
    mov [buffer], a
    mov a, 0xaa
    mov [buffer+1], a
```

---

### `.byte <values>`

**Byte Data.** Initializes one or more bytes. Values can be decimal, hex, binary, or character literals.

```asm
data:
    .byte 0x41, 0x42, 0x43
    .byte 'H', 'i', 0x00
    .byte 0b11110000, 255
```

---

### `.word <values>`

**Word Data.** Initializes 16-bit words (2 bytes each).

```asm
offsets:
    .word 0x1000, 0x2000, 0x3000
```

---

### `.asciiz "<string>"`

**ASCII Zero-Terminated String.** Initializes a string with automatic null terminator.

```asm
message:
    .asciiz "Hello, World!"
```

---

### `.ascii "<string>"`

**ASCII String (no null terminator).**

```asm
label:
    .ascii "No terminator"
```

---

### `.org <address>`

**Origin.** Sets the memory origin for subsequent code/data.

```asm
.org 0x6000
    mov a, 10
    brk
```

---

### `.segment "<name>"`

**Segment Declaration.** Defines a named memory segment for code organization.

```asm
.segment "SEG_CODE"
    mov a, 5

.segment "SEG_DATA"
    .byte 0x01, 0x02
```

**Placement Flexibility:** Directives can appear at any point in your code - at the beginning, middle, or end of code blocks. They are assembled in order with adjacent instructions:

```asm
start:
    mov a, 0x10

.byte 0xAA, 0xBB    ; data in the middle

    mov b, 0x20

.resb 4             ; reservation at the end
```

---

## Access Modifiers

### `global <label>`

Exports a label for external linkage (visible to other modules).

```asm
    global main
    global process_data
```

---

### `extern <label>`

Declares an external label defined in another module.

```asm
    extern __PUT_CHAR__
    extern buffer_start
```

---

## Instructions

### Data Movement

**`mov <dest>, <src>`** - Move data between registers and memory.

```asm
    mov a, 0x10
    mov [0x100], a
    mov a, [0x200]
    mov a, b
```

---

### Arithmetic

**`add <reg>, <value>`** - Add (with carry).
**`adc <reg>, <value>`** - Add with carry flag.
**`sub <reg>, <value>`** - Subtract.
**`sbb <reg>, <value>`** - Subtract with borrow.
**`inc <reg>`** - Increment by 1.
**`dec <reg>`** - Decrement by 1.

```asm
    add a, 0x10
    adc a, 0x01
    dec b
    inc c
```

---

### Bitwise Operations

**`and <reg>, <value>`** - Bitwise AND.
**`or <reg>, <value>`** - Bitwise OR.
**`xor <reg>, <value>`** - Bitwise XOR.
**`not <reg>`** - Bitwise NOT.
**`shl <reg>`** - Shift left.
**`shr <reg>`** - Shift right.
**`rol <reg>`** - Rotate left.
**`ror <reg>`** - Rotate right.

```asm
    and a, 0x0f
    or a, 0xf0
    shl a
    rol b
```

---

### Comparison & Flags

**`cmp <reg>, <value>`** - Compare (sets flags).
**`clc`** - Clear carry flag.

```asm
    cmp a, b
    jz equal_label
    clc
```

---

### Control Flow

**`jmp <label>`** - Unconditional jump.
**`jz <label>`** - Jump if zero.
**`jnz <label>`** - Jump if not zero.
**`js <label>`** - Jump if sign.
**`jns <label>`** - Jump if not sign.
**`jc <label>`** - Jump if carry.
**`jnc <label>`** - Jump if not carry.

```asm
    cmp a, 0
    jz done
    jnz loop

done:
    brk
```

---

### Subroutines

**`call <label>`** - Call subroutine (push return address).
**`ret`** - Return from subroutine (pop return address).

```asm
    call print_value
    brk

print_value:
    mov a, [buffer]
    ret
```

---

### Stack Operations

**`push <value>`** - Push onto stack.
**`pop <reg>`** - Pop from stack.

```asm
    push a
    push 0x1000
    pop b
```

---

### System

**`brk`** - Break (halt execution).
**`nop`** - No operation.
**`rti`** - Return from interrupt.

```asm
    mov a, 10
    brk         ; halt
    nop         ; ignored
```

---

## Registers

### 8-bit Registers
- `a`, `b`, `c`, `d`, `e`
- `ipl`, `iph` (instruction pointer low/high)
- `bpl`, `bph` (base pointer low/high)

### 16-bit Registers
- `ip` (instruction pointer)
- `bp` (base pointer)

### Special Registers
- `sp` (stack pointer)
- `sr` (status register)

---

## Addressing Modes

**Immediate:** `mov a, 0x10`

**Direct (absolute):** `mov a, [0x100]`

**Indexed:** `mov a, [0x100+b]`

**Register indirect:** `mov a, [ip]`

**Indirect offset:** `mov a, [ip+0x10]`

---

## Value Formats

- **Hexadecimal:** `0x10`, `0xABCD`
- **Decimal:** `10`, `255`
- **Binary:** `0b1010`, `0b11110000`
- **Character:** `'A'`, `','`
