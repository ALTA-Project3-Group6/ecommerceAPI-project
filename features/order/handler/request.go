package handler

type OrderReq struct {
	TotalPrice float64 `json:"total_price" form:"total_price"`
}
