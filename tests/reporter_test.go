package tests

import (
	"net/http"
	"testing"
	"time"

	"pitempsensor/api"
	"pitempsensor/config"
	"pitempsensor/model"
	"pitempsensor/sensor"
	"pitempsensor/service"
)

func TestReporter_SendsTemperature(t *testing.T) {
	config, err := config.LoadConfig("../config/config.json")
	if err != nil {
		t.Fatalf("could not load config file. Error: %v", err)
	}

	fakeSensor := &sensor.FakeSensor{
		Readings: []float64{72.1},
	}

	client := &api.Client{
		HTTP: &http.Client{Timeout: 10 * time.Second},
		URL:  config.ApiURL,
	}

	r := service.Reporter{
		Sensor:      fakeSensor,
		API:         client,
		Secrets:     config.Secrets,
		Datasource:  "Rockford",
		MaxFailures: 10,
	}

	response, err := r.ProcessTemperature()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !response.Success {
		t.Error("response returned unsuccessful")
	}

	if response.StatusCode != 200 {
		t.Errorf("received incorrect status code: %d", response.StatusCode)
	}

}

func TestReporter_RetriesBeforeSuccess(t *testing.T) {
	fakeSensor := &sensor.FakeSensor{
		Errors:   3,
		Readings: []float64{19.8},
	}

	fakeAPI := &api.FakeClient{}

	r := service.Reporter{
		Sensor:      fakeSensor,
		API:         fakeAPI,
		Secrets:     model.Secrets{},
		Datasource:  "test",
		MaxFailures: 10,
	}

	_, err := r.ProcessTemperature()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fakeAPI.SentTemperature == nil {
		t.Fatal("temperature was not sent after retries")
	}
}

func TestReporter_SendsErrorAfterMaxFailures(t *testing.T) {
	fakeSensor := &sensor.FakeSensor{
		Errors: 10,
	}

	fakeAPI := &api.FakeClient{}

	r := service.Reporter{
		Sensor:      fakeSensor,
		API:         fakeAPI,
		Secrets:     model.Secrets{},
		Datasource:  "test",
		MaxFailures: 10,
	}

	_, err := r.ProcessTemperature()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fakeAPI.SentError == "" {
		t.Fatal("expected error message to be sent")
	}
}
