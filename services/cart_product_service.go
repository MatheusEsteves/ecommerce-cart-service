package services

import (
	"context"

	"github.com/ecommerce-cart-service/models"
	"github.com/ecommerce-cart-service/repositories"
)

type CartProductsService interface {
	GetCartProducts(ctx context.Context) ([]models.CartProductData, error)
	SaveCartProduct(ctx context.Context, product *models.CartProductData) error
}

type CartProductsServiceImp struct {
	repository repositories.CartProductsRepository
}

func NewCartProductsService(repository repositories.CartProductsRepository) CartProductsService {
	return &CartProductsServiceImp{
		repository: repository,
	}
}

func (c *CartProductsServiceImp) GetCartProducts(ctx context.Context) ([]models.CartProductData, error) {
	return c.repository.GetCartProducts(ctx)
}

func (c *CartProductsServiceImp) SaveCartProduct(ctx context.Context, product *models.CartProductData) error {
	return c.repository.SaveCartProduct(ctx, product)
}
