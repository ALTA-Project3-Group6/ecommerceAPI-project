package main

import (
	"ecommerceapi/config"
	pd "ecommerceapi/features/product/data"
	ph "ecommerceapi/features/product/handler"
	ps "ecommerceapi/features/product/services"
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

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "- method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	//user

	//product
	e.POST("/products", prodHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products", prodHdl.GetAllProducts())
	e.PUT("/products/:id_product", prodHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products/:id_product", prodHdl.GetProductById())
	e.DELETE("/products/:id_product", prodHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
