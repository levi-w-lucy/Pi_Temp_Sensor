package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"pitempsensor/api"
	"pitempsensor/config"
	"pitempsensor/model"
	"pitempsensor/sensor"
	"pitempsensor/service"

	_ "github.com/MichaelS11/go-dht"
)

func main() {
	commandArgs := ReadCommandArgs()

	cfg, err := config.LoadConfig(commandArgs.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	sensor := sensor.DHT22{Pin: "GPIO17", Dht: nil}
	sensor.InitializeSensor()

	reporter := service.Reporter{
		Sensor: &sensor,
		API: &api.Client{
			HTTP: &http.Client{Timeout: 10 * time.Second},
			URL:  cfg.ApiURL,
		},
		Secrets:     cfg.Secrets,
		Datasource:  "Rockford",
		MaxFailures: 10,
	}

	for {
		_, err = reporter.ProcessTemperature()
		if err != nil {
			log.Println("run failed:", err)
		}
		time.Sleep(time.Duration(commandArgs.ReadInterval) * time.Minute)
	}
}

func ReadCommandArgs() model.CmdArgs {
	configFile := flag.String("c", "./config/config.json", "Path to config file for temperature sensor")
	interval := flag.Int("i", 5, "Interval to sleep between temperature reads")
	flag.Parse()

	return model.CmdArgs{
		ConfigFilePath: *configFile,
		ReadInterval:   *interval,
	}
}
