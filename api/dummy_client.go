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
) error {
	f.SentTemperature = &model.TemperaturePayload{
		Temperature:    temp,
		DatasourceName: datasource,
		ClientID:       secrets.ClientID,
		ClientSecret:   secrets.ClientSecret,
	}
	return nil
}

func (f *FakeClient) SendError(message string) error {
	f.SentError = message
	return nil
}
