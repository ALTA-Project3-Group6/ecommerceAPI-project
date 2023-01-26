package services

import (
	"ecommerceapi/features/order"
	"ecommerceapi/helper"
	"ecommerceapi/mocks"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewOrderData(t)

	totalPrice := float64(1000000)

	resData := order.Core{
		ID:            1,
		BuyerId:       1,
		SellerId:      2,
		BuyerName:     "Fauzan",
		SellerName:    "Alfian",
		TotalPrice:    totalPrice,
		CreatedAt:     time.Now(),
		OrderStatus:   "waiting for payment",
		TransactionId: "Transaction-1",
	}
	url := "https://app.sandbox.midtrans.com/snap/v3/redirection/b95128fb-f2ef-4e57-bd89-c4a187cca536"

	t.Run("success add order", func(t *testing.T) {
		repo.On("Add", uint(1), totalPrice).Return(resData, url, nil).Once()

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, redirectURL, err := srv.Add(pToken, totalPrice)
		assert.Nil(t, err)
		assert.Equal(t, res, resData)
		assert.Equal(t, redirectURL, url)
		repo.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		res, url, err := srv.Add(token, totalPrice)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", url)
		repo.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		repo.On("Add", uint(1), totalPrice).Return(order.Core{}, "", errors.New("bad request")).Once()

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, redirectURL, err := srv.Add(pToken, totalPrice)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", redirectURL)
		assert.Contains(t, err.Error(), "bad")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("Add", uint(1), totalPrice).Return(order.Core{}, "", errors.New("server problem")).Once()

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, redirectURL, err := srv.Add(pToken, totalPrice)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", redirectURL)
		assert.Contains(t, err.Error(), "server")
		repo.AssertExpectations(t)
	})
}

func TestGetOrderHistory(t *testing.T) {
	repo := mocks.NewOrderData(t)
	userId := uint(1)
	resData := []order.Core{
		{
			ID:            1,
			BuyerId:       1,
			BuyerName:     "Fauzan",
			SellerId:      2,
			SellerName:    "Alfian",
			TotalPrice:    500000,
			CreatedAt:     time.Now(),
			OrderStatus:   "waiting for payment",
			TransactionId: "transaction-1",
		},
	}

	t.Run("success get order history", func(t *testing.T) {
		repo.On("GetOrderHistory", userId).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetOrderHistory(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(res), len(resData))
		repo.AssertExpectations(t)
	})

	t.Run("error extract token", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		res, err := srv.GetOrderHistory(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetOrderHistory", userId).Return([]order.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetOrderHistory(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetOrderHistory", userId).Return([]order.Core{}, errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetOrderHistory(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetSellingHistory(t *testing.T) {
	repo := mocks.NewOrderData(t)
	userId := uint(2)
	resData := []order.Core{
		{
			ID:            1,
			BuyerId:       1,
			BuyerName:     "Fauzan",
			SellerId:      2,
			SellerName:    "Alfian",
			TotalPrice:    500000,
			CreatedAt:     time.Now(),
			OrderStatus:   "waiting for payment",
			TransactionId: "transaction-1",
		},
	}

	t.Run("success get selling history", func(t *testing.T) {
		repo.On("GetSellingHistory", userId).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetSellingHistory(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(res), len(resData))
		repo.AssertExpectations(t)
	})

	t.Run("error extract token", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		res, err := srv.GetSellingHistory(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetSellingHistory", userId).Return([]order.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetSellingHistory(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetSellingHistory", userId).Return([]order.Core{}, errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetSellingHistory(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestNotificationTransaction(t *testing.T) {
	repo := mocks.NewOrderData(t)
	// transactionId := "transaction-1"
	// var test *coreapi.TransactionStatusResponse
	t.Run("error check transaction status", func(t *testing.T) {
		srv := New(repo)

		err := srv.NotificationTransactionStatus("xxxxx")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error check")
	})

	// t.Run("error calling data", func(t *testing.T) {
	// 	repo.On("CheckTransaction", "Transaction-1").Return(test, nil)
	// 	repo.On("NotificationTransactionStatus", "xxxx", "failure").Return(errors.New("error update order status"))
	// 	srv := New(repo)

	// 	err := srv.NotificationTransactionStatus("xxxx")
	// 	assert.NotNil(t, err)
	// 	assert.Contains(t, err.Error(), "error calling")
	// })

	// t.Run("error calling data", func(t *testing.T) {
	// 	repo.On("CheckTransaction", "Transaction-1").Return(test, nil)
	// 	repo.On("NotificationTransactionStatus", "xxxx", "failure").Return(nil)
	// 	srv := New(repo)

	// 	err := srv.NotificationTransactionStatus("xxxx")
	// 	assert.Nil(t, err)
	// 	repo.AssertExpectations(t)
	// })
}

func TestUpdateStatus(t *testing.T) {
	repo := mocks.NewOrderData(t)
	orderId := uint(1)

	t.Run("success update status", func(t *testing.T) {
		repo.On("UpdateStatus", orderId, "canceled").Return(nil).Once()

		srv := New(repo)
		err := srv.UpdateStatus(orderId, "canceled")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("UpdateStatus", orderId, "canceled").Return(errors.New("not found")).Once()

		srv := New(repo)
		err := srv.UpdateStatus(orderId, "canceled")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "bad request")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("UpdateStatus", orderId, "canceled").Return(errors.New("server problem")).Once()

		srv := New(repo)
		err := srv.UpdateStatus(orderId, "canceled")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "server problem")
		repo.AssertExpectations(t)
	})
}
