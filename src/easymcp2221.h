#ifndef EASYMCP2221_H_
#define EASYMCP2221_H_

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <libmcp2221.h>

#define sleep_ms(a) { usleep( a * 1000); }

int init_app(void);
void i2cWrite2byte(uint8_t, uint8_t, uint8_t);

mcp2221_t*   myMCP2221;

#endif /* EASYMCP2221_H_ */