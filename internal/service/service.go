package service

import (
	"github.com/salesforceanton/products-data-manager/internal/repository"
)

type Service struct {
	Products
}

type Products interface {
	// MergeData(ctx context.Context, data []model.Product) ([]model.Product, error)
	// GetSortedData(ctx context.Context, paginOpts model.PaginOptions, sortOpts model.SortOptions) ([]model.Product, error)
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Products: NewProductsService(repos.Products),
	}
}
