package data

import (
	"ecommerceapi/features/product"
	"errors"
	"log"

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

func (pq *productQuery) Add(userId uint, newProduct product.Core) (product.Core, error) {
	cnvP := CoreToData(newProduct)
	cnvP.UserId = userId

	err := pq.db.Create(&cnvP).Error
	if err != nil {
		log.Println("\tadd product query error: ", err.Error())
		return product.Core{}, errors.New("server problem")
	}

	newProduct.ID = cnvP.ID
	newProduct.UserId = cnvP.UserId

	return newProduct, nil
}
func (pq *productQuery) Update(userId, productId uint, updProduct product.Core) (product.Core, error) {
	return product.Core{}, nil
}
func (pq *productQuery) Delete(userId, productId uint) error {
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
