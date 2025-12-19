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
)

func main() {
	commandArgs := ReadCommandArgs()

	cfg, err := config.LoadConfig(commandArgs.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	reporter := service.Reporter{
		Sensor: &sensor.DHT22{Pin: "GPIO17", Dht: nil},
		API: &api.Client{
			HTTP: &http.Client{Timeout: 10 * time.Second},
			URL:  cfg.ApiURL,
		},
		Secrets:     cfg.Secrets,
		Datasource:  "RockfordHome",
		MaxFailures: 10,
	}

	for {
		if _, err := reporter.ProcessTemperature(); err != nil {
			log.Println("run failed:", err)
		}

		time.Sleep(time.Duration(commandArgs.ReadInterval) * time.Minute)
	}
}

func ReadCommandArgs() model.CmdArgs {
	var commandArgs model.CmdArgs
	flag.StringVar(&commandArgs.ConfigFilePath, "c", "./config/config.json", "Path to config file for temperature sensor")
	flag.IntVar(&commandArgs.ReadInterval, "i", 5, "Interval to sleep between temperature reads")
	return commandArgs
}
