package main

import (
	"apiServerTls/src/pb/product"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	product.ProductServiceServer
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("./src/cert/server-cert.pem", "./src/cert/server-key.pem")
	if err != nil {
		return nil, fmt.Errorf("error on load server certification. error: %v\n", err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}

func (s *Server) FindAll(ctx context.Context, req *product.ListProductListRequest) (*product.ProductListResponse, error) {
	var productList []*product.Product
	productList = append(productList, &product.Product{
		Id:    1,
		Title: "Mac book M2",
	})
	return &product.ProductListResponse{Products: productList}, nil
}

func main() {
	fmt.Println("starting gRPC server, port: 9090")
	listenner, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("error on listenn. erro:", err)
	}
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	product.RegisterProductServiceServer(srv, &Server{})
	if err := srv.Serve(listenner); err != nil {
		log.Fatal("error on server. error:", err)
	}
}
