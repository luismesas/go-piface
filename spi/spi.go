package spi

import (
	"fmt"
	"os"
	"unsafe"
	// "syscall"
	"time"
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

	//watchs for spi messages
	go func(){
		for{
			r := make([]byte,1)
			spi.fd.Read(r)
			if r[0] != 0 {
				fmt.Printf("%q",r[0])
			}
			time.Sleep(time.Millisecond)			
		}
	}()
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
	n, err := spi.fd.Write(bytes_to_send)
	if err != nil {
		fmt.Printf("Error writting: %s", err)
	} else {
		fmt.Printf("Sent %d bytes: %q", n, bytes_to_send)
	}

	return make([]byte, len(bytes_to_send))
	/*
	wBuffer := bytes_to_send
	rBuffer := make([]byte, len(bytes_to_send))

	transfer := SpiIOcTransfer{}
	transfer.txBuf = uint64( uintptr( unsafe.Pointer(&wBuffer)))
	transfer.rxBuf = uint64( uintptr( unsafe.Pointer(&rBuffer)))
	transfer.length = uint32(len(bytes_to_send))

	fmt.Printf("sent %d bytes: %q\n", len(bytes_to_send), wBuffer)
	_,_,ep := syscall.Syscall(syscall.SYS_IOCTL, spi.fd.Fd(), SpiIOcMessage(1), uintptr(unsafe.Pointer(&transfer)))
	if ep != 0 {
		fmt.Printf("Error on syscall: %s\n", syscall.Errno(ep))
	}
	fmt.Printf("read %d bytes: %q\n", len(bytes_to_send), rBuffer)
	return rBuffer
	*/
}

type SpiIOcTransfer struct{
	txBuf uint64
	rxBuf uint64
	length uint32
	speedHz uint32
	delayUsecs uint16
	bitsPerWord uint8
	csChange uint8
	pad uint32
}

const SPI_IOC_MAGIC = 107

const (
	IOC_NONE = 0
	IOC_WRITE = 1
	IOC_READ = 2

	IOC_NRBITS = 8
	IOC_TYPEBITS = 8

	IOC_SIZEBITS = 14
	IOC_DIRBITS = 2

	IOC_NRSHIFT = 0
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS
	IOC_DIRSHIFT = IOC_SIZESHIFT + IOC_SIZEBITS
)

func SpiIOcMessage(n uintptr) uintptr{
	var b byte
	return IOW(SPI_IOC_MAGIC, 0, unsafe.Sizeof(b) * SpiMessageSize(n))
}

func SpiMessageSize(n uintptr) uintptr{
	if (n * unsafe.Sizeof(SpiIOcTransfer{})) < (1 << IOC_SIZEBITS) {
		return (n * unsafe.Sizeof(SpiIOcTransfer{}))
	} else {
		return 0
	}
}

func IOW(t, nr, size uintptr) uintptr{
	return IOC(IOC_WRITE, t, nr, unsafe.Sizeof(size))
}

func IOC(dir, t, nr, size uintptr) uintptr{
	return (dir << IOC_DIRSHIFT) | (t << IOC_TYPESHIFT) | (nr << IOC_NRSHIFT) | (size << IOC_SIZESHIFT)
}
