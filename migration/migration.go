package migration

import (
	product "ecommerceapi/features/product/data"
	user "ecommerceapi/features/user/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(product.Product{})
}
