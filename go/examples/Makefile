.PHONY: all clean deps

all: \
	bin/input \
	bin/output \
	bin/input_interrupt \
	bin/blinkbtn \
	bin/gpio_output \
	bin/gpio_input

clean:
	rm -rf bin

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

bin/blinkbtn: blinkbtn.go
	mkdir -p bin
	go build -o $@ ./blinkbtn.go

bin/gpio_output: gpio_output.go
	mkdir -p bin
	go build -o $@ ./gpio_output.go

bin/gpio_input: gpio_input.go
	mkdir -p bin
	go build -o $@ ./gpio_input.go

