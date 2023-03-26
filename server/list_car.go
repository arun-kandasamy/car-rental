package server

import (
	"context"
	"strings"

	pb "github.com/example/car-rental-service/proto/generated/carrental"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	GARAGE_ID_LEN = 10
)

// START OMIT
func (s *CarRentalServiceServer) ListCarAvailability(ctx context.Context,
	req *pb.ListCarAvailabilityRequest) (*pb.ListCarAvailabilityResponse, error) {

	if req.GetGarageId() == "" { // HL_MANDATORY
		return nil, status.Error(codes.InvalidArgument, "missing garage id") // HL_MANDATORY
	} // HL_MANDATORY

	if !strings.HasPrefix(req.GetGarageId(), "gar_") { // HL_FORMAT
		return nil, status.Error(codes.InvalidArgument, "invalid garage id format") // HL_FORMAT
	} // HL_FORMAT

	if len(req.GetGarageId()) != GARAGE_ID_LEN { // HL_LEN
		return nil, status.Error(codes.InvalidArgument, "incorrect garage id length") // HL_LEN
	} // HL_LEN

	// BUSINESS LOGIC HERE!

	return &pb.ListCarAvailabilityResponse{}, nil
}

//END OMIT
