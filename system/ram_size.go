package system

import "golang.org/x/sys/unix"

func GetRAMSize() uint64 {
	var stats unix.Sysinfo_t
	err := unix.Sysinfo(&stats)
	if err != nil {
		panic(err)
	}
	return stats.Totalram * uint64(stats.Unit) / (1024 * 1024 * 1024)
}
