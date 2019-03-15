.PHONY: all deps clean

PROGRAMS=$(addprefix ./bin/,$(notdir $(basename $(shell find . -maxdepth 1 -type f -name '*.go'))))

all: $(PROGRAMS)

deps:
	go get -v github.com/davecgh/go-spew/spew
	go get -v github.com/mamemomonga/MCP2221A-MCP23017/go/mcp2221

clean:
	rm -rf ./bin

bin/%: %.go
	mkdir -p ./bin
	go build -o $@ $<

