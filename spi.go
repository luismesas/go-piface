package main

import (
	"fmt"
	"os"
	"log"
)

const SPIDEV = "/dev/spidev"
const SPI_HELP_LINK = "http://piface.github.io/pifacecommon/installation.html#enable-the-spi-module"

const DEFAULT_BUS = 0
const DEFAULT_CHIP = 0

type SPIDevice struct{
	bus int // 0
	chip_select int // 0
	spi_callback *interface{} // nil
	fd  *os.File // nil 
	spi_device string
}

// An SPI Device at /dev/spi<bus>.<chip_select>.
func NewSPIDevice(bus int, chip_select int, spi_callback *interface{}) *SPIDevice{
	spi := new(SPIDevice)
	spi.bus = bus
	spi.chip_select = chip_select
	spi.spi_callback = spi_callback
	spi.fd = nil

	spi_device := fmt.Sprintf("%s%d.%d", SPIDEV, spi.bus, spi.chip_select)
	spi.open_fd(spi_device)

	return spi
}

func (spi *SPIDevice) open_fd(spi_device string){
	var err error
	spi.fd, err = os.OpenFile(spi_device, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalf("I can't see %s. Have you enabled the SPI module? (%s)", spi_device, SPI_HELP_LINK)
	}
}

func (spi *SPIDevice) close_fd(){
	err := spi.fd.Close()
	if err != nil {
		log.Fatalf("Error closing spi", err)
	}
}

func main(){
	device := NewSPIDevice(DEFAULT_BUS, DEFAULT_CHIP, nil)
	device.close_fd()
}