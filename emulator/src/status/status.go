package status

const (
	STATUS_FLAG_ZERO uint8 = 1 << iota
	STATUS_FLAG_SIGN
	STATUS_FLAG_CARRY
	STATUS_FLAG_INTERRUPT_ENABLE
)

type StatusRegister struct {
	flag uint8
}

func NewStatusRegister() *StatusRegister {
	return &StatusRegister{flag: 0}
}

func (s *StatusRegister) Set(value uint8) {
	s.flag = value
}

func (s *StatusRegister) Clear() { s.flag = 0 }

func (s *StatusRegister) SetZeroFlag() {
	s.flag |= STATUS_FLAG_ZERO
}

func (s *StatusRegister) SetSignFlag() {
	s.flag |= STATUS_FLAG_SIGN
}

func (s *StatusRegister) SetCarryFlag() {
	s.flag |= STATUS_FLAG_CARRY
}

func (s *StatusRegister) SetInterruptEnableFlag() {
	s.flag |= STATUS_FLAG_INTERRUPT_ENABLE
}

func (s *StatusRegister) ClearZeroFlag() {
	s.flag &= ^STATUS_FLAG_ZERO
}

func (s *StatusRegister) ClearSignFlag() {
	s.flag &= ^STATUS_FLAG_SIGN
}

func (s *StatusRegister) ClearCarryFlag() {
	s.flag &= ^STATUS_FLAG_CARRY
}

func (s *StatusRegister) ClearInterruptEnableFlag() {
	s.flag &= ^STATUS_FLAG_INTERRUPT_ENABLE
}

func (s *StatusRegister) IsSet(mask uint8) bool {
	return s.flag&mask > 0
}

func (s *StatusRegister) Flag() uint8 {
	return s.flag
}

func (s *StatusRegister) String() string {
	out := ""
	for i := 0; i < 4; i++ {
		bit := s.flag & (1 << i) >> i
		out += string(48 + bit)
		if i != 7 {
			out += " "
		}
	}
	return out
}
