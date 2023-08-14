package service

import (
	"context"

	"github.com/salesforceanton/products-data-manager/internal/repository"
	product_manager "github.com/salesforceanton/products-data-manager/pkg/domain"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) MergeData(ctx context.Context, data []product_manager.Product) error {
	return s.repo.MergeData(ctx, data)
}
func (s *ProductsService) GetSortedData(ctx context.Context, request *product_manager.ListRequest) ([]product_manager.Product, error) {
	return s.repo.GetSortedData(ctx,
		product_manager.PaginOptions{
			Limit:  int(request.GetPaginOpts().GetLimit()),
			Offset: int(request.GetPaginOpts().GetOffset()),
		},
		product_manager.SortOptions{
			Field: request.GetSortOpts().GetField(),
			Order: product_manager.SortingOrder(request.GetSortOpts().GetOrder().String()),
		},
	)
}
