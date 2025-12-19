package api

import (
	"bytes"
	"encoding/json"
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
) error {

	payload := model.TemperaturePayload{
		Temperature:    temp,
		DatasourceName: datasource,
		ClientID:       secrets.ClientID,
		ClientSecret:   secrets.ClientSecret,
	}

	return c.post(payload)
}

func (c *Client) SendError(
	message string,
) error {

	payload := map[string]string{}

	return c.post(payload)
}

func (c *Client) post(body any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.URL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = c.HTTP.Do(req)
	return err
}
