package spi

import (
	"syscall"
	"unsafe"
)

const(
	_IOC_NRBITS = 8
	_IOC_TYPEBITS = 8

	_IOC_SIZEBITS = 14
	_IOC_DIRBITS = 2

	_IOC_NRMASK =   (1 << _IOC_NRBITS) - 1
	_IOC_TYPEMASK = (1 << _IOC_TYPEBITS) - 1
	_IOC_SIZEMASK = (1 << _IOC_SIZEBITS) - 1
	_IOC_DIRMASK =  (1 << _IOC_DIRBITS) - 1

	_IOC_NRSHIFT =0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT = _IOC_SIZESHIFT + _IOC_SIZEBITS

	// Direction bits
	_IOC_NONE = 0
	_IOC_WRITE = 1
	_IOC_READ = 2
)

//...and for the drivers/sound files...
const (
	IOC_IN = _IOC_WRITE << _IOC_DIRSHIFT
	IOC_OUT = _IOC_READ << _IOC_DIRSHIFT
	IOC_INOUT = (_IOC_WRITE|_IOC_READ) << _IOC_DIRSHIFT
	IOCSIZE_MASK = _IOC_SIZEMASK << _IOC_SIZESHIFT
	IOCSIZE_SHIFT = _IOC_SIZESHIFT
)


func IOC(dir, t, nr, size uintptr) uintptr{
	return (dir  << _IOC_DIRSHIFT) | (t << _IOC_TYPESHIFT) | (nr   << _IOC_NRSHIFT) | (size << _IOC_SIZESHIFT)
}

// used to create ioctl numbers

func IO(t,nr uintptr) uintptr{
	return IOC(_IOC_NONE, t, nr, 0)
}

func IOR(t, nr, size uintptr) uintptr{
	return IOC(_IOC_READ, t, nr, size)
}

func IOW(t, nr, size uintptr) uintptr{
	return IOC(_IOC_WRITE, t, nr, size)
}

func IOWR(t,nr,size uintptr) uintptr{
	return IOC(_IOC_READ|_IOC_WRITE, t, nr, size)
}

func IOR_BAD(t,nr,size uintptr) uintptr{
	return IOC(_IOC_READ, t, nr, size)
}

func IOW_BAD(t,nr,size uintptr) uintptr{
	return IOC(_IOC_WRITE,t,nr, size)
}

func IOWR_BAD(t,nr,size uintptr) uintptr{
	return IOC(_IOC_READ|_IOC_WRITE, t, nr, size)
}


func IOCTL(fd, op, arg uintptr) error{
	_,_,ep := syscall.Syscall(syscall.SYS_IOCTL, fd, op, arg)
	if ep != 0 {
		return syscall.Errno(ep)
	}
	return nil
}

const SPI_IOC_MAGIC = 107

/* Read / Write of SPI mode (SPI_MODE_0..SPI_MODE_3) */
func SPI_IOC_RD_MODE() uintptr {
	return IOR(SPI_IOC_MAGIC, 1, 1)
}

func SPI_IOC_WR_MODE() uintptr {
	return IOW(SPI_IOC_MAGIC, 1, 1)
}

/* Read / Write SPI bit justification */
func SPI_IOC_RD_LSB_FIRST() uintptr {
	return IOR(SPI_IOC_MAGIC, 2, 1)
}

func SPI_IOC_WR_LSB_FIRST() uintptr {
	return IOW(SPI_IOC_MAGIC, 2, 1)
}

/* Read / Write SPI device word length (1..N) */
func SPI_IOC_RD_BITS_PER_WORD() uintptr {
	return IOR(SPI_IOC_MAGIC, 3, 1)
}

func SPI_IOC_WR_BITS_PER_WORD() uintptr {
	return IOW(SPI_IOC_MAGIC, 3, 1)
}

/* Read / Write SPI device default max speed hz */
func SPI_IOC_RD_MAX_SPEED_HZ() uintptr {
	return IOR(SPI_IOC_MAGIC, 4, 4)
}

func SPI_IOC_WR_MAX_SPEED_HZ() uintptr {
	return IOW(SPI_IOC_MAGIC, 4, 4)
}


func SPI_IOC_MESSAGE(n uintptr) uintptr{
	bytes := make([]byte,SPI_MESSAGE_SIZE(n))
	return IOW(SPI_IOC_MAGIC, 0 , uintptr(unsafe.Pointer(&bytes)))
}

func SPI_MESSAGE_SIZE(n uintptr) uintptr{
	if (n * unsafe.Sizeof(SPI_IOC_TRANSFER{})) < (1 << _IOC_SIZEBITS) {
		return (n * unsafe.Sizeof(SPI_IOC_TRANSFER{}))
	} else {
		return 0
	}
}

type SPI_IOC_TRANSFER struct{
	txBuf uint64
	rxBuf uint64
	length uint32
	speedHz uint32
	delayUsecs uint16
	bitsPerWord uint8
	csChange uint8
	pad uint32
}





