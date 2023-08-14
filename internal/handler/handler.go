package handler

import (
	"context"
	"encoding/csv"
	"net/http"
	"strconv"

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
	if err != nil {
		logger.LogHandlerIssue("fetch-merge", err)
		return &products_manager.Empty{}, err
	}
	return &products_manager.Empty{}, nil
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

	// Configure reader
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Populate structures from csv data
	var convertedData []products_manager.Product
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		price, _ := strconv.ParseFloat(row[1], 64)
		convertedData = append(convertedData, products_manager.Product{Name: row[0], Price: price})
	}
	return convertedData, nil
}
