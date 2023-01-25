package data

import (
	"ecommerceapi/features/order"
	user "ecommerceapi/features/user/data"
	"time"
)

type Order struct {
	ID            uint
	BuyerId       uint
	SellerId      uint
	TotalPrice    float64
	CreatedAt     time.Time
	OrderStatus   string
	TransactionId string

	Seller user.User `gorm:"foreignkey:SellerId;association_foreignkey:ID"`
	Buyer  user.User `gorm:"foreignkey:BuyerId;association_foreignkey:ID"`
}

func DataToCore(data Order) order.Core {
	return order.Core{
		ID:            data.ID,
		BuyerId:       data.BuyerId,
		SellerId:      data.SellerId,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
		OrderStatus:   data.OrderStatus,
		TransactionId: data.TransactionId,
	}
}

func CoreToData(data order.Core) Order {
	return Order{
		ID:            data.ID,
		BuyerId:       data.BuyerId,
		SellerId:      data.SellerId,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
		OrderStatus:   data.OrderStatus,
		TransactionId: data.TransactionId,
	}
}
