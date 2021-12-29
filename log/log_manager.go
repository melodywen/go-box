package log

import (
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/contracts/log"
)

type LogManager struct {
	app        foundation.ApplicationInterface
	dateFormat string
	channels   map[string]log.LoggerInterface
}

func newLogManager(app foundation.ApplicationInterface) *LogManager {
	logManager := &LogManager{app: app}
	logManager.dateFormat = "2222-02-22 22:22:22"

	logManager.channels = map[string]log.LoggerInterface{
		"default": NewLoggerOfLogrus(),
	}

	return logManager
}

func (l *LogManager) Driver(message string) log.LoggerInterface {
	return l.channels[message]
}

func (l *LogManager) Trace(message string, fields map[string]interface{}) {
	l.Driver("default").Trace(message, fields)
}

func (l *LogManager) Debug(message string, fields map[string]interface{}) {
	l.Driver("default").Debug(message, fields)
}

func (l *LogManager) Info(message string, fields map[string]interface{}) {
	l.Driver("default").Info(message, fields)
}

func (l *LogManager) Warn(message string, fields map[string]interface{}) {
	l.Driver("default").Warn(message, fields)
}

func (l *LogManager) Error(message string, fields map[string]interface{}) {
	l.Driver("default").Error(message, fields)
}

func (l *LogManager) Fatal(message string, fields map[string]interface{}) {
	l.Driver("default").Fatal(message, fields)
}

func (l *LogManager) Panic(message string, fields map[string]interface{}) {
	l.Driver("default").Panic(message, fields)
}
