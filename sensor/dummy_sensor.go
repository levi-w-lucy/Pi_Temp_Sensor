package sensor

import "errors"

type FakeSensor struct {
	Readings []float64
	Errors   int
	calls    int
}

func (f *FakeSensor) ReadTemperature() (float64, error) {
	f.calls++

	if f.Errors > 0 {
		f.Errors--
		return 0, errors.New("sensor error")
	}

	if len(f.Readings) == 0 {
		return 0, errors.New("no readings")
	}

	return f.Readings[0], nil
}
