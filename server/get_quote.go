package server

import (
	"context"
	"strings"

	pb "github.com/example/car-rental-service/proto/generated/carrental"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CarRentalServiceServer) GetQuote(ctx context.Context,
	req *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	if req.GetGarageId() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing garage id")
	}

	if !strings.HasPrefix(req.GetGarageId(), "gar_") {
		return nil, status.Error(codes.InvalidArgument, "invalid garage id")
	}

	if req.GetCarId() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing car id")
	}

	if !strings.HasPrefix(req.GetCarId(), "car_") {
		return nil, status.Error(codes.InvalidArgument, "invalid car id")
	}

	return &pb.GetQuoteResponse{}, nil
}
