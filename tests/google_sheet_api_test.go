package tests

import (
	"net/http"
	"pitempsensor/api"
	"pitempsensor/config"
	"testing"
	"time"
)

func TestSendTemperature(t *testing.T) {
	config, err := config.LoadConfig("../config/config.json")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	client := &api.Client{
		HTTP: &http.Client{Timeout: 10 * time.Second},
		URL:  config.ApiURL,
	}

	temperaturePayload, err := client.SendTemperature(64, "Rockford", config.Secrets)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !temperaturePayload.Success {
		t.Errorf("api call unsuccessful")
	}

	if temperaturePayload.StatusCode != 200 {
		t.Errorf("Wrong status code %d", temperaturePayload.StatusCode)
	}
}

func TestSendError(t *testing.T) {
	config, err := config.LoadConfig("../config/config.json")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	client := &api.Client{
		HTTP: &http.Client{Timeout: 10 * time.Second},
		URL:  config.ApiURL,
	}

	errorPayload, err := client.SendError("Test Error from GO")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !errorPayload.Success {
		t.Errorf("api call unsuccessful")
	}

	if errorPayload.StatusCode != 200 {
		t.Errorf("Wrong status code %d", errorPayload.StatusCode)
	}
}
