package data

import (
	"ecommerceapi/features/product"
	user "ecommerceapi/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserId       uint
	User         user.User
	Name         string
	ProductImage string
	Description  string
	Stock        int
	Price        float64
}

func DataToCore(data Product) product.Core {
	return product.Core{
		ID:           data.ID,
		UserId:       data.UserId,
		Name:         data.Name,
		ProductImage: data.ProductImage,
		Description:  data.Description,
		Stock:        data.Stock,
		Price:        data.Price,
	}
}

func CoreToData(data product.Core) Product {
	return Product{
		Model:        gorm.Model{ID: data.ID},
		UserId:       data.UserId,
		Name:         data.Name,
		ProductImage: data.ProductImage,
		Description:  data.Description,
		Stock:        data.Stock,
		Price:        data.Price,
	}
}
