package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-test/context"
	"go-gin-test/controller"
	"go-gin-test/middleware"
	"go-gin-test/response"
)

func InitCommonRouter(app *gin.RouterGroup, conf *context.Config) {

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 探针
	app.GET("/", func(c *gin.Context) {
		response.Ok(c)
	}, middleware.Logger(conf))

	// 测试用
	goFePrefix := app.Group("/test", middleware.Logger(conf))
	{
		controller.Test(goFePrefix, conf)
	}
}
