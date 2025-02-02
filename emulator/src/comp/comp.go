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
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	DefaultMemorySize = 1 << 16
)

var (
	StackStart uint16
)

type Comp struct {
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

	codeLines          map[uint16]string
	activeCodeLine     int
	lastActiveCodeLine int
	programLoaded      bool
	stopChan           chan struct{}
	running            bool
	pause              bool
}

func New(conf *Config) (*Comp, error) {
	c := &Comp{
		Config:        conf,
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
	err := c.CreateDevices()
	if err != nil {
		return nil, err
	}
	err = c.LoadProgram()
	if err != nil {
		return nil, err
	}
	err = c.ReadSymbolFile()
	if err != nil {
		return nil, err
	}
	return c, nil
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
	c.Config.Debug = enable
}

func (c *Comp) SetDelay(delay time.Duration) {
	c.Config.Delay = delay
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

func (c *Comp) Run() {

	if c.Config.Test {
		c.SetPause(false)
		c.run()
		return
	}

	if c.Config.Headless {
		c.SetPause(false)
		c.SetDebug(true)
	} else {
		c.SetPause(true)
		c.Help()
	}

	go c.run()

	if !c.Config.Headless {
		c.LogWithStyle("TIP: type 's' and press 'ENTER' to start the program", DefaultStyle6)
	}

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	if !c.Config.Headless {
		terminal.ResetScreen()
	}
}

func (c *Comp) run() {
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
		keep := c.execute(mi)
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

func (c *Comp) execute(mi uint64) bool {
	f, ok := microinstructionFunctions[mi]
	if ok {
		f(c, mi)
		return true
	}
	return false
}

func (c *Comp) Reset(excludeROM bool, startWithPause bool) {
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
