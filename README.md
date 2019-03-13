# MCP2221A + MCP23017

MCP2221A + MCP23017 でLEDチカチカ

ホストマシン: Linux(Debian 9)

# 回路図

そのうち書く

# ビルドと実行

	$ cd libmcp2221
	$ sudo apt install libudev-dev libusb-1.0-0-dev
	$ make
	$ sudo make install
	$ cd ..

blink

	$ cd c/blink
	$ make
	$ sudo ./blink

# LICENSE

GNU GPL v3
