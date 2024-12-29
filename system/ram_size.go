package system

import "github.com/shirou/gopsutil/mem"

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
