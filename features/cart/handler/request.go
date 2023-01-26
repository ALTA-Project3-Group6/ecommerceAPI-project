package handler

import "ecommerceapi/features/cart"

type AddCartReq struct {
	UserID    uint    `json:"id_user" form:"id_user"`
	ProductID uint    `json:"id_product" form:"id_product"`
	Quantity  int     `json:"quantity" form:"quantity"`
	Price     float64 `json:"price" form:"price"`
}

type UpdCartReq struct {
	Quantity int `json:"quantity" form:"quantity"`
}

func ToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddCartReq:
		cnv := data.(AddCartReq)
		res.UserID = cnv.UserID
		res.ProductID = cnv.ProductID
		res.Quantity = cnv.Quantity
		res.Price = cnv.Price
	case UpdCartReq:
		cnv := data.(UpdCartReq)
		res.Quantity = cnv.Quantity
	default:
		return nil
	}
	return &res
}
