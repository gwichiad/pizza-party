package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
)

func readAndDecode(r io.Reader) (SatelliteResponse, error) {
	var data SatelliteResponse

	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&data); err != nil {
		log.Fatal(err)
	}

	if data.Specs.Name == "" {
		return SatelliteResponse{}, errors.New("satellite name is required")
	}
	if data.Time == "" {
		return SatelliteResponse{}, errors.New("time is required")
	}

	data.SatelliteName = data.Specs.Name
	return data, nil
}
