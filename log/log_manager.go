package log

import (
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/contracts/log"
)

// LoggerManager struct
type LoggerManager struct {
	app        foundation.ApplicationInterface
	dateFormat string
	channels   map[string]log.LoggerInterface
}

// NewLoggerManager new log manager instance
func NewLoggerManager(app foundation.ApplicationInterface) *LoggerManager {
	LoggerManager := &LoggerManager{app: app}
	LoggerManager.dateFormat = "2222-02-22 22:22:22"

	LoggerManager.channels = map[string]log.LoggerInterface{
		"default": NewLoggerOfLogrus(),
	}

	return LoggerManager
}

// Driver set Driver
func (l *LoggerManager) Driver(message string) log.LoggerInterface {
	return l.channels[message]
}

// Trace Something very low level.
func (l *LoggerManager) Trace(message string, fields map[string]interface{}) {
	l.Driver("default").Trace(message, fields)
}

// Debug Useful debugging information.
func (l *LoggerManager) Debug(message string, fields map[string]interface{}) {
	l.Driver("default").Debug(message, fields)
}

// Info Something noteworthy happened!
func (l *LoggerManager) Info(message string, fields map[string]interface{}) {
	l.Driver("default").Info(message, fields)
}

// Warn You should probably take a look at this.
func (l *LoggerManager) Warn(message string, fields map[string]interface{}) {
	l.Driver("default").Warn(message, fields)
}

// Error Something failed but I'm not quitting.
func (l *LoggerManager) Error(message string, fields map[string]interface{}) {
	l.Driver("default").Error(message, fields)
}

// Fatal Calls os.Exit(1) after logging
func (l *LoggerManager) Fatal(message string, fields map[string]interface{}) {
	l.Driver("default").Fatal(message, fields)
}

// Panic Calls panic() after logging
func (l *LoggerManager) Panic(message string, fields map[string]interface{}) {
	l.Driver("default").Panic(message, fields)
}
