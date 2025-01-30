package instruction

const (
	REGISTER_SYMBOL_A  string = "a"
	REGISTER_SYMBOL_B  string = "b"
	REGISTER_SYMBOL_C  string = "c"
	REGISTER_SYMBOL_D  string = "d"
	REGISTER_SYMBOL_SP string = "sp"
)

var (
	REGISTER_OPCODE = map[string]uint8{
		REGISTER_SYMBOL_A:  0,
		REGISTER_SYMBOL_B:  1,
		REGISTER_SYMBOL_C:  2,
		REGISTER_SYMBOL_D:  3,
		REGISTER_SYMBOL_SP: 4,
	}
)
