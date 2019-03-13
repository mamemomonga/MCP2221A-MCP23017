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

func input_interrupt() {
	gm.WriteMCP23017(0x01,0xFF) // IODIRB
	gm.WriteMCP23017(0x0D,0xFF) // GPPUB

	gm.WriteMCP23017(0x05,0xFF); // GPINTENB
	gm.WriteMCP23017(0x07,0xFF); // DEFVALB
	gm.WriteMCP23017(0x09,0xFF); // INTCONB

	var prev uint8 = 0x00

	cb := func() {
		val := gm.ReadMCP23017(0x13) // GPIOB
		if prev != val {
			fmt.Printf("Value: %02x\n",val)
		}
		prev = val
	}

	gm.GpioI2CInterrupt(cb)

}

func main() {
	gm = gomcp2221.NewGoMCP2221()
	fmt.Println("Running...")
	input_interrupt()
}

