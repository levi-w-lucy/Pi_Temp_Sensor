package tests

import (
	"testing"

	"pitempsensor/api"
	"pitempsensor/model"
	"pitempsensor/sensor"
	"pitempsensor/service"
)

func TestReporter_SendsTemperature(t *testing.T) {
	fakeSensor := &sensor.FakeSensor{
		Readings: []float64{22.3},
	}

	fakeAPI := &api.FakeClient{}

	r := service.Reporter{
		Sensor:      fakeSensor,
		API:         fakeAPI,
		Secrets:     model.Secrets{ClientID: "id", ClientSecret: "secret"},
		Datasource:  "test",
		MaxFailures: 10,
	}

	err := r.ProcessTemperature()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fakeAPI.SentTemperature == nil {
		t.Fatal("temperature was not sent")
	}

	if fakeAPI.SentTemperature.Temperature != 22.3 {
		t.Errorf("wrong temperature")
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

	err := r.ProcessTemperature()
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

	err := r.ProcessTemperature()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fakeAPI.SentError == "" {
		t.Fatal("expected error message to be sent")
	}
}
