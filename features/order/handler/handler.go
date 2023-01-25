package handler

import (
	"ecommerceapi/features/order"
	"ecommerceapi/helper"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHandle struct {
	srv order.OrderService
}

func New(os order.OrderService) order.OrderHandler {
	return &OrderHandle{
		srv: os,
	}
}

func (oh *OrderHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var totalPrice float64
		if err := c.Bind(&totalPrice); err != nil {
			log.Println("error bind totalrpice: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (bad request)"))
		}

		token := c.Get("user")
		if token == nil {
			log.Println("error get token JWT")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (bad request)"))
		}

		res, url, err := oh.srv.Add(token, totalPrice)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (bad request)"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    CoreToOrderResp(res, url),
			"message": "order payment created",
		})
	}
}
func (oh *OrderHandle) GetOrderHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := oh.srv.GetOrderHistory(token)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (data not found)"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show order history",
		})
	}
}
func (oh *OrderHandle) GetSellingHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := oh.srv.GetOrderHistory(token)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (data not found)"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show order history",
		})
	}
}
func (oh *OrderHandle) GetTransactionStatus() echo.HandlerFunc {
	return nil
}
