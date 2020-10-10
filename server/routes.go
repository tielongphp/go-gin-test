package server

import (
	"github.com/gin-gonic/gin"

	"go-gin-test/context"
	"go-gin-test/controller"
	"go-gin-test/controller/v1"
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

	shopOrderRouter := app.Group("shopOrder")
	//Use(middleware.JWTAuth())
	{
		shopOrderRouter.GET("getInfo", v1.GetShopOrderInfoByOrderId) // 查询
		shopOrderRouter.PUT("updateOne", v1.UpdateShopOrder)         // 更新
		shopOrderRouter.POST("addOne", v1.AddShopOrder)              // 添加
		shopOrderRouter.DELETE("delOne", v1.DelShopOrder)            // 删除
	}
}
