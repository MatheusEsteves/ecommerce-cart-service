package repositories

import (
	"context"

	"github.com/ecommerce-cart-service/clients"
	"github.com/ecommerce-cart-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

const CartProductsCollection = "products"

type CartProductsRepository interface {
	GetCartProducts(ctx context.Context) ([]models.CartProductData, error)
	SaveCartProduct(ctx context.Context, product *models.CartProductData) error
}

type CartProductsRepositoryImp struct {
	mongoClient clients.MongoClient
}

func NewCartProductsRepository(mongoClient clients.MongoClient) CartProductsRepository {
	return &CartProductsRepositoryImp{
		mongoClient: mongoClient,
	}
}

func (c *CartProductsRepositoryImp) GetCartProducts(ctx context.Context) ([]models.CartProductData, error) {
	result, err := c.mongoClient.Find(ctx, CartProductsCollection, bson.M{})

	if err != nil {
		return []models.CartProductData{}, err
	}

	products := new([]models.CartProductData)
	if err := result.All(ctx, products); err != nil {
		return []models.CartProductData{}, err
	}

	return *products, nil
}

func (c *CartProductsRepositoryImp) SaveCartProduct(ctx context.Context, product *models.CartProductData) error {
	if _, err := c.mongoClient.InsertOne(ctx, CartProductsCollection, product); err != nil {
		return err
	}
	return nil
}
