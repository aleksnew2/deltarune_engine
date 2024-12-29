package system

import "strconv"

// BitSystem determines the bit size of the system.
// It returns 32 if the system is 32-bit, and 64 if the system is 64-bit.
func BitSystem() int8 {
	if strconv.IntSize == 64 {
		return 64
	}
	return 32
}

// ConvertBitSystemIntoString converts the result of the BitSystem function
// into a string representation. It assumes that BitSystem returns an integer
// value which is then converted to a string using strconv.Itoa.
func ConvertBitSystemIntoString() string {
	return strconv.Itoa(int(BitSystem()))
}
