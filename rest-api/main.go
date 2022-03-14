package main

import (
	"AirportApi/rest-api/handlers"
	"AirportApi/rest-api/middleware"
	"AirportApi/rest-api/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	airCtx := &utils.AirContext{}
	db := utils.ConnectDB()
	router := gin.Default()
	if os.Getenv("ENV") == "local" {
		router.Use(middleware.CORSMiddleware())
	}
	router.Use(middleware.ValidateToken(db, airCtx))

	// Auth handlers
	authHandlers := handlers.AuthHandler{DB: db}
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandlers.Login())
		auth.POST("/logout", authHandlers.Logout(airCtx))
	}

	// gRPC Handlers
	grpcGroup := router.Group("/grpc")
	{
		grpcGroup.POST("/details", handlers.GetAirportDetails(conn))
		grpcGroup.POST("/distance", handlers.GetAirportDistance(conn))
	}

	log.Fatal(router.Run(":8000"))
}
