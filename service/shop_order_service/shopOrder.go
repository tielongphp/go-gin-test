package shop_order_service

import (
	"fmt"
	"unsafe"

	"gorm.io/plugin/dbresolver"

	"go-gin-test/global"
	"go-gin-test/model"
	"go-gin-test/model/request"
)

/**
 * 查询示列
 */
func GetShopOrder(orderId int) (err error, info interface{}) {
	var shopOrderApi model.ShopOrder

	// 指定使用从库
	//db := global.DB.Clauses(dbresolver.Read).Model(&model.ShopOrder{})

	// 指定使用主库
	//db := global.DB.Clauses(dbresolver.Write).Model(&model.ShopOrder{})

	// 自动读写分离（写：选择主库，读：选择从库）
	db := global.DB.Model(&model.ShopOrder{})

	//if orderId > 0 {
	db = db.Where("order_id = ?", orderId)
	//}

	err = db.First(&shopOrderApi).Error
	return err, shopOrderApi
}

/**
 * 查询示列
 */
func GetShopOrderList(info request.PageInfo) (err error, list interface{}, total int) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 指定使用从库
	//db := global.DB.Clauses(dbresolver.Read).Model(&model.ShopOrder{})
	// 指定使用主库
	db := global.DB.Clauses(dbresolver.Write).Model(&model.ShopOrder{})

	// 自动读写分离（写：选择主库，读：选择从库）
	//db := global.DB.Model(&model.ShopOrder{})
	var shop []model.ShopOrder
	db.Count((*int64)(unsafe.Pointer(&total)))
	err = db.Limit(limit).Offset(offset).Find(&shop).Error
	if len(shop) > 0 {
		for k := range shop {
			fmt.Println(k)
		}
	}
	return err, shop, total

}

/**
 * 更新示列
 */
func UpDateShopOrder(shop model.ShopOrder) (err error, rows int64) {
	// 自动读写分离（写：选择主库，读：选择从库）
	db := global.DB.Model(&model.ShopOrder{})
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
	db := global.DB.Model(&model.ShopOrder{})
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
	db := global.DB.Model(&model.ShopOrder{})
	result := db.Where("order_id = ?", shop.OrderID).Delete(shop)
	rows = result.RowsAffected // 删除的记录数
	err = result.Error         // 更新的错误
	return err, rows
}
