package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"pitempsensor/model"
)

type Client struct {
	HTTP *http.Client
	URL  string
}

func (c *Client) SendTemperature(
	temp float64,
	datasource string,
	secrets model.Secrets,
) (model.GoogleSheetAPIResponse, error) {

	payload := model.TemperaturePayload{
		Temperature:    temp,
		DatasourceName: datasource,
		ClientID:       secrets.ClientID,
		ClientSecret:   secrets.ClientSecret,
	}

	return c.post(payload)
}

func (c *Client) SendError(message string) (model.GoogleSheetAPIResponse, error) {

	payload := map[string]string{
		"error": message,
	}

	return c.post(payload)
}

func (c *Client) post(body any) (model.GoogleSheetAPIResponse, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return model.GoogleSheetAPIResponse{}, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.URL,
		bytes.NewBuffer(data),
	)

	if err != nil {
		return model.GoogleSheetAPIResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := c.HTTP.Do(req)
	if err != nil {
		return model.GoogleSheetAPIResponse{}, err
	}

	responseAsBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return model.GoogleSheetAPIResponse{}, err
	}
	var temperatureResponse model.GoogleSheetAPIResponse
	err = json.Unmarshal(responseAsBytes, &temperatureResponse)

	if err != nil {
		return model.GoogleSheetAPIResponse{}, err
	}

	return temperatureResponse, err
}
