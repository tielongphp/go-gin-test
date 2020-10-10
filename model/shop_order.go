package model

//// 基本模型的定义
//type ShopOrderModel struct {
//	OrderID int `gorm:"primary_key" json:"order_id"`
//	// CreateTime 订单生成时间
//	CreateTime int `json:"create_time"`
//	// UpdateTime 订单修改时间
//	UpdateTime int `json:"update_time"`
//	//CreateTime time.Time
//	//UpdateTime time.Time
//	//DeleteTime *time.Time
//}
type ShopOrder struct {
	//gorm.Model
	OrderID     int    `gorm:"primary_key;column:order_id" json:"order_id"` //设置一个普通的索引，没有设置索引名，gorm会自动命名
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime  int64  `gorm:"column:update_time" json:"update_time"`
	MainOrderID int    `gorm:"column:main_order_id" json:"main_order_id"`
	OrderSn     string `gorm:"column:order_sn" json:"order_sn"`
	UserID      string `gorm:"column:user_id" json:"user_id"`
	OrderStatus int    `gorm:"column:order_status" json:"order_status"`
	Consignee   string `gorm:"column:consignee" json:"consignee"`
}

// ShopOrder ...
//type ShopOrder struct {
//	OrderID int `gorm:"primary_key" json:"order_id"`
//	// CreateTime 订单生成时间
//	CreateTime int `json:"create_time"`
//	// UpdateTime 订单修改时间
//	UpdateTime int `json:"update_time"`
//	// MainOrderID 区分主订单作用（0为主订单，大于0是子订单）
//	MainOrderID int `json:"main_order_id"`
//	// OrderSn 订单编号
//	OrderSn string `json:"order_sn"`
//	// UserID 会员ID
//	UserID int `json:"user_id"`
//	// OrderStatus 支付状态：0-未支付，1-支付成功，2-订单取消,3-退款
//	OrderStatus int `json:"order_status"`
//	// Consignee 收货人地址
//	Consignee string `json:"consignee"`
//	// Province 省份
//	Province int `json:"province"`
//	// City 城市
//	City int `json:"city"`
//	// District 地区
//	District int `json:"district"`
//	// Street 街道
//	Street int `json:"street"`
//	// Address 收货地址
//	Address string `json:"address"`
//	// Mobile 手机号码
//	Mobile string `json:"mobile"`
//	// Email 邮箱
//	Email string `json:"email"`
//	// Postscript 订单附言
//	Postscript string `json:"postscript"`
//	// OrderAmount 应付款金额
//	OrderAmount float32 `json:"order_amount"`
//	// Closed 删除：0-正常，1-删除
//	Closed int `json:"closed"`
//	// CouponID 优惠卷ID
//	CouponID int `json:"coupon_id"`
//	// CouponCode 优惠卷码
//	CouponCode string `json:"coupon_code"`
//	// Integral 积分数
//	Integral int `json:"integral"`
//	// CashDeducted 积分抵扣金额
//	CashDeducted float32 `json:"cash_deducted"`
//}

// 设置ShopOrder的表名为`gz_shop_order`
//func (GzShopOrder) TableName() string {
//return "gz_shop_order"
//}

//func (ShopOrder *ShopOrder) BeforeCreate(scope *gorm.Scope) error {
// scope.SetColumn("create_time", time.Now().Unix())
// return nil
//}
//
//func (ShopOrder *ShopOrder) BeforeUpdate(scope *gorm.Scope) error {
// scope.SetColumn("create_time", time.Now().Unix())
// return nil
//}
//type ShopOrderApi struct {
//	OrderID int    `json:"order_id"`
//	OrderSn string `json:"order_sn"`
//}
//type User struct {
//}

//func GetShopOrderByOrderId(orderId int) (interface{}, error) {
//	shopOrderApi := &ShopOrderApi{OrderID: orderId, OrderSn: ""}
//	DB.Table("gz_shop_order").Clauses(dbresolver.Read).First(&shopOrderApi)
//	//err := DB.Table("gz_shop_order").First(&shopOrderApi).Error
//	//err := DB.Table("gz_shop_order").Select("order_id, order_sn").Where("order_id = ?", orderId).Scan(&ShopOrder).Error
//	//err := DB.Table("gz_shop_order").Select("order_id, order_sn").Where("order_id = ?", orderId).Find(&ShopOrderApi{}).Error
//	//err := DB.Table("gz_shop_order").Clauses(dbresolver.Write).Where("order_id = ?", orderId).First(&shopOrderApi).Error
//	//DB.Model(&ShopOrder{}).Limit(10).Find(&ShopOrderApi{})
//	//DB.Table("gz_shop_order").Clauses(dbresolver.Write).First(&ShopOrderApi{})
//	//fmt.Println(err, 222222222)
//	//if err != nil || err == gorm.ErrRecordNotFound {
//	//	return nil, err
//	//}
//	return ShopOrderApi, nil
//}
//func GetShopOrders(pageNum, pageSize int, maps interface{}) ([]*ShopOrder, error) {
//	var ShopOrders []*ShopOrder
//	err := DB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&ShopOrders).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//	return ShopOrders, nil
//}
//func EditShopOrder(id int, data interface{}) error {
//	if err := DB.Model(&ShopOrder{}).Where("id = ? AND closed = ?", id, 0).Save(data).Error; err != nil {
//		return err
//	}
//	return nil
//}
//func AddShopOrder(data map[string]interface{}) (int, error) {
//	ShopOrder := ShopOrder{
//		OrderID:    data["order_id"].(int),
//		OrderSn:    data["order_sn"].(string),
//		CreateTime: data["create_time"].(int),
//	}
//	if err := DB.Create(&ShopOrder).Error; err != nil {
//		return 0, err
//	}
//	return ShopOrder.OrderID, nil
//}
//func DeleteGzShopOrder(id int) error {
//	if err := DB.Where("id = ?", id).Delete(ShopOrder{}).Error; err != nil {
//		return err
//	}
//	return nil
//}
//func CleanAllGzShopOrder() error {
//	if err := DB.Unscoped().Where("closed = ?", 0).Delete(&ShopOrder{}).Error; err != nil {
//		return err
//	} // 硬删除使用 Unscoped()，GORM 的约定
//	return nil
//}
