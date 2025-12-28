package sensor

import (
	"math"

	"github.com/MichaelS11/go-dht"
)

type Reader interface {
	ReadTemperature() (float64, error)
}

// Pin naming should follow standards found in the go-dht package
type DHT22 struct {
	Pin string
	Dht *dht.DHT
}

func (d *DHT22) InitializeSensor() error {
	err := dht.HostInit()
	if err != nil {
		return err
	}

	dht, err := dht.NewDHT(d.Pin, dht.Fahrenheit, "")
	if err != nil {
		return err
	}

	d.Dht = dht
	return nil
}

func (d *DHT22) ReadTemperature() (float64, error) {
	_, temperature, err := d.Dht.ReadRetry(11)
	if err != nil {
		return -1, err
	}

	return math.Round(temperature*100) / 100, nil
}
