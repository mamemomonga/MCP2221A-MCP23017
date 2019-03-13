package gomcp2221

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

const MCP23017_ADDR = 0x20

type GoMCP2221 struct {
	myDev *C.mcp2221_t
}

func NewGoMCP2221()(*GoMCP2221) {
	this := new(GoMCP2221)

	C.mcp2221_find(C.MCP2221_DEFAULT_VID, C.MCP2221_DEFAULT_PID, nil, nil, nil)
	this.myDev = C.mcp2221_open()

	if this.myDev == nil {
		log.Fatal("Device open failed")
	}

	// Set divider from 12MHz
	C.mcp2221_i2cDivider(this.myDev, 26);
	return this
}

func (this *GoMCP2221) i2cWaitState(i2c_state C.mcp2221_i2c_state_t) {
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

func (this *GoMCP2221) I2CWrite1byte(adr int, val uint8) {
	C.mcp2221_i2cWrite(this.myDev, (C.int)(adr), unsafe.Pointer(&val), 1, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_IDLE)

}

func (this *GoMCP2221) I2CWrite2byte(adr int, val1 uint8, val2 uint8) {
	v := []uint8{val1,val2}
	C.mcp2221_i2cWrite(this.myDev, (C.int)(adr), unsafe.Pointer(&v[0]), 2, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_IDLE)
}

func (this *GoMCP2221) I2CRead1byte(adr int) (uint8) {
	C.mcp2221_i2cRead(this.myDev, (C.int)(adr), 1, C.MCP2221_I2CRW_NORMAL)
	this.i2cWaitState(C.MCP2221_I2C_DATAREADY)
	var d uint8 = 0
	C.mcp2221_i2cGet(this.myDev, unsafe.Pointer(&d), 1);
	return d;
}

func (this *GoMCP2221) WriteMCP23017(val1 uint8, val2 uint8) {
	this.I2CWrite2byte(MCP23017_ADDR,val1,val2)
}

func (this *GoMCP2221) ReadMCP23017(val uint8) (uint8) {
	this.I2CWrite1byte(MCP23017_ADDR,val)
	return this.I2CRead1byte(MCP23017_ADDR)
}

