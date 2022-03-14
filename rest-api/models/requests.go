package models

type LoginReq struct {
	UserName string
	Password string
}

type AirportDetailsReq struct {
	AirportID string
}

type AirportDistanceReq struct {
	FirstAirportID string
	SecondAirportID string
}
