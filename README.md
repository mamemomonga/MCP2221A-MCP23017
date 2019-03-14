# MCP2221A + MCP23017

MCP2221A + MCP23017 でLEDチカチカ

ホストマシン: Linux(Debian 9)

# 回路図


![schematics.png](resource/schematics.png)

![photo.hpg](resource/photo.jpg)


# ビルドと実行

	$ cd libmcp2221
	$ sudo apt install libudev-dev libusb-1.0-0-dev
	$ make
	$ sudo make install
	$ cd ..


* [C](/c/)
* [Go](/go/)
* [libmcp2221](/libmcp2221/)
	
# LICENSE

GNU GPL v3
