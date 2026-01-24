package main

import (
	"apiProducts/src/pb/products"
	"apiProducts/src/repositories"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	products.ProductServiceServer
	productcRepo repositories.ProductRepository
}

func (s *Server) Create(ctx context.Context, req *products.Product) (*products.Product, error) {
	err := s.productcRepo.Create(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func (s *Server) FindAll(ctx context.Context, req *products.Product) (*products.ProductList, error) {
	productsList, err := s.productcRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return productsList, nil
}
func main() {
	fmt.Println("Starting gRPC server on port 9090...")
	srv := Server{productcRepo: repositories.ProductRepository{}}

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	products.RegisterProductServiceServer(s, &srv)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
