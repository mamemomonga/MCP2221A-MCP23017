# MCP2221A + MCP23017 サンプル

ビルド方法

	$ make deps
	$ make
	$ sudo ./bin/PROGRAM

## bin/mcp2221\_gpio\_input

MCP2221のGPIOで入力

## bin/mcp2221\_gpio\_output

MCP2221のGPIOで出力

## bin/mcp23017\_output

MCP23017で出力

## bin/mcp23017\_input

MCP23017で入力

## bin/mcp23107\_interrupt

MCP23017で割込入力

(MCP23017の割込機能のみ使用し、MCP2221の割込機能は使っていない)

## bin/mcp23017\_blink\_btn

MCP23017でピカピカ点滅しつつ、ボタンで挙動の変更

## bin/74138-1, bin/74138-2

[TC74HC138AP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC74HC138AP.html)(3 to 8 Line Decoder)で、任意のピンを1つだけLowにする

## bin/4501-1

[TC4051BP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC4051BP.html)(Single 8ch Multiplexer/Demultiplexer)で、入出力信号の切り替え。


