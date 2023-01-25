package handler

import (
	"ecommerceapi/features/order"
	"time"
)

type OrderResp struct {
	ID            uint      `json:"id"`
	BuyerId       uint      `json:"buyer_id"`
	SellerId      uint      `json:"seller_id"`
	TotalPrice    float64   `json:"total_price"`
	CreatedAt     time.Time `json:"created_at"`
	OrderStatus   string    `json:"order_status"`
	TransactionId string    `json:"transaction_id"`
	RedirectURL   string    `json:"redirect_url"`
}

func CoreToOrderResp(data order.Core, url string) OrderResp {
	return OrderResp{
		ID:            data.ID,
		BuyerId:       data.BuyerId,
		SellerId:      data.SellerId,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
		OrderStatus:   data.OrderStatus,
		TransactionId: data.TransactionId,
		RedirectURL:   url,
	}
}
