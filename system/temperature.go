package system

import (
	"github.com/aleksnew2/deltarune_engine/exception"

	"github.com/shirou/gopsutil/host"
)

func GetCurrentTemperature() (float64, error) {
	var temperature float64
	temps, err := host.SensorsTemperatures()

	if err != nil {
		return 0.0, exception.CreateException(-1, err.Error())
	}

	for _, t := range temps {
		temperature = t.Temperature
	}

	return temperature, nil
}
