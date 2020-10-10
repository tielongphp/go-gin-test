package initialize

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/context"
	"go-gin-test/middleware"
	"go-gin-test/router"
)

// 初始化总路由
func RoutersInit(app *gin.Engine, conf *context.Config) *gin.Engine {
	//var Router = gin.Default()

	//shopOrderRouter := app.Group("shopOrder").
	//	//Use(middleware.JWTAuth()).
	//	Use(middleware.Cors()) // 跨域
	//{
	//	shopOrderRouter.GET("getInfo", v1.GetShopOrderInfoByOrderId) // 查询one
	//	shopOrderRouter.GET("getList", v1.GetShopOrderList)          // 查询列表
	//	shopOrderRouter.PUT("updateOne", v1.UpdateShopOrder)         // 更新
	//	shopOrderRouter.POST("addOne", v1.AddShopOrder)              // 添加
	//	shopOrderRouter.DELETE("delOne", v1.DelShopOrder)            // 删除
	//}

	// app.Use(middleware.LoadTls())  // 打开就能玩https

	// 方便统一添加路由组前缀
	ApiGroup := app.Group("", middleware.Logger(conf))
	router.InitCommonRouter(ApiGroup, conf)
	router.InitShopOrderRouter(ApiGroup, conf) // shopOrder路由

	return app
}
