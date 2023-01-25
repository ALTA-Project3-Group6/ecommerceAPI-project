package handler

import "ecommerceapi/features/cart"

type SellerProductsRes struct {
	UserId   uint   `json:"id_user" form:"id_user"`
	UserName string `json:"user_name" form:"user_name"`
}

type ProductsSellerRes struct {
	ProductId   uint   `json:"id_product" form:"id_product"`
	ProductName string `json:"product_name" form:"product_name"`
}

type CartsResp struct {
	ID        uint              `json:"id" form:"id"`
	UserID    uint              `json:"id_user" form:"id_user"`
	ProductID uint              `json:"id_product" form:"id_product"`
	Quantity  int               `json:"quantity" form:"quantity"`
	Price     float64           `json:"price" form:"price"`
	Product   ProductsSellerRes `json:"product"`
	Seller    SellerProductsRes `json:"seller"`
}

func CoreToCartsResp(data []cart.Core) []CartsResp {
	res := []CartsResp{}
	for _, val := range data {
		res = append(res, CoreToCartResp(val))
	}
	return res
}

func CoreToCartResp(data cart.Core) CartsResp {
	return CartsResp{
		ID:        data.ID,
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Price:     data.Price,
		Product:   ProductsSellerRes{ProductId: data.ProductID, ProductName: data.ProductName},
		Seller:    SellerProductsRes{UserId: data.UserID, UserName: data.SellerName},
	}
}