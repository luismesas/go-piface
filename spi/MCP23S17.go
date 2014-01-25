package spi

//	Microchip's MCP23S17: A 16-Bit I/O Expander with Serial Interface.
type MCP23S17 struct{
	Device *SPIDevice
	HardwareAddress byte
	
	// Controls the direction of the data I/O.
	IODIRa *MCP23S17Register
	IODIRb *MCP23S17Register

	// This register allows the user to configure the polarity on the corresponding GPIO port bits.
	IPOLa *MCP23S17Register
	IPOLb *MCP23S17Register

	// The GPINTEN register controls the interrupt-onchange feature for each pin.
	GPINTENa *MCP23S17Register
	GPINTENb *MCP23S17Register

	// The default comparison value is configured in the DEFVAL register.
	DEFVALa *MCP23S17Register
	DEFVALb *MCP23S17Register

	// The INTCON register controls how the associated pin value is compared for the interrupt-on-change feature.
	INTCONa *MCP23S17Register
	INTCONb *MCP23S17Register

	// The IOCON register contains several bits for configuring the device.
	IOCON *MCP23S17Register

	// The GPPU register controls the pull-up esistors for the port pins.
	GPPUa *MCP23S17Register
	GPPUb *MCP23S17Register

	// The INTF register reflects the interrupt condition on the port pins of any pin that is enabled for interrupts via the GPINTEN register.
	INTFa *MCP23S17Register
	INTFb *MCP23S17Register

	// The INTCAP register captures the GPIO port value at the time the interrupt occurred.
	INTCAPa *MCP23S17Register
	INTCAPb *MCP23S17Register

	// The GPIO register reflects the value on the port.
	GPIOa *MCP23S17Register
	GPIOb *MCP23S17Register

	// The OLAT register provides access to the output latches.
	OLATa *MCP23S17Register
	OLATb *MCP23S17Register
}

func NewMCP23S17(hardwareAddress byte, bus int, chip_select int) *MCP23S17{
	mcp := new(MCP23S17)
	mcp.Device = NewSPIDevice(bus, chip_select)
	mcp.HardwareAddress = hardwareAddress

   	mcp.IODIRa = NewMCP23S17Register(IODIRA, mcp)
   	mcp.IODIRb = NewMCP23S17Register(IODIRB, mcp)
   	mcp.IPOLa = NewMCP23S17Register(IPOLA, mcp)
   	mcp.IPOLb = NewMCP23S17Register(IPOLB, mcp)
   	mcp.GPINTENa = NewMCP23S17Register(GPINTENA, mcp)
   	mcp.GPINTENb = NewMCP23S17Register(GPINTENB, mcp)
   	mcp.DEFVALa = NewMCP23S17Register(DEFVALA, mcp)
   	mcp.DEFVALb = NewMCP23S17Register(DEFVALB, mcp)
   	mcp.INTCONa = NewMCP23S17Register(INTCONA, mcp)
   	mcp.INTCONb = NewMCP23S17Register(INTCONB, mcp)
   	mcp.IOCON = NewMCP23S17Register(IOCON, mcp)
   	mcp.GPPUa = NewMCP23S17Register(GPPUA, mcp)
   	mcp.GPPUb = NewMCP23S17Register(GPPUB, mcp)
   	mcp.INTFa = NewMCP23S17Register(INTFA, mcp)
   	mcp.INTFb = NewMCP23S17Register(INTFB, mcp)
   	mcp.INTCAPa = NewMCP23S17Register(INTCAPA, mcp)
   	mcp.INTCAPb = NewMCP23S17Register(INTCAPB, mcp)
   	mcp.GPIOa = NewMCP23S17Register(GPIOA, mcp)
   	mcp.GPIOb = NewMCP23S17Register(GPIOB, mcp)
   	mcp.OLATa = NewMCP23S17Register(OLATA, mcp)
   	mcp.OLATb = NewMCP23S17Register(OLATB, mcp)

	return mcp
}

func (mcp *MCP23S17) Open(spi_device string) error{
	return mcp.Device.Open(spi_device)
}

func (mcp *MCP23S17) Close() error{
	return mcp.Device.Close()
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
	board_addr_pattern := (mcp.HardwareAddress << 0x01) & 0xE
	rw_cmd_pattern := read_write_cmd & 0x01  // make sure it's just 1 bit long
	return 0x40 | board_addr_pattern | rw_cmd_pattern
}

// Returns the value of the address specified.
func (mcp *MCP23S17) Read(address byte) byte{
	ctrl_byte := mcp.getSPIControlByte(READ_CMD)
	return mcp.Device.Send([]byte{ctrl_byte, address, 0})[2]
}

// Writes data to the address specified.
func (mcp *MCP23S17) Write(data byte, address byte){
	ctrl_byte := mcp.getSPIControlByte(WRITE_CMD)
	mcp.Device.Send([]byte{ctrl_byte, address, data})
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
