# MCP2221A + MCP23017 サンプル

ビルド方法

	$ make deps
	$ make
	$ sudo ./bin/PROGRAM

### MCP2221-gpio-input.go

MCP2221のGPIOで入力

### MCP2221-gpio-output.go

MCP2221のGPIOで出力

### MCP23017-output.go

MCP23017で出力

### MCP23017-input.go

MCP23017で入力

### MCP23107-interrupt.go

MCP23017で割込入力

(MCP23017の割込機能のみ使用し、MCP2221の割込機能は使っていない)

### MCP23017-blink-btn.go

MCP23017でピカピカ点滅しつつ、ボタンで挙動の変更

## TC74HC138AP-1.go,  TC74HC138AP-2.go

[TC74HC138AP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC74HC138AP.html)(3 to 8 Line Decoder)で、任意のピンを1つだけLowにする

## TC4501BP-1.go

[TC4051BP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC4051BP.html)(Single 8ch Multiplexer/Demultiplexer)で、入出力信号の切り替え。


