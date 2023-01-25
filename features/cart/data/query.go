package data

import (
	"ecommerceapi/features/cart"
	"errors"
	"log"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

func (cq *cartQuery) AddCart(userId uint, productId uint, newCart cart.Core) (cart.Core, error) {
	cnvC := CoreToData(newCart)
	cnvC.UserID = userId
	cnvC.ProductID = productId

	err := cq.db.Create(&cnvC).Error
	if err != nil {
		log.Println("\tadd cart query error: ", err.Error())
		return cart.Core{}, errors.New("server problem")
	}
	return DataToCore(cnvC), nil
}

func (cq *cartQuery) ShowCart(userId uint) ([]cart.Core, error) {
	allcart := []cart.Core{}
	err := cq.db.Raw("SELECT carts.user_id, carts.product_id, products.price, carts.quantity FROM carts JOIN users ON cart.user_id = users.id JOIN products ON carts.product_id = products.id").Scan(&allcart).Error
	if err != nil {
		log.Println("\terror query get all cart: ", err.Error())
		return []cart.Core{}, err
	}

	return allcart, nil
}

func (cq *cartQuery) UpdateCart(userId uint, cartId uint, updCart cart.Core) (cart.Core, error) {
	cnvC := CoreToData(updCart)
	cnvC.UserID = userId
	cnvC.ID = cartId

	qry := cq.db.Where("id = ? AND user_id = ?", cartId, userId).Updates(&cnvC)
	if qry.RowsAffected <= 0 {
		log.Println("\tupdate cart query error: data not found")
		return cart.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tupdate cart query error: ", err.Error())
		return cart.Core{}, errors.New("not found")
	}
	return DataToCore(cnvC), nil
}

func (cq *cartQuery) DeleteCart(userId uint, cartId uint) error {
	qry := cq.db.Where("user_id = ?", userId).Delete(&Cart{}, cartId)

	if aff := qry.RowsAffected; aff <= 0 {
		log.Println("\tno rows affected: data not found")
		return errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tdelete query error: ", err.Error())
		return err
	}
	return nil
}