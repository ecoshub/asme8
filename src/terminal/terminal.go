package terminal

import (
	"emu/src/bus"
	"emu/src/connectable"
	"emu/src/utils"
)

var _ connectable.Connectable = &Terminal{}

type Terminal struct {
	addressBus *bus.Bus
	dataBus    *bus.Bus
	rangeStart uint16
	rangeEnd   uint16
}

func New(size uint16) *Terminal {
	return &Terminal{}
}

func (t *Terminal) Attach(addrBus, dataBus *bus.Bus, rangeStart, rangeEnd uint16) {
	t.addressBus = addrBus
	t.dataBus = dataBus
	t.rangeStart = rangeStart
	t.rangeEnd = rangeEnd
}

func (t *Terminal) Tick(rw uint8) {
	if rw == utils.IO_READ {
		t.ReadRequest()
		return
	}
	if rw == utils.IO_WRITE {
		t.WriteRequest()
	}
}

func (t *Terminal) ReadRequest() {

}

func (t *Terminal) WriteRequest() {

}

// const (
// 	CHAR_ADDR       uint16 = 0x8000
// 	CHAR_READY_ADDR uint16 = 0x8001
// )

// func ConfigureBIOS(c *Comp) {
// 	ConfigureCharIO(c)
// }

// func ConfigureCharIO(c *Comp) {
// 	c.ram.AttachPort(CHAR_READY_ADDR)
// 	c.ram.AttachPortListener(func(addr uint16, data uint8, rw mem.PortEvent) {
// 		if rw == mem.PORT_W {
// 			if addr == CHAR_READY_ADDR {
// 				row := c.registers.Read(uint8(register.RegA))
// 				column := c.registers.Read(uint8(register.RegB))
// 				char := c.ram.Read(CHAR_ADDR)
// 				peripheral.TerminalCharOut(int(row), int(column), rune(char))
// 			}
// 		}
// 	})
// }
