# Memory Layout



## Init (0x0000 - 0x000f)
- Entry point Address (EP_LOW, EP_HIGH)
```
0x0000 = 0x03           ; INST_JMP_IMPL_IMM16
0x0001 = EP_LOW
0x0002 = EP_HIGH
0x0003 = 0x3f           ; hlt
0x0004 = 0x3f           ; hlt
0x0005 = 0x3f           ; hlt
0x0006 = 0x3f           ; hlt
0x0007 = 0x3f           ; hlt
0x0008 = 0x3f           ; hlt
0x0009 = 0x3f           ; hlt
0x000a = 0x3f           ; hlt
0x000b = 0x3f           ; hlt
0x000c = 0x3f           ; hlt
0x000d = 0x3f           ; hlt
0x000e = 0x3f           ; hlt
0x000f = 0x3f           ; hlt
```


## Interrupt vectors (0x0010 - 0x001f)

```
0x0010 = VEC_1_LOW
0x0011 = VEC_1_HIGH
0x0012 = VEC_2_LOW
0x0013 = VEC_2_HIGH
0x0014 = VEC_3_LOW
0x0015 = VEC_3_HIGH
0x0016 = VEC_4_LOW
0x0017 = VEC_4_HIGH
0x0018 = VEC_5_HIGH
0x0019 = VEC_5_HIGH
0x001a = VEC_6_HIGH
0x001b = VEC_6_HIGH
0x001c = VEC_7_HIGH
0x001d = VEC_7_HIGH
0x001e = VEC_8_HIGH
0x001f = VEC_8_HIGH
```









| Address              | Type |
| :---------------- | :------: |
| 0x0000 |  1   |
| ... |
| 0x000f |  1   |


| Address              | Type |
| :---------------- | :------: |
| 0x0010 |  1   |
| ... |
| 0x001f |  1   |
