package server

import (
	"fmt"
	"net"

	products_manager "github.com/salesforceanton/products-data-manager/pkg/domain"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv *grpc.Server
	handler products_manager.ProductManagerServiceServer
}

func New(handler products_manager.ProductManagerServiceServer) *Server {
	return &Server{
		grpcSrv: grpc.NewServer(),
		handler: handler,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	// Initilialize listener on target port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Start cycle as a http.ListenAndServe
	products_manager.RegisterProductManagerServiceServer(s.grpcSrv, s.handler)
	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
