package main

import (
	"log"

    "fmt"

    "net/http"

    "strconv"

    "strings"

    "time"
 "github.com/gin-gonic/gin"

)
type ChargingStation struct {

    ID               uint      `gorm:"primaryKey" json:"id"`

    StationID        int       `gorm:"unique" json:"stationID"`

    EnergyOutput     string    `json:"energyOutput"`

    Type             string    `json:"type"`

    Status           string    `json:"status"`

    AvailabilityTime time.Time `json:"time"`

}

type StartChargingRequest struct {

    StationID              int    `json:"stationID"`

    VehicleBatteryCapacity string `json:"vehicleBatteryCapacity"`

    CurrentVehicleCharge   string `json:"currentVehicleCharge"`

}

type ChargingSession struct {

    ID                     uint      `gorm:"primaryKey" json:"id"`

    StationID              int       `json:"stationID"`

    VehicleBatteryCapacity string    `json:"vehicleBatteryCapacity"`

    CurrentVehicleCharge   string    `json:"currentVehicleCharge"`

    ChargingStartTime      time.Time `json:"chargingStartTime"`

    AvailabilityTime       time.Time `json:"availabilityTime"`

}

func AddChargingStation(c *gin.Context) {
    db, _ := ConnectDatabase()

    var chargingStation ChargingStation

    if err := c.ShouldBindJSON(&chargingStation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    chargingStation.Status = "available"
    chargingStation.AvailabilityTime = time.Now()

    // Check if the station ID already exists
    var existingStation ChargingStation
    result := db.First(&existingStation, "station_id = ?", chargingStation.StationID)
    if result.RowsAffected > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Charging station already exists"})
        log.Println("charging station already existed")
        return
    }

    err := db.Create(&chargingStation).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, chargingStation)
}
func parseKWh(value string) float64 {

    val := strings.TrimSuffix(value, "kwh")

    result, err := strconv.ParseFloat(val, 64)

    if err != nil {

        return 0.0

    }

    return result

}
func StartCharging(c *gin.Context) {

    db, _ := ConnectDatabase()

    var request StartChargingRequest

    if err := c.ShouldBindJSON(&request); err != nil {

        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

        return

    }

   var station ChargingStation

    err := db.First(&station, "station_id = ?", request.StationID).Error

    if err != nil {

        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

        return

    }

    batteryCapacity := parseKWh(request.VehicleBatteryCapacity)
    
    currentCharge := parseKWh(request.CurrentVehicleCharge)

    energyOutput := parseKWh(station.EnergyOutput)

    remainingCharge := batteryCapacity - currentCharge


fmt.Println(batteryCapacity,currentCharge,energyOutput,station)

    chargingStartTime := time.Now() // Use current local time as the charging start time
 // Calculate the charging duration based on energy output and remaining charge

    chargingDuration := time.Duration(remainingCharge/energyOutput) * time.Hour

    fmt.Println(chargingDuration)

    // Calculate the availability time based on charging start time and charging duration

    availabilityTime := chargingStartTime.Add(chargingDuration)

    station.AvailabilityTime = availabilityTime

    db.Save(&station)

    loc,err:= time.LoadLocation("Asia/Kolkata")

    if err != nil {

        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

        return

    }

    availabilityTime = availabilityTime.In(loc)

    chargingSession := ChargingSession{

        StationID:              request.StationID,

        VehicleBatteryCapacity: request.VehicleBatteryCapacity,

        CurrentVehicleCharge:   request.CurrentVehicleCharge,

        ChargingStartTime:      chargingStartTime,

        AvailabilityTime:       availabilityTime,

    }

    err = db.Create(&chargingSession).Error

    if err != nil {

        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

        return

    }
     response := struct {

        ChargingSession

        AvailabilityTimeFormatted string `json:"availabilityTime"`

    }{

        ChargingSession:           chargingSession,

        AvailabilityTimeFormatted: availabilityTime.Format("2006-01-02 15:04:05"),

    }
     c.JSON(http.StatusOK, response)

    db.Model(&station).Update("status", "occupied")

}

func GetChargingStationByID(c *gin.Context) {
// Get the charging station ID from the request parameters

    stationID := c.Param("id")

// Try to get the charging station from the cache

    cachedStation, found := GetFromCache("charging_station_" + stationID)

    if found {

        // Charging station found in cache
        chargingStation := cachedStation.(ChargingStation)
         source := "cache"
        c.JSON(http.StatusOK, gin.H{"data": chargingStation, "source": source})
        
        return

    }

// Charging station not found in cache, query the database
db, err := ConnectDatabase()
if err != nil {
     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
     }
    var chargingStation ChargingStation

    if err := db.First(&chargingStation, "id = ?", stationID).Error; err != nil {

        c.JSON(http.StatusNotFound, gin.H{"error": "Charging station not found"})

        return

    }
// Store the fetched charging station in cache for future use
     SetToCache("charging_station_"+stationID, chargingStation)

    c.JSON(http.StatusOK, chargingStation)

}