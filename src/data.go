package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *api) dataHandler(w http.ResponseWriter, r *http.Request) {
	field := chi.URLParam(r, "field")
	subfield := chi.URLParam(r, "subfield")

	data := readAndDecode()

	if subfield == "" {
		switch field {
		case "sensor_name":
			fmt.Fprint(w, data.Sensor_name)
		case "temperature":
			fmt.Fprint(w, data.Temperature)
		case "pressure":
			fmt.Fprint(w, data.Pressure)
		case "time":
			fmt.Fprint(w, data.Time)
		case "info":
			fmt.Fprint(w, data.Info)
		default:
			http.Error(w, "unknown field", http.StatusNotFound)

		}
	} else {
		switch field {
		case "position":
			switch subfield {
			case "city":
				fmt.Fprint(w, data.Position.City)
			case "height":
				fmt.Fprint(w, data.Position.Height)
			default:
				http.Error(w, "unknown field", http.StatusNotFound)
			}
		case "specs":
			switch subfield {
			case "name":
				fmt.Fprint(w, data.Specs.Name)
			case "model":
				fmt.Fprint(w, data.Specs.Model)
			case "launch_date":
				fmt.Fprint(w, data.Specs.LaunchDate)
			case "sensor":
				fmt.Fprint(w, data.Specs.Sensors)
			case "nation":
				fmt.Fprint(w, data.Specs.Nation)
			default:
				http.Error(w, "unknown field", http.StatusNotFound)
			}
		}
	}
}
