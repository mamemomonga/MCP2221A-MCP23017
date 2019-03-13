.PHONY: all clean deps

all: bin/input bin/output bin/input_interrupt

clean:
	rm -f bin

deps:
	go get -v github.com/davecgh/go-spew/spew

bin/input: input.go
	mkdir -p bin
	go build -o $@ ./input.go

bin/output: output.go
	mkdir -p bin
	go build -o $@ ./output.go

bin/input_interrupt: input_interrupt.go
	mkdir -p bin
	go build -o $@ ./input_interrupt.go

