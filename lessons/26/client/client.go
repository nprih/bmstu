package main

import (
	pb "26/order"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req := &pb.CreateOrderRequest{
		UserId:   "user123",
		Product:  "Coffee",
		Quantity: 2,
	}
	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Order created, ur number: %s", res.OrderId)
}
