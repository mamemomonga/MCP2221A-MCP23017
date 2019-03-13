package mcp2221

/*
#cgo LDFLAGS: -lmcp2221
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <libmcp2221.h>
*/
import "C"
import (
	"log"
	"unsafe"
//	"github.com/davecgh/go-spew/spew"
)

type MCP2221 struct {
	myDev *C.mcp2221_t
}

func NewMCP2221()(*MCP2221) {
	this := new(MCP2221)

	C.mcp2221_find(C.MCP2221_DEFAULT_VID, C.MCP2221_DEFAULT_PID, nil, nil, nil)
	this.myDev = C.mcp2221_open()

	if this.myDev == nil {
		log.Fatal("Device open failed")
	}

	// Set divider from 12MHz
	C.mcp2221_i2cDivider(this.myDev, 26);
	return this
}

func (this *MCP2221) i2cWaitState(i2c_state C.mcp2221_i2c_state_t) {
	var state C.mcp2221_i2c_state_t
	state = C.MCP2221_I2C_IDLE
	for true {
		if C.mcp2221_i2cState(this.myDev, &state) != C.MCP2221_SUCCESS {
			log.Fatal("mcp2221_i2cState Not Success")
		}
		if state == i2c_state {
			break
		}
	}
}

func (this *MCP2221) I2CWrite1byte(adr int, val uint8) {
	C.mcp2221_i2cWrite(this.myDev, (C.int)(adr), unsafe.Pointer(&val), 1, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_IDLE)

}

func (this *MCP2221) I2CWrite2byte(adr int, val1 uint8, val2 uint8) {
	v := []uint8{val1,val2}
	C.mcp2221_i2cWrite(this.myDev, (C.int)(adr), unsafe.Pointer(&v[0]), 2, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_IDLE)
}

func (this *MCP2221) I2CRead1byte(adr int) (uint8) {
	C.mcp2221_i2cRead(this.myDev, (C.int)(adr), 1, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_DATAREADY)
	var d uint8 = 0
	C.mcp2221_i2cGet(this.myDev, unsafe.Pointer(&d), 1)
	return d
}

func (this *MCP2221) GpioI2CInterrupt(cb func()) {

	var gpioConf C.mcp2221_gpioconfset_t
	gpioConf = C.mcp2221_GPIOConfInit()

	gpioConf.conf[0].gpios = C.MCP2221_GPIO1
	gpioConf.conf[0].mode  = C.MCP2221_GPIO_MODE_ALT3
	C.mcp2221_setGPIOConf(this.myDev, &gpioConf)
	C.mcp2221_setInterrupt(this.myDev, C.MCP2221_INT_TRIG_FALLING, 1)

	var res C.mcp2221_error
	for true {
		var interrupt C.int
		res = C.mcp2221_readInterrupt(this.myDev, &interrupt)
		if res != C.MCP2221_SUCCESS {
			break
		}
		if interrupt != 0 {
			res = C.mcp2221_clearInterrupt(this.myDev)
			if res != C.MCP2221_SUCCESS {
				break
			}
			cb()
		}
	}
	switch res {
		case C.MCP2221_SUCCESS:
			log.Println("No error")
		case C.MCP2221_ERROR:
			log.Println("General error")
		case C.MCP2221_INVALID_ARG:
			log.Println("Invalid argument, probably null pointer")
		case C.MCP2221_ERROR_HID:
			log.Println("USB HID Error")
		default:
			log.Printf("Unknown error %d\n", res)
	}
}

