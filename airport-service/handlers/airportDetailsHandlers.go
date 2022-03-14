package handlers

import (
	"AirportApi/airport-service/data"
	"AirportApi/airport-service/models"
	pb "AirportApi/airport-service/proto"
	"context"
	"database/sql"
	"math"
)

type AirportDataConfig struct {
	pb.UnimplementedAirportDataServer
	DB *sql.DB
}

func (c *AirportDataConfig) GetAirportDetails(_ context.Context, in *pb.AirportDetailsReq) (*pb.AirportDetailsRes, error) {
	airportDetails, err := data.SelectAirportDetails(c.DB, in.AirportID)
	if err != nil {
		return nil, err
	}
	return airportDetails, nil
}

func (c *AirportDataConfig) GetDistance(_ context.Context, in *pb.AirportDistanceReq) (*pb.AirportDistanceRes, error) {
	res := &pb.AirportDistanceRes{}
	firstLongLat, err := data.SelectLongLatForAirport(c.DB, in.FirstAirportID)
	if err != nil {
		return nil, err
	}

	secondLongLat, err := data.SelectLongLatForAirport(c.DB, in.SecondAirportID)
	if err != nil {
		return nil, err
	}

	distance := calculateDistance(firstLongLat, secondLongLat)

	res.Distance = distance

	return res, nil
}

func calculateDistance(firstLongLat *models.LongLat, secondLongLat *models.LongLat) float64 {
	long1 := firstLongLat.Longitude / (180 / math.Pi)
	lat1 := firstLongLat.Latitude / (180 / math.Pi)
	long2 := secondLongLat.Longitude / (180 / math.Pi)
	lat2 := secondLongLat.Latitude / (180 / math.Pi)

	dlon := long2 - long1
	dlat := lat2 - lat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	t := math.Asin(math.Sqrt(a))

	r := 6371.0
	distance := t * r
	return distance
}
