package assembler

type OperandType int

const (
	OperandTypeUndefined OperandType = 0
	OperandTypeRegister  OperandType = 1
	OperandTypeValue     OperandType = 2
	OperandTypeAddress   OperandType = 3
)

type Operand struct {
	_type OperandType
	raw   string
	value int
}
