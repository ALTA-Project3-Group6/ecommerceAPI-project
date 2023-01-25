package data

import (
	product "ecommerceapi/features/product/data"
	user "ecommerceapi/features/user/data"
)

type Cart struct {
	ID        uint
	UserId    uint
	ProductId uint
	Quantity  int
	Price     float64

	Seller  user.User       `gorm:"foreignkey:UserId;association_foreignkey:ID"`
	Product product.Product `gorm:"foreignkey:ProductId;association_foreignkey:ID"`
}
