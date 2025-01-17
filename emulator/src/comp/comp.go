package comp

import (
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/instruction"
	"asme8/emulator/src/register"
	"asme8/emulator/src/status"
	"asme8/emulator/src/terminal"
	"asme8/emulator/utils"
	"fmt"
	"time"
)

const (
	DefaultMemorySize = 1 << 16

	StackStart uint16 = 0x80ff
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

	aluDirectOut   bool
	aluEnable      bool
	aluStoreEnable bool
	bridgeEnable   bool
	bridgeDir      uint8

	busX    *bus.Bus
	busY    *bus.Bus
	dataBus *bus.Bus
	addrBus *bus.Bus
	store   uint8
	rw      uint8 // read 1 write 0
	devices []connectable.Connectable

	terminalComponents     *terminal.Components
	singleTicker           chan struct{}
	breakPoints            []uint16
	inspectionMemoryOffset uint16

	stopChan chan struct{}
	running  bool
	pause    bool
	debug    bool
	delay    time.Duration
}

func New() *Comp {
	c := &Comp{
		registers:    register.NewModule(),
		status:       status.NewStatusRegister(),
		stackPointer: StackStart,
		busX:         bus.New(),
		busY:         bus.New(),
		dataBus:      bus.New(),
		addrBus:      bus.New(),
		devices:      make([]connectable.Connectable, 0, 1),
		stopChan:     make(chan struct{}, 1),
		singleTicker: make(chan struct{}, 1),
	}
	return c
}

func (c *Comp) AttachTerminalComponents(terminalComponents *terminal.Components) {
	c.terminalComponents = terminalComponents
	terminalComponents.Screen.CommandPalette.ListenKeyEventEnter(func(input string) {
		c.HandleCommands(input)
	})
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
	fmt.Printf("busses, data: %x, x: %x, y: %x\n", c.dataBus.Read_16(), c.busX.Read_16(), c.busY.Read_16())
}

func (c *Comp) ReadRegister(index uint8) uint8 {
	return c.registers.Read(index)
}

func (c *Comp) ConnectDevice(dev connectable.Connectable, rangeStart uint16, rangeEnd uint16) {
	if dev == nil {
		return
	}
	dev.Attach(c.addrBus, c.dataBus, rangeStart, rangeEnd)
	c.devices = append(c.devices, dev)
}

func (c *Comp) Run() {

	c.dataBus.AttachBusChange(func() {
		c.LogMemory()
	})

	c.LogMemory()
	c.LogState()

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
	c.registers.Write(register.IndexRegA, 0)
	c.registers.Write(register.IndexRegB, 0)
	c.registers.Write(register.IndexRegC, 0)
	c.registers.Write(register.IndexRegD, 0)
	c.status.Clear()
	c.step = 0
	c.instructionRegister = 0
	c.operandRegister = 0
	c.programCounter = 0
	c.memoryAddressRegister = 0
	c.stackPointer = StackStart
	c.memoryAddressRegister = 0
	c.store = 0
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
	c.dataBus.Clear()
	c.busX.Clear()
	c.busY.Clear()
	c.addrBus.Clear()
	c.aluEnable = false
	c.aluStoreEnable = false
	c.aluDirectOut = false
	c.bridgeEnable = false
	c.bridgeDir = BRIDGE_DIR_IN
	c.rw = utils.IO_READ
}
