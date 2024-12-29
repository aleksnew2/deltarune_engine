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

func Format(message_type Type) string {
	return convertEnumIntoString(message_type)
}
