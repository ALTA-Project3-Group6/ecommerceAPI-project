package migration

import (
	order "ecommerceapi/features/order/data"
	product "ecommerceapi/features/product/data"
	user "ecommerceapi/features/user/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(order.Order{})
}
