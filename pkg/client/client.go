package grpc_client

import (
	"context"
	"fmt"

	products_manager "github.com/salesforceanton/products-data-manager/pkg/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	DefaultSortOptions  = &products_manager.SortingOptions{Field: "name", Order: products_manager.SortingOptions_ASC}
	DefaultPaginOptions = &products_manager.PaginationOptions{Limit: 10, Offset: 0}
)

type ProductsManagerClient struct {
	conn   *grpc.ClientConn
	client products_manager.ProductManagerServiceClient
}

func NewClient(port int) (*ProductsManagerClient, error) {
	// Create connection trought grpc
	var conn *grpc.ClientConn

	addr := fmt.Sprintf(":%d", port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Construct client instance defined in generated files from .proto
	return &ProductsManagerClient{
		conn:   conn,
		client: products_manager.NewProductManagerServiceClient(conn),
	}, nil
}

func (c *ProductsManagerClient) CloseConnection() error {
	return c.conn.Close()
}

func (c *ProductsManagerClient) List(paginOpts *products_manager.PaginOptions, sortOpts *products_manager.SortOptions) (*products_manager.ListResponse, error) {
	var request products_manager.ListRequest
	if paginOpts == nil {
		request.PaginOpts = DefaultPaginOptions
	} else {
		request.PaginOpts = &products_manager.PaginationOptions{
			Limit:  int64(paginOpts.Limit),
			Offset: int64(paginOpts.Offset),
		}
	}
	if sortOpts == nil {
		request.SortOpts = DefaultSortOptions
	} else {
		sortOrder, _ := products_manager.ToPbSortOrder(sortOpts.Order)
		request.SortOpts = &products_manager.SortingOptions{
			Field: sortOpts.Field,
			Order: sortOrder,
		}
	}

	return c.client.List(context.TODO(), &request)
}

func (c *ProductsManagerClient) Fetch(url string) error {
	_, err := c.client.Fetch(context.TODO(), &products_manager.FetchRequest{Url: url})
	return err
}
