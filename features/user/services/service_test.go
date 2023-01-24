package service

import (
	"ecommerceapi/features/user"
	"ecommerceapi/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)

	newUser := user.Core{
		Email:       "mfauzanptra@gmail.com",
		Name:        "Muhamad Fauzan Putra",
		PhoneNumber: "085659171799",
		Password:    "paupau99",
		Address:     "Jln. Lembayung No 24, Bantul, Yogyakarta",
	}
	expectedData := user.Core{
		Email:       "mfauzanptra@gmail.com",
		Name:        "Muhamad Fauzan Putra",
		PhoneNumber: "085659171799",
		Password:    "paupau99",
		Address:     "Jln. Lembayung No 24, Bantul, Yogyakarta",
	}

	t.Run("success register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(expectedData, nil).Once()
		srv := New(data)
		res, err := srv.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		assert.Equal(t, expectedData.Name, res.Name)
		data.AssertExpectations(t)
	})

	t.Run("duplicate", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()
		srv := New(data)
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user already exist")
		assert.Equal(t, res.Name, "")
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(data)
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Name, "")
	})
}
