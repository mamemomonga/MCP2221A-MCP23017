#include "easymcp2221.h"

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

void i2cWrite1byte(uint8_t adr, uint8_t val) {
	mcp2221_i2cWrite(myMCP2221, adr, &val, 1, MCP2221_I2CRW_NORMAL);
	mcp2221_i2c_state_t state = MCP2221_I2C_IDLE;
	while(1) {
		if(mcp2221_i2cState(myMCP2221, &state) != MCP2221_SUCCESS) puts("ERROR");
		if(state == MCP2221_I2C_IDLE) break;
	}
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

uint8_t i2cRead1byte(uint8_t adr) {
	mcp2221_i2cRead(myMCP2221, adr, 1, MCP2221_I2CRW_NORMAL);
	mcp2221_i2c_state_t state = MCP2221_I2C_IDLE;
	while(1) {
		mcp2221_i2cState(myMCP2221, &state);
		if(state == MCP2221_I2C_DATAREADY) break;
	}
	uint8_t d;
	mcp2221_i2cGet(myMCP2221, &d, 1);
	return d;
}

void gpioInterrupt(void(*cb)(void)) {

	mcp2221_gpioconfset_t gpioConf = mcp2221_GPIOConfInit();
	gpioConf.conf[0].gpios = MCP2221_GPIO1;
	gpioConf.conf[0].mode  = MCP2221_GPIO_MODE_ALT3;
	mcp2221_setGPIOConf(myMCP2221, &gpioConf);
	mcp2221_setInterrupt(myMCP2221, MCP2221_INT_TRIG_FALLING, 1);

	mcp2221_error res;
	while(1) {
		int interrupt;
		res = mcp2221_readInterrupt(myMCP2221, &interrupt);
		if(res != MCP2221_SUCCESS) break;
		if(interrupt) {
			res = mcp2221_clearInterrupt(myMCP2221);
			if(res != MCP2221_SUCCESS) break;
			(*cb)();
		}
	}
	switch(res) {
		case MCP2221_SUCCESS:
			puts("No error");
			break;
		case MCP2221_ERROR:
			puts("General error");
			break;
		case MCP2221_INVALID_ARG:
			puts("Invalid argument, probably null pointer");
			break;
		case MCP2221_ERROR_HID:
			puts("USB HID Error");
			break;
		default:
			printf("Unknown error %d\n", res);
			break;
	}
}

