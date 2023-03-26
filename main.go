package main

import (
	"log"
	"net"

	pb "github.com/example/car-rental-service/proto/generated/carrental"
	"github.com/example/car-rental-service/server"
	"github.com/example/car-rental-service/validate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	gserver := grpc.NewServer(grpc.UnaryInterceptor(validate.ValidateInboundRequest))
	// gserver := grpc.NewServer()

	// Register the car rental service with the server
	pb.RegisterCarRentalServiceServer(gserver, server.New())

	// Register grpcui server reflection
	reflection.Register(gserver)

	// Start the server in a goroutine
	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
