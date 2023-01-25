package handler

import (
	"ecommerceapi/features/cart"
	"ecommerceapi/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		token := c.Get("user")
		input := AddCartReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("\tbind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		res, err := cc.srv.AddCart(token, *ToCore(input))
		if err != nil {
			log.Println("\terror running add product service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success add product",
		})
	}
}
func (cc *cartControl) ShowCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := cc.srv.ShowCart(token)
		if err != nil {
			log.Println("error running GetAllProducts service: ", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    CoreToCartsResp(res),
			"message": "success show all products",
		})
	}
}
// func (cc *cartControl) UpdateCart() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }
func (cc *cartControl) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id_product")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			log.Println("\tRead param error: ", err.Error())
			return c.JSON(http.StatusBadRequest, "wrong product id parameter")
		}

		err = cc.srv.DeleteCart(token, uint(cnv))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("error calling delete product service: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("product not found"))
			} else {
				log.Println("error calling delete product service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete product",
		})
	}
}
