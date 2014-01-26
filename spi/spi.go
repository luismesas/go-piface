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
	file *os.File // nil

	mode uint8
	bpw uint8
	speed uint32
}

// An SPI Device at /dev/spi<bus>.<chip_select>.
func NewSPIDevice(bus int, chipSelect int) *SPIDevice{
	spi := new(SPIDevice)
	spi.Bus = bus
	spi.Chip = chipSelect

	return spi
}

// Opens SPI device
func (spi *SPIDevice) Open() error{
	spiDevice := fmt.Sprintf("%s%d.%d", SPIDEV, spi.Bus, spi.Chip)

	var err error
	spi.file, err = os.OpenFile(spiDevice, os.O_RDWR, 0)
	// spi.file, err = os.Create(spiDevice)
	if err != nil {
		return fmt.Errorf("I can't see %s. Have you enabled the SPI module? (%s)", spiDevice, SPI_HELP_LINK)
	}

	log.Println("SPI Open")
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
	rBuffer := make([]byte, unsafe.Sizeof(wBuffer))

	log.Printf("Size of sent buffer is: %d", unsafe.Sizeof(wBuffer))
	log.Printf("Len of sent buffer is: %d", len(wBuffer))
	log.Printf("Size of receive buffer is: %d", unsafe.Sizeof(rBuffer))
	log.Printf("Len of receive buffer is: %d", len(rBuffer))

	transfer := SPI_IOC_TRANSFER{}
	transfer.txBuf = uint64( uintptr( unsafe.Pointer(&wBuffer)))
	transfer.rxBuf = uint64( uintptr( unsafe.Pointer(&rBuffer)))
	transfer.length = uint32(unsafe.Sizeof(wBuffer))
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
