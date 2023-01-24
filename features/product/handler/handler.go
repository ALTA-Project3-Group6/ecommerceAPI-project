package handler

import (
	"ecommerceapi/features/product"
	"ecommerceapi/helper"
	"log"
	"mime/multipart"
	"net/http"

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
			log.Println("error running add product service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success add product",
		})
	}
}
func (pc *productControl) Update() echo.HandlerFunc
func (pc *productControl) Delete() echo.HandlerFunc
func (pc *productControl) GetAllProducts() echo.HandlerFunc
func (pc *productControl) GetUserProducts() echo.HandlerFunc
func (pc *productControl) GetProductById() echo.HandlerFunc
