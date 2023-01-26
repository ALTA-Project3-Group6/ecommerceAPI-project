package data

import (
	"ecommerceapi/features/order"
	product "ecommerceapi/features/product/data"
	user "ecommerceapi/features/user/data"
	"time"
)

type Order struct {
	ID            uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	BuyerId       uint
	SellerId      uint
	TotalPrice    float64
	CreatedAt     time.Time
	OrderStatus   string
	TransactionId string
	PaymentURL    string

	Seller user.User `gorm:"foreignkey:SellerId;association_foreignkey:ID"`
	Buyer  user.User `gorm:"foreignkey:BuyerId;association_foreignkey:ID"`
}

type OrderProduct struct {
	ID        uint    `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`

	Order   Order           `gorm:"foreignkey:OrderId;association_foreignkey:ID"`
	Product product.Product `gorm:"foreignkey:ProductId;association_foreignkey:ID"`
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
