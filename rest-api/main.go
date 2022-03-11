package main

import (
	"AirportApi/rest-api/handlers"
	"AirportApi/rest-api/middleware"
	"AirportApi/rest-api/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
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

	log.Fatal(router.Run(":8000"))
}
