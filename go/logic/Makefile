.PHONY: all

all: bin/74138-1 bin/74138-2

clean:
	rm -rf bin

bin/74138-1: 74138-1.go
	mkdir -p bin
	go build -o $@ ./$<

bin/74138-2: 74138-2.go
	mkdir -p bin
	go build -o $@ ./$<

