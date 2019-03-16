# pi-reset

* 8台のRaspberry Pi をリセットできます
* 8台のRaspberry Pi のUARTコンソールを切り替えます

# ビルド

	$ make deps
	$ make
	$ bin/pi-hwctl --help

# 使い方

ポート0 のRaspberry Pi に接続

	$ bin/pi-hwctl -p 0

ポート0 のRaspberry Pi に接続して、GNU Screenを起動

	$ bin/pi-hwctl -p 0 -s

ポート0 のRaspberry Pi をリセット

	$ bin/pi-hwctl -r 0

ポート0 のRaspberry Pi に接続、ポート0のRaspberry Piを接続、GNU Screenを起動

	$ bin/pi-hwctl -p 0 -r 0 -s

