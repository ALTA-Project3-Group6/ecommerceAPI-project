package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint    `json:"id" form:"id"`
	UserId       uint    `json:"user_id" form:"user_id"`
	UserName     string  `json:"user_name" form:"user_name"`
	Name         string  `json:"name" form:"name"`
	ProductImage string  `json:"product_image" form:"product_image"`
	Description  string  `json:"description" form:"description"`
	Stock        int     `json:"stock" form:"stock"`
	Price        float64 `json:"price" form:"price"`
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
	GetProductById(token interface{}, productId uint) (Core, error)
}

type ProductData interface {
	Add(userId uint, newProduct Core, productImage *multipart.FileHeader) (Core, error)
	Update(userId, productId uint, updProduct Core, productImage *multipart.FileHeader) (Core, error)
	Delete(userId, productId uint) error
	GetAllProducts() ([]Core, error)
	GetUserProducts(userId uint) ([]Core, error)
	GetProductById(userId, productId uint) (Core, error)
}
