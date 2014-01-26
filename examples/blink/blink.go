package main

import (
	"fmt"
	"github.com/luismesas/go-piface"
	"github.com/luismesas/go-rpi/spi"
	"time"
)

func main(){
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)
	
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}

	fmt.Println("Blinking led 7 each second")

	for{
		pfd.Leds[7].Toggle()
		time.Sleep(time.Second)
	}
}