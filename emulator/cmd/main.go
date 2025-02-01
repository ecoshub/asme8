package main

import (
	"asme8/assembler/src/utils"
	"asme8/common/config"
	"asme8/emulator/src/comp"
	"asme8/emulator/src/ram"
	"asme8/emulator/src/rom"
	"asme8/emulator/src/terminal"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	flagFileBin          = flag.String("load-bin", "", "load bin file to rom")
	flagFileAsm          = flag.String("load-asm", "", "assemble and load asm file to rom")
	flagDelay            = flag.Duration("delay", 10*time.Millisecond, "delay between instruction execution cycle")
	flagHeadless         = flag.Bool("headless", false, "run computer as 'headless'")
	flagMemoryConfigPath = flag.String("memory-config", "", "Path to the memory config file")
)

func main() {
	flag.Parse()

	if *flagMemoryConfigPath == "" {
		fmt.Println("Memory config path required")
		return
	}

	memoryConfig, err := config.ParseMemoryConfig(*flagMemoryConfigPath)
	if err != nil {
		fmt.Printf("Memory config parse failed. err: %s", err.Error())
		return
	}

	var program []uint8

	program, err = utils.ResolveProgram(*flagFileBin, *flagFileAsm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := Setup(memoryConfig, program, *flagHeadless)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *flagHeadless {
		HeadlessMode(c)
		return
	}

	EmulatorMode(c)
}

func Setup(memoryConfig *config.MemoryConfig, program []byte, headless bool) (*comp.Comp, error) {
	c := comp.New()

	err := config.ResolveMemoryLayout(memoryConfig.Configs)
	if err != nil {
		return nil, err
	}

	buffer := make([]uint8, 0xffff)
	copy(buffer[:], program[:])

	var term *terminal.Terminal
	var deviceRom *rom.Rom
	var deviceRam *ram.Ram
	var ramStart uint16
	for _, mc := range memoryConfig.Configs {
		if strings.HasPrefix(mc.Name, "ROM") {
			r := rom.New(mc.Size.Value)
			c.ConnectDevice(r, mc.Start.Value, mc.Size.Value)
			start := mc.Start.Value
			end := mc.Start.Value + mc.Size.Value
			r.Load(0, buffer[start:end])
			if deviceRom == nil {
				deviceRom = r
			}
			c.ProgramLoaded()
		}
		if strings.HasPrefix(mc.Name, "RAM") {
			r := ram.New(mc.Size.Value)
			c.ConnectDevice(r, mc.Start.Value, mc.Size.Value)
			if deviceRam == nil {
				ramStart = mc.Start.Value
				deviceRam = r
			}
		}
		if strings.HasPrefix(mc.Name, "VIDEO") {
			if headless {
				continue
			}
			term, err = terminal.New(int(mc.Size.Value))
			if err != nil {
				return nil, err
			}
			c.ConnectDevice(term.Screen, mc.Start.Value, mc.Size.Value)
			// screen components such as panels attaching to computer
			c.AttachTerminal(term)
		}
		if strings.HasPrefix(mc.Name, "KEYBOARD") {
			if headless {
				continue
			}
			if term == nil {
				return nil, fmt.Errorf("can not define KEYBOARD without define VIDEO first")
			}
			c.ConnectDevice(term.Keyboard, mc.Start.Value, mc.Size.Value)
		}
		// c.Logf("DEVICE CONNECTED. name: %s, start: %04x, size: %04x, end: %04x\n", mc.Name, mc.Start.Value, mc.Size.Value, mc.Start.Value+mc.Size.Value)
		// fmt.Printf("DEVICE CONNECTED. name: %s, start: %04x, size: %04x, end: %04x\n", mc.Name, mc.Start.Value, mc.Size.Value, mc.Start.Value+mc.Size.Value)
	}

	if deviceRom == nil {
		return nil, fmt.Errorf("memory config must define ROM")
	}

	if deviceRam == nil {
		return nil, fmt.Errorf("memory config must define RAM")
	}

	// setting base clock speed
	c.SetDelay(*flagDelay)

	c.SetStackStart(ramStart + 0xff)

	return c, nil
}

func EmulatorMode(c *comp.Comp) {

	c.SetPause(true)
	c.Help()

	go c.Run()
	c.LogWithStyle("TIP: type 's' and press 'ENTER' to start the program", comp.DefaultStyle6)

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	terminal.ResetScreen()
}

func HeadlessMode(c *comp.Comp) {

	c.SetDelay(*flagDelay)
	c.SetPause(false)
	c.SetDebug(true)

	go c.Run()
	c.LogWithStyle("TIP: type 's' and press 'ENTER' to start the program", comp.DefaultStyle6)

	// standard main thread blocker
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}
