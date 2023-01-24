package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	UserId       uint
	Name         string
	ProductImage string
	Description  string
	Stock        int
	Price        float64
}

type ProductHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAllProducts() echo.HandlerFunc
	GetUserProducts() echo.HandlerFunc
	GetProductById() echo.HandlerFunc
}

type ProductService interface {
	Add(token interface{}, newProduct Core, productImage *multipart.FileHeader) (Core, error)
	Update(token interface{}, productId uint, updProduct Core, productImage *multipart.FileHeader) (Core, error)
	Delete(token interface{}, productId uint) error
	GetAllProducts() ([]Core, error)
	GetUserProducts(token interface{}) ([]Core, error)
	GetProductById(token interface{}, productId uint)
}

type ProductData interface {
	Add(userId uint, newProduct Core) (Core, error)
	Update(userId, productId uint, updProduct Core) (Core, error)
	Delete(userId, productId uint) error
	GetAllProducts() ([]Core, error)
	GetUserProducts(userId uint) ([]Core, error)
	GetProductById(userId, productId uint) (Core, error)
}
