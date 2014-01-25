package interrupts

// A device that interrupts using the GPIO pins.
type GPIOInterruptDevice struct{
}

func NewGPIOInterruptDevice() *GPIOInterruptDevice{
	return new(GPIOInterruptDevice)
}