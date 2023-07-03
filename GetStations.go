package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAvailableChargingStations(c *gin.Context) {
	var chargingStations []ChargingStation
	db, err := ConnectDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in connecting databse": err.Error()})
		return
	}

	if err := db.Where("availability_time  > ?", time.Now()).Find(&chargingStations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Output generated from database")

	c.JSON(http.StatusOK, chargingStations)
}

func GetOccupiedChargingStations(c *gin.Context) {
	var chargingStations []ChargingStation
	db, err := ConnectDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	
		}else {
		if err := db.Where("availability_time  > ?", time.Now()).Find(&chargingStations).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, chargingStations)
}
