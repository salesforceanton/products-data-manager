package repository

import (
	"context"

	model "github.com/salesforceanton/products-data-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Products
}

type Products interface {
	MergeData(ctx context.Context, data []model.Product) ([]model.Product, error)
	GetSortedData(ctx context.Context, paginOpts model.PaginOptions, sortOpts model.SortOptions) ([]model.Product, error)
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Products: NewProductsRepository(db),
	}
}
