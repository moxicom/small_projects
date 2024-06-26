package mongo_storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_lrn/internal/config"
	"mongo_lrn/internal/models"
	"mongo_lrn/internal/storage/mongo_storage/category_rep"
	"mongo_lrn/internal/storage/mongo_storage/product_rep"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const DBName = "db_test"

type Category interface {
	NewCategory(context.Context, models.Category) (insertedID primitive.ObjectID, err error)
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id interface{}) (models.Category, error)
	DeleteCategories(ctx context.Context) error
	DeleteCategoryByID(ctx context.Context, id primitive.ObjectID) error
}

type Product interface {
	// Product interface
	// TODO: implement
}

type Repository struct {
	Category
	Product
}

func NewRepository(c *mongo.Client) *Repository {
	db := c.Database(DBName)
	return &Repository{
		Category: category_rep.NewCategoryRep(db.Collection(CategoriesCol)),
		Product:  product_rep.NewProductRep(db.Collection(ProductsCol)),
	}
}

func Connect(cfg config.Config) (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, client, cancel := mustSetupMongoDB(cfg)
	return client, ctx, cancel
	//if err := migrateCollections(ctx, client.Database(DBName)); err != nil {
	//	panic(err)
	//}
}

func mustSetupMongoDB(cfg config.Config) (context.Context, *mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return ctx, client, cancel
}
