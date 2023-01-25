package handler

import (
	"ecommerceapi/features/cart"

	"github.com/labstack/echo/v4"
)

type cartControl struct {
	srv cart.CartService
}

func New(srv cart.CartService) cart.CartHandler {
	return &cartControl{
		srv: srv,
	}
}

func (cc *cartControl) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		
	}
}
func (cc *cartControl) ShowCart() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
func (cc *cartControl) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
func (cc *cartControl) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
