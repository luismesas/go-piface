package spi

import (
	"log"
	"fmt"
	"os"
	"unsafe"
)

const SPIDEV = "/dev/spidev"
const SPI_HELP_LINK = "http://piface.github.io/pifacecommon/installation.html#enable-the-spi-module"
const SPI_DELAY = 0

type SPIDevice struct{
	Bus int // 0
	Chip int // 0
	file  *os.File // nil 
	spi_device string

	mode uint8
	bpw uint8
	speed uint32
}

// An SPI Device at /dev/spi<bus>.<chip_select>.
func NewSPIDevice(bus int, chipSelect int) *SPIDevice{
	spi := new(SPIDevice)
	spi.Bus = bus
	spi.Chip = chipSelect

	spiDevice := fmt.Sprintf("%s%d.%d", SPIDEV, spi.Bus, spi.Chip)
	spi.Open(spiDevice)

	return spi
}

// Opens SPI device
func (spi *SPIDevice) Open(spi_device string) error{
	log.Println("SPI Open")

	var err error
	// spi.fd, err = os.OpenFile(spi_device, os.O_RDWR|os.O_SYNC, 0)
	spi.file, err = os.Create(spi_device)
	if err != nil {
		return fmt.Errorf("I can't see %s. Have you enabled the SPI module? (%s)", spi_device, SPI_HELP_LINK)
	}
	return nil
}

// Closes SPI device
func (spi *SPIDevice) Close() error{
	err := spi.file.Close()
	if err != nil {
		return fmt.Errorf("Error closing spi", err)
	}
	return nil
}

// Sends bytes over SPI channel and returns []byte response
func (spi *SPIDevice) Send(bytes_to_send []byte) ([]byte, error){
	wBuffer := bytes_to_send
	var rBuffer []byte
	rBuffer = make([]byte,len(bytes_to_send))

	transfer := SPI_IOC_TRANSFER{}
	transfer.txBuf = uint64( uintptr( unsafe.Pointer(&wBuffer)))
	transfer.rxBuf = uint64( uintptr( unsafe.Pointer(&rBuffer)))
	transfer.length = uint32(len(bytes_to_send))
	transfer.delayUsecs = SPI_DELAY
	transfer.bitsPerWord = spi.bpw
	transfer.speedHz = spi.speed

	log.Printf("sent %d bytes: %q\n", len(bytes_to_send), wBuffer)
	err := IOCTL(spi.file.Fd(), SPI_IOC_MESSAGE(1), uintptr(unsafe.Pointer(&transfer)))
	if err != nil {
		return nil, fmt.Errorf("Error on sending: %s\n", err)
	}
	return rBuffer, nil
}

func (spi *SPIDevice) SetMode(mode uint8) error{
	spi.mode = mode
	err := IOCTL(spi.file.Fd(), SPI_IOC_WR_MODE(), uintptr(unsafe.Pointer(&mode)))
	if err != nil {
		return fmt.Errorf("Error setting mode: %s\n", err)
	}
	return nil
}

func (spi *SPIDevice) SetBitsPerWord(bpw uint8) error{
	spi.bpw = bpw
	err := IOCTL(spi.file.Fd(), SPI_IOC_WR_BITS_PER_WORD(), uintptr(unsafe.Pointer(&bpw)))
	if err != nil {
		return fmt.Errorf("Error setting bits per word: %s\n", err)
	}
	return nil
}

func (spi *SPIDevice) SetSpeed(speed uint32) error{
	spi.speed = speed
	err := IOCTL(spi.file.Fd(), SPI_IOC_WR_MAX_SPEED_HZ(), uintptr(unsafe.Pointer(&speed)))
	if err != nil {
		return fmt.Errorf("Error setting speed: %s\n", err)
	}
	return nil
}
