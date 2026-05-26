package main

import (
	"20/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type userServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUserById(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	fmt.Println("I got a request!")
	return &user.UserResponse{
		Id:      request.Id,
		Name:    "Vasya",
		Surname: "Ivanov",
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userServiceServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
