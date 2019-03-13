package main

import (
	"../mcp2221"
	"fmt"
	"time"
	"github.com/davecgh/go-spew/spew"
)

const DELAY1=30

var iox *mcp2221.MCP23017

func blink() {
	iox.DirectionA(iox.AllLow())
	for true {
		for i:=0; i<8; i++ {
			v := iox.AllLow()
			v[i]=1
			iox.LatchA(v)
			time.Sleep(DELAY1 * time.Millisecond)
		}
		for i:=0; i<3; i++ {
			iox.LatchA(iox.AllLow())
			time.Sleep(DELAY1 * time.Millisecond)
			iox.LatchA(iox.AllHigh())
			time.Sleep(DELAY1 * time.Millisecond)
		}
		for i:=7; i>=0; i-- {
			v := iox.AllLow()
			v[i]=1
			iox.LatchA(v)
			time.Sleep(DELAY1 * time.Millisecond)
		}
	}
}

func buttons() {
	iox.DirectionB(iox.AllHigh())
	iox.PullUpB(iox.AllHigh())
	iox.InterruptB(func(val []uint8) {
		spew.Dump(val)
	})
}

func main() {
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	fmt.Println("Running...")

	go blink()
	buttons()

}

