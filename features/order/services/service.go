package services

import (
	"ecommerceapi/config"
	"ecommerceapi/features/order"
	"ecommerceapi/helper"
	"errors"
	"log"
	"strings"
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
		msg := ""
		if strings.Contains(err.Error(), "bad request") {
			msg = "bad request"
		} else {
			msg = "server problem"
		}
		log.Println("error add order query in service : ", err.Error())
		return order.Core{}, "", errors.New(msg)
	}

	return res, redirectURL, nil
}
func (os *orderSvc) GetOrderHistory(token interface{}) ([]order.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token")
		return []order.Core{}, errors.New("user not found")
	}

	res, err := os.qry.GetOrderHistory(uint(userId))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error calling getorderhistory data in service: ", err.Error())
		return []order.Core{}, errors.New(msg)
	}

	return res, nil
}
func (os *orderSvc) GetSellingHistory(token interface{}) ([]order.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token")
		return []order.Core{}, errors.New("user not found")
	}

	res, err := os.qry.GetSellingHistory(uint(userId))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error calling getorderhistory data in service: ", err.Error())
		return []order.Core{}, errors.New(msg)
	}

	return res, nil
}
func (os *orderSvc) NotificationTransactionStatus(transactionId string) error {
	c := config.MidtransCoreAPIClient()

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := c.CheckTransaction(transactionId)
	if e != nil {
		return errors.New("error check transaction status")
	}

	err := os.qry.NotificationTransactionStatus(transactionId, transactionStatusResp.TransactionStatus)
	if err != nil {
		return errors.New("error calling NotificationTransactionStatus data in service")
	}

	return nil
}
func (os *orderSvc) UpdateStatus(orderId uint, status string) error {
	err := os.qry.UpdateStatus(orderId, status)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "bad request"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}
	return nil
}
