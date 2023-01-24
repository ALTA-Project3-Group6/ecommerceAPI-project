package handler

import (
	"ecommerceapi/features/product"
	"ecommerceapi/helper"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type productControl struct {
	srv product.ProductService
}

func New(srv product.ProductService) product.ProductHandler {
	return &productControl{
		srv: srv,
	}
}

func (pc *productControl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddProductReq{}
		var productImage *multipart.FileHeader

		if err := c.Bind(&input); err != nil {
			log.Println("\tbind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		file, err := c.FormFile("product_image")
		if file != nil && err == nil {
			productImage = file
		} else if file != nil && err != nil {
			log.Println("\terror read product image: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		}

		res, err := pc.srv.Add(token, *ToCore(input), productImage)
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
func (pc *productControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		var prodImg *multipart.FileHeader

		productId := c.Param("id_product")
		cProdId, _ := strconv.Atoi(productId)

		input := AddProductReq{}
		err := c.Bind(&input)
		if err != nil {
			log.Println("bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		file, err := c.FormFile("product_image")
		if file != nil && err == nil {
			prodImg = file
		} else if file != nil && err != nil {
			log.Println("\terror read product image: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		}

		res, err := pc.srv.Update(token, uint(cProdId), *ToCore(input), prodImg)
		if err != nil {
			log.Println("\terror running update post service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success update product",
		})
	}
}
func (pc *productControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id")
		cnv, _ := strconv.Atoi(input)

		err := pc.srv.Delete(token, uint(cnv))
		if err != nil {
			msg := ""
			if strings.Contains(err.Error(), "not found") {
				msg = "product not found"
			} else {
				msg = "server problem"
			}
			log.Println("error calling delete product service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(msg))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete product",
		})
	}
}
func (pc *productControl) GetAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := pc.srv.GetAllProducts()
		if err != nil {
			log.Println("error running GetAllProducts service: ", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show all products",
		})
	}
}
func (pc *productControl) GetUserProducts() echo.HandlerFunc
func (pc *productControl) GetProductById() echo.HandlerFunc
