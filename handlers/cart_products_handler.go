package handlers

import (
	"net/http"

	"github.com/ecommerce-cart-service/models"
	"github.com/ecommerce-cart-service/services"
	"github.com/labstack/echo/v4"
)

type CartProductsHandler interface {
	GetCartProducts(echoContext echo.Context) error
	SaveCartProduct(echoContext echo.Context) error
}

type CartProductsHandlerImp struct {
	service services.CartProductsService
}

func NewCartProductsHandler(service services.CartProductsService) CartProductsHandler {
	return &CartProductsHandlerImp{
		service: service,
	}
}

func (c *CartProductsHandlerImp) GetCartProducts(echoContext echo.Context) error {
	products, err := c.service.GetCartProducts(echoContext.Request().Context())

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}
	return echoContext.JSON(http.StatusOK, products)
}

func (c *CartProductsHandlerImp) SaveCartProduct(echoContext echo.Context) error {
	product := models.CartProductData{}
	if err := echoContext.Bind(&product); err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.SaveCartProduct(echoContext.Request().Context(), &product); err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}

	return echoContext.JSON(http.StatusCreated, "Product saved sucessfully")
}
