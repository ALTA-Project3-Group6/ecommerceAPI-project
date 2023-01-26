package handler

type OrderReq struct {
	TotalPrice float64 `json:"total_price" form:"total_price"`
}

type StatusReq struct {
	OrderStatus string `json:"order_status" form:"order_status"`
}
