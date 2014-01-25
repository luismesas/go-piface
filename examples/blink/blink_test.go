package main

import (
	"testing"
	"unsafe"
	"github.com/luismesas/go-piface/spi"
)

func TestSome(t *testing.T){
	var b byte
	t.Errorf("Size of byte is %d", unsafe.Sizeof(b))

	t.Errorf("Size of SpiIOcTransfer is %d", unsafe.Sizeof(SpiIOcTransfer{}))

	t.Errorf("Value of SpiIOcMessage(1) is %d", SpiIOcMessage(1))
}