package service

import (
	"fmt"
	"time"

	"pitempsensor/api"
	"pitempsensor/model"
	"pitempsensor/sensor"
)

type Reporter struct {
	Sensor      sensor.Reader
	API         api.ReporterClient
	Secrets     model.Secrets
	Datasource  string
	MaxFailures int
}

func (r *Reporter) ProcessTemperature() (model.GoogleSheetAPIResponse, error) {
	failures := 0

	for {
		temp, err := r.Sensor.ReadTemperature()
		if err == nil {
			return r.API.SendTemperature(
				temp,
				r.Datasource,
				r.Secrets,
			)
		}

		failures++
		time.Sleep(5 * time.Second)
		if failures >= r.MaxFailures {
			return r.API.SendError(
				fmt.Sprintf("temperature read failed %d times", failures),
			)
		}
	}
}
