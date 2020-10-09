package shop_order_service

import (
	"go-gin-test/model"
)

//func Add() (int, error) {
//	shopOrder := map[string]interface{}{
//		"order_id":    a.OrderID,
//		"order_sn":    a.OrderSn,
//		"create_time": a.CreateTime,
//		// ...
//	}
//	orderId, err := models.AddShopOrder(shopOrder)
//	if orderId == 0 || err != nil {
//		return 0, err
//	}
//	return orderId, nil
//}

func GetShopOrder(orderId int) (err error, info interface{}) {
	var shopOrderApi model.ShopOrder

	// 指定使用从库
	//db := model.DB.Clauses(dbresolver.Read).Model(&model.ShopOrder{})

	// 指定使用主库
	//db := model.DB.Clauses(dbresolver.Write).Model(&model.ShopOrder{})

	// 自动读写分离（写：选择主库，读：选择从库）
	db := model.DB.Model(&model.ShopOrder{})

	//if orderId > 0 {
	db = db.Where("order_id = ?", orderId)
	//}

	err = db.First(&shopOrderApi).Error
	return err, shopOrderApi
}
