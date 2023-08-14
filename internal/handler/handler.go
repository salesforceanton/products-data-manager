package handler

import (
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/jszwec/csvutil"
	"github.com/salesforceanton/products-data-manager/internal/logger"
	"github.com/salesforceanton/products-data-manager/internal/service"
	products_manager "github.com/salesforceanton/products-data-manager/pkg/domain"
)

type Handler struct {
	service *service.Service
	products_manager.UnimplementedProductManagerServiceServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Fetch(ctx context.Context, request *products_manager.FetchRequest) (*products_manager.Empty, error) {
	// Get data from side and convert to Product structs
	url := request.GetUrl()
	convertedData, err := h.getDataFromSide(url)
	if err != nil {
		logger.LogHandlerIssue("fetch-convertion", err)
		return &products_manager.Empty{}, err
	}

	// Merge data from request to data in db
	err = h.service.Products.MergeData(ctx, convertedData)
	logger.LogHandlerIssue("fetch-merge", err)
	return &products_manager.Empty{}, err
}
func (h *Handler) List(ctx context.Context, request *products_manager.ListRequest) (*products_manager.ListResponse, error) {
	data, err := h.service.Products.GetSortedData(ctx, request)
	if err != nil {
		logger.LogHandlerIssue("list", err)
		return nil, err
	}
	var responseData []*products_manager.ProductItem
	for _, p := range data {
		responseData = append(responseData, &products_manager.ProductItem{
			Name:  p.Name,
			Price: float32(p.Price),
		})
	}
	return &products_manager.ListResponse{Data: responseData}, nil

}
func (h *Handler) getDataFromSide(url string) ([]products_manager.Product, error) {
	// Make Http callout to getting actual data in csv
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse csv data into slice of Product struct
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	var convertedData []products_manager.Product
	if err := csvutil.Unmarshal(dump, &convertedData); err != nil {
		return nil, err
	}

	return convertedData, nil
}
