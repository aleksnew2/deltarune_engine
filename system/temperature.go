package system

import (
	"strconv"

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

// ConvertTemperatureIntoString retrieves the current temperature and converts it into a string representation.
// It returns the temperature as a string and any error encountered during the process.
// If an error occurs while getting the temperature, it wraps the error in an exception with code -1.
func ConvertTemperatureIntoString() (string, error) {
	temperature, err := GetCurrentTemperature()
	if err != nil {
		return "", exception.CreateException(-1, err.Error())
	}

	return strconv.Itoa(int(temperature)), nil
}
