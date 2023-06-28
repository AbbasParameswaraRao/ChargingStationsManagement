package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	InitLogger()
	

	err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration")
	}

	// the database connection
	_, err = ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database")
	}

	// Initialize the cache
	InitCache()
	db, err := ConnectDatabase()
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %s", err))
	}

	// Auto-migrate the model structs to create the tables
	db.AutoMigrate(&ChargingStation{})
	db.AutoMigrate(&ChargingSession{})

	// Create a new Gin router
	router := gin.Default()
	// Use the logger middleware
	router.Use(LoggerMiddleware())

	// Define the API routes
	router.POST("/charging-stations", AddChargingStation)
	router.POST("/charging-sessions", StartCharging)
	router.GET("/charging-stations/available", GetAvailableChargingStations)
	router.GET("/charging-stations/occupied", GetOccupiedChargingStations)
	router.GET("/charging-stations/:id", GetChargingStationByID)

	// Run the server
	router.Run(":8080")
}
