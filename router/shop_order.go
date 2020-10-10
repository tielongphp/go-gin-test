package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/context"
	"go-gin-test/controller/v1"
	"go-gin-test/middleware"
)

func InitShopOrderRouter(app *gin.RouterGroup, conf *context.Config) {
	shopOrderRouter := app.Group("shopOrder").
		//Use(middleware.JWTAuth()).
		Use(middleware.Cors()) // 跨域
	{
		shopOrderRouter.GET("getInfo", v1.GetShopOrderInfoByOrderId) // 查询one
		shopOrderRouter.GET("getList", v1.GetShopOrderList)          // 查询列表
		shopOrderRouter.PUT("updateOne", v1.UpdateShopOrder)         // 更新
		shopOrderRouter.POST("addOne", v1.AddShopOrder)              // 添加
		shopOrderRouter.DELETE("delOne", v1.DelShopOrder)            // 删除
	}
}
