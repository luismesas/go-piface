package main

import (
	"testing"
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
			t.Errorf("GetBitMask(%d) must be %d", x, exp2)
		}
	}
}