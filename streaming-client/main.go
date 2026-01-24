package main

import (
	"context"
	"fmt"
	"log"
	"streaming_client/src/pb/calc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":9090", grpc.
		WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("erro on get connection. erro: ", err)
	}
	defer conn.Close()
	calClient := calc.NewCalServiceClient(conn)
	stream, err := calClient.Calc(context.Background())
	if err != nil {
		log.Fatalln("erro on instantiate stream connection. erro: ", err)
	}
	nums := []int32{3, 4, 7, 2, 1}
	for i, v := range nums {
		i += 1
		err := stream.Send(&calc.Input{
			Value: v,
			Index: int32(i),
		})
		if err != nil {
			log.Fatalln("erro to data to stream. erro:", err)
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("err to close stream. erro: ", err)
	}
	fmt.Printf("stream response: %+v\n", response)
}
