package main

import (
	"github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221"
	"log"
	"flag"
	"os"
	"time"
	"os/exec"
)

const DELAY=100
const BIN_SCREEN="screen"
const SERIAL_PORT="/dev/ttyACM0"

var iox *mcp2221.MCP23017
var gpio []uint8

func gpio_clear() {
	gpio = iox.AllLow()
}

func gpio_set() {
	iox.LatchA(gpio)
}

// 0=All High, 1~8: Set Low /Y0~/Y7
func gpio_tc74hc138_set(num uint8) {
	val := uint8(0)
	// 0x00 だとすべてのピンがHIGHになる
	// 0x08 から0x0fまでをいれると特定のピンがLOWになる
	if num == 0 { val = 0x00 } else { val = 0x07 + num }
	// 末尾4ビットをGPAにセット
	for i:=uint8(0); i<4; i++ {
		if val >> i & 1 > 0 { gpio[i]=1 } else { gpio[i]=0 }
	}
}

// 0=All Off, 1~8: On 0~7
func gpio_tc4051_set(num uint8) {
	val := uint8(0)
	// 0x08 だとすべてのピンがOFFになる
	// 0x00 から0x07までをいれると特定のピンがONになる
	if num == 0 { val = 0x08 } else { val = uint8(num-1) }
	// 末尾4ビットをGPAにセット
	for i:=uint8(0); i<4; i++ {
		if val >> i & 1 > 0 { gpio[i+4]=1 } else { gpio[i+4]=0 }
	}
}

func run_command(c string, p... string) error {
	cmd := exec.Command(c, p...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin  = os.Stdin
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func main() {
	p := flag.Int("p",-1, "select serial(0-7) / -1 Off")
	r := flag.Int("r",-1, "select reset(0-7) / -1 Do nothing")
	s := flag.Bool("s",false,"run screen")

	flag.Parse()

	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	iox.DirectionA(iox.AllLow())
	gpio_clear()

	if *p > 7 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *r > 7 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *p == -1 {
		gpio_tc4051_set(0)
	} else {
		n := uint8(*p)
		log.Printf("switch serial port %d\n", n)
		n += 1
		gpio_tc4051_set(n)
	}

	if *r == -1 {
		gpio_tc74hc138_set(0)
	} else {
		n := uint8(*r)
		log.Printf("reset %d\n", n)
		n += 1
		gpio_tc74hc138_set(n)
	}
	gpio_set()
	time.Sleep(100 * time.Millisecond)
	gpio_tc74hc138_set(0)
	gpio_set()

	if *s {
		log.Println("start screen")
		time.Sleep(1000 * time.Millisecond)
		run_command(BIN_SCREEN,SERIAL_PORT,"115200")
		gpio_tc4051_set(0)
		gpio_set()
	}
}

