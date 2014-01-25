package main

import (
	"log"
	"github.com/luismesas/go-piface/spi"
)

func main(){
	log.Println("MCP23S17")
	device := spi.NewMCP23S17(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)
	log.Println("MCP23S17 open")
	device.Close()
	log.Println("MCP23S17 close")
}