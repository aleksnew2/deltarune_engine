package message

func convertEnumIntoString(message_type Type) string {
	switch message_type {
	case Error:
		return "deltarune_error"
	case Warning:
		return "deltarune_warning"
	case Info:
		return "deltarune_info"
	}

	return ""
}

// Format converts a message type enum into its corresponding string representation.
// It takes a parameter of type `Type` and returns the string representation of that type.
//
// Parameters:
//   - message_type: The message type enum to be converted.
//
// Returns:
//   - A string representing the message type.
func Format(message_type Type) string {
	return convertEnumIntoString(message_type)
}
