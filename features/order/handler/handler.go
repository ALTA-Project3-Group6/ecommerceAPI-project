package handler

import (
	"ecommerceapi/features/order"
	"ecommerceapi/helper"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		input := OrderReq{}
		if err := c.Bind(&input); err != nil {
			log.Println("error bind totalrpice: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (bad request)"))
		}

		token := c.Get("user")
		if token == nil {
			log.Println("error get token JWT")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (bad request)"))
		}

		res, url, err := oh.srv.Add(token, input.TotalPrice)
		if err != nil {
			if strings.Contains(err.Error(), "bad request") || strings.Contains(err.Error(), "not found") {
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

		res, err := oh.srv.GetSellingHistory(token)
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
func (oh *OrderHandle) NotificationTransactionStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 1. Initialize empty map
		var notificationPayload map[string]interface{}

		// 2. Parse JSON request body and use it to set json to payload
		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)
		if err != nil {
			// do something on error when decode
			return c.JSON(http.StatusBadRequest, err)
		}

		// 3. Get order-id from payload
		transactionId, exists := notificationPayload["order_id"].(string)
		if !exists {
			// do something when key `order_id` not found
			return c.JSON(http.StatusBadRequest, err)
		}

		err = oh.srv.NotificationTransactionStatus(transactionId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(c.Response().Write([]byte("ok")))
	}
}
func (oh *OrderHandle) UpdateStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		status := StatusReq{}
		err := c.Bind(&status)
		if err != nil {
			log.Println("bind order status error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		oid := c.Param("order_id")
		orderId, err := strconv.Atoi(oid)
		if err != nil {
			log.Println("error read parameter: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("fail to read parameter"))
		}

		err = oh.srv.UpdateStatus(uint(orderId), status.OrderStatus)
		if err != nil {
			if strings.Contains(err.Error(), "bad request") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input (bad request)"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete order",
		})
	}
}
