package main

import (
	"../mcp2221"
	"fmt"
)

func main() {
	var iox *mcp2221.MCP23017
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	fmt.Println("Running...")
	iox.DirectionA(iox.AllLow())
	iox.LatchA([]uint8{1,1,0,0,1,1,0,0})
}

