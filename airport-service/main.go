package main

import (
	"AirportApi/airport-service/handlers"
	pb "AirportApi/airport-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAirportDataServer(s, &handlers.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}