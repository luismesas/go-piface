package spi

// Defaults
const (
	DEFAULT_HARDWARE_ADDR = 0
	DEFAULT_BUS = 0
	DEFAULT_CHIP = 0
)

// Register addresses
const (
	IODIRA = 0x0  // I/O direction A
	IODIRB = 0x1  // I/O direction B
	IPOLA = 0x2  // I/O polarity A
	IPOLB = 0x3  // I/O polarity B
	GPINTENA = 0x4  // interupt enable A
	GPINTENB = 0x5  // interupt enable B
	DEFVALA = 0x6  // register default value A (interupts)
	DEFVALB = 0x7  // register default value B (interupts)
	INTCONA = 0x8  // interupt control A
	INTCONB = 0x9  // interupt control B
	IOCON = 0xA  // I/O config (also 0xB)
	GPPUA = 0xC  // port A pullups
	GPPUB = 0xD  // port B pullups
	INTFA = 0xE  // interupt flag A (where the interupt came from)
	INTFB = 0xF  // interupt flag B
	INTCAPA = 0x10  // interupt capture A (value at interupt is saved here)
	INTCAPB = 0x11  // interupt capture B
	GPIOA = 0x12  // port A
	GPIOB = 0x13  // port B
	OLATA = 0x14  // output latch A
	OLATB = 0x15  // output latch B
) 

const (
	WRITE_CMD = 0
	READ_CMD = 1
)

const (
	LOWER_NIBBLE = 0
	UPPER_NIBBLE = 1
)