package mcp2221

import (
	"log"
	"time"
	//	"github.com/davecgh/go-spew/spew"
)

const DEBUG = false
const MCP23017_DEFAULT_ADDR = 0x20

type MCP23017 struct {
	devAddr         int
	MCP2221         *MCP2221
	intervalPrevVal uint8
}

func NewMCP23017(addr int) *MCP23017 {
	this := new(MCP23017)
	this.MCP2221 = NewMCP2221()
	this.devAddr = addr
	return this
}

func (this *MCP23017) Write(val1 uint8, val2 uint8) {
	this.MCP2221.I2CWrite2byte(this.devAddr, val1, val2)
}

func (this *MCP23017) Read(val uint8) uint8 {
	this.MCP2221.I2CWrite1byte(this.devAddr, val)
	return this.MCP2221.I2CRead1byte(this.devAddr)
}

func (this *MCP23017) a2bv(in []uint8) uint8 {
	var v uint8
	var i uint8
	for i = 0; i < 8; i++ {
		v = v | (in[i] << i)
	}
	if DEBUG {
		log.Printf("BV: 0x%02X\n", v)
	}
	return v
}

func (this *MCP23017) bv2a(in uint8) []uint8 {
	v := this.AllLow()
	var i uint8
	for i = 0; i < 8; i++ {
		if (in & (1 << i)) > 0 {
			v[i] = 1
		} else {
			v[i] = 0
		}
	}
	return v
}

func (this *MCP23017) AllHigh() []uint8 {
	return []uint8{1, 1, 1, 1, 1, 1, 1, 1}
}

func (this *MCP23017) AllLow() []uint8 {
	return []uint8{0, 0, 0, 0, 0, 0, 0, 0}
}

// ICON.BANK=0 専用

func (this *MCP23017) DirectionA(v []uint8) {
	this.Write(0x00, this.a2bv(v)) // IODIRA
}
func (this *MCP23017) DirectionB(v []uint8) {
	this.Write(0x01, this.a2bv(v)) // IODIRB
}
func (this *MCP23017) LatchA(v []uint8) {
	this.Write(0x14, this.a2bv(v)) // OLATA
}
func (this *MCP23017) LatchB(v []uint8) {
	this.Write(0x15, this.a2bv(v)) // OLATB
}
func (this *MCP23017) PullUpA(v []uint8) {
	this.Write(0x0c, this.a2bv(v)) // GPPUA
}
func (this *MCP23017) PullUpB(v []uint8) {
	this.Write(0x0d, this.a2bv(v)) // GPPUB
}
func (this *MCP23017) GpioA() (v []uint8) {
	return this.bv2a(this.Read(0x12)) // GPIOA
}
func (this *MCP23017) GpioB() (v []uint8) {
	return this.bv2a(this.Read(0x13)) // GPIOB
}

// MCP2221の割り込み機能を使わず、MCP23017の割り込みのみ使う
// MCP2221 GP0 - MCP23017 INTB
func (this *MCP23017) InterruptB(cb func([]uint8)) {
	this.Write(0x05, 0xFF) // GPINTENB
	this.Write(0x07, 0xFF) // DEFVALB
	this.Write(0x09, 0x00) // INTCONB

	this.MCP2221.GpioDirection(0, 1) // GP0 Input
	this.MCP2221.GpioSet(0, 1)       // GP0 High

	for true {
		intr := this.MCP2221.GpioGet(0)
		if intr == 0 {
			time.Sleep(1 * time.Millisecond) // 誤検出防止
			val := this.Read(0x13)           // GPIOB
			if this.intervalPrevVal != val {
				cb(this.bv2a(val))
			}
			this.intervalPrevVal = val
		}
	}
}
