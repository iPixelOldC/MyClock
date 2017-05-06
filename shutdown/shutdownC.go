package main

import (
	"fmt"
	sak "github.com/hanbang-wang/SAKS-SDK-GO"
	"github.com/stianeikeland/go-rpio"
	"os"
	"os/exec"
	"time"
)

var (
	// Use mcu pin 10,  corresponds to physical pin 19 on the pi
	pin = rpio.Pin(22)
)

func main() {
	// Open and map memory to access gpio,  check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Input()
	led := sak.LEDRow
	led.Off()
	// Toggle pin 20 times
	for true {
		a := pin.Read()
		if a == 1 {
			display := sak.DigitalDisplay
			display.Off()

			LED1 := [8]bool{false, false, false, false, false, false, false, true}
			led.SetRow(LED1)
			cmd := exec.Command("sudo", "init", "0")
			cmd.Start()
		}
		time.Sleep(time.Second / 5)
	}
}
