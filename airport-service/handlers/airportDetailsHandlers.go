package handlers

import (
	pb "AirportApi/airport-service/proto"
	"context"
)

type Server struct {
	pb.UnimplementedAirportDataServer
}

func (s *Server) GetAirportDetails(ctx context.Context, in *pb.AirportDetailsReq) (*pb.AirportDetailsRes, error) {
	return &pb.AirportDetailsRes{}, nil
}
