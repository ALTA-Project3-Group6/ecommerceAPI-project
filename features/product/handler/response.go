package handler

import "ecommerceapi/features/product"

type UserProductRes struct {
	UserId   uint   `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
}
type GetProductsResp struct {
	ID           uint    `json:"id" form:"id"`
	Name         string  `json:"name" form:"name"`
	Description  string  `json:"description" form:"description"`
	Stock        int     `json:"stock" form:"stock"`
	Price        float64 `json:"price" form:"price"`
	ProductImage string  `json:"product_image" form:"product_image"`
	User         UserProductRes
}

func CoreToGetProductsResp(data []product.Core) []GetProductsResp {
	res := []GetProductsResp{}
	for i, val := range data {
		res[i].ID = val.ID
		res[i].Name = val.Name
		res[i].Description = val.Name
		res[i].Stock = val.Stock
		res[i].Price = val.Price
		res[i].ProductImage = val.ProductImage
		res[i].User.UserId = val.UserId
		res[i].User.UserName = val.UserName
	}
	return res
}

func CoreToGetProductResp(data product.Core) GetProductsResp {
	return GetProductsResp{
		ID:           data.ID,
		Name:         data.Name,
		Description:  data.Description,
		Stock:        data.Stock,
		Price:        data.Price,
		ProductImage: data.ProductImage,
		User:         UserProductRes{UserId: data.UserId, UserName: data.UserName},
	}
}
