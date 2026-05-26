package main

import (
	"20/user"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userResp, err := client.GetUserById(ctx, &user.UserRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("I got user info: %s %s\n", userResp.Name, userResp.Surname)
}
