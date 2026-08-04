package main

import (
	"encoding/json"
	"log"
	"os"
)

type SatelliteResponse struct {
	Sensor_name string `json:"sensor_name"`
	Temperature float32 `json:"temperature"`
	Pressure float32 `json:"pressure"`
	Position []struct {
		City string `json:"city"`
		Height float32 `json:"height"`
	}
	Time string `json:"time"`
	Info string `json:"info"`
	Specs []struct {
		Name string `json:"name"`
		Model string `json:"model"`
		LaunchDate string `json:"launchDate"`
		Sensors string `json:"sensors"`
		Nation string `json:"nation"`
	}
}

func readAndDecode() SatelliteResponse {
	fileName := "data.json"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var satelliteResponse SatelliteResponse
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&satelliteResponse); err != nil {
		log.Fatal(err)
	}

	return satelliteResponse

}
