package data

import (
	"ecommerceapi/features/cart"
	product "ecommerceapi/features/product/data"
	user "ecommerceapi/features/user/data"
)

type Cart struct {
	ID        uint
	UserID    uint
	ProductID uint
	Quantity  int
	Price     float64
	Seller    user.User       `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Product   product.Product `gorm:"foreignkey:ProductID;association_foreignkey:ID"`
}

func DataToCore(data Cart) cart.Core {
	return cart.Core{
		ID:        data.ID,
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Price:     data.Price,
	}
}

func CoreToData(data cart.Core) Cart {
	return Cart{
		ID:        data.ID,
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Price:     data.Price,
	}
}
