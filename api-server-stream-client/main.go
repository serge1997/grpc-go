package main

import (
	"apiServerStreamClient/src/pb/department"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("erro on connection to department service. error: %v\n", err)
	}
	client := department.NewDepartmentServiceClient(conn)
	stream, err := client.ListDepartmentMember(context.Background(), &department.DeparmentMembersRequest{DepartmentId: 4})
	if err != nil {
		log.Fatalf("erro on list department members. error: %v\n", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error on receive department mambers. error: %v\n", err)
		}
		fmt.Printf("Members: %+v\n", response)
	}
}
