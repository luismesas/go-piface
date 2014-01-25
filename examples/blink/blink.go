package main

import (
	"github.com/luismesas/go-piface/pifacedigital"
	"github.com/luismesas/go-piface/spi"
	"time"
)

func main(){
	pfd := pifacedigital.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP, true)

	for{
		pfd.Leds[7].Toggle()
		time.Sleep(time.Second)
	}
}