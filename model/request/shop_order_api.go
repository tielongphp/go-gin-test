package request

//
type ShopOrderApiInfoParams struct {
	OrderID int `form:"order_id"`
}

type OrderIdJsonParams struct {
	OrderID int `json:"order_id"`
}
