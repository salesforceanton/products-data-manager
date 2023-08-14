package repository

import (
	"context"

	model "github.com/salesforceanton/products-data-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PRODUCTS_COLLECTION_NAME = "products"

type ProductsRepository struct {
	db *mongo.Database
}

func NewProductsRepository(db *mongo.Database) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) MergeData(ctx context.Context, data []model.Product) error {
	// Upsert records
	var updetex []mongo.WriteModel
	for _, product := range data {
		var updetexItem mongo.UpdateOneModel
		updetexItem.SetFilter(bson.D{primitive.E{Key: "name", Value: product.Name}})
		updetexItem.SetUpdate(bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "price", Value: product.Price}}}})
		updetexItem.SetUpsert(true)

		updetex = append(updetex, &updetexItem)
	}

	opts := options.BulkWrite()
	_, err := r.db.Collection(PRODUCTS_COLLECTION_NAME).BulkWrite(context.TODO(), updetex, opts)
	if err != nil {
		return err
	}

	return nil
}
func (r *ProductsRepository) GetSortedData(ctx context.Context, paginOpts model.PaginOptions, sortOpts model.SortOptions) ([]model.Product, error) {
	// Setup aggregation options from request
	sortStage := bson.D{
		{Key: "$sort", Value: bson.D{{Key: sortOpts.Field, Value: sortOpts.Order}}},
	}

	paginStage := bson.D{
		{Key: "$skip", Value: paginOpts.Offset},
		{Key: "$limit", Value: paginOpts.Limit},
	}

	// Retain sorted results from db and parse
	cursor, err := r.db.Collection(PRODUCTS_COLLECTION_NAME).Aggregate(ctx, mongo.Pipeline{sortStage, paginStage})
	if err != nil {
		return nil, err
	}

	var result []model.Product
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
