package data

import (
	"ecommerceapi/features/product"
	"ecommerceapi/helper"
	"errors"
	"log"
	"mime/multipart"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productQuery{
		db: db,
	}
}

func (pq *productQuery) Add(userId uint, newProduct product.Core, productImage *multipart.FileHeader) (product.Core, error) {
	cnvP := CoreToData(newProduct)
	cnvP.UserId = userId

	if productImage != nil {
		path, err := helper.UploadProductPhotoS3(*productImage, int(cnvP.ID))
		if err != nil {
			log.Println("\terror upload product photo: ", err.Error())
			return product.Core{}, err
		}
		cnvP.ProductImage = path
	}

	err := pq.db.Create(&cnvP).Error
	if err != nil {
		log.Println("\tadd product query error: ", err.Error())
		return product.Core{}, errors.New("server problem")
	}

	newProduct.ID = cnvP.ID
	newProduct.UserId = cnvP.UserId

	return newProduct, nil
}
func (pq *productQuery) Update(userId, productId uint, updProduct product.Core, productImage *multipart.FileHeader) (product.Core, error) {
	cnvP := CoreToData(updProduct)
	cnvP.UserId = userId

	if productImage != nil {
		path, err := helper.UploadProductPhotoS3(*productImage, int(cnvP.ID))
		if err != nil {
			log.Println("\terror upload product photo: ", err.Error())
			return product.Core{}, err
		}
		cnvP.ProductImage = path
	}

	qry := pq.db.Where("id = ? AND user_id = ?", productId, userId).Updates(&cnvP)
	if qry.RowsAffected <= 0 {
		log.Println("\tupdate product query error: data not found")
		return product.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tupdate peoduct query error: ", err.Error())
		return product.Core{}, errors.New("not found")
	}

	return DataToCore(cnvP), nil
}
func (pq *productQuery) Delete(userId, productId uint) error {
	qry := pq.db.Where("user_id = ?", userId).Delete(&Product{}, productId)

	if aff := qry.RowsAffected; aff <= 0 {
		log.Println("\tno rows affected: data not found")
		return errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tdelete query error: ", err.Error())
		return err
	}

	return nil
}
func (pq *productQuery) GetAllProducts() ([]product.Core, error) {
	return []product.Core{}, nil
}
func (pq *productQuery) GetUserProducts(userId uint) ([]product.Core, error) {
	return []product.Core{}, nil
}
func (pq *productQuery) GetProductById(userId, productId uint) (product.Core, error) {
	return product.Core{}, nil
}
