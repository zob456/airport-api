package data

import (
	"AirportApi/airport-service/models"
	pb "AirportApi/airport-service/proto"
	"database/sql"
)

func SelectAirportDetails(db *sql.DB, airportID string) (*pb.AirportDetailsRes, error) {
	airportDetails := &pb.AirportDetailsRes{}
	/*language=PostgreSQL*/
	const query = `SELECT
       "AirportID",
       "Name",
       "City",
       "Country",
       "IATACode",
       "ICAOCode",
       "Latitude",
       "Longitude",
       "Altitude",
       "TimeZone"
	FROM "AirportData"."vw_Airport"
	WHERE "AirportID" = $1 LIMIT 1;`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(airportID)
	err = row.Scan(
		&airportDetails.AirportID,
		&airportDetails.AirportName,
		&airportDetails.City,
		&airportDetails.Country,
		&airportDetails.IATACode,
		&airportDetails.ICAOCode,
		&airportDetails.Latitude,
		&airportDetails.Longitude,
		&airportDetails.Altitude,
		&airportDetails.Altitude,
		&airportDetails.TimeZone,
	)
	if err != nil {
		return nil, err
	}
	return airportDetails, nil
}

func SelectLongLatForAirport(db *sql.DB, airportID string) (*models.LongLat, error) {
	longLat := &models.LongLat{}
	/*language=PostgreSQL*/
	const query = `SELECT
       "Latitude",
       "Longitude"
	FROM "AirportData"."vw_Airport"
	WHERE "AirportID" = $1 LIMIT 1;`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(airportID)
	err = row.Scan(
		&longLat.Latitude,
		&longLat.Longitude,
	)
	if err != nil {
		return nil, err
	}
	return longLat, nil
}