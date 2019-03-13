package main

import (
	"../mcp2221"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var iox *mcp2221.MCP23017
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	fmt.Println("Running...")

	iox.DirectionB([]uint8{1,1,1,1,1,1,1,1})
	iox.PullUpB([]uint8{1,1,1,1,1,1,1,1})
	iox.InterruptB(func(val []uint8) {
		spew.Dump(val)
	})
}

