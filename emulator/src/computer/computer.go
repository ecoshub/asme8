package computer

import (
	"asme8/common/instruction"
	"asme8/emulator/src/bus"
	"asme8/emulator/src/connectable"
	"asme8/emulator/src/register"
	"asme8/emulator/src/status"
	"asme8/emulator/src/terminal"
	"asme8/emulator/utils"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	StackSize uint16 = 0xff
)

type Computer struct {
	Config *Config

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

	aluBus     *bus.Bus
	outputBus  *bus.Bus
	inputBus   *bus.Bus
	addrBus    *bus.Bus
	rw         uint8 // read 1 write 0
	devices    []connectable.Connectable
	stackStart uint16
	ramStart   uint16

	terminal         *terminal.Terminal
	tickEventHandle  func(pc uint16, step uint8)
	breakEventHandle func(pc uint16)
	singleTicker     chan struct{}
	stopChan         chan struct{}
	running          bool
	pause            bool
	programLoaded    bool
}

func New(conf *Config) (*Computer, error) {
	if conf.Delay == 0 {
		conf.Delay = time.Millisecond * 50
	}
	c := &Computer{
		Config:        conf,
		registers:     register.NewModule(16),
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
	err := c.CreateDevices()
	if err != nil {
		return c, err
	}
	err = c.LoadProgram()
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c *Computer) AttachTickEventHandle(f func(pc uint16, step uint8)) {
	c.tickEventHandle = f
}

func (c *Computer) AttachBreakEventHandle(f func(pc uint16)) {
	c.breakEventHandle = f
}

func (c *Computer) GetTerminal() *terminal.Terminal {
	return c.terminal
}

func (c *Computer) SetProgramLoaded() {
	c.programLoaded = true
}

func (c *Computer) IsProgramLoaded() bool {
	return c.programLoaded
}

func (c *Computer) SetStackStart(start uint16) {
	c.stackStart = start
	c.stackPointer = start
}

func (c *Computer) SetPause(enable bool) {
	c.pause = enable
}

func (c *Computer) IsPause() bool {
	return c.pause
}

func (c *Computer) IsRunning() bool {
	return c.running
}

func (c *Computer) SetDelay(delay time.Duration) {
	c.Config.Delay = delay
}

func (c *Computer) ReadRegister(index uint8) uint8 {
	return c.registers.Read_8(index)
}

func (c *Computer) GetStatusRegister() *status.StatusRegister {
	return c.status
}

func (c *Computer) GetInstructionRegister() instruction.Type {
	return c.instructionRegister
}

func (c *Computer) GetStep() uint8 {
	return c.step
}

func (c *Computer) GetRegisters() register.Module {
	return c.registers
}

func (c *Computer) GetProgramCounter() uint16 {
	return c.programCounter
}

func (c *Computer) GetStackPointer() uint16 {
	return c.stackPointer
}

func (c *Computer) GetStartOfStack() uint16 {
	return c.stackStart
}

func (c *Computer) GetRamStart() uint16 {
	return c.ramStart
}

func (c *Computer) GetMemoryAddressRegister() uint16 {
	return c.memoryAddressRegister
}

func (c *Computer) GetMemoryDataRegister() uint8 {
	return c.memoryDataRegister
}

func (c *Computer) GetOperandRegister() uint8 {
	return c.operandRegister
}

func (c *Computer) GetAluBus() *bus.Bus {
	return c.aluBus
}

func (c *Computer) GetOutputBus() *bus.Bus {
	return c.outputBus
}

func (c *Computer) GetInputBus() *bus.Bus {
	return c.inputBus
}

func (c *Computer) GetAddrBus() *bus.Bus {
	return c.addrBus
}

func (c *Computer) GetRw() uint8 {
	return c.rw
}

func (c *Computer) GetBridgeEnable() bool {
	return c.bridgeEnable
}

func (c *Computer) GetDevices() []connectable.Connectable {
	return c.devices
}

func (c *Computer) Tick() {
	c.singleTicker <- struct{}{}
}

func (c *Computer) Run() {

	if c.Config.TestMode {
		c.SetPause(false)
		c.run()
		return
	}

	if c.Config.Headless {
		c.SetPause(false)
		c.run()
		return
	} else {
		c.SetPause(true)
	}

	go c.run()

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

}

func (c *Computer) Stop() {
	if !c.running {
		return
	}
	c.stopChan <- struct{}{}
	c.running = false
}

func (c *Computer) Reset(excludeROM bool, startWithPause bool) {
	c.Stop()
	for _, dev := range c.devices {
		if excludeROM {
			if strings.HasPrefix(dev.GetName(), "ROM") {
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
	c.stackPointer = c.stackStart
	c.memoryAddressRegister = 0
	c.memoryAddressRegister = 0
	c.rw = utils.IO_READ
	c.stopChan = make(chan struct{}, 1)
	c.running = false
	c.pause = startWithPause
	c.clear()
}

func (c *Computer) Restart(startWithPause bool) {
	c.Reset(true, true)
	go c.Run()
}

func (c *Computer) tick() bool {
	defer c.clear()
	defer func() {
		if c.tickEventHandle != nil {
			c.tickEventHandle(c.programCounter, c.step)
		}
	}()

	microinstructions := CONTROL_SIGNALS[c.instructionRegister][c.step]
	for _, mi := range microinstructions {
		keep := c.execute(mi)
		if !keep {
			break
		}
	}

	if len(microinstructions) == 0 || c.instructionRegister == instruction.Type(MI_BRK) {
		if c.breakEventHandle != nil {
			c.breakEventHandle(c.programCounter)
		}
		c.running = false
		return false
	}
	return true
}

func (c *Computer) run() {
	if c.running {
		return
	}

	if c.Config.TestMode {
		c.running = true
		for {
			keep := c.tick()
			if !keep {
				return
			}
		}
	}

	t := time.NewTicker(c.Config.Delay)
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
			if !c.running {
				return
			}
			keep := c.tick()
			if !keep {
				return
			}
		case <-c.stopChan:
			return
		}
	}
}

func (c *Computer) execute(mi uint64) bool {
	f, ok := microinstructionFunctions[mi]
	if ok {
		f(c, mi)
		return true
	}
	return false
}

func (c *Computer) clear() {
	c.inputBus.Clear()
	c.aluBus.Clear()
	c.outputBus.Clear()
	c.addrBus.Clear()
	c.aluOut = false
	c.bridgeEnable = false
	c.rw = utils.IO_READ
}

func (c *Computer) getDevice(namePrefix string) (connectable.Connectable, bool) {
	for _, dev := range c.devices {
		if strings.HasPrefix(dev.GetName(), namePrefix) {
			return dev, true
		}
	}
	return nil, false
}
