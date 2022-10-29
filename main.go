package main

import (
	"os"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(19)
	pin.Mode(rpio.Pwm)
	pin.Freq(25000)
	pin.DutyCycle(0, 25)
}
