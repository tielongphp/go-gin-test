package shop_order_service

import (
	"go-gin-test/models"
)

type ShopOrder struct {
	OrderID int
	// MainOrderID 区分主订单作用（0为主订单，大于0是子订单）
	MainOrderID int
	// OrderSn 订单编号
	OrderSn string
	// UserID 会员ID
	UserID int
	// OrderStatus 支付状态：0-未支付，1-支付成功，2-订单取消,3-退款
	OrderStatus int
	// Consignee 收货人地址
	Consignee string
	// Province 省份
	Province   int
	CreateTime int
	// ...
}

func (a *ShopOrder) Add() (int, error) {
	shopOrder := map[string]interface{}{
		"order_id":    a.OrderID,
		"order_sn":    a.OrderSn,
		"create_time": a.CreateTime,
		// ...
	}
	orderId, err := models.AddShopOrder(shopOrder)
	if orderId == 0 || err != nil {
		return 0, err
	}
	return orderId, nil
}

func (a *ShopOrder) GetShopOrder(orderId int) (*models.ShopOrder, error) {
	shopOrder, err := models.GetShopOrderByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	return shopOrder, nil
}
