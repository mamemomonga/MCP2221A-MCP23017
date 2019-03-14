package main

import (
	"../mcp2221"
	"fmt"
	"time"
	"sync"
	"github.com/davecgh/go-spew/spew"
)

const DELAY1=30
const DELAY2=60

var iox *mcp2221.MCP23017

var blinkm   *sync.Mutex
var patterns [][]uint8
var patternc []uint8

func blink() {
	iox.DirectionA(iox.AllLow())

	styles := [5]func(){}

	styles[0] = func() {
		time.Sleep(DELAY1 * time.Millisecond)
	}

	styles[1] = func() {
		for i:=0; i<8; i++ {
			v := iox.AllLow()
			v[i]=1
			iox.LatchA(v)
			time.Sleep(DELAY1 * time.Millisecond)
		}
	}

	styles[2] = func() {
		for i:=7; i>=0; i-- {
			v := iox.AllLow()
			v[i]=1
			iox.LatchA(v)
			time.Sleep(DELAY1 * time.Millisecond)
		}
	}

	styles[3] = func() {
		iox.LatchA(iox.AllLow())
		time.Sleep(DELAY2 * time.Millisecond)
	}

	styles[4] = func() {
		iox.LatchA(iox.AllHigh())
		time.Sleep(DELAY2 * time.Millisecond)
	}

	for true {

		blinkm.Lock()
		pt := patternc
		blinkm.Unlock()

		for _,v := range pt {
			styles[v]()
		}
	}
}

func buttons() {
	iox.DirectionB(iox.AllHigh())
	iox.PullUpB(iox.AllHigh())
	iox.InterruptB(func(val []uint8) {
		spew.Dump(val)
		if val[0] == 0 {
			blinkm.Lock()
			patternc = patterns[0]
			blinkm.Unlock()
		}
		if val[1] == 0 {
			blinkm.Lock()
			patternc = patterns[1]
			blinkm.Unlock()
		}
	})
}

func main() {
	iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	fmt.Println("Running...")

	blinkm  = new(sync.Mutex)

	patterns = make([][]uint8,10)
	patterns[0] = []uint8{ 0,0,1,0,0,2,0,0, 3,4,3,4 }
	patterns[1] = []uint8{ 3,4,3,4 }

	patternc = patterns[0]

	go buttons()
	blink()

}

