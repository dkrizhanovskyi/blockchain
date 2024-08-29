package utils

import (
	"log"
	"os"
)

// Logger is a custom logger that wraps Go's built-in log package.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new instance of Logger.
// Parameters:
// - prefix: A string that will be prefixed to each log message.
// - flag: Log flags that define the properties of the log (e.g., log.Ldate, log.Ltime).
// Returns:
// - A new Logger instance.
func NewLogger(prefix string, flag int) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, prefix, flag),
	}
}

// Info logs an informational message.
// Parameters:
// - v: The message or variables to log.
func (l *Logger) Info(v ...interface{}) {
	l.SetPrefix("[INFO] ")
	l.Println(v...)
}

// Warn logs a warning message.
// Parameters:
// - v: The message or variables to log.
func (l *Logger) Warn(v ...interface{}) {
	l.SetPrefix("[WARN] ")
	l.Println(v...)
}

// Error logs an error message.
// Parameters:
// - v: The message or variables to log.
func (l *Logger) Error(v ...interface{}) {
	l.SetPrefix("[ERROR] ")
	l.Println(v...)
}

// Debug logs a debug message.
// Parameters:
// - v: The message or variables to log.
func (l *Logger) Debug(v ...interface{}) {
	l.SetPrefix("[DEBUG] ")
	l.Println(v...)
}
