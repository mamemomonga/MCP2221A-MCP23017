# MCP2221A + MCP23017 + LogicIC

MCP2221A + MCP23017 と ロジックICの組み合わせでいろいろ

ビルド

	$ make

[TC74HC138AP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC74HC138AP.html)(3 to 8 Line Decoder)で、任意のピンを1つだけLowにする

	$ sudo bin/74138-1
	$ sudo bin/74138-2

[TC4051BP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC4051BP.html)(Single 8ch Multiplexer/Demultiplexer)で、入出力信号の切り替え。

	$ sudo bin/4051-1

