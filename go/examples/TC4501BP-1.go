package main

import (
	"github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221"
	"log"
	"flag"
	"os"
)

const DELAY=100

var iox *mcp2221.MCP23017

var gpio []uint8

func gpio_clear() {
	gpio = iox.AllLow()
}

func gpio_set() {
	iox.LatchA(gpio)
	log.Printf("GPA: 0x%02x\n", iox.S2BV(gpio))
}

// 0=All High, 1~8: Set Low /Y0~/Y7
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

func main() {

	var portnum uint8
	var alloff bool = false
	{
		p := flag.Int("p",-1,"select port(0-7)")
		x := flag.Bool("x",false,"all off")
		flag.Parse()

		if *x == false && *p == -1 {
			flag.PrintDefaults()
			os.Exit(1)
		}

		if *p > 7 {
			flag.PrintDefaults()
			os.Exit(1)
		}

		alloff = *x
		portnum = uint8(*p)
	}

	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	iox.DirectionA(iox.AllLow())
	gpio_clear()

	if alloff {
		log.Println("All port off")
		gpio_tc4051_set(0)
	} else {
		log.Printf("switched port %d\n", portnum)
		portnum += 1
		gpio_tc4051_set(portnum)
	}
	gpio_set()
}

