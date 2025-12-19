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
		if err := reporter.ProcessTemperature(); err != nil {
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

// func main() {
// 	var minTemp float64
// 	retryCount := 0
// 	for {

// 		err := dht.HostInit()
// 		if err != nil {
// 			fmt.Println("HostInit error:", err)
// 			return
// 		}

// 		dht, err := dht.NewDHT("GPIO17", dht.Fahrenheit, "")
// 		if err != nil {
// 			fmt.Println("NewDHT error:", err)
// 		}

// 		humidity, temperature, err := dht.ReadRetry(5)
// 		if err != nil {
// 			retryCount++
// 			fmt.Println("Read error:", err)
// 			if retryCount > 5 {
// 				//send email
// 			}
// 			continue
// 		}

// 		retryCount = 0

// 		if minTemp == 0.0 || temperature < minTemp {
// 			minTemp = temperature
// 			os.WriteFile("./min_temp.txt", []byte(fmt.Sprintf("Min temp was %.2f", minTemp)), 0644)
// 		}

// 		fmt.Printf("humidity: %v\n", humidity)
// 		fmt.Printf("temperature: %v\n", math.Round(10*temperature)/10)

// 		time.Sleep(4 * time.Second)
// 	}
// }
