package comp

import (
	"emu/src/mem"
	"emu/src/peripheral"
	"emu/src/register"
)

const (
	CHAR_ADDR       uint16 = 0x8000
	CHAR_READY_ADDR uint16 = 0x8001
)

func ConfigureBIOS(c *Comp) {
	ConfigureCharIO(c)
}

func ConfigureCharIO(c *Comp) {
	c.ram.AttachPort(CHAR_READY_ADDR)
	c.ram.AttachPortListener(func(addr uint16, data uint8, rw mem.PortEvent) {
		if rw == mem.PORT_W {
			if addr == CHAR_READY_ADDR {
				row := c.registers.Read(uint8(register.RegA))
				column := c.registers.Read(uint8(register.RegB))
				char := c.ram.Read(CHAR_ADDR)
				peripheral.TerminalCharOut(int(row), int(column), rune(char))
			}
		}
	})
}
