package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID           uint    `json:"id" form:"id"`
	UserID       uint    `json:"user_id" form:"user_id"`
	SellerID     uint    `json:"seller_id" form:"seller_id"`
	ProductID    uint    `json:"product_id" form:"product_id"`
	ProductImage string  `json:"product_image" form:"product_image"`
	Price        float64 `json:"price" form:"price"`
	Quantity     int     `json:"quantity" form:"quantity"`
	ProductName  string  `json:"product_name" form:"product_name"`
	SellerName   string  `json:"seller_name" form:"seller_name"`
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	ShowCart() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}

type CartService interface {
	AddCart(token interface{}, productId uint, newCart Core) (Core, error)
	ShowCart(token interface{}) ([]Core, error)
	UpdateCart(token interface{}, cartId uint, updCart Core) (Core, error)
	DeleteCart(token interface{}) error
}

type CartData interface {
	AddCart(userId uint, productId uint, newCart Core) (Core, error)
	ShowCart(userId uint) ([]Core, error)
	UpdateCart(userId uint, cartId uint, updCart Core) (Core, error)
	DeleteCart(userId uint) error
}
