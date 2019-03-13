#include "../easymcp2221.h"

uint8_t prev_val = 0x00;

void check_state(void) {
	uint8_t data;
	mcp23017Read(0x13,&data); // GPIOB
	if(prev_val != data) {
		printf("DATA: 0x%02x\n",data);
	}
	prev_val = data;
}

int main(void) {
	if(init_app()){ return 1; }

	// GPIO B setup
	mcp23017Write(0x01,0xFF); // IODIRB
	mcp23017Write(0x0D,0xFF); // GPPUB

	// Interrupt setup
	mcp23017Write(0x05,0xFF); // GPINTENB
	mcp23017Write(0x07,0xFF); // DEFVALB
	mcp23017Write(0x09,0xFF); // INTCONB

	gpioInterrupt(&check_state);

	return 0;
}
