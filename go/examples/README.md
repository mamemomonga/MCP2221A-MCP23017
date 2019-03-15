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

### TC74HC138AP-1.go,  TC74HC138AP-2.go

[TC74HC138AP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC74HC138AP.html)(3 to 8 Line Decoder)で、任意のピンを1つだけLowにする

配線

 Misc | MCP23017 | TC74HC138
------|----------|----------
      | GPA0     | A
      | GPA1     | B
      | GPA2     | C
      | GPA3     | G1
 GND  |          | /G2A
 GND  |          | /G2B

### TC4501BP-1.go

[TC4051BP](https://toshiba.semicon-storage.com/jp/product/logic/cmos-logic/detail.TC4051BP.html)(Single 8ch Multiplexer/Demultiplexer)で、入出力信号の切り替え。

配線

シリアルポートの切り替えを想定しているので、2つ使用します。

0~1 をラズベリーパイのTx,Rx、COMをMCP2221のURx, UTxへ。

 Misc  | MCP23017 | TC4501BP(1)| TC4501BP(2)
-------|----------|------------|-----
       | GPA4     | A          | A
       | GPA5     | B          | B
       | GPA6     | C          | C
       | GPA7     | INH        | INH
 +3.3V |          | VDD        | VDD
 GND   |          | VEE        | VEE
 GND   |          | VSS        | VSS


