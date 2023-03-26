package server

import (
	pb "github.com/example/car-rental-service/proto/generated/carrental"
)

type CarRentalServiceServer struct {
	pb.UnimplementedCarRentalServiceServer
}

func New() *CarRentalServiceServer {
	return &CarRentalServiceServer{}
}
