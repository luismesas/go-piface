package spi

import (
	"fmt"
	"os"
	"log"
)

const SPIDEV = "/dev/spidev"
const SPI_HELP_LINK = "http://piface.github.io/pifacecommon/installation.html#enable-the-spi-module"

type SPIDevice struct{
	Bus int // 0
	Chip int // 0
	fd  *os.File // nil 
	spi_device string
}

// An SPI Device at /dev/spi<bus>.<chip_select>.
func NewSPIDevice(bus int, chipSelect int) *SPIDevice{
	spi := new(SPIDevice)
	spi.Bus = bus
	spi.Chip = chipSelect
	spi.fd = nil

	spiDevice := fmt.Sprintf("%s%d.%d", SPIDEV, spi.Bus, spi.Chip)
	spi.Open(spiDevice)

	return spi
}

// Opens SPI device
func (spi *SPIDevice) Open(spi_device string) error{
	var err error
	spi.fd, err = os.OpenFile(spi_device, os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
		return fmt.Errorf("I can't see %s. Have you enabled the SPI module? (%s)", spi_device, SPI_HELP_LINK)
	}
	return nil
}

// Closes SPI device
func (spi *SPIDevice) Close() error{
	err := spi.fd.Close()
	if err != nil {
		return fmt.Errorf("Error closing spi", err)
	}
	return nil
}

// Sends bytes over SPI channel and returns []byte response
func (spi *SPIDevice) Send(bytes_to_send []byte) []byte{
	//sends command
	count, err := spi.fd.Write(bytes_to_send)
	if err != nil {
		log.Fatalf("Error sending bytes", err)
	}
	fmt.Printf("sent %d bytes: %q\n", count, bytes_to_send)

	err = spi.fd.Sync()
	if err != nil {
		log.Fatalf("Error syncing bytes", err)
	}

	data := make([]byte, 100)
	count, err = spi.fd.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	return data
}
