package shop_order_service

import (
	"go-gin-test/model"
)

/**
 * 查询示列
 */
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

/**
 * 更新示列
 */
func UpDateShopOrder(shop model.ShopOrder) (err error, rows int64) {
	// 自动读写分离（写：选择主库，读：选择从库）
	db := model.DB.Model(&model.ShopOrder{})
	db = db.Where("order_id = ?", shop.OrderID)
	result := db.First(&model.ShopOrder{}).Updates(&shop)

	rows = result.RowsAffected // 更新的记录数
	err = result.Error         // 更新的错误
	return err, rows
}

/**
 * 更新示列
 */
func AddShopOrder(shop model.ShopOrder) (err error, rows int64) {
	// 自动读写分离（写：选择主库，读：选择从库）
	db := model.DB.Model(&model.ShopOrder{})
	result := db.Create(&shop)

	rows = result.RowsAffected // 更新的记录数
	err = result.Error         // 更新的错误
	return err, rows
}

/**
 * 删除示列
 */
func DelShopOrder(shop model.ShopOrder) (err error, rows int64) {
	// 自动读写分离（写：选择主库，读：选择从库）
	db := model.DB.Model(&model.ShopOrder{})
	result := db.Where("order_id = ?", shop.OrderID).Delete(shop)
	rows = result.RowsAffected // 删除的记录数
	err = result.Error         // 更新的错误
	return err, rows
}
