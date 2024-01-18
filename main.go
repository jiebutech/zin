package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebutech/app"
	"github.com/jiebutech/datasource"
	"github.com/jiebutech/zin/conf"
	"github.com/jiebutech/zin/route"
	"github.com/jiebutech/zin/wire"
)

func main() {
	cfg := conf.GetConfig()
	orm := datasource.Open(*cfg.Database)
	app.RunApp(conf.GetConfig().Configuration,
		app.BeforeRegister(func() {

		}),
		app.RouteRegister(func(c *gin.Engine) {
			demoCtrl := wire.NewDemoController(orm)
			route.RegisterDemoRoute(c, demoCtrl)
		}),
		app.AfterRegister(func() {

		}),
	)
}
