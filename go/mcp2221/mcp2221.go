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
	"sync"
//	"github.com/davecgh/go-spew/spew"
)

type MCP2221 struct {
	myDev *C.mcp2221_t
	m *sync.Mutex
}

func NewMCP2221()(*MCP2221) {
	this := new(MCP2221)
	this.m = new(sync.Mutex)

	C.mcp2221_find(C.MCP2221_DEFAULT_VID, C.MCP2221_DEFAULT_PID, nil, nil, nil)
	this.myDev = C.mcp2221_open()

	if this.myDev == nil {
		log.Fatal("Device open failed")
	}

	// Set divider from 12MHz
	C.mcp2221_i2cDivider(this.myDev, 26);
	return this
}
func (this *MCP2221) Lock() {
	this.m.Lock()
}
func (this *MCP2221) Unlock() {
	this.m.Unlock()
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
	this.Lock()
	C.mcp2221_i2cWrite(this.myDev, (C.int)(adr), unsafe.Pointer(&val), 1, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_IDLE)
	this.Unlock()
}

func (this *MCP2221) I2CWrite2byte(adr int, val1 uint8, val2 uint8) {
	v := []uint8{val1,val2}
	this.Lock()
	C.mcp2221_i2cWrite(this.myDev, (C.int)(adr), unsafe.Pointer(&v[0]), 2, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_IDLE)
	this.Unlock()
}

func (this *MCP2221) I2CRead1byte(adr int) (uint8) {
	this.Lock()
	C.mcp2221_i2cRead(this.myDev, (C.int)(adr), 1, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_DATAREADY)
	var d uint8 = 0
	C.mcp2221_i2cGet(this.myDev, unsafe.Pointer(&d), 1)
	this.Unlock()
	return d
}

func (this *MCP2221) GpioDirection(gpio uint8, direction uint8 ) {
	this.Lock()
	var gpioConf C.mcp2221_gpioconfset_t
	gpioConf = C.mcp2221_GPIOConfInit()
	switch gpio {
		case 0: gpioConf.conf[0].gpios = C.MCP2221_GPIO0
		case 1: gpioConf.conf[0].gpios = C.MCP2221_GPIO1
		case 2: gpioConf.conf[0].gpios = C.MCP2221_GPIO2
		case 3: gpioConf.conf[0].gpios = C.MCP2221_GPIO3
	}
	switch direction {
		case 0: gpioConf.conf[0].direction = C.MCP2221_GPIO_DIR_OUTPUT
		case 1: gpioConf.conf[0].direction = C.MCP2221_GPIO_DIR_INPUT
	}
	gpioConf.conf[0].mode = C.MCP2221_GPIO_MODE_GPIO
	gpioConf.conf[0].value = C.MCP2221_GPIO_VALUE_LOW

	C.mcp2221_setGPIOConf(this.myDev, &gpioConf)
	this.Unlock()
}

func (this *MCP2221) GpioSet(gpio uint8, value uint8 ) {
	pin := C.MCP2221_GPIO0
	switch gpio {
		case 0: pin = C.MCP2221_GPIO0
		case 1: pin = C.MCP2221_GPIO1
		case 2: pin = C.MCP2221_GPIO2
		case 3: pin = C.MCP2221_GPIO3
	}
	v := C.MCP2221_GPIO_VALUE_LOW
	switch value {
		case 0: v = C.MCP2221_GPIO_VALUE_LOW
		case 1: v = C.MCP2221_GPIO_VALUE_HIGH
	}
	this.Lock()
	res := C.mcp2221_setGPIO(this.myDev, (C.mcp2221_gpio_t)(pin), (C.mcp2221_gpio_value_t)(v));
	if res != C.MCP2221_SUCCESS {
		this.handle_mcp2221_error(res)
	}
	this.Unlock()
}

func (this *MCP2221) GpioGet(gpio uint8) uint8 {
	this.Lock()
	var values [4]C.mcp2221_gpio_value_t
	res := C.mcp2221_readGPIO(this.myDev,&values[0])

	if res != C.MCP2221_SUCCESS {
		this.handle_mcp2221_error(res)
	}
	this.Unlock()

	return (uint8)(values[gpio])
}

func (this *MCP2221) GpioI2CInterrupt(cb func()) {

	this.Lock()
	var gpioConf C.mcp2221_gpioconfset_t
	gpioConf = C.mcp2221_GPIOConfInit()

	gpioConf.conf[0].gpios = C.MCP2221_GPIO1
	gpioConf.conf[0].mode  = C.MCP2221_GPIO_MODE_ALT3
	C.mcp2221_setGPIOConf(this.myDev, &gpioConf)
	C.mcp2221_setInterrupt(this.myDev, C.MCP2221_INT_TRIG_FALLING, 1)
	this.Unlock()

	var res C.mcp2221_error
	for true {
		var interrupt C.int
		this.Lock()
		res = C.mcp2221_readInterrupt(this.myDev, &interrupt)
		this.Unlock()
		if res != C.MCP2221_SUCCESS {
			break
		}
		if interrupt != 0 {
			this.Lock()
			res = C.mcp2221_clearInterrupt(this.myDev)
			this.Unlock()
			if res != C.MCP2221_SUCCESS {
				break
			}
			cb()
		}
	}
	this.handle_mcp2221_error(res)
}

func (this *MCP2221) handle_mcp2221_error(res C.mcp2221_error) {
	switch res {
		case C.MCP2221_SUCCESS:
			log.Println("No error")

		case C.MCP2221_ERROR:
			log.Fatal("General error")

		case C.MCP2221_INVALID_ARG:
			log.Fatal("Invalid argument, probably null pointer")

		case C.MCP2221_ERROR_HID:
			log.Fatal("USB HID Error")

		default:
			log.Fatal("Unknown error %d\n", res)
	}
}


