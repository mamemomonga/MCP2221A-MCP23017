
all: bin/input bin/output bin/input_interrupt

bin/input: input/main.go
	mkdir -p bin
	cd input; make
	cp -f input/input bin/

bin/output: output/main.go
	mkdir -p bin
	cd output; make
	cp -f output/output bin/

bin/input_interrupt: input_interrupt/main.go
	mkdir -p bin
	cd input_interrupt; make
	cp -f input_interrupt/input_interrupt bin/

clean:
	rm -f bin/*

deps:
	go get -v github.com/davecgh/go-spew/spew

.PHONY: clean deps

