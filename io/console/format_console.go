// Package console provides functionalities for input and output operations
// in the console, including formatting and displaying text.
package console

import (
	"fmt"
	"github.com/aleksnew2/deltarune_engine/io/console/message"
	"github.com/aleksnew2/deltarune_engine/utils"
)

func formatExceptionCode(settings Settings, error_code int32) string {
	if settings.IsWithExceptionCode {
		return fmt.Sprintf(" [%d]", error_code)
	}

	return ""
}

func formatTimeStamp(settings Settings) string {
	if settings.IsWithTimestamp {
		return fmt.Sprintf("[%s] ", utils.ReturnTimeStamp())
	}

	return ""
}

func format(settings Settings, message_type message.Type, args ...any) string {
	formattedMessageType := message.Format(message_type)
	formattedTimeStamp := formatTimeStamp(settings)

	return fmt.Sprintf("%s%s: %s", formattedTimeStamp, formattedMessageType, utils.Join(args, ""))
}

func formatException(settings Settings, error_code int32, args ...any) string {
	formattedMessageType := message.Format(message.Error)
	formattedTimeStamp := formatTimeStamp(settings)
	formattedExceptionCode := formatExceptionCode(settings, error_code)
	return fmt.Sprintf("%s%s%s: %s", formattedTimeStamp, formattedMessageType, formattedExceptionCode, utils.Join(args, ""))
}
