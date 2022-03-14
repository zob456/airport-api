package main

import (
	"AirportApi/airport-service/handlers"
	pb "AirportApi/airport-service/proto"
	"AirportApi/airport-service/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9000"
)

func main() {
	db := utils.ConnectDB()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAirportDataServer(s, &handlers.AirportDataConfig{DB: db})
	fmt.Println(fmt.Sprintf("Starting Airport Service gRPC server on port: %s", port))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
