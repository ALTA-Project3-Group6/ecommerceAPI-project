package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint      `json:"id"`
	BuyerId       uint      `json:"buyer_id"`
	BuyerName     string    `json:"buyer_name"`
	SellerId      uint      `json:"seller_id"`
	SellerName    string    `json:"seller_name"`
	TotalPrice    float64   `json:"total_price"`
	CreatedAt     time.Time `json:"created_at"`
	OrderStatus   string    `json:"order_status"`
	TransactionId string    `json:"transaction_id"`
}

type OrderHandler interface {
	Add() echo.HandlerFunc
	GetOrderHistory() echo.HandlerFunc
	GetSellingHistory() echo.HandlerFunc
	GetTransactionStatus() echo.HandlerFunc
}

type OrderService interface {
	Add(token interface{}, totalPrice float64) (Core, string, error)
	GetOrderHistory(token interface{}) ([]Core, error)
	GetSellingHistory(token interface{}) ([]Core, error)
	GetTransactionStatus(orderId uint) (string, error)
}

type OrderData interface {
	Add(userId uint, totalPrice float64) (Core, string, error)
	GetOrderHistory(userId uint) ([]Core, error)
	GetSellingHistory(userId uint) ([]Core, error)
	GetTransactionStatus(orderId uint) (string, error)
}
