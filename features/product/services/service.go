package services

import (
	"ecommerceapi/features/product"
	"ecommerceapi/features/product/data"
	"ecommerceapi/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
)

type productSvc struct {
	qry product.ProductData
}

func New(data product.ProductData) product.ProductService {
	return &productSvc{
		qry: data,
	}
}

func (ps *productSvc) Add(token interface{}, newProduct product.Core, productImage *multipart.FileHeader) (product.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("\t error extract token add product")
		return product.Core{}, errors.New("user not found")
	}

	cnvP := data.CoreToData(newProduct)

	if productImage != nil {
		path, err := helper.UploadProductPhotoS3(*productImage, int(cnvP.ID))
		if err != nil {
			log.Println("\terror upload product photo: ", err.Error())
			return product.Core{}, err
		}
		cnvP.ProductImage = path
	}

	res, err := ps.qry.Add(uint(userId), data.DataToCore(cnvP))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror add query in service: ", err.Error())
		return product.Core{}, errors.New(msg)
	}

	return res, nil
}
func (ps *productSvc) Update(token interface{}, productId uint, updProduct product.Core, productImage *multipart.FileHeader) (product.Core, error)
func (ps *productSvc) Delete(token interface{}, productId uint) error
func (ps *productSvc) GetAllProducts() ([]product.Core, error)
func (ps *productSvc) GetUserProducts(token interface{}) ([]product.Core, error)
func (ps *productSvc) GetProductById(token interface{}, productId uint)
