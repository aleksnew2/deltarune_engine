package system

import "fmt"

// PrintPCInfo prints information about the current PC, including the bit system,
// username, user home directory, current temperature, RAM size, and operating system.
// It retrieves this information using various helper functions and prints it to the console.
// If any error occurs while retrieving the information, the function will panic.
func PrintPCInfo() {
	username, err := GetCurrentUserName()

	if err != nil {
		panic(err)
	}

	home, err := Homedir()
	if err != nil {
		panic(err)
	}

	temperature, err := GetCurrentTemperature()
	if err != nil {
		panic(err)
	}

	currentOS, err := Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bit System: %s\nUsername: %s\nUser home directory: %s\n", ConvertBitSystemIntoString(), username, home)
	fmt.Printf("Temperature: %f\nRAM: %s\nOS: %s", temperature, ConvertRAMSizeIntoString(), currentOS)
}
