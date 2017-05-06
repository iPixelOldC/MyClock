package main

import (
	"fmt"
	"time"

	saks "github.com/hanbang-wang/SAKS-SDK-GO"
	"github.com/imroc/req"
	//	rpio  "github.com/stianeikeland/go-rpio"
)

func main() {
	showDisplay()
}

func showDisplay() {
	display := saks.DigitalDisplay
	for true {
		checkBasicVar := req.Get("http://localhost:3000/basic/get").String()
		if checkBasicVar == "1000" {
			hour, mins, sec := time.Now().Clock()
			_ = sec
			timeNow := fmt.Sprintf("%02d.%02d", hour, mins)
			display.Show(timeNow)
			time.Sleep(20 * time.Second)
		} else {
			display.Off()
			time.Sleep(5 * time.Second)
		}
	}

}
