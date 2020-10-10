package v1

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go-gin-test/model"
	"go-gin-test/model/request"
	"go-gin-test/response"
	"go-gin-test/service/shop_order_service"
	"go-gin-test/utils"
)

// @Tags ShopOrder
// @Summary 查询ShopOrder
// @Router /shopOrder/getInfo [get]
func GetShopOrderInfoByOrderId(c *gin.Context) {
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
}

// @Tags ShopOrder
// @Summary 查询ShopOrder列表
// @Router /shopOrder/getList [get]
func GetShopOrderList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.BindQuery(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMsg(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := shop_order_service.GetShopOrderList(pageInfo)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags ShopOrder
// @Summary 更新ShopOrder
// @Router /shopOrder/updateOne [put]
func UpdateShopOrder(c *gin.Context) {
	var so model.ShopOrder
	//var so request.OrderIdJsonParams
	_ = c.ShouldBindJSON(&so) // post json参数
	// 参数非空验证
	ShopOrderVerify := utils.Rules{
		"OrderID":     {utils.NotEmpty()},
		"MainOrderID": {utils.NotEmpty()},
	}
	ShopOrderVerifyErr := utils.Verify(so, ShopOrderVerify)

	if ShopOrderVerifyErr != nil {
		response.FailWithMsg(ShopOrderVerifyErr.Error(), c)
		return
	}
	// model层读写分离实现
	err, _ := shop_order_service.UpDateShopOrder(so)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("更新失败：%v", err), c)
	} else {
		response.Ok(c)
	}
}

// @Tags ShopOrder
// @Summary 添加
// @Router /shopOrder/addOne [post]
func AddShopOrder(c *gin.Context) {
	var so model.ShopOrder
	//var so request.OrderIdJsonParams
	_ = c.ShouldBindJSON(&so) // post json参数
	// 参数非空验证
	ShopOrderVerify := utils.Rules{
		"OrderID":     {utils.NotEmpty()},
		"MainOrderID": {utils.NotEmpty()},
		"OrderSn":     {utils.NotEmpty()},
		"UserID":      {utils.NotEmpty()},
	}
	ShopOrderVerifyErr := utils.Verify(so, ShopOrderVerify)
	timestamp := time.Now().Unix()
	so.CreateTime = timestamp
	so.UpdateTime = timestamp

	if ShopOrderVerifyErr != nil {
		response.FailWithMsg(ShopOrderVerifyErr.Error(), c)
		return
	}
	// model层读写分离实现
	err, _ := shop_order_service.AddShopOrder(so)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("添加失败：%v", err), c)
	} else {
		response.Ok(c)
	}
}

// @Tags ShopOrder
// @Summary 删除
// @Router /shopOrder/delOne [delete]
func DelShopOrder(c *gin.Context) {
	var so model.ShopOrder
	//var so request.OrderIdJsonParams
	_ = c.ShouldBindJSON(&so) // post json参数
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
	err, _ := shop_order_service.DelShopOrder(so)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.Ok(c)
	}
}
