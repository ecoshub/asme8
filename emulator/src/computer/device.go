package computer

import (
	"asme8/common/config"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"asme8/emulator/src/terminal"
	"errors"
	"fmt"
	"strings"
)

func (c *Computer) LoadProgram(program []byte) error {
	if len(program) == 0 {
		return errors.New("program not found")
	}

	if len(program) > 0x10000 {
		return errors.New("program size is bigger than 64K")
	}

	buffer := make([]byte, 0x10000)
	copy(buffer[:], program[:])

	for _, dev := range c.devices {
		start, end := dev.GetRange()
		dev.Load(0, buffer[start:end])
	}

	c.programLoaded = true
	return nil
}

func (c *Computer) CreateDevices() error {

	err := config.ResolveMemoryLayout(c.Config.MemorySegment.Memory.Configs)
	if err != nil {
		return err
	}

	var term *terminal.Terminal
	roDevicesCount := 0
	rwDevicesCount := 0
	var ramStart uint16

	for _, mc := range c.Config.MemorySegment.Memory.Configs {
		switch mc.Type {
		case config.MemoryTypeReadOnly:
			r := rom.New(mc.Size)
			r.SetName(mc.Name)
			c.ConnectDevice(r, mc.Start.Value, mc.Size)
			roDevicesCount += 1
			continue
		case config.MemoryTypeReadWrite:
			r := ram.New(mc.Size)
			r.SetName(mc.Name)
			c.ConnectDevice(r, mc.Start.Value, mc.Size)
			if rwDevicesCount == 0 {
				ramStart = mc.Start.Value
			}
			rwDevicesCount += 1
			continue
		case config.MemoryTypeSerial:
			if c.Config.Headless {
				continue
			}
			term, err = terminal.New(int(mc.Size), true, c.InterruptRequest)
			if err != nil {
				return err
			}
			c.ConnectDevice(term.Screen, mc.Start.Value, 1)
			c.ConnectDevice(term.Keyboard, mc.Start.Value+1, 2)
			// screen components such as panels attaching to computer
			c.terminal = term
			continue
		}
	}

	if roDevicesCount == 0 {
		return fmt.Errorf("memory config must define ROM")
	}

	if rwDevicesCount == 0 {
		return fmt.Errorf("memory config must define RAM")
	}

	StackStart = ramStart + 0xff

	if rwDevicesCount > 1 && c.terminal != nil {
		term.Components.SysLogPanel.Push(fmt.Sprintf("you have more than one 'RAM' section. stack starts from: %04x", StackStart))
	}

	c.SetStackStart(StackStart)
	c.ramStart = ramStart
	return nil
}

func (c *Computer) tickDevices() {
	for _, dev := range c.devices {
		dev.Tick(c.rw)
	}
}

func (c *Computer) ConnectDevice(dev connectable.Connectable, rangeStart uint16, size uint16) {
	if dev == nil {
		return
	}
	dev.Attach(c.addrBus, c.outputBus, rangeStart, size)
	c.devices = append(c.devices, dev)
}

func (c *Computer) GetDevice(name string) (connectable.Connectable, bool) {
	for _, d := range c.devices {
		if strings.HasPrefix(d.GetName(), name) {
			return d, true
		}
	}
	return nil, false
}

func (c *Computer) GetDeviceInterface(name string) (interface{}, bool) {
	for _, d := range c.devices {
		if strings.HasPrefix(d.GetName(), name) {
			return d, true
		}
	}
	return nil, false
}

func (c *Computer) GetRom() (*rom.Rom, bool) {
	i, ok := c.GetDeviceInterface("ROM")
	if !ok {
		return nil, false
	}
	r, ok := i.(*rom.Rom)
	if ok {
		return r, true
	}
	return nil, false
}

func (c *Computer) GetRam() (*ram.Ram, bool) {
	i, ok := c.GetDeviceInterface("RAM")
	if !ok {
		return nil, false
	}
	r, ok := i.(*ram.Ram)
	if ok {
		return r, true
	}
	return nil, false
}
