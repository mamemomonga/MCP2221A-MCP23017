package main

/*
  MCP2221A + MCP23017 + TC74HC138
  GPAの先頭4ビットのみ使い、残りを別の用途として使えるようにする。

  $ go get ./74138-2.go
  $ go build ./74138-2.go
  $ sudo ./74138-2

Wireing:
  MCP23017 GPA0 - TC74HC138 A
  MCP23017 GPA1 - TC74HC138 B
  MCP23017 GPA2 - TC74HC138 C
  MCP23017 GPA3 - TC74HC138 G1
  GND           - TC74HC138 /G2A
  GND           - TC74HC138 /G2B
*/

import (
	"github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221"
	"time"
	"log"
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
func gpio_74138_set(num uint8) {
	val := uint8(0)
	// 0x00 だとすべてのピンがHIGHになる
	// 0x08 から0x0fまでをいれると特定のピンがLOWになる
	if num == 0 { val = 0x00 } else { val = 0x07 + num }
	// 末尾4ビットをGPAにセット
	if val >> 0 & 1 > 0 { gpio[0]=1 } else { gpio[0]=0 }
	if val >> 1 & 1 > 0 { gpio[1]=1 } else { gpio[1]=0 }
	if val >> 2 & 1 > 0 { gpio[2]=1 } else { gpio[2]=0 }
	if val >> 3 & 1 > 0 { gpio[3]=1 } else { gpio[3]=0 }
}

func main() {
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	iox.DirectionA(iox.AllLow())

	iox.LatchA(iox.AllLow())
	time.Sleep(1000 * time.Millisecond)

	for true {
		for i:=uint8(0);i<=8;i++ {
			log.Printf("74138: %d\n",i)
			gpio_clear()
			gpio_74138_set(i)
			gpio_set()
			time.Sleep(DELAY * time.Millisecond)
		}
	}
}

