package main

import (
	"log"
	"os"
	"time"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

var (
	relayBlockA1 = rpio.Pin(17)
	relayBlockA2 = rpio.Pin(27)
	relayBlockB1 = rpio.Pin(22)
	relayBlockB2 = rpio.Pin(10)
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	relayBlockA1.Output()
	relayBlockA2.Output()
	relayBlockB1.Output()
	relayBlockB2.Output()

	log.Println("Wait 5 sec to turn on relays")
	time.Sleep(time.Second * 5)
	relayBlockA1.High()
	relayBlockA2.High()
	relayBlockB1.High()
	relayBlockB2.High()
	log.Println("relays turned on!")

	log.Println("Wait 5 sec to turn off relays")
	time.Sleep(time.Second * 10)
	relayBlockA1.Low()
	relayBlockA2.Low()
	relayBlockB1.Low()
	relayBlockB2.Low()
	log.Println("relays turned on!")
}
