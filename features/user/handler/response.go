package handler

import "ecommerceapi/features/user"

type UserResponse struct {
	ID           uint   `json:"id" form:"id"`
	Email        string `json:"email" form:"email"`
	Name         string `json:"name" form:"name"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"`
	ProfilePhoto string `json:"profile_photo" form:"profile_photo"`
	Address      string `json:"address" form:"address"`
}

type LoginResp struct {
	ID           uint   `json:"id" form:"id"`
	Email        string `json:"email" form:"email"`
	Name         string `json:"name" form:"name"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"`
	ProfilePhoto string `json:"profile_photo" form:"profile_photo"`
	Address      string `json:"address" form:"address"`
	Token        string `json:"token" form:"token"`
}

func ToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:           data.ID,
		Email:        data.Email,
		Name:         data.Name,
		PhoneNumber:  data.PhoneNumber,
		ProfilePhoto: data.ProfilePhoto,
		Address:      data.Address,
	}
}

func ToLoginResp(data user.Core, token string) LoginResp {
	return LoginResp{
		ID:           data.ID,
		Email:        data.Email,
		Name:         data.Name,
		PhoneNumber:  data.PhoneNumber,
		ProfilePhoto: data.ProfilePhoto,
		Address:      data.Address,
		Token:        token,
	}
}
