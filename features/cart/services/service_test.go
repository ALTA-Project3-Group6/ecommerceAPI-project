package services

import (
	"ecommerceapi/features/cart"
	"ecommerceapi/helper"
	"ecommerceapi/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddCart(t *testing.T) {
	repo := mocks.NewCartData(t)

	inputData := cart.Core{
		ProductID: 1,
		Quantity:  5,
	}

	resData := cart.Core{
		ID:        1,
		ProductID: 1,
		Quantity:  5,
	}

	t.Run("success add product to cart", func(t *testing.T) {

		repo.On("AddCart", uint(1), uint(1), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.AddCart(pToken, uint(1), inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Cart not found", func(t *testing.T) {
		repo.On("AddCart", uint(5), uint(1), inputData).Return(cart.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(5)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddCart(useToken, 1, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server error", func(t *testing.T) {
		repo.On("AddCart", uint(1), uint(1), inputData).Return(cart.Core{}, errors.New("server error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddCart(useToken, 1, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("AddCart", uint(1), uint(1), inputData).Return(cart.Core{}, errors.New("data not found")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AddCart(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})
}

func TestShowCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	resData := []cart.Core{
		{ID: 1,
			ProductName: "Surya 16",
			SellerName:  "fauzan",
			Quantity:    1,
			Price:       20000,
		},
	}

	t.Run("Success show cart", func(t *testing.T) {
		repo.On("ShowCart", uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ShowCart(pToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("ShowCart", uint(1)).Return([]cart.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ShowCart(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("server error", func(t *testing.T) {
		repo.On("ShowCart", uint(1)).Return([]cart.Core{}, errors.New("server error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ShowCart(pToken)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server problem")
		repo.AssertExpectations(t)
	})
}

func TestUpdateCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	resData := cart.Core{Quantity: 2}

	t.Run("Success update data", func(t *testing.T) {
		repo.On("UpdateCart", uint(1), uint(1), cart.Core{}).Return(resData, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.UpdateCart(useToken, 1, cart.Core{})
		assert.Nil(t, err)
		assert.Equal(t, resData.Quantity, res.Quantity)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("UpdateCart", uint(1), uint(5), cart.Core{}).Return(cart.Core{}, errors.New("repo not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.UpdateCart(useToken, 5, cart.Core{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, res.Quantity)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("Tourble in server", func(t *testing.T) {
		repo.On("UpdateCart", uint(1), uint(1), cart.Core{}).Return(cart.Core{}, errors.New("server error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.UpdateCart(useToken, 1, cart.Core{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, res.Quantity)
		assert.ErrorContains(t, err, "server problem")
		repo.AssertExpectations(t)
	})
}

func TestDeleteCart(t *testing.T) {
	repo := mocks.NewCartData(t)

	t.Run("Success delete data", func(t *testing.T) {
		repo.On("DeleteCart", uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.DeleteCart(useToken)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("DeleteCart", uint(5)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(5)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.DeleteCart(useToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("invalid JWT", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.DeleteCart(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
	})

}
