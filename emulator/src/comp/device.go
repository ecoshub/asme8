package comp

import (
	"asme8/common/config"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"asme8/emulator/src/terminal"
	"fmt"
	"strings"
)

func (c *Comp) LoadProgram() error {
	if len(c.Config.Program) == 0 {
		return nil
	}

	buffer := make([]uint8, 0xffff)
	copy(buffer[:], c.Config.Program[:])

	for _, mc := range c.Config.MemoryConfig.Configs {
		if strings.HasPrefix(mc.Name, "ROM") {
			start := mc.Start.Value
			end := mc.Start.Value + mc.Size.Value
			i, ok := c.GetDeviceInterface(mc.Name)
			if !ok {
				return fmt.Errorf("fatal error. device not found. dev: %s", mc.Name)
			}
			r := i.(*rom.Rom)
			r.Load(0, buffer[start:end])
			c.ProgramLoaded()
		}
	}

	return nil
}

func (c *Comp) CreateDevices() error {

	err := config.ResolveMemoryLayout(c.Config.MemoryConfig.Configs)
	if err != nil {
		return err
	}

	var term *terminal.Terminal
	var deviceRom *rom.Rom
	var deviceRam *ram.Ram
	var ramStart uint16

	for _, mc := range c.Config.MemoryConfig.Configs {
		if strings.HasPrefix(mc.Name, "ROM") {
			r := rom.New(mc.Size.Value)
			r.SetName(mc.Name)
			c.ConnectDevice(r, mc.Start.Value, mc.Size.Value)
			if deviceRom == nil {
				deviceRom = r
			}
		}
		if strings.HasPrefix(mc.Name, "RAM") {
			r := ram.New(mc.Size.Value)
			r.SetName(mc.Name)
			c.ConnectDevice(r, mc.Start.Value, mc.Size.Value)
			if deviceRam == nil {
				ramStart = mc.Start.Value
				deviceRam = r
			}
		}
		if strings.HasPrefix(mc.Name, "VIDEO") {
			if c.Config.Headless {
				continue
			}
			term, err = terminal.New(int(mc.Size.Value))
			if err != nil {
				return err
			}
			c.ConnectDevice(term.Screen, mc.Start.Value, mc.Size.Value)
			// screen components such as panels attaching to computer
			c.AttachTerminal(term)
		}
		if strings.HasPrefix(mc.Name, "KEYBOARD") {
			if c.Config.Headless {
				continue
			}
			if term == nil {
				return fmt.Errorf("can not define KEYBOARD without define VIDEO first")
			}
			c.ConnectDevice(term.Keyboard, mc.Start.Value, mc.Size.Value)
		}
	}

	if deviceRom == nil {
		return fmt.Errorf("memory config must define ROM")
	}

	if deviceRam == nil {
		return fmt.Errorf("memory config must define RAM")
	}

	c.SetStackStart(ramStart + 0xff)

	return nil
}

func (c *Comp) ConnectDevice(dev connectable.Connectable, rangeStart uint16, size uint16) {
	if dev == nil {
		return
	}
	dev.Attach(c.addrBus, c.outputBus, rangeStart, size)
	c.devices = append(c.devices, dev)
}

func (c *Comp) tickDevices() {
	for _, dev := range c.devices {
		dev.Tick(c.rw)
	}
}

func (c *Comp) GetDevice(name string) (connectable.Connectable, bool) {
	for _, d := range c.devices {
		if strings.HasPrefix(d.GetName(), name) {
			return d, true
		}
	}
	return nil, false
}

func (c *Comp) GetDeviceInterface(name string) (interface{}, bool) {
	for _, d := range c.devices {
		if strings.HasPrefix(d.GetName(), name) {
			return d, true
		}
	}
	return nil, false
}

func (c *Comp) GetRom() (*rom.Rom, bool) {
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

func (c *Comp) GetRam() (*ram.Ram, bool) {
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
