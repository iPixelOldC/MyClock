package main

import (
	"github.com/go-martini/martini"
)

// BasicVar:
// 1000 = Time Run
// 2000 = Clock Run
// ShutdownVar 0 = False, 1 =True(Raspberry Pi Shutdown)
var (
	ShutdownVar string = "0"
	BasicVar    string = "1000"
)

func main() {
	mart := martini.Classic()

	// Shutdown URL
	mart.Get("/shut/change/:bool", func(parmas martini.Params) string {
		ShutdownVar = parmas["bool"]
		return ShutdownVar
	})

	mart.Get("/shut/get", func(parmas martini.Params) string {
		return ShutdownVar
	})

	// Basic URL
	mart.Get("/basic/change/:var", func(parmas martini.Params) string {
		BasicVar = parmas["var"]
		return BasicVar
	})

	mart.Get("/basic/get", func(parmas martini.Params) string {
		return BasicVar
	})

	// Run Server
	mart.Run()
}
