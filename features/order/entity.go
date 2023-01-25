package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	BuyerId       uint
	SellerId      uint
	TotalPrice    float64
	CreatedAt     time.Time
	OrderStatus   string
	TransactionId string
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
