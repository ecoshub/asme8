package instruction

const (
	ADDR_MODE_NONE                       uint8 = iota
	ADDR_MODE_REG_TO_MEM_INDEXED               // <inst> [0x1234+b], a
	ADDR_MODE_REG_TO_MEM_DIRECT                // <inst> [0x1234], a
	ADDR_MODE_REG_TO_MEM_INDIRECT_OFFSET       // <inst> [ip+8], a
	ADDR_MODE_REG_TO_MEM_REG16_INDEXED         // <inst> [ip+b], a
	ADDR_MODE_REG_TO_MEM_INDIRECT              // <inst> [ip], a
	ADDR_MODE_IMPLIED_REG8                     // <inst> a
	ADDR_MODE_REG8_IMM8                        // <inst> a, 0x10
	ADDR_MODE_MEM_TO_REG_INDEXED               // <inst> a, [0x1234+b]
	ADDR_MODE_MEM_TO_REG_DIRECT                // <inst> a, [0x1234]
	ADDR_MODE_MEM_TO_REG_INDIRECT_OFFSET       // <inst> a, [ip+8]
	ADDR_MODE_MEM_TO_REG_REG16_INDEXED         // <inst> a, [ip+b]
	ADDR_MODE_MEM_TO_REG_INDIRECT              // <inst> a, [ip]
	ADDR_MODE_REG8_REG8                        // <inst> a, b
	ADDR_MODE_IMPLIED                          // <inst> brk
	ADDR_MODE_IMPLIED_IMM16                    // <inst> jmp 0x1234
	ADDR_MODE_IMPLIED_REG16                    // <inst> ip
	ADDR_MODE_REG16_IMM8                       // <inst> ip, 0x10
	ADDR_MODE_REG16_IMM16                      // <inst> ip, 0x1234
	ADDR_MODE_REG16_REG16                      // <inst> ip, bp
	ADDR_MODE_REG16_STACK                      // <inst> ip, sp
	ADDR_MODE_IMPLIED_STACK                    // <inst> sp
	ADDR_MODE_STACK_IMM8                       // <inst> sp, 0x10
	ADDR_MODE_STACK_REG16                      // <inst> sp, ip
)
