package handlers

import (
	pb "AirportApi/airport-service/proto"
	"context"
)

func (s *pb.AirportDataServer) SayHello(ctx context.Context, in *pb.AirportDetailsReq) (*pb.AirportDetailsRes, error) {
	return &pb.AirportDetailsRes{}, nil
}
