package server

import (
	"github.com/gin-gonic/gin"

	"go-gin-test/context"
	"go-gin-test/controller"
	"go-gin-test/response"
)

func registerRoutes(app *gin.Engine, conf *context.Config) {
	//routes
	// 探针
	app.GET("/", func(c *gin.Context) {
		response.Ok(c)
	}, Logger(conf))

	goFePrefix := app.Group("/test", Logger(conf))
	{
		controller.Test(goFePrefix, conf)
	}
}
