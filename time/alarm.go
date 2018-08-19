package main

import (
	"fmt"
	"time"

	//	yaml "gopkg.in/yaml.v2"
	saks "github.com/hanbang-wang/SAKS-SDK-GO"
	"github.com/imroc/req"
	//	rpio "github.com/stianeikeland/go-rpio"
)

var (
	BUTTON1 = saks.TactRow
	BUTTON0 = saks.TactRow
	SWITCH0 = saks.DipSwitch
	//	config = `time: []`
	button1Time  int  = 0
	switch0Statu bool = false
	alarmHour         = 0
	alarmSec          = 0
	defLED            = saks.LEDRow
	defDisplay        = saks.DigitalDisplay
)

func main() {
	for true {
		time.Sleep(5 * time.Second)
		getSwitch0Statu()
	}
}

func getButton1Statu() (button1Time int) {
	for true {
		if BUTTON1.IsOn(1) == true {
			fmt.Print("+1: ")
			button1Time += 1
			fmt.Println(button1Time)
			defDisplay.Show(fmt.Sprintf("-%02d-", button1Time))
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("0", button1Time)
			time.Sleep(1 * time.Second)
			return
		}
	}
	return
}

func getSwitch0Statu() (switch0Statu bool) {
	if SWITCH0.IsOn(0) == true {
		req.Get("http://localhost:3000/basic/change/2000").Do()
		setAlarmTime()
		req.Get("http://localhost:3000/basic/change/1000").Do()
		time.Sleep(3 * time.Second)
	} else {
		time.Sleep(10 * time.Second)
	}
	return
}

func setAlarmTime() {
	defDisplay.Off()
	defLED.SetRow([8]bool{true, true, false, false, false, false, false, false})
	alarmHour = getButton1Statu()
	time.Sleep(5 * time.Second)
	defDisplay.Off()
	defLED.SetRow([8]bool{true, true, true, true, false, false, false, false})
	alarmSec = getButton1Statu()
	time.Sleep(5 * time.Second)
	defDisplay.Off()
	defLED.SetRow([8]bool{true, true, true, true, true, true, false, false})
	alarmTime := fmt.Sprintf("%02d.%02d", alarmHour, alarmSec)
	defDisplay.Show(alarmTime)
	time.Sleep(5 * time.Second)
	defDisplay.Off()
	defLED.Off()
	//	saveToFile(alarmTime)
}

/*
func saveToFile(alarmT string) {
	return
}
*/
