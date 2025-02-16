package instruction

const (
	ADDRESSING_MODE_NONE uint8 = iota
	ADDRESSING_MODE_IMPL
	ADDRESSING_MODE_IMPL_REG
	ADDRESSING_MODE_REG_IMM
	ADDRESSING_MODE_IMPL_IMM_16
	ADDRESSING_MODE_REG_REG
	ADDRESSING_MODE_REG_MEM
	ADDRESSING_MODE_MEM_REG
	ADDRESSING_MODE_MEM_REG_OFFSET
	ADDRESSING_MODE_REG_MEM_OFFSET
	ADDRESSING_MODE_IMPL_SP
	ADDRESSING_MODE_SP_IMM
	ADDRESSING_MODE_SP_REG
	ADDRESSING_MODE_REG_SP
	ADDRESSING_MODE_SP_REG_OFFSET
	ADDRESSING_MODE_REG_SP_OFFSET
	ADDRESSING_MODE_SP_REG_OFFSET_REG
	ADDRESSING_MODE_REG_SP_OFFSET_REG
	ADDRESSING_MODE_IMPL_IP
	ADDRESSING_MODE_IP_IMM
	ADDRESSING_MODE_IP_IMM_16
	ADDRESSING_MODE_IP_REG
	ADDRESSING_MODE_REG_IP
	ADDRESSING_MODE_IP_REG_OFFSET
	ADDRESSING_MODE_REG_IP_OFFSET
	ADDRESSING_MODE_IP_REG_OFFSET_REG
	ADDRESSING_MODE_REG_IP_OFFSET_REG
	ADDRESSING_MODE_MEM_IMM
)
