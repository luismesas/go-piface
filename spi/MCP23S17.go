package spi

//	Microchip's MCP23S17: A 16-Bit I/O Expander with Serial Interface.
type MCP23S17 struct{
	device *SPIDevice
	hardware_addr byte
	
	// Controls the direction of the data I/O.
	iodira *MCP23S17Register
	iodirb *MCP23S17Register

	// This register allows the user to configure the polarity on the corresponding GPIO port bits.
	ipola *MCP23S17Register
	ipolb *MCP23S17Register

	// The GPINTEN register controls the interrupt-onchange feature for each pin.
	gpintena *MCP23S17Register
	gpintenb *MCP23S17Register

	// The default comparison value is configured in the DEFVAL register.
	defvala *MCP23S17Register
	defvalb *MCP23S17Register

	// The INTCON register controls how the associated pin value is compared for the interrupt-on-change feature.
	intcona *MCP23S17Register
	intconb *MCP23S17Register

	// The IOCON register contains several bits for configuring the device.
	iocon *MCP23S17Register

	// The GPPU register controls the pull-up esistors for the port pins.
	gppua *MCP23S17Register
	gppub *MCP23S17Register

	// The INTF register reflects the interrupt condition on the port pins of any pin that is enabled for interrupts via the GPINTEN register.
	intfa *MCP23S17Register
	intfb *MCP23S17Register

	// The INTCAP register captures the GPIO port value at the time the interrupt occurred.
	intcapa *MCP23S17Register
	intcapb *MCP23S17Register

	// The GPIO register reflects the value on the port.
	gpioa *MCP23S17Register
	gpiob *MCP23S17Register

	// The OLAT register provides access to the output latches.
	olata *MCP23S17Register
	olatb *MCP23S17Register
}

func NewMCP23S17(hardware_addr byte, bus int, chip_select int) *MCP23S17{
	mcp := new(MCP23S17)
	mcp.device = NewSPIDevice(bus, chip_select)
	mcp.hardware_addr = hardware_addr

   	mcp.iodira = NewMCP23S17Register(IODIRA, mcp)
   	mcp.iodirb = NewMCP23S17Register(IODIRB, mcp)
   	mcp.ipola = NewMCP23S17Register(IPOLA, mcp)
   	mcp.ipolb = NewMCP23S17Register(IPOLB, mcp)
   	mcp.gpintena = NewMCP23S17Register(GPINTENA, mcp)
   	mcp.gpintenb = NewMCP23S17Register(GPINTENB, mcp)
   	mcp.defvala = NewMCP23S17Register(DEFVALA, mcp)
   	mcp.defvalb = NewMCP23S17Register(DEFVALB, mcp)
   	mcp.intcona = NewMCP23S17Register(INTCONA, mcp)
   	mcp.intconb = NewMCP23S17Register(INTCONB, mcp)
   	mcp.iocon = NewMCP23S17Register(IOCON, mcp)
   	mcp.gppua = NewMCP23S17Register(GPPUA, mcp)
   	mcp.gppub = NewMCP23S17Register(GPPUB, mcp)
   	mcp.intfa = NewMCP23S17Register(INTFA, mcp)
   	mcp.intfb = NewMCP23S17Register(INTFB, mcp)
   	mcp.intcapa = NewMCP23S17Register(INTCAPA, mcp)
   	mcp.intcapb = NewMCP23S17Register(INTCAPB, mcp)
   	mcp.gpioa = NewMCP23S17Register(GPIOA, mcp)
   	mcp.gpiob = NewMCP23S17Register(GPIOB, mcp)
   	mcp.olata = NewMCP23S17Register(OLATA, mcp)
   	mcp.olatb = NewMCP23S17Register(OLATB, mcp)

	return mcp
}

func (mcp *MCP23S17) Open(spi_device string) {
	mcp.device.Open(spi_device)
}

func (mcp *MCP23S17) Close(){
	mcp.device.Close()
}


// Returns an SPI control byte.
// The MCP23S17 is a slave SPI device. The slave address contains
// four fixed bits and three user-defined hardware address bits
// (if enabled via IOCON.HAEN) (pins A2, A1 and A0) with the
// read/write bit filling out the control byte::

// 	+--------------------+
// 	|0|1|0|0|A2|A1|A0|R/W|
// 	+--------------------+
// 	 7 6 5 4 3  2  1   0

// :param read_write_cmd: Read or write command.
// :type read_write_cmd: int	
func (mcp *MCP23S17) getSPIControlByte(read_write_cmd byte) byte {
    // board_addr_pattern = (self.hardware_addr & 0b111) << 1
	board_addr_pattern := (mcp.hardware_addr << 0x01) & 0xE
	rw_cmd_pattern := read_write_cmd & 0x01  // make sure it's just 1 bit long
	return 0x40 | board_addr_pattern | rw_cmd_pattern
}

// Returns the value of the address specified.
func (mcp *MCP23S17) Read(address byte) byte{
	ctrl_byte := mcp.getSPIControlByte(READ_CMD)
	return mcp.device.Send([]byte{ctrl_byte, address, 0})[0]
}

// Writes data to the address specified.
func (mcp *MCP23S17) Write(data byte, address byte){
	ctrl_byte := mcp.getSPIControlByte(WRITE_CMD)
	mcp.device.Send([]byte{ctrl_byte, address, data})
}

// Returns the bit specified from the address.
func (mcp *MCP23S17) ReadBit(bit_num uint, address byte) byte{
	value := mcp.Read(address)
	bit_mask := GetBitMask(bit_num)
	if value & bit_mask > 0 {
		return 1
	} else {
		return 0
	}
}

// Writes the value given to the bit in the address specified.
func (mcp *MCP23S17) WriteBit(value byte, bit_num uint, address byte){
	bit_mask := GetBitMask(bit_num)
	old_byte := mcp.Read(address)
	newbyte := old_byte | bit_mask
	mcp.Write(newbyte, address)
}

// Clears the interrupt flags by reading the capture register.
func (mcp *MCP23S17) ClearInterrupts(port int){
	var address byte
	address = INTCAPA
	if port == GPIOA {
		address = INTCAPB
	}
	mcp.Read(address)
}
