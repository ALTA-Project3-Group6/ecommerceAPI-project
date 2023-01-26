package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	UserID      uint
	ProductID   uint
	Price       float64
	Quantity    int
	ProductName string
	SellerName  string
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
