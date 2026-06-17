package main

import (
	pb "26/order"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `bson:"user_id"`
	Product   string             `bson:"product"`
	Quantity  int32              `bson:"quantity"`
	CreatedAt time.Time          `bson:"created_at"`
}
type server struct {
	pb.UnimplementedOrderServiceServer
	db *mongo.Collection
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := Order{
		UserId:    req.UserId,
		Product:   req.Product,
		Quantity:  req.Quantity,
		CreatedAt: time.Now(),
	}
	res, err := s.db.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	Id := res.InsertedID.(primitive.ObjectID).Hex()
	return &pb.CreateOrderResponse{
		OrderId: Id,
	}, nil
}

func main() {
	//mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	collection := client.Database("orders").Collection("orders")
	//gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &server{db: collection})
	log.Println("All good and connected, Listening...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Println(err)
		return
	}
}
