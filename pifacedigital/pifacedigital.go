package pifacedigital

import (
	"fmt"
	"github.com/luismesas/go-piface/spi"
	// "github.com/luismesas/go-piface/interrupts"
)

// A PiFace Digital board.
type PiFaceDigital struct{
	mcp *spi.MCP23S17
	// device *interrupts.GPIOInterruptDevice

	InputPins []*spi.MCP23S17RegisterBitNeg
	InputPort *spi.MCP23S17RegisterNeg
	OutputPins []*spi.MCP23S17RegisterBit
	OutputPort *spi.MCP23S17Register
	Leds []*spi.MCP23S17RegisterBit
	Relays []*spi.MCP23S17RegisterBit
	Switches []*spi.MCP23S17RegisterBit
}

func NewPiFaceDigital(hardware_addr byte, bus int, chip_select int, init bool) *PiFaceDigital{
	pfd := new(PiFaceDigital)
	pfd.mcp = spi.NewMCP23S17(hardware_addr, bus, chip_select)
	// pfd.device = interrupts.NewGPIOInterruptDevice()

	pfd.InputPins = make([]*spi.MCP23S17RegisterBitNeg,8)
	for i := range(pfd.InputPins){
		pfd.InputPins[i] = spi.NewMCP23S17RegisterBitNeg(uint(i), spi.GPIOB, pfd.mcp)
	}

	pfd.InputPort = spi.NewMCP23S17RegisterNeg(spi.GPIOB, pfd.mcp)

	pfd.OutputPins = make([]*spi.MCP23S17RegisterBit,8)
	for i := range(pfd.OutputPins){
		pfd.OutputPins[i] = spi.NewMCP23S17RegisterBit(uint(i), spi.GPIOA, pfd.mcp)
	}

	pfd.OutputPort = spi.NewMCP23S17Register(spi.GPIOA, pfd.mcp)

	pfd.Leds = make([]*spi.MCP23S17RegisterBit,8)
	for i := range(pfd.Leds){
		pfd.Leds[i] = spi.NewMCP23S17RegisterBit(uint(i), spi.GPIOA, pfd.mcp)
	}

	pfd.Relays = make([]*spi.MCP23S17RegisterBit,2)
	for i := range(pfd.Relays){
		pfd.Relays[i] = spi.NewMCP23S17RegisterBit(uint(i), spi.GPIOA, pfd.mcp)
	}

	pfd.Switches = make([]*spi.MCP23S17RegisterBit,4)
	for i := range(pfd.Switches){
		pfd.Switches[i] = spi.NewMCP23S17RegisterBit(uint(i), spi.GPIOA, pfd.mcp)
	}

	if init {
		pfd.InitBoard()
	}

	return pfd
}

func (pfd *PiFaceDigital) InitBoard() error{

	var ioconfig byte
	ioconfig = (
		spi.BANK_OFF |
		spi.INT_MIRROR_OFF |
		spi.SEQOP_OFF |
		spi.DISSLW_OFF |
		spi.HAEN_ON |
		spi.ODR_OFF |
		spi.INTPOL_LOW)

	pfd.mcp.IOCON.SetValue(ioconfig)
	// if pfd.mcp.IOCON.Value() != ioconfig {
	// 	return fmt.Errorf("No PiFace Digital board detected (hardware_addr=%d, bus=%b, chip_select=%b).", pfd.mcp.HardwareAddress, pfd.mcp.Device.Bus, pfd.mcp.Device.Chip)
	// }

	pfd.mcp.GPIOa.SetValue(0)
	pfd.mcp.IODIRa.SetValue(0) // GPIOA as outputs
	pfd.mcp.IODIRb.SetValue(0xFF) // GPIOB as inputs
	pfd.mcp.GPPUb.SetValue(0xFF) // input pullups on
	// pfd.EnableInterrupts()

	return nil
}

func (pfd *PiFaceDigital) EnableInterrupts() error{
	return fmt.Errorf("EnableInterrupts() Not implemented")
}

func (pfd *PiFaceDigital) Open() error{
	return fmt.Errorf("Open() Not implemented")
}

func (pfd *PiFaceDigital) Close() error{
	return fmt.Errorf("Close() Not implemented")
}