package main

import (
	"../mcp2221"
	"fmt"
	"time"
//	"github.com/davecgh/go-spew/spew"
)

func main() {
	var iox *mcp2221.MCP2221
	iox = mcp2221.NewMCP2221()
	fmt.Println("Running...")

	iox.GpioDirection(0,0)

	for true {
		iox.GpioSet(0,0)
		time.Sleep(500 * time.Millisecond)

		iox.GpioSet(0,1)
		time.Sleep(500 * time.Millisecond)
	}
}

