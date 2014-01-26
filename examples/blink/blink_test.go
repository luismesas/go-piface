package main

import (
	"testing"
	"unsafe"
)

func TestSome(t *testing.T){
	var b byte
	t.Errorf("Size of byte is %d", unsafe.Sizeof(b))
	var u8 uint8
	t.Errorf("Size of uint8 is %d", unsafe.Sizeof(u8))
	ba3 := []byte{0,0,0}
	t.Errorf("Size of []byte{0,0,0} is %d", unsafe.Sizeof(ba3))
	ba1 := []byte{0}
	t.Errorf("Size of []byte{0} is %d", unsafe.Sizeof(ba1))
	bam := make([]byte,3)
	t.Errorf("Size of []byte{0,0,0} with make is %d", unsafe.Sizeof(bam))
}