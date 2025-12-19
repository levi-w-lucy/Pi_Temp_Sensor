package api

import "pitempsensor/model"

type ReporterClient interface {
	SendTemperature(float64, string, model.Secrets) (model.GoogleSheetAPIResponse, error)
	SendError(string) (model.GoogleSheetAPIResponse, error)
}
