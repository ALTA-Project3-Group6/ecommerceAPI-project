package main

import (
	"ecommerceapi/config"
	cd "ecommerceapi/features/cart/data"
	ch "ecommerceapi/features/cart/handler"
	cs "ecommerceapi/features/cart/services"
	od "ecommerceapi/features/order/data"
	oh "ecommerceapi/features/order/handler"
	os "ecommerceapi/features/order/services"
	pd "ecommerceapi/features/product/data"
	ph "ecommerceapi/features/product/handler"
	ps "ecommerceapi/features/product/services"
	ud "ecommerceapi/features/user/data"
	uh "ecommerceapi/features/user/handler"
	us "ecommerceapi/features/user/services"
	"ecommerceapi/migration"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	migration.Migrate(db)

	prodData := pd.New(db)
	prodSrv := ps.New(prodData)
	prodHdl := ph.New(prodSrv)

	userData := ud.New(db)
	userSrv := us.New(userData)
	userHdl := uh.New(userSrv)

	orderData := od.New(db)
	orderSrv := os.New(orderData)
	orderHdl := oh.New(orderSrv)

	cartData := cd.New(db)
	cartSrv := cs.New(cartData)
	cartHdl := ch.New(cartSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "- method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	//user
	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	//product
	e.POST("/products", prodHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products", prodHdl.GetAllProducts())
	e.PUT("/products/:id_product", prodHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products/:id_product", prodHdl.GetProductById())
	e.DELETE("/products/:id_product", prodHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	//cart
	e.POST("/carts", cartHdl.AddCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/carts", cartHdl.ShowCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/carts", cartHdl.DeleteCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/carts/:id_carts", cartHdl.UpdateCart(), middleware.JWT([]byte(config.JWT_KEY)))

	//order
	e.POST("/orders", orderHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/orders", orderHdl.GetOrderHistory(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/sales", orderHdl.GetSellingHistory(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/paymentnotification", orderHdl.NotificationTransactionStatus())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
