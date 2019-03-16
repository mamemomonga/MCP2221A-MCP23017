.PHONY: all deps clean

PROGNAME=pi-hwctl

all: bin/$(PROGNAME)

deps:
	go get -v github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221

clean:
	rm -f ./bin/$(PROGNAME)

bin/$(PROGNAME): main.go
	mkdir -p ./bin
	go build -o $@ $<

