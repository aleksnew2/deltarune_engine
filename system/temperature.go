package system

import (
	"github.com/aleksnew2/deltarune_engine/exception"

	"github.com/shirou/gopsutil/host"
)

// GetCurrentTemperature retrieves the current system temperature reading from hardware sensors.
// It returns the temperature value as a float64 in Celsius and an error if the reading fails.
// If multiple temperature sensors are present, returns the last sensor's reading.
// Returns 0.0 and an error if unable to access temperature sensors.
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
