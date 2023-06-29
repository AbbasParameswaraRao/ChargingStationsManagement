# ChargingStationsManagement

Created a Charging Station Management application which has the following functionalities

Add Charging Stations Start Charging Available Charging Stations Occupied ChargingStations Along With Their Availability Time.

Steps to Run The code

=> Install GOLANG and set path correctly. => Create mod file with command " go mod init => Install all the necessary dependencies with command "go mod tidy" 4.To run the code use : " go run . => This will start the server in port 8080

Check whether the server is running or not by simply opening the new chrome tab and enter the following URL: "localhost:8080" it should give some response. Since we have not passed any data it should give "error 404".

You can use API testing tools like "POSTMAN" or "ThundeClient(vscode extension)" to check its functioning. Sample Requests and Responses.

1.Add Charging Station:

POST localhost:8080/charging-stations

Request Body: { "stationID": 01, "energyOutput": "600kWh", "type": "DC", "status":Available }

2.Start Charging: Request Body :{"stationID" :1 ,"vehicleBatteryCapacity" :30Kwh, "currentVehicleCharge": 30kw }
3.Available Charging Stations 
4.Occupied Charging Station
5.Get charging station by ID.

[ https://localhost:8080/charging-stations, AddChargingStation ], 
[https://localhost:8080/charging-sessions, StartCharging], 
[ https://localhost:8080/charging-stations/available, GetAvailableChargingStations ], 
[ https://localhost:8080/charging-stations/occupied, GetOccupiedChargingStations ], 
[ https://localhost:8080/charging-stations/:id, GetChargingStationByID ]
