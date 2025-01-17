package status

const (
	STATUS_FLAG_ZERO uint8 = 1 << iota
	STATUS_FLAG_SIGN
	STATUS_FLAG_PARITY
	STATUS_FLAG_CARRY
	STATUS_FLAG_OVERFLOW
	STATUS_FLAG_INTERRUPT
)

type StatusRegister struct {
	flag uint8
}

func NewStatusRegister() *StatusRegister {
	return &StatusRegister{flag: 0}
}

func (s *StatusRegister) Clear() { s.flag = 0 }

func (s *StatusRegister) SetZeroFlag() {
	s.flag |= STATUS_FLAG_ZERO
	// fmt.Println("setting STATUS_FLAG_ZERO")
}
func (s *StatusRegister) SetSignFlag() {
	s.flag |= STATUS_FLAG_SIGN
	// fmt.Println("setting STATUS_FLAG_SIGN")
}
func (s *StatusRegister) SetParityFlag() {
	s.flag |= STATUS_FLAG_PARITY
	// fmt.Println("setting STATUS_FLAG_PARITY")
}
func (s *StatusRegister) SetCarryFlag() {
	s.flag |= STATUS_FLAG_CARRY
	// fmt.Println("setting STATUS_FLAG_CARRY")
}
func (s *StatusRegister) SetOverflowFlag() {
	s.flag |= STATUS_FLAG_OVERFLOW
	// fmt.Println("setting STATUS_FLAG_OVERFLOW")
}
func (s *StatusRegister) SetInterruptFlag() {
	s.flag |= STATUS_FLAG_INTERRUPT
	// fmt.Println("setting STATUS_FLAG_INTERRUPT")
}

func (s *StatusRegister) ClearZeroFlag() {
	s.flag &= ^STATUS_FLAG_ZERO
	// fmt.Println("clearing STATUS_FLAG_ZERO")
}
func (s *StatusRegister) ClearSignFlag() {
	s.flag &= ^STATUS_FLAG_SIGN
	// fmt.Println("clearing STATUS_FLAG_SIGN")
}
func (s *StatusRegister) ClearParityFlag() {
	s.flag &= ^STATUS_FLAG_PARITY
	// fmt.Println("clearing STATUS_FLAG_PARITY")
}
func (s *StatusRegister) ClearCarryFlag() {
	s.flag &= ^STATUS_FLAG_CARRY
	// fmt.Println("clearing STATUS_FLAG_CARRY")
}
func (s *StatusRegister) ClearOverflowFlag() {
	s.flag &= ^STATUS_FLAG_OVERFLOW
	// fmt.Println("clearing STATUS_FLAG_OVERFLOW")
}
func (s *StatusRegister) ClearInterruptFlag() {
	s.flag &= ^STATUS_FLAG_INTERRUPT
	// fmt.Println("clearing STATUS_FLAG_INTERRUPT")
}

func (s *StatusRegister) IsSet(mask uint8) bool {
	// fmt.Printf("flag: %08b, mask: %08b, result: %08b\n", s.flag, mask, s.flag&mask)
	return s.flag&mask > 0
}

func (s *StatusRegister) Flag() uint8 {
	return s.flag
}
func (s *StatusRegister) String() string {
	out := ""
	for i := 0; i < 8; i++ {
		bit := s.flag & (1 << i) >> i
		out += string(48 + bit)
		if i != 7 {
			out += " "
		}
	}
	return out
}
