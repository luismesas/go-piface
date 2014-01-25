package main

import (
	"testing"
	"github.com/luismesas/go-piface/pifacedigital"
	"github.com/luismesas/go-piface/spi"
)

func TestBitMask(t *testing.T){
	var x,y uint
	var exp2 byte
	for x=0;x<64;x++{
		exp2 = 0x01
		for y=0;y<x;y++{
			exp2 = exp2 * 0x02
		}

		if spi.GetBitMask(x) != exp2 {
			t.Errorf("SPI: GetBitMask(%d) must be %d", x, exp2)
		}
	}
}

func TestPiFaceDigital(t *testing.T){
	pfd := pifacedigital.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP, false)
	err := pfd.Open()
	if err != nil {
		t.Errorf("TestPiFaceDigital: %s", err)
	}
	err = pfd.Close()
	if err != nil {
		t.Errorf("TestPiFaceDigital: %s", err)
	}
}
