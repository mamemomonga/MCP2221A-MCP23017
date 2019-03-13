package main

import (
	"./gomcp2221"
	"fmt"
	"time"
)

var gm *gomcp2221.GoMCP2221

func output() {
	gm.WriteMCP23017(0x00,0x00)
	gm.WriteMCP23017(0x14,0x15)
}

func input() {
	gm.WriteMCP23017(0x01,0xFF) // IODIRB
	gm.WriteMCP23017(0x0D,0xFF) // GPPUB

	for true {
		val := gm.ReadMCP23017(0x13) // GPIOB
		fmt.Printf("Value: %02x\n",val)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	gm = gomcp2221.NewGoMCP2221()
	input()
	fmt.Println("Running...")
}

