package services

import (
	"ecommerceapi/features/product"
	// "ecommerceapi/features/product/data"
	"ecommerceapi/helper"
	"ecommerceapi/mocks"
	"mime/multipart"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewProductData(t)

	inputData := product.Core{
		Name:        "Indomie goreng",
		Description: "mie instan terenak didunia",
		Stock:       20,
		// ProductImage: "Indomie.jpg",
	}
	resData := product.Core{
		ID:          1,
		Name:        "Indomie goreng",
		Description: "mie instan terenak didunia",
		Stock:       20,
		// ProductImage: "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/indomie-photo.jpeg",
	}
	var a *multipart.FileHeader

	t.Run("success add post", func(t *testing.T) {

		repo.On("Add", uint(1), inputData, mock.Anything).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, a)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})
}
