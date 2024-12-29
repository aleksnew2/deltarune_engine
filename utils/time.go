package utils

import (
	"fmt"
	"time"
)

// ReturnTimeStamp returns the current time as a formatted string in the format "hour:minute".
// The hour and minute are extracted from the current local time.
func ReturnTimeStamp() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d:%d", currentTime.Hour(), currentTime.Minute())
}
