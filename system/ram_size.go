package system

import (
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

// GetRAMSize returns the total amount of RAM installed in the system in gigabytes (GB).
// It uses the system's virtual memory information to calculate the total RAM.
// The function will panic if it fails to retrieve the memory information.
func GetRAMSize() uint64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	return v.Total / (1024 * 1024 * 1024)
}

// ConvertRAMSizeIntoString converts the RAM size from bytes into a string representation.
// It retrieves the RAM size using GetRAMSize function and returns it as a string.
// Returns a string containing the RAM size value.
func ConvertRAMSizeIntoString() string {
	return strconv.Itoa(int(GetRAMSize()))
}
