package services

import (
	"ecommerceapi/features/cart"
	"ecommerceapi/helper"
	"errors"
	"log"
	"strings"
)

type cartSvc struct {
	qry cart.CartData
}

func New(data cart.CartData) cart.CartService {
	return &cartSvc{
		qry: data,
	}
}

func (cs *cartSvc) AddCart(token interface{}, productId uint, newCart cart.Core) (cart.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("\t error extract token add cart")
		return cart.Core{}, errors.New("user not found")
	}
	res, err := cs.qry.AddCart(uint(userId), productId, newCart)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror add query in service: ", err.Error())
		return cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *cartSvc) ShowCart(token interface{}) ([]cart.Core, error) {
	userId := helper.ExtractToken(token)
	res, err := cs.qry.ShowCart(uint(userId))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *cartSvc) UpdateCart(token interface{}, cartId uint, updCart cart.Core) (cart.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("\t error extract token add cart")
		return cart.Core{}, errors.New("user not found")
	}

	res, err := cs.qry.UpdateCart(uint(userId), cartId, updCart)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "cart data not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror update data in service: ", err.Error())
		return cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *cartSvc) DeleteCart(token interface{}, cartId uint) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		log.Println("\terror extract token delete cart service")
		return errors.New("user not found")
	}

	err := cs.qry.DeleteCart(uint(userID), cartId)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror calling delete data in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
}
