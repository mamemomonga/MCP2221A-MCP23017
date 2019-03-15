package main

import (
	"github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221"
	"time"
	"log"
)

const DELAY=100

var iox *mcp2221.MCP23017

func main() {
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	iox.DirectionA(iox.AllLow())

	iox.LatchA(iox.AllLow())
	time.Sleep(1000 * time.Millisecond)

	s := func (title string, v []uint8) {
		iox.LatchA(v)
		log.Printf("[%s] 0x%02x\n",title, iox.S2BV(v))
		time.Sleep(DELAY * time.Millisecond)
	}

	for true {
		s("All High",iox.Byte(0, 0, 0, 0, 0, 0, 0, 0)) // 0x00
		s("0", iox.Byte(0, 0, 0, 1, 0, 0, 0, 0)) // 0x08
		s("1", iox.Byte(1, 0, 0, 1, 0, 0, 0, 0)) // 0x09
		s("2", iox.Byte(0, 1, 0, 1, 0, 0, 0, 0)) // 0x0a
		s("3", iox.Byte(1, 1, 0, 1, 0, 0, 0, 0)) // 0x0b
		s("4", iox.Byte(0, 0, 1, 1, 0, 0, 0, 0)) // 0x0c
		s("5", iox.Byte(1, 0, 1, 1, 0, 0, 0, 0)) // 0x0d
		s("6", iox.Byte(0, 1, 1, 1, 0, 0, 0, 0)) // 0x0e
		s("7", iox.Byte(1, 1, 1, 1, 0, 0, 0, 0)) // 0x0f
	}
}

