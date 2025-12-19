package tests

import (
	"pitempsensor/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	received, err := config.LoadConfig("../config/config_test.json")

	if err != nil {
		t.Fatalf("PostTemperature returned error: %v", err)
	}

	if received.ApiURL != "https://google.com" {
		t.Errorf("ApiURL mismatch: %v", received.ApiURL)
	}

	if received.Secrets.ClientID != "abc" {
		t.Errorf("ApiURL mismatch: %v", received.Secrets.ClientID)
	}

	if received.Secrets.ClientSecret != "def" {
		t.Errorf("ApiURL mismatch: %v", received.Secrets.ClientSecret)
	}
}
