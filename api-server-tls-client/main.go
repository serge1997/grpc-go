package main

import (
	"apiServerTlsClient/src/pb/product"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadCertificate() (credentials.TransportCredentials, error) {
	serverCA, err := os.ReadFile("./src/cert/ca-cert.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read server CA certificate: %w", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCA) {
		return nil, fmt.Errorf("failed to append server CA certificate to cert pool")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func main() {
	tlsCert, err := loadCertificate()
	if err != nil {
		log.Fatalf("error while loading TLS certificate: %v", err)
	}
	client, err := grpc.NewClient("0.0.0.0:9090", grpc.WithTransportCredentials(tlsCert))
	if err != nil {
		log.Fatalf("error while creating gRPC client: %v", err)
	}
	defer client.Close()
	productclient := product.NewProductServiceClient(client)
	prodcuctList, err := productclient.FindAll(context.Background(), &product.ListProductListRequest{})
	if err != nil {
		log.Fatalf("error while calling FindAll: %v", err)
	}
	fmt.Printf("%+v\n", prodcuctList)
}
