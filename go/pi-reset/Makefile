.PHONY: all deps clean

all: bin/pi-reset

deps:
	go get -v github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221

clean:
	rm -rf ./bin

bin/pi-reset: main.go
	mkdir -p ./bin
	go build -o $@ $<

