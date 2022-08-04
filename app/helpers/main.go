package helpers

import (
	"fmt"
	"time"
)

// FormatLog formats the log message to be unified across all the packages.
func FormatLog(message string) string {
	return fmt.Sprintf(time.Now().Format(time.UnixDate) + " - " + message + "\n")
}
