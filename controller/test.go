package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-gin-test/common"
	"go-gin-test/context"
	"go-gin-test/service/shop_order_service"
)

func Test(router *gin.RouterGroup, conf *context.Config) {
	router.GET("/a", func(c *gin.Context) {
		// 查询
		articleService := shop_order_service.ShopOrder{
			OrderID: 1,
			// MainOrderID 区分主订单作用（0为主订单，大于0是子订单）

		}
		data, _ := articleService.GetShopOrder(1)

		// 添加
		articleServiceAdd := shop_order_service.ShopOrder{
			//OrderID:    5,
			OrderSn:    "202020200202",
			CreateTime: 15512121212,
		}
		flag1, _ := articleServiceAdd.Add()

		fmt.Println("添加：", flag1)

		if data != nil {
			common.FormatResponse(c, 10000, "成功", data)
		} else {
			common.GetCodeMsg("DATA_NOT_FIND", c)
			//common.FormatResponseWithoutData(c, common.MsgCode{Code: }, "成功")
		}

	})
}
