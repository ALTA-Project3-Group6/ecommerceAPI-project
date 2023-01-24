package data

import (
	"ecommerceapi/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Address      string
	PhoneNumber  string
	Email        string
	Password     string
	ProfilePhoto string
	// Product      []Product
	// Orders       []Order
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:           data.ID,
		Email:        data.Email,
		Name:         data.Name,
		Address:      data.Address,
		PhoneNumber:  data.PhoneNumber,
		Password:     data.Password,
		ProfilePhoto: data.ProfilePhoto,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:        gorm.Model{ID: data.ID},
		Email:        data.Email,
		Name:         data.Name,
		Address:      data.Address,
		PhoneNumber:  data.PhoneNumber,
		Password:     data.Password,
		ProfilePhoto: data.ProfilePhoto,
	}
}
