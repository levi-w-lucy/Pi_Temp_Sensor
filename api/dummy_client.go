package api

import "pitempsensor/model"

type FakeClient struct {
	SentTemperature *model.TemperaturePayload
	SentError       string
}

func (f *FakeClient) SendTemperature(
	temp float64,
	datasource string,
	secrets model.Secrets,
) (model.GoogleSheetAPIResponse, error) {
	f.SentTemperature = &model.TemperaturePayload{
		Temperature:    temp,
		DatasourceName: datasource,
		ClientID:       secrets.ClientID,
		ClientSecret:   secrets.ClientSecret,
	}

	return model.GoogleSheetAPIResponse{
		Success:    true,
		StatusCode: 200,
		Message:    "Successfully updated temperature",
	}, nil
}

func (f *FakeClient) SendError(message string) (model.GoogleSheetAPIResponse, error) {
	f.SentError = message
	return model.GoogleSheetAPIResponse{
		Success:    true,
		StatusCode: 200,
		Message:    "Successfully updated temperature",
	}, nil
}
