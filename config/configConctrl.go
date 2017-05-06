package main

import (
	"github.com/ivpusic/golog"
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo/middlewares/logger"
)

var (
	log      = golog.GetLogger("application")
	basicVar = "1000"
)

func main() {
	log.Info("My Clock Config Contorl")

	app := neo.App()
	app.Use(logger.Log)

	regionBasic := app.Region().Prefix("/basic")
	regionBasic.Get("/get", func(ctx *neo.Ctx) (int, error) {
		return 200, ctx.Res.Text(basicVar)
	})
	regionBasic.Get("/change/:bvar", func(ctx *neo.Ctx) (int, error) {
		basicVar = ctx.Req.Params.Get("bvar")
		return 200, ctx.Res.Text("ok")
	})

	app.Start()
}
