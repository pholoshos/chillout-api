// internal/logger/logger.go

package logger

import (
	"log"
	"os"
)

// Logger is a struct that contains a logger instance.
type Logger struct {
	*log.Logger
}

// NewLogger returns a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "", log.Ldate|log.Ltime)}
}

// Error logs an error message to the logger.
func (l *Logger) Error(v ...interface{}) {
	l.Logger.Println("[ERROR]", v)
}

// Info logs an info message to the logger.
func (l *Logger) Info(v ...interface{}) {
	l.Logger.Println("[INFO]", v)
}
