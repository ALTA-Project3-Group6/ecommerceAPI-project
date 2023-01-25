package services

import (
	"ecommerceapi/features/order"
	"ecommerceapi/helper"
	"errors"
	"log"
)

type orderSvc struct {
	qry order.OrderData
}

func New(data order.OrderData) order.OrderService {
	return &orderSvc{
		qry: data,
	}
}

func (os *orderSvc) Add(token interface{}, totalPrice float64) (order.Core, string, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token add order")
		return order.Core{}, "", errors.New("user not found")
	}

	res, redirectURL, err := os.qry.Add(uint(userId), totalPrice)
	if err != nil {
		log.Println("error add order query in service : ", err.Error())
		return order.Core{}, "", errors.New("server problem")
	}

	return res, redirectURL, nil
}
func (os *orderSvc) GetOrderHistory(token interface{}) ([]order.Core, error) {
	return []order.Core{}, nil
}
func (os *orderSvc) GetSellingHistory(token interface{}) ([]order.Core, error) {
	return []order.Core{}, nil
}
func (os *orderSvc) GetTransactionStatus(orderId uint) (string, error) {
	return "", nil
}
