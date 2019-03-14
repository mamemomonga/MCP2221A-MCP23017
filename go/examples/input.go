package main

import (
	"../mcp2221"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"time"
)

func main() {
	var iox *mcp2221.MCP23017
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	fmt.Println("Running...")

	iox.DirectionB(iox.AllHigh())
	iox.PullUpB(iox.AllHigh())

	for true {
		spew.Dump(iox.GpioB())
		time.Sleep(10 * time.Millisecond)
	}

}
