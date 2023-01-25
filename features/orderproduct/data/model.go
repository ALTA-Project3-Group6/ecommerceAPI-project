package data

import (
	order "ecommerceapi/features/order/data"
	product "ecommerceapi/features/product/data"
)

type OrderProduct struct {
	ID        uint    `json:"id"`
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`

	Order   order.Order     `gorm:"foreignkey:OrderId;association_foreignkey:ID"`
	Product product.Product `gorm:"foreignkey:ProductId;association_foreignkey:ID"`
}
