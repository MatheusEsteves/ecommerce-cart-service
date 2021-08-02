package clients

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var onceDatabase sync.Once
var client *mongo.Client

const DatabaseName = "ecommerce-cart-products"

type MongoFindResult interface {
	Err() error
	All(c context.Context, v interface{}) error
}

type MongoClient interface {
	Find(ctx context.Context, collection string, filter interface{}) (MongoFindResult, error)
	InsertOne(ctx context.Context, collection string, document interface{}) (string, error)
}

type MongoClientImp struct {
	client *mongo.Client
}

func NewMongoClient() (MongoClient, error) {
	onceDatabase.Do(func() {
		credentials := &options.Credential{
			Username:   "ecommerce_cart_user",
			Password:   "ecommerce_cart_pass",
			AuthSource: DatabaseName,
		}
		options := &options.ClientOptions{
			Auth: credentials,
		}

		var err error
		if client, err = mongo.NewClient(options); err == nil && client != nil {
			ctx := context.Background()
			if err := client.Connect(ctx); err != nil {
				log.Fatal(err)
			} else {
				if err = client.Ping(ctx, readpref.Primary()); err != nil {
					log.Fatal(err)
				}
			}
		}
	})

	return &MongoClientImp{
		client: client,
	}, nil
}

func (c *MongoClientImp) Find(ctx context.Context, collection string, filter interface{}) (MongoFindResult, error) {
	mongoCollection := c.client.Database(DatabaseName).Collection(collection)
	cursor, err := mongoCollection.Find(ctx, filter)
	return cursor, err
}

func (c *MongoClientImp) InsertOne(ctx context.Context, collection string, document interface{}) (string, error) {
	mongoCollection := c.client.Database(DatabaseName).Collection(collection)
	result, err := mongoCollection.InsertOne(ctx, document)

	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.String(), nil
	} else {
		return "", nil
	}
}
