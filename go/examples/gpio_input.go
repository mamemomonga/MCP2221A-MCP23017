package main

import (
	"../mcp2221"
	"fmt"
	"time"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var iox *mcp2221.MCP2221
	iox = mcp2221.NewMCP2221()
	fmt.Println("Running...")

	iox.GpioDirection(1,1)

	for true {
		spew.Dump( iox.GpioGet(1) )
		time.Sleep(500 * time.Millisecond)
	}
}

