package main

import (
	"log"

	"github.com/ecommerce-cart-service/clients"
	"github.com/ecommerce-cart-service/handlers"
	"github.com/ecommerce-cart-service/repositories"
	"github.com/ecommerce-cart-service/services"
	"github.com/labstack/echo/v4"
)

func main() {
	echoInstance := echo.New()

	mongoClient, err := clients.NewMongoClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	cartProductsRepository := repositories.NewCartProductsRepository(mongoClient)
	cartProductsService := services.NewCartProductsService(cartProductsRepository)
	cartProductsHandler := handlers.NewCartProductsHandler(cartProductsService)

	cartProductsRoute := echoInstance.Group("/cart-products")

	cartProductsRoute.GET("", cartProductsHandler.GetCartProducts)
	cartProductsRoute.POST("", cartProductsHandler.SaveCartProduct)

	echoInstance.Start(":8081")
}
