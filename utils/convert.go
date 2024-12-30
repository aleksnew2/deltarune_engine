package utils

// ConvertAnyIntoString converts a value of any type to a string.
// If the value is already a string, it returns the value as is.
// Otherwise, it returns an empty string.
//
// Parameters:
// 	- value (any): The value to be converted.
func ConvertAnyIntoString(value any) string {
	if str, ok := value.(string); ok {
		return str
	}

	return ""
}
