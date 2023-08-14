package service

import (
	"context"

	"github.com/salesforceanton/products-data-manager/internal/repository"
	product_manager "github.com/salesforceanton/products-data-manager/pkg/domain"
)

type Service struct {
	Products
}

type Products interface {
	MergeData(ctx context.Context, data []product_manager.Product) error
	GetSortedData(ctx context.Context, request *product_manager.ListRequest) ([]product_manager.Product, error)
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Products: NewProductsService(repos.Products),
	}
}
