package main

import (
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
	var airportDataServer pb.AirportDataServer
	s := grpc.NewServer()
	pb.RegisterAirportDataServer(s, airportDataServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}