#include "../easymcp2221.h"

void check_state(void) {
	uint8_t data;
	mcp23017Read(0x13,&data); // GPIOB
	printf("DATA: 0x%02x\n",data);
}

int main(void) {
	if(init_app()){ return 1; }

	// GPIO B setup
	mcp23017Write(0x01,0xFF); // IODIRB
	mcp23017Write(0x0D,0xFF); // GPPUB

	while(1) {
		check_state();
		sleep_ms(100);
	}
	return 0;
}
