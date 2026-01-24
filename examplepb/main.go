package main

import (
	"examplebp/src/pb/users"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := users.User{
		Id:       1,
		Name:     "John Doe",
		Email:    "serge@email.com",
		Password: "password123",
	}
	out, err := proto.Marshal(&user)
	if err != nil {
		fmt.Println("Error marshaling user:", err)
		return
	}
	createData(out)
	readData()

}

func createData(user []byte) {
	err := os.WriteFile("user.txt", user, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v\n", err)
	}
}

func readData() {
	var user users.User
	data, err := os.ReadFile("user.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	err = proto.Unmarshal(data, &user)
	if err != nil {
		log.Fatalf("Error unmarshaling data: %v\n", err)
	}
	fmt.Printf("User: %+v\n", user)
}
