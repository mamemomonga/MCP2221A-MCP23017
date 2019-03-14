# Go

	$ cd examples
	$ make deps
	$ make
	$ sudo bin/output


# ライブラリとして使う

事前に [libmcp2221](../libmcp2221) のインストールが必要。

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

ビルドと実行

	$ go get ./main.go
	$ go build ./main
	$ sudo ./main


