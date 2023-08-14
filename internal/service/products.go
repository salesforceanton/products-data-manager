package service

import "github.com/salesforceanton/products-data-manager/internal/repository"

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}
