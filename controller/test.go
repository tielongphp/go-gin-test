package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-gin-test/context"
	"go-gin-test/model/request"
	"go-gin-test/response"
	"go-gin-test/service/shop_order_service"
	"go-gin-test/utils"
)

func Test(router *gin.RouterGroup, conf *context.Config) {
	router.GET("/a", func(c *gin.Context) {
		// 查询
		var so request.ShopOrderApiInfoParams
		_ = c.BindQuery(&so) // get参数
		// 参数非空验证
		ShopOrderVerify := utils.Rules{
			"OrderID": {utils.NotEmpty()},
		}
		ShopOrderVerifyErr := utils.Verify(so, ShopOrderVerify)

		if ShopOrderVerifyErr != nil {
			response.FailWithMsg(ShopOrderVerifyErr.Error(), c)
			return
		}
		// model层读写分离实现
		err, info := shop_order_service.GetShopOrder(so.OrderID)
		if err != nil {
			response.FailWithMsg(fmt.Sprintf("获取失败：%v", err), c)
		} else {
			response.OkWithData(info, c)
		}

	})
}
