package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type LoggerOfLogrus struct {
	contextLogger *log.Entry
}

func NewLoggerOfLogrus() *LoggerOfLogrus {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	contextLogger := log.WithFields(log.Fields{})
	return &LoggerOfLogrus{
		contextLogger: contextLogger,
	}
}

func (l *LoggerOfLogrus) Trace(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Trace(message)
}

func (l *LoggerOfLogrus) Debug(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Debug(message)
}

func (l *LoggerOfLogrus) Info(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Info(message)
}

func (l *LoggerOfLogrus) Warn(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Warn(message)
}

func (l *LoggerOfLogrus) Error(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Error(message)
}

func (l *LoggerOfLogrus) Fatal(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Fatal(message)
}

func (l *LoggerOfLogrus) Panic(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Panic(message)
}
