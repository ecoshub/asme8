package comp

import (
	"asme8/common/instruction"
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/register"
	"asme8/emulator/src/status"
	"asme8/emulator/src/terminal"
	"asme8/emulator/utils"
	"fmt"
	"time"
)

const (
	DefaultMemorySize = 1 << 16
)

var (
	StackStart uint16
)

type Comp struct {
	registers register.Module
	status    *status.StatusRegister
	step      uint8

	instructionRegister instruction.Type
	operandRegister     uint8

	programCounter        uint16
	memoryAddressRegister uint16
	stackPointer          uint16
	memoryDataRegister    uint8

	aluOut       bool
	bridgeEnable bool

	aluBus    *bus.Bus
	outputBus *bus.Bus
	inputBus  *bus.Bus
	addrBus   *bus.Bus
	rw        uint8 // read 1 write 0
	devices   []connectable.Connectable

	terminal               *terminal.Terminal
	singleTicker           chan struct{}
	breakPoints            []uint16
	inspectionMemoryOffset uint16

	programLoaded bool
	stopChan      chan struct{}
	running       bool
	pause         bool
	debug         bool
	delay         time.Duration
}

func New() *Comp {
	c := &Comp{
		registers:     register.NewModule(6),
		status:        status.NewStatusRegister(),
		aluBus:        bus.New(),
		outputBus:     bus.New(),
		inputBus:      bus.New(),
		addrBus:       bus.New(),
		devices:       make([]connectable.Connectable, 0, 1),
		stopChan:      make(chan struct{}, 1),
		singleTicker:  make(chan struct{}, 1),
		programLoaded: false,
	}
	return c
}

func (c *Comp) AttachTerminal(terminal *terminal.Terminal) {
	c.terminal = terminal
	c.terminal.Components.Screen.CommandPalette.ListenKeyEventEnter(func(input string) {
		c.HandleCommands(input)
	})
}

func (c *Comp) ProgramLoaded() {
	c.programLoaded = true
}

func (c *Comp) SetStackStart(start uint16) {
	StackStart = start
	c.stackPointer = start
}

func (c *Comp) SetPause(enable bool) {
	c.pause = enable
}

func (c *Comp) SetDebug(enable bool) {
	c.debug = enable
}

func (c *Comp) SetDelay(delay time.Duration) {
	c.delay = delay
}

func (c *Comp) PrintRegisters() {
	fmt.Println(c.registers)
}

func (c *Comp) PrintBusses() {
	fmt.Printf("busses, data: %x, x: %x, y: %x\n", c.inputBus.Read_16(), c.aluBus.Read_16(), c.outputBus.Read_16())
}

func (c *Comp) ReadRegister(index uint8) uint8 {
	return c.registers.Read_8(index)
}

func (c *Comp) GetStatusRegister() uint8 {
	return c.status.Flag()
}

func (c *Comp) ConnectDevice(dev connectable.Connectable, rangeStart uint16, size uint16) {
	if dev == nil {
		return
	}
	dev.Attach(c.addrBus, c.outputBus, rangeStart, size)
	c.devices = append(c.devices, dev)
}

func (c *Comp) Run() {
	if c.running {
		return
	}

	if c.terminal != nil {
		go c.terminal.Run()
	}

	c.inputBus.AttachBusChange(func(rw uint8) {
		if rw == utils.IO_WRITE {
			c.LogMemory()
		}
	})

	c.LogMemory()
	c.LogState()

	if !c.programLoaded {
		c.LogWithStyle("no program load. load program with 'load-bin' or 'load-asm' command", DefaultWarningStyle)
		return
	}

	t := time.NewTicker(c.delay)
	c.running = true
	for {
		select {
		case <-t.C:
			if c.pause {
				continue
			}
			keep := c.tick()
			if !keep {
				return
			}
		case <-c.singleTicker:
			keep := c.tick()
			if !keep {
				return
			}
		case <-c.stopChan:
			return
		}
	}
}

func (c *Comp) Stop() {
	if !c.running {
		return
	}
	c.stopChan <- struct{}{}
	c.running = false
}

func (c *Comp) tick() bool {
	defer c.clear()
	microinstructions := CONTROL_SIGNALS[c.instructionRegister][c.step]
	for _, mi := range microinstructions {
		keep := c.run(mi)
		if !keep {
			break
		}
	}
	c.LogState()
	if len(microinstructions) == 0 || c.instructionRegister == instruction.Type(MI_BRK) {
		c.LogWithStyle(" ## BREAK ## ", DefaultWarningStyle)
		return false
	}
	return true
}

func (c *Comp) run(mi uint64) bool {
	f, ok := microinstructionFunctions[mi]
	if ok {
		f(c, mi)
		return true
	}
	return false
}

func (c *Comp) tickDevices() {
	for _, dev := range c.devices {
		dev.Tick(c.rw)
	}
}

func (c *Comp) Reset(excludeROM bool, startWithPause bool) {
	c.Stop()
	for _, dev := range c.devices {
		if excludeROM {
			if dev.GetName() == "ROM" {
				continue
			}
		}
		dev.Clear()
	}
	c.registers.Clear()
	c.status.Clear()
	c.step = 0
	c.instructionRegister = 0
	c.operandRegister = 0
	c.programCounter = 0
	c.stackPointer = StackStart
	c.memoryAddressRegister = 0
	c.memoryAddressRegister = 0
	c.rw = utils.IO_READ
	c.stopChan = make(chan struct{}, 1)
	c.running = false
	c.pause = startWithPause
	c.clear()
}

func (c *Comp) Restart(startWithPause bool) {
	c.Reset(true, true)
	go c.Run()
}

func (c *Comp) clear() {
	c.inputBus.Clear()
	c.aluBus.Clear()
	c.outputBus.Clear()
	c.addrBus.Clear()
	c.aluOut = false
	c.bridgeEnable = false
	c.rw = utils.IO_READ
}
