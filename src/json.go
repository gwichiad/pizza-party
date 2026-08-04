package main

import "encoding/json"

type SatelliteResponse struct {
	sensor_name string `json:"sensor_name"`
	temperature float32 `json:"temperature"`
	pressure float32 `json:"pressure"`
	position []struct {
		city string `json:"city"`
		height float32 `json:"height"`
	}
	time string `json:"time"`
	info string `json:"info"`
	specs []struct {
		name string `json:"name"`
		model string `json:"model"`
		launchDate string `json:"launchDate`
		sensors string `json:"sensors"`
		nation string `json:nation`
	}
}
