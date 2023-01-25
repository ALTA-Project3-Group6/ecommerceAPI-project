package migration

import (
	cart "ecommerceapi/features/cart/data"
	order "ecommerceapi/features/order/data"
	orderProduct "ecommerceapi/features/orderproduct/data"
	product "ecommerceapi/features/product/data"
	user "ecommerceapi/features/user/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(order.Order{})
	db.AutoMigrate(cart.Cart{})
	db.AutoMigrate(orderProduct.OrderProduct{})
}
