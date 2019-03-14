# Go

	$ cd examples
	$ make deps
	$ make
	$ sudo bin/output


# $B%i%$%V%i%j$H$7$F;H$&(B

$B;vA0$K(B libmcp2221 $B$N%$%s%9%H!<%k$,I,MW!#(B

main.go

	package main
	
	import (
	        "github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221"
	        "time"
	)
	
	func main() {
	        var iox *mcp2221.MCP23017
	        iox = mcp2221.NewMCP23017(mcp2221.MCP23017_DEFAULT_ADDR)
	        iox.DirectionA(iox.AllLow())
	
	        for true {
	                iox.LatchA(iox.AllLow())
	                time.Sleep(500 * time.Millisecond)
	
	                iox.LatchA(iox.AllHigh())
	                time.Sleep(500 * time.Millisecond)
	
	                iox.LatchA(iox.AllLow())
	                time.Sleep(500 * time.Millisecond)
	
	                iox.LatchA([]uint8{1, 0, 1, 0, 1, 0, 1, 0})
	                time.Sleep(500 * time.Millisecond)
	
	                iox.LatchA(iox.AllLow())
	                time.Sleep(500 * time.Millisecond)
	
	                iox.LatchA([]uint8{0, 1, 0, 1, 0, 1, 0, 1})
	                time.Sleep(500 * time.Millisecond)
	
	        }
	}

$B%S%k%I$H<B9T(B

	$ go get ./main.go
	$ go build ./main
	$ sudo ./main


