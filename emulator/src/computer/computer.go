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

var (
	// set by device
	StackStart uint16 = 0
	// interrupt high bytes are always zero
	IntVec0AddrLow uint8 = 0x20
	IntVec1AddrLow uint8 = 0x22
	IntVec2AddrLow uint8 = 0x24
	IntVec3AddrLow uint8 = 0x26
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
	// stackPointer          uint16
	memoryDataRegister uint8

	aluOut       bool
	bridgeEnable bool

	irqLine        bool
	irqAckFlipFlop bool

	aluBus     *bus.Bus
	outputBus  *bus.Bus
	inputBus   *bus.Bus
	addrBus    *bus.Bus
	rw         uint8 // read 1 write 0
	devices    []connectable.Connectable
	stackStart uint16
	ramStart   uint16

	terminal        *terminal.Terminal
	tickEventHandle func(pc uint16, step uint8)
	haltEventHandle func(pc uint16)
	singleTicker    chan struct{}
	stopChan        chan struct{}
	running         bool
	pause           bool
	programLoaded   bool
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
	c.haltEventHandle = f
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
	c.registers.Write(instruction.REGISTER_OPCODE_SPL, uint8(start))
	c.registers.Write(instruction.REGISTER_OPCODE_SPH, uint8(start>>8))
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
	stack := uint16(0)
	stack += uint16(c.registers.Read_8(instruction.REGISTER_OPCODE_SPH))
	stack = (stack << 8)
	stack += uint16(c.registers.Read_8(instruction.REGISTER_OPCODE_SPL))
	return stack
}

func (c *Computer) SetStackPointer(val uint16) {
	c.registers.Write(instruction.REGISTER_OPCODE_SPL, uint8(val))
	c.registers.Write(instruction.REGISTER_OPCODE_SPH, uint8(val>>8))
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

func (c *Computer) InterruptRequest() {
	c.irqLine = true
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
	c.SetStackStart(StackStart)
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

	if c.tickEventHandle != nil {
		c.tickEventHandle(c.programCounter, c.step)
	}

	irq := checkIrq(c.irqLine, c.status.Flag())
	if irq && c.step == 0 {
		c.irqAckFlipFlop = true
	}
	jmp := checkJmp(c.instructionRegister, c.status.Flag())

	var microinstructions []uint64
	if c.irqAckFlipFlop && irq {
		microinstructions = IRQ_MICRO_INSTRUCTIONS[c.step]
	} else {
		if jmp {
			microinstructions = GENERIC_CONDITIONAL_JMP[c.step]
		} else {
			microinstructions = CONTROL_SIGNALS[c.instructionRegister][c.step]
		}
	}
	for _, mi := range microinstructions {
		keep := c.execute(mi)
		if !keep {
			break
		}
	}

	if len(microinstructions) == 0 || c.instructionRegister == instruction.INST_HLT_IMPL {
		if c.haltEventHandle != nil {
			c.haltEventHandle(c.programCounter)
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

func checkJmp(ir uint8, sf uint8) bool {
	statusMask := uint8(0)
	not := false
	switch ir {
	case instruction.INST_JZ_IMM16:
		statusMask = status.STATUS_FLAG_ZERO
	case instruction.INST_JNZ_IMM16:
		statusMask = status.STATUS_FLAG_ZERO
		not = true
	case instruction.INST_JC_IMM16:
		statusMask = status.STATUS_FLAG_CARRY
	case instruction.INST_JNC_IMM16:
		statusMask = status.STATUS_FLAG_CARRY
		not = true
	case instruction.INST_JS_IMM16:
		statusMask = status.STATUS_FLAG_SIGN
	case instruction.INST_JNS_IMM16:
		statusMask = status.STATUS_FLAG_SIGN
		not = true
	default:
		return false
	}

	if (sf&statusMask > 0) == !not {
		return true
	}

	return false
}

func checkIrq(irq bool, sf uint8) bool {
	if !irq {
		return false
	}
	if (sf & status.STATUS_FLAG_INTERRUPT_ENABLE) == 0 {
		return false
	}
	return true
}
