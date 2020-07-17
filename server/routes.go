package server

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/common"
	"go-gin-test/context"
	"go-gin-test/controller"
)

func registerRoutes(app *gin.Engine, conf *context.Config) {
	//routes
	// 探针
	app.GET("/", func(c *gin.Context) {
		common.FormatResponseWithoutData(c, 10000, "成功")
	},Logger(conf))

	goFePrefix := app.Group("/test", Logger(conf))
	{
		controller.Test(goFePrefix, conf)
	}
}
