package service

import (
	"ecommerceapi/features/user"
	"ecommerceapi/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
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

func TestLogn(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("succcess login", func(t *testing.T) {
		email := "mfauzanptra@gmail.com"
		password := "paupau99"
		hashed, _ := bcrypt.GenerateFromPassword([]byte("paupau99"), bcrypt.DefaultCost)

		expectedData := user.Core{
			Email:       "mfauzanptra@gmail.com",
			Name:        "Muhamad Fauzan Putra",
			PhoneNumber: "085659171799",
			Password:    string(hashed),
			Address:     "Jln. Lembayung No 24, Bantul, Yogyakarta",
		}
		data.On("Login", email).Return(expectedData, nil)
		srv := New(data)
		token, res, err := srv.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.Email, res.Email)
		assert.NotNil(t, token)
		data.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		inputEmail := "mfp@gmail.com"
		data.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found"))

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1422")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "mfauzanptra@gmail.com"
		hashed, _ := bcrypt.GenerateFromPassword([]byte("paupau99"), bcrypt.DefaultCost)
		expData := user.Core{
			Email:       "mfauzanptra@gmail.com",
			Name:        "Muhamad Fauzan Putra",
			PhoneNumber: "085659171799",
			Password:    string(hashed),
			Address:     "Jln. Lembayung No 24, Bantul, Yogyakarta",
		}
		data.On("Login", inputEmail).Return(expData, nil)

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1423")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong password")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
