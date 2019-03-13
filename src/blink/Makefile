PROJECT=blink

SOURCES= \
	../easymcp2221.c main.c

CFLAGS= \
	-c \
	-Wall \
	-Wextra \
	-Wstrict-prototypes \
	-Wunused-result \
	-O3 \
	-std=gnu99 \
	-fmessage-length=0

LDFLAGS= \
	-s

LDLIBS= \
	-lmcp2221

EXECUTABLE=$(PROJECT)

CC=gcc

OBJECTS=$(SOURCES:.c=.o)


all: $(SOURCES) $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
	$(CC) $(LDFLAGS) $(OBJECTS) $(LDLIBS) -o $@

.c.o:
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm -rf *.o $(EXECUTABLE)

.PHONY: clean all
