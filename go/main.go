package main

import (
	"./mcp2221"
	"fmt"
	"time"
	"github.com/davecgh/go-spew/spew"
)

var iox *mcp2221.MCP23017

func output() {
	iox.DirectionA([]uint8{0,0,0,0,0,0,0,0})
	iox.LatchA([]uint8{1,1,0,0,1,1,0,0})
}
func input() {
	iox.DirectionB([]uint8{1,1,1,1,1,1,1,1})
	iox.PullUpB([]uint8{1,1,1,1,1,1,1,1})
	for true {
		spew.Dump( iox.GpioB() )
		time.Sleep(100 * time.Millisecond)
	}
}
func input_interrupt() {
	iox.DirectionB([]uint8{1,1,1,1,1,1,1,1})
	iox.PullUpB([]uint8{1,1,1,1,1,1,1,1})
	iox.InterruptB(func(val []uint8) {
		spew.Dump(val)
	})
}

func main() {
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	fmt.Println("Running...")
	//output()
	//input()
	input_interrupt()
}

