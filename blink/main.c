#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <libmcp2221.h>

#define sleep_ms(a) { usleep( a * 1000); }

int main(void);
int init_app(void);
void i2cWrite2byte(uint8_t, uint8_t, uint8_t);

mcp2221_t*   myMCP2221;

/* ---------------------------------- */

int init_app(void) {
	mcp2221_find(MCP2221_DEFAULT_VID, MCP2221_DEFAULT_PID, NULL, NULL, NULL);
	myMCP2221 = mcp2221_open();

	if(!myMCP2221) {
		mcp2221_exit();
		puts("Device opening failed");
		return 1;
	}

	// Set divider from 12MHz
	mcp2221_i2cDivider(myMCP2221, 26);
	return 0;
}

void i2cWrite2byte(uint8_t adr, uint8_t val1, uint8_t val2) {
	uint8_t v[]={ val1, val2 };
	mcp2221_i2cWrite(myMCP2221, adr, v, 2, MCP2221_I2CRW_NORMAL);
	mcp2221_i2c_state_t state = MCP2221_I2C_IDLE;
	while(1) {
		if(mcp2221_i2cState(myMCP2221, &state) != MCP2221_SUCCESS) puts("ERROR");
		if(state == MCP2221_I2C_IDLE) break;
	}
}

int main(void) {
	if(init_app()){ return 1; }

	uint8_t sweep[]={0x00, 0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80};
	uint8_t szSweep=sizeof(sweep);

	i2cWrite2byte(0x20,0x00,0x00);
	while(1) {
		for(uint8_t i=0; i<szSweep; i++) {
			i2cWrite2byte(0x20,0x14,sweep[i]);
			sleep_ms(40);
		}
		for(uint8_t i=0; i<3; i++) {
			i2cWrite2byte(0x20,0x14,0xFF); sleep_ms(50);
			i2cWrite2byte(0x20,0x14,0x00); sleep_ms(50);
		}
		for(uint8_t i=szSweep; i>0; i--) {
			i2cWrite2byte(0x20,0x14,sweep[i]);
			sleep_ms(40);
		}
		for(uint8_t i=0; i<3; i++) {
			i2cWrite2byte(0x20,0x14,0xFF); sleep_ms(50);
			i2cWrite2byte(0x20,0x14,0x00); sleep_ms(50);
		}
	}
	return 0;
}

