package api

import "pitempsensor/model"

type ReporterClient interface {
	SendTemperature(float64, string, model.Secrets) error
	SendError(string) error
}
