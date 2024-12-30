package system

import (
	"github.com/aleksnew2/deltarune_engine/exception"

	"github.com/shirou/gopsutil/host"
)

func GetCurrentTemperature() (string, float64, error) {
	var temperature float64
	var sensorKey string

	temps, err := host.SensorsTemperatures()

	if err != nil {
		return "", 0.0, exception.CreateException(-1, err.Error())
	}

	for _, t := range temps {
		temperature = t.Temperature
		sensorKey = t.SensorKey
	}

	return sensorKey, temperature, nil
}
