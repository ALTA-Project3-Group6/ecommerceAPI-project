package handler

import (
	"ecommerceapi/features/product"
)

type AddProductReq struct {
	Name         string  `json:"product_name" form:"product_name"`
	ProductImage string  `json:"product_image" form:"product_image"`
	Description  string  `json:"description" form:"description"`
	Stock        int     `json:"stock" form:"stock"`
	Price        float64 `json:"price" form:"price"`
}

func ToCore(data interface{}) *product.Core {
	res := product.Core{}

	switch data.(type) {
	case AddProductReq:
		cnv := data.(AddProductReq)
		res.Name = cnv.Name
		res.ProductImage = cnv.ProductImage
		res.Description = cnv.Description
		res.Stock = cnv.Stock
		res.Price = cnv.Price
	default:
		return nil
	}
	return &res
}
